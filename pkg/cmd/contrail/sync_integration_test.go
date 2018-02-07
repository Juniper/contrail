package contrail

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"path"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	etcdJSONPrefix             = "json"
	virtualNetworkKey          = "virtual_network"
	virtualNetworkSingularPath = "/virtual-network"
	virtualNetworkPluralPath   = "/virtual-networks"
	etcdWatchTimeout           = 10 * time.Second
)

func TestSyncSynchronizesPostgresDataToEtcdUsingJSONStorage(t *testing.T) {
	t.Skip("Skipping because of not sufficient CI test env") // TODO(Daniel): resolve

	s := integration.NewRunningAPIServer(t, "../../..")
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())
	ec := integration.NewEtcdClient(t)
	defer ec.Close(t)

	testID := generateTestID(t)
	vnRedUUID := testID + "-red-uuid"
	vnGreenUUID := testID + "-green-uuid"
	vnBlueUUID := testID + "-blue-uuid"
	uuids := []string{vnRedUUID, vnGreenUUID, vnBlueUUID}

	checkNoVirtualNetworksInAPIServer(t, hc, uuids)
	checkNoVirtualNetworksInEtcd(t, ec, uuids)

	vnRedWatch, redCtx := watchVirtualNetworkInEtcd(ec, vnRedUUID)
	vnGreenWatch, greenCtx := watchVirtualNetworkInEtcd(ec, vnGreenUUID)
	vnBlueWatch, blueCtx := watchVirtualNetworkInEtcd(ec, vnBlueUUID)

	createThreeVirtualNetworks(t, hc, uuids)

	vnRed := readVirtualNetwork(t, hc, vnRedUUID)
	vnGreen := readVirtualNetwork(t, hc, vnGreenUUID)
	vnBlue := readVirtualNetwork(t, hc, vnBlueUUID)

	sync := integration.NewRunningSync(t)
	defer sync.Close(t)

	wrRed := retrieveWatchResponse(redCtx, t, vnRedWatch)
	wrGreen := retrieveWatchResponse(greenCtx, t, vnGreenWatch)
	wrBlue := retrieveWatchResponse(blueCtx, t, vnBlueWatch)

	checkSyncedVirtualNetwork(t, wrRed, vnRed)
	checkSyncedVirtualNetwork(t, wrGreen, vnGreen)
	checkSyncedVirtualNetwork(t, wrBlue, vnBlue)

	hc.DeleteResource(t, path.Join(virtualNetworkSingularPath, vnRedUUID))
	hc.DeleteResource(t, path.Join(virtualNetworkSingularPath, vnGreenUUID))
	hc.DeleteResource(t, path.Join(virtualNetworkSingularPath, vnBlueUUID))

	waitForSyncToFinishDump() // TODO(Daniel): change Sync not to throw error on Dump context cancellation
}

func generateTestID(t *testing.T) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v-%v", t.Name(), rand.Uint64())
}

func checkNoVirtualNetworksInAPIServer(t *testing.T, hc *integration.HTTPAPIClient, uuids []string) {
	log.WithField("uuids", uuids).Debug("Checking that there are no test-specific resources in API Server's DB")
	for _, uuid := range uuids {
		hc.CheckResourceDoesNotExist(t, path.Join(virtualNetworkSingularPath, uuid))
	}
}

func checkNoVirtualNetworksInEtcd(t *testing.T, ec *integration.EtcdClient, uuids []string) {
	log.WithField("uuids", uuids).Debug("Checking that there are no test-specific resources in etcd")
	for _, uuid := range uuids {
		ec.CheckKeyDoesNotExist(t, virtualNetworkEtcdKey(uuid))
	}
}

func watchVirtualNetworkInEtcd(ec *integration.EtcdClient, uuid string) (clientv3.WatchChan, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), etcdWatchTimeout)
	w := ec.Watch(ctx, virtualNetworkEtcdKey(uuid))
	return w, ctx
}

func virtualNetworkEtcdKey(uuid string) string {
	return path.Join(etcdJSONPrefix, virtualNetworkKey, uuid)
}

func createThreeVirtualNetworks(t *testing.T, hc *integration.HTTPAPIClient, uuids []string) {
	hc.CreateResource(
		t,
		virtualNetworkPluralPath,
		&models.CreateVirtualNetworkRequest{VirtualNetwork: &models.VirtualNetwork{
			UUID:                uuids[0],
			FQName:              []string{"default", "admin", uuids[0] + "-fq-name"},
			Perms2:              &models.PermType2{Owner: "admin"},
			MacLearningEnabled:  true,
			PortSecurityEnabled: true,
			FabricSnat:          true,
		}},
	)

	hc.CreateResource(
		t,
		virtualNetworkPluralPath,
		&models.CreateVirtualNetworkRequest{VirtualNetwork: &models.VirtualNetwork{
			UUID:           uuids[1],
			FQName:         []string{"default", "admin", uuids[1] + "-fq-name"},
			Perms2:         &models.PermType2{Owner: "admin"},
			FabricSnat:     true,
			PBBEtreeEnable: true,
			IsShared:       true,
		}},
	)

	hc.CreateResource(
		t,
		virtualNetworkPluralPath,
		&models.CreateVirtualNetworkRequest{VirtualNetwork: &models.VirtualNetwork{
			UUID:   uuids[2],
			FQName: []string{"default", "admin", uuids[2] + "-fq-name"},
			Perms2: &models.PermType2{Owner: "admin"},
		}},
	)
}

func readVirtualNetwork(t *testing.T, hc *integration.HTTPAPIClient, uuid string) *models.VirtualNetwork {
	var responseData models.GetVirtualNetworkResponse
	hc.GetResource(t, virtualNetworkSingularPath+"/"+uuid, &responseData)
	return responseData.VirtualNetwork
}

func retrieveWatchResponse(ctx context.Context, t *testing.T, watch clientv3.WatchChan) *clientv3.WatchResponse {
	wr := <-watch
	assert.NoError(t, wr.Err(), "watching virtual network failed")
	assert.NotEqual(t, ctx.Err(), context.DeadlineExceeded, "watching virtual network timed out")
	return &wr
}

func checkSyncedVirtualNetwork(t *testing.T, wr *clientv3.WatchResponse, expectedVN *models.VirtualNetwork) {
	log.Debug("Checking that virtual network synced to etcd has correct value")
	assert.Equal(t, 1, len(wr.Events))
	assert.True(t, wr.Events[0].IsCreate())
	syncedVN := decodeVirtualNetworkJSON(t, wr.Events[0].Kv.Value)
	assert.Equal(t, expectedVN, syncedVN, "synced VN does not match created VN")
}

func decodeVirtualNetworkJSON(t *testing.T, vnBytes []byte) *models.VirtualNetwork {
	var vn models.VirtualNetwork
	assert.NoError(t, json.Unmarshal(vnBytes, &vn))
	return &vn
}

func waitForSyncToFinishDump() {
	time.Sleep(100 * time.Millisecond)
}
