package sync_test

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
	etcdJSONPrefix    = "json"
	etcdWatchTimeout  = 10 * time.Second
	projectType       = "project"
	virtualNetworkKey = "virtual_network"
)

func TestSyncSynchronizesExistingPostgresDataToEtcdUsingJSONStorage(t *testing.T) {
	s := integration.NewRunningAPIServer(t, "../../..")
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())
	ec := integration.NewEtcdClient(t)
	defer ec.Close(t)

	testID := generateTestID(t)
	projectUUID := testID + "-project"
	networkIPAMUUID := "-network-ipam"
	vnRedUUID := testID + "-red-vn"
	vnGreenUUID := testID + "-green-vn"
	vnBlueUUID := testID + "-blue-vn"
	vnUUIDs := []string{vnRedUUID, vnGreenUUID, vnBlueUUID}

	checkNoSuchVirtualNetworksInAPIServer(t, hc, vnUUIDs)
	checkNoSuchVirtualNetworksInEtcd(t, ec, vnUUIDs)

	vnRedWatch, redCtx, cancelRedCtx := watchVirtualNetworkInEtcd(ec, vnRedUUID)
	defer cancelRedCtx()
	vnGreenWatch, greenCtx, cancelGreenCtx := watchVirtualNetworkInEtcd(ec, vnGreenUUID)
	defer cancelGreenCtx()
	vnBlueWatch, blueCtx, cancelBlueCtx := watchVirtualNetworkInEtcd(ec, vnBlueUUID)
	defer cancelBlueCtx()

	hc.CreateProject(t, project(projectUUID))
	defer hc.DeleteProject(t, projectUUID)

	hc.CreateNetworkIPAM(t, networkIPAM(networkIPAMUUID, projectUUID))
	defer hc.DeleteNetworkIPAM(t, networkIPAMUUID)

	hc.CreateVirtualNetwork(t, virtualNetworkRed(vnRedUUID, projectUUID, networkIPAMUUID))
	hc.CreateVirtualNetwork(t, virtualNetworkGreen(vnGreenUUID, projectUUID, networkIPAMUUID))
	hc.CreateVirtualNetwork(t, virtualNetworkBlue(vnBlueUUID, projectUUID, networkIPAMUUID))
	defer deleteVirtualNetworks(t, hc, vnUUIDs)

	vnRed := hc.GetVirtualNetwork(t, vnRedUUID)
	vnGreen := hc.GetVirtualNetwork(t, vnGreenUUID)
	vnBlue := hc.GetVirtualNetwork(t, vnBlueUUID)

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
		hc.CheckResourceDoesNotExist(t, path.Join(integration.VirtualNetworkSingularPath, uuid))
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

func project(uuid string) *models.Project {
	return &models.Project{
		UUID:   uuid,
		FQName: []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
	}
}

func networkIPAM(uuid string, parentUUID string) *models.NetworkIpam {
	return &models.NetworkIpam{
		UUID:       uuid,
		ParentType: projectType,
		ParentUUID: parentUUID,
		FQName:     []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
	}
}

func virtualNetworkRed(uuid, parentUUID, networkIPAMUUID string) *models.VirtualNetwork {
	return &models.VirtualNetwork{
		UUID:       uuid,
		ParentType: projectType,
		ParentUUID: parentUUID,
		FQName:     []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Perms2:     &models.PermType2{Owner: integration.AdminUserID},
		RouteTargetList: &models.RouteTargetList{
			RouteTarget: []string{"100:200"},
		},
		DisplayName:        "red",
		MacLearningEnabled: true,
		//NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{{
		//	UUID: networkIPAMUUID,
		//	To: []string{integration.DefaultDomainID, integration.AdminProjectID, networkIPAMUUID + "-fq-name"},
		//}},
		// TODO: add refs
	}
}

func virtualNetworkGreen(uuid, parentUUID, networkIPAMUUID string) *models.VirtualNetwork {
	return &models.VirtualNetwork{
		UUID:                uuid,
		ParentType:          projectType,
		ParentUUID:          parentUUID,
		FQName:              []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Perms2:              &models.PermType2{Owner: integration.AdminUserID},
		DisplayName:         "green",
		PortSecurityEnabled: true,
		// TODO: add refs
	}
}

func virtualNetworkBlue(uuid, parentUUID, networkIPAMUUID string) *models.VirtualNetwork {
	return &models.VirtualNetwork{
		UUID:        uuid,
		ParentType:  projectType,
		ParentUUID:  parentUUID,
		FQName:      []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Perms2:      &models.PermType2{Owner: integration.AdminUserID},
		DisplayName: "blue",
		FabricSnat:  true,
		// TODO: add refs
	}
}

func deleteVirtualNetworks(t *testing.T, hc *integration.HTTPAPIClient, uuids []string) {
	for _, uuid := range uuids {
		hc.DeleteVirtualNetwork(t, uuid)
	}
}

func retrieveEtcdCreateEvent(ctx context.Context, t *testing.T, watch clientv3.WatchChan) *clientv3.WatchResponse {
	wr := <-watch
	assert.NoError(t, wr.Err(), "watching virtual network failed")
	assert.NotEqual(t, ctx.Err(), context.DeadlineExceeded, "watching virtual network timed out")
	if assert.Equal(t, 1, len(wr.Events)) {
		assert.True(t, wr.Events[0].IsCreate())
	}
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
