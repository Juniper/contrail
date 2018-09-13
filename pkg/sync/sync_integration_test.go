package sync_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"path"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestSyncSynchronizesExistingPostgresDataToEtcd(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: false,
	})
	defer s.CloseT(t)
	hc := integration.NewTestingHTTPClient(t, s.URL())
	ec := integration.NewEtcdClient(t)
	defer ec.Close(t)

	testID := generateTestID(t)
	projectUUID := testID + "-project"
	networkIPAMUUID := testID + "-network-ipam"
	vnRedUUID := testID + "-red-vn"
	vnGreenUUID := testID + "-green-vn"
	vnBlueUUID := testID + "-blue-vn"
	vnUUIDs := []string{vnRedUUID, vnGreenUUID, vnBlueUUID}

	checkNoSuchVirtualNetworksInAPIServer(t, hc, vnUUIDs)
	checkNoSuchVirtualNetworksInEtcd(t, ec, vnUUIDs)

	vnRedWatch, redCtx, cancelRedCtx := ec.WatchResource(integration.VirtualNetworkSchemaID, vnRedUUID)
	defer cancelRedCtx()
	vnGreenWatch, greenCtx, cancelGreenCtx := ec.WatchResource(integration.VirtualNetworkSchemaID, vnGreenUUID)
	defer cancelGreenCtx()
	vnBlueWatch, blueCtx, cancelBlueCtx := ec.WatchResource(integration.VirtualNetworkSchemaID, vnBlueUUID)
	defer cancelBlueCtx()

	hc.CreateRequiredProject(t, project(projectUUID))
	defer hc.RemoveProject(t, projectUUID)
	defer ec.DeleteProject(t, projectUUID)

	hc.CreateRequiredNetworkIPAM(t, networkIPAM(networkIPAMUUID, projectUUID))
	defer hc.RemoveNetworkIPAM(t, networkIPAMUUID)
	defer ec.DeleteNetworkIPAM(t, networkIPAMUUID)

	hc.CreateRequiredVirtualNetwork(t, virtualNetworkRed(vnRedUUID, projectUUID, networkIPAMUUID))
	hc.CreateRequiredVirtualNetwork(t, virtualNetworkGreen(vnGreenUUID, projectUUID, networkIPAMUUID))
	hc.CreateRequiredVirtualNetwork(t, virtualNetworkBlue(vnBlueUUID, projectUUID, networkIPAMUUID))
	defer deleteVirtualNetworksFromAPIServer(t, hc, vnUUIDs)
	defer ec.DeleteKey(t, integration.JSONEtcdKey(integration.VirtualNetworkSchemaID, ""),
		clientv3.WithPrefix()) // delete all VNs

	vnRed := hc.FetchVirtualNetwork(t, vnRedUUID)
	vnGreen := hc.FetchVirtualNetwork(t, vnGreenUUID)
	vnBlue := hc.FetchVirtualNetwork(t, vnBlueUUID)

	closeSync := integration.RunSyncService(t)
	defer closeSync()

	redEvent := integration.RetrieveCreateEvent(redCtx, t, vnRedWatch)
	greenEvent := integration.RetrieveCreateEvent(greenCtx, t, vnGreenWatch)
	blueEvent := integration.RetrieveCreateEvent(blueCtx, t, vnBlueWatch)

	checkSyncedVirtualNetwork(t, redEvent, vnRed)
	checkSyncedVirtualNetwork(t, greenEvent, vnGreen)
	checkSyncedVirtualNetwork(t, blueEvent, vnBlue)
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
		ec.CheckKeyDoesNotExist(t, integration.JSONEtcdKey(integration.VirtualNetworkSchemaID, uuid))
	}
}

func project(uuid string) *models.Project {
	return &models.Project{
		UUID:       uuid,
		ParentType: integration.DomainType,
		ParentUUID: integration.DefaultDomainUUID,
		FQName:     []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Quota:      &models.QuotaType{},
	}
}

func networkIPAM(uuid string, parentUUID string) *models.NetworkIpam {
	return &models.NetworkIpam{
		UUID:       uuid,
		ParentType: integration.ProjectType,
		ParentUUID: parentUUID,
		FQName:     []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
	}
}

func virtualNetworkRed(uuid, parentUUID, networkIPAMUUID string) *models.VirtualNetwork {
	return &models.VirtualNetwork{
		UUID:       uuid,
		ParentType: integration.ProjectType,
		ParentUUID: parentUUID,
		FQName:     []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Perms2:     &models.PermType2{Owner: integration.AdminUserID},
		RouteTargetList: &models.RouteTargetList{
			RouteTarget: []string{"100:200"},
		},
		DisplayName:        "red",
		MacLearningEnabled: true,
		NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{{
			UUID: networkIPAMUUID,
			To:   []string{integration.DefaultDomainID, integration.AdminProjectID, networkIPAMUUID + "-fq-name"},
		}},
	}
}

func virtualNetworkGreen(uuid, parentUUID, networkIPAMUUID string) *models.VirtualNetwork {
	return &models.VirtualNetwork{
		UUID:                uuid,
		ParentType:          integration.ProjectType,
		ParentUUID:          parentUUID,
		FQName:              []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Perms2:              &models.PermType2{Owner: integration.AdminUserID},
		DisplayName:         "green",
		PortSecurityEnabled: true,
		NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{{
			UUID: networkIPAMUUID,
			To:   []string{integration.DefaultDomainID, integration.AdminProjectID, networkIPAMUUID + "-fq-name"},
		}},
	}
}

func virtualNetworkBlue(uuid, parentUUID, networkIPAMUUID string) *models.VirtualNetwork {
	return &models.VirtualNetwork{
		UUID:        uuid,
		ParentType:  integration.ProjectType,
		ParentUUID:  parentUUID,
		FQName:      []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Perms2:      &models.PermType2{Owner: integration.AdminUserID},
		DisplayName: "blue",
		FabricSnat:  true,
		NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{{
			UUID: networkIPAMUUID,
			To:   []string{integration.DefaultDomainID, integration.AdminProjectID, networkIPAMUUID + "-fq-name"},
		}},
	}
}

func deleteVirtualNetworksFromAPIServer(t *testing.T, hc *integration.HTTPAPIClient, uuids []string) {
	for _, uuid := range uuids {
		hc.RemoveVirtualNetwork(t, uuid)
	}
}

func checkSyncedVirtualNetwork(t *testing.T, event *clientv3.Event, expectedVN *models.VirtualNetwork) {
	syncedVN := decodeVirtualNetworkJSON(t, event.Kv.Value)
	assert.Equal(t, expectedVN, syncedVN, "synced VN does not match created VN")
}

func decodeVirtualNetworkJSON(t *testing.T, vnBytes []byte) *models.VirtualNetwork {
	var vn models.VirtualNetwork
	assert.NoError(t, json.Unmarshal(vnBytes, &vn))
	return &vn
}
