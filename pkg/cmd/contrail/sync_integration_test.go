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
	s := integration.NewRunningAPIServer(t, "../../..")
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())
	ec := integration.NewEtcdClient(t)
	defer ec.Close(t)

	testID := generateTestID(t)
	vnRedUUID := testID + "-red-vn"
	vnGreenUUID := testID + "-green-vn"
	vnBlueUUID := testID + "-blue-vn"
	uuids := []string{vnRedUUID, vnGreenUUID, vnBlueUUID}

	checkNoSuchVirtualNetworksInAPIServer(t, hc, uuids)
	checkNoSuchVirtualNetworksInEtcd(t, ec, uuids)

	vnRedWatch, redCtx, cancelRedCtx := watchVirtualNetworkInEtcd(ec, vnRedUUID)
	defer cancelRedCtx()
	vnGreenWatch, greenCtx, cancelGreenCtx := watchVirtualNetworkInEtcd(ec, vnGreenUUID)
	defer cancelGreenCtx()
	vnBlueWatch, blueCtx, cancelBlueCtx := watchVirtualNetworkInEtcd(ec, vnBlueUUID)
	defer cancelBlueCtx()

	createThreeVirtualNetworks(t, hc, uuids)
	defer deleteThreeVirtualNetworks(t, hc, uuids)

	vnRed := readVirtualNetwork(t, hc, vnRedUUID)
	vnGreen := readVirtualNetwork(t, hc, vnGreenUUID)
	vnBlue := readVirtualNetwork(t, hc, vnBlueUUID)

	closeSync := integration.RunSync(t)
	defer closeSync()

	wrRed := retrieveEtcdCreateEvent(redCtx, t, vnRedWatch)
	wrGreen := retrieveEtcdCreateEvent(greenCtx, t, vnGreenWatch)
	wrBlue := retrieveEtcdCreateEvent(blueCtx, t, vnBlueWatch)

	checkSyncedVirtualNetwork(t, wrRed, vnRed)
	checkSyncedVirtualNetwork(t, wrGreen, vnGreen)
	checkSyncedVirtualNetwork(t, wrBlue, vnBlue)

	waitForSyncToFinishDump() // TODO(Daniel): change Sync not to throw error on Dump context cancellation
}

// generateTestID creates pseudo-random string and is used to create resources with
// unique UUIDs and FQNames.
func generateTestID(t *testing.T) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v-%v", t.Name(), rand.Uint64())
}

func checkNoSuchVirtualNetworksInAPIServer(t *testing.T, hc *integration.HTTPAPIClient, uuids []string) {
	for _, uuid := range uuids {
		hc.CheckResourceDoesNotExist(t, path.Join(virtualNetworkSingularPath, uuid))
	}
}

func checkNoSuchVirtualNetworksInEtcd(t *testing.T, ec *integration.EtcdClient, uuids []string) {
	for _, uuid := range uuids {
		ec.CheckKeyDoesNotExist(t, virtualNetworkEtcdKey(uuid))
	}
}

func watchVirtualNetworkInEtcd(ec *integration.EtcdClient, uuid string) (clientv3.WatchChan, context.Context,
	context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), etcdWatchTimeout)
	w := ec.Watch(ctx, virtualNetworkEtcdKey(uuid))
	return w, ctx, cancel
}

func virtualNetworkEtcdKey(uuid string) string {
	return path.Join(etcdJSONPrefix, virtualNetworkKey, uuid)
}

//func createProject(t *testing.T, hc *integration.HTTPAPIClient) {
//
//}
//
//func createNetworkPolicy() {
//
//}

func createThreeVirtualNetworks(t *testing.T, hc *integration.HTTPAPIClient, uuids []string) {
	hc.CreateResource(
		t,
		virtualNetworkPluralPath,
		&models.CreateVirtualNetworkRequest{VirtualNetwork: &models.VirtualNetwork{
			UUID:       uuids[0],
			ParentType: "project",
			ParentUUID: "admin_project_uuid",
			FQName:     []string{"default", "admin", uuids[0] + "-fq-name"},
			Perms2:     &models.PermType2{Owner: "admin"},
			RouteTargetList: &models.RouteTargetList{
				RouteTarget: []string{"100:200"},
			},
			DisplayName:        "red",
			MacLearningEnabled: true,
			//NetworkPolicyRefs: []*models.VirtualNetworkNetworkPolicyRef{
			//	{UUID: "network_policy_uuid_red"},
			//},
		}},
	)

	hc.CreateResource(
		t,
		virtualNetworkPluralPath,
		&models.CreateVirtualNetworkRequest{VirtualNetwork: &models.VirtualNetwork{
			UUID:                uuids[1],
			ParentType:          "project",
			ParentUUID:          "admin_project_uuid",
			FQName:              []string{"default", "admin", uuids[1] + "-fq-name"},
			Perms2:              &models.PermType2{Owner: "admin"},
			DisplayName:         "green",
			PortSecurityEnabled: true,
			//NetworkPolicyRefs: []*models.VirtualNetworkNetworkPolicyRef{
			//	{UUID: "network_policy_uuid"},
			//},
		}},
	)

	hc.CreateResource(
		t,
		virtualNetworkPluralPath,
		&models.CreateVirtualNetworkRequest{VirtualNetwork: &models.VirtualNetwork{
			UUID:        uuids[2],
			ParentType:  "project",
			ParentUUID:  "admin_project_uuid",
			FQName:      []string{"default", "admin", uuids[2] + "-fq-name"},
			Perms2:      &models.PermType2{Owner: "admin"},
			DisplayName: "blue",
			FabricSnat:  true,
			//NetworkPolicyRefs: []*models.VirtualNetworkNetworkPolicyRef{
			//	{UUID: "network_policy_uuid"},
			//},
		}},
	)
}

func deleteThreeVirtualNetworks(t *testing.T, hc *integration.HTTPAPIClient, uuids []string) {
	for _, uuid := range uuids {
		hc.DeleteResource(t, path.Join(virtualNetworkSingularPath, uuid))
	}
}

func readVirtualNetwork(t *testing.T, hc *integration.HTTPAPIClient, uuid string) *models.VirtualNetwork {
	var responseData models.GetVirtualNetworkResponse
	hc.GetResource(t, virtualNetworkSingularPath+"/"+uuid, &responseData)
	return responseData.VirtualNetwork
}

func retrieveEtcdCreateEvent(ctx context.Context, t *testing.T, watch clientv3.WatchChan) *clientv3.WatchResponse {
	wr := <-watch
	assert.NoError(t, wr.Err(), "watching virtual network failed")
	assert.NotEqual(t, ctx.Err(), context.DeadlineExceeded, "watching virtual network timed out")
	assert.Equal(t, 1, len(wr.Events))
	assert.True(t, wr.Events[0].IsCreate())
	return &wr
}

func checkSyncedVirtualNetwork(t *testing.T, wr *clientv3.WatchResponse, expectedVN *models.VirtualNetwork) {
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
