package sync_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"path"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogo/protobuf/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/sync"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

func TestSyncService(t *testing.T) {
	tests := []struct {
		name     string
		ops      func(*testing.T, services.WriteService)
		watchers integration.Watchers
	}{
		{
			name: "some initial resources are dumped",
			watchers: integration.Watchers{
				"/test/virtual_network/5720afd0-d5a6-46ef-bd81-3be7f715cd27": []integration.Event{
					{
						"uuid":        "5720afd0-d5a6-46ef-bd81-3be7f715cd27",
						"parent_uuid": "beefbeef-beef-beef-beef-beefbeef0003",
						"parent_type": "project",
						"fq_name":     []interface{}{"default-domain", "default-project", "default-virtual-network"},
						"id_perms": map[string]interface{}{
							"enable":        true,
							"created":       "2018-05-23T17:29:57.559916",
							"user_visible":  true,
							"last_modified": "2018-05-23T17:29:57.559916",
							"permissions": map[string]interface{}{
								"owner":        "cloud-admin",
								"owner_access": 7,
								"other_access": 7,
								"group":        "cloud-admin-group",
								"group_access": 7,
							},
							"uuid": map[string]interface{}{
								"uuid_mslong": 6278211400000000000,
								"uuid_lslong": float64(13655262000000000000)},
						},
						"perms2": map[string]interface{}{
							"owner":        "cloud-admin",
							"owner_access": 7,
						},
						"virtual_network_network_id": 1,
						"routing_instances": []interface{}{
							map[string]interface{}{
								"uuid": "d59c5934-1dbd-4865-b8e9-ff9d7f3f16d0",
								"fq_name": []interface{}{
									"default-domain",
									"default-project",
									"default-virtual-network",
									"default-virtual-network",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "create and update virtual network",
			ops: func(t *testing.T, sv services.WriteService) {
				ctx := context.Background()
				_, err := sv.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{
					VirtualNetwork: &models.VirtualNetwork{
						UUID: "vn-blue",
						Name: "vn_blue",
					},
				})
				require.NoError(t, err, "create virtual network failed")

				_, err = sv.UpdateVirtualNetwork(ctx, &services.UpdateVirtualNetworkRequest{
					VirtualNetwork: &models.VirtualNetwork{
						UUID: "vn-blue",
						Name: "vn_bluuee",
					},
					FieldMask: types.FieldMask{Paths: []string{"name"}},
				})
				assert.NoError(t, err, "update virtual network failed")

				_, err = sv.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: "vn-blue"})
				assert.NoError(t, err, "delete virtual network failed")
			},
			watchers: integration.Watchers{
				"/test/virtual_network/vn-blue": []integration.Event{
					{
						"name": "vn_blue",
					},
					{
						"name": "vn_bluuee",
					},
					nil,
				},
			},
		},
		{
			name: "create and delete reference from virtual network to network IPAM",
			ops: func(t *testing.T, sv services.WriteService) {
				ctx := context.Background()
				_, err := sv.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{
					VirtualNetwork: &models.VirtualNetwork{
						UUID: "vn-blue",
						Name: "vn_blue",
					},
				})
				assert.NoError(t, err, "create virtual network failed")

				_, err = sv.CreateNetworkIpam(ctx, &services.CreateNetworkIpamRequest{
					NetworkIpam: &models.NetworkIpam{
						UUID: "ni-blue",
						Name: "ni_blue",
					},
				})
				assert.NoError(t, err, "create network IPAM failed")

				_, err = sv.CreateVirtualNetworkNetworkIpamRef(ctx,
					&services.CreateVirtualNetworkNetworkIpamRefRequest{
						ID: "vn-blue",
						VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
							UUID: "ni-blue",
							Attr: &models.VnSubnetsType{HostRoutes: &models.RouteTableType{
								Route: []*models.RouteType{{Prefix: "test_prefix", NextHop: "1.2.3.5"}},
							}},
						},
					},
				)
				assert.NoError(t, err, "create vn-ni reference failed")

				_, err = sv.DeleteVirtualNetworkNetworkIpamRef(ctx,
					&services.DeleteVirtualNetworkNetworkIpamRefRequest{
						ID: "vn-blue",
						VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
							UUID: "ni-blue",
						}})
				assert.NoError(t, err, "delete vn-ni reference failed")

				_, err = sv.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: "vn-blue"})
				assert.NoError(t, err, "delete virtual network failed")

				_, err = sv.DeleteNetworkIpam(ctx, &services.DeleteNetworkIpamRequest{ID: "ni-blue"})
				assert.NoError(t, err, "delete network IPAM failed")
			},
			watchers: integration.Watchers{
				"/test/virtual_network/vn-blue": []integration.Event{
					{
						"name":              "vn_blue",
						"network_ipam_refs": "$null",
					},
					{
						"name": "vn_blue",
						"network_ipam_refs": []interface{}{map[string]interface{}{
							"uuid": "ni-blue",
							"attr": map[string]interface{}{
								"ipam_subnets": nil,
								"host_routes": map[string]interface{}{
									"route": []interface{}{map[string]interface{}{
										"next_hop": "1.2.3.5",
										"prefix":   "test_prefix",
									}},
								},
							},
						}},
					},
					{
						"name":              "vn_blue",
						"network_ipam_refs": "$null",
					},
					nil,
				},
				"/test/network_ipam/ni-blue": []integration.Event{
					{
						"name": "ni_blue",
						"virtual_network_back_refs": "$null",
					},
					{
						"name": "ni_blue",
						"virtual_network_back_refs": []interface{}{
							map[string]interface{}{
								"uuid": "vn-blue",
							},
						},
					},
					{
						"name": "ni_blue",
						"virtual_network_back_refs": "$null",
					},
					nil,
				},
			},
		},
	}

	etcdPath := "test"
	viper.Set("etcd.path", etcdPath)
	integration.SetDefaultSyncConfig()

	dbService, err := db.NewServiceFromConfig()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := integrationetcd.NewEtcdClient(t)
			defer ec.Close(t)

			ec.Clear(t)

			check := integration.StartWatchers(t, tt.watchers)

			sync, err := sync.NewService()
			require.NoError(t, err)

			defer integration.RunNoError(t, sync)(t)

			if tt.ops != nil {
				tt.ops(t, dbService)
			}

			check(t)
		})
	}
}

func TestSyncSynchronizesExistingPostgresDataToEtcd(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: false,
	})
	defer s.CloseT(t)
	hc := integration.NewTestingHTTPClient(t, s.URL())
	ec := integrationetcd.NewEtcdClient(t)
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

	vnRedWatch, redCtx, cancelRedCtx := ec.WatchResource(integrationetcd.VirtualNetworkSchemaID, vnRedUUID)
	defer cancelRedCtx()
	vnGreenWatch, greenCtx, cancelGreenCtx := ec.WatchResource(integrationetcd.VirtualNetworkSchemaID, vnGreenUUID)
	defer cancelGreenCtx()
	vnBlueWatch, blueCtx, cancelBlueCtx := ec.WatchResource(integrationetcd.VirtualNetworkSchemaID, vnBlueUUID)
	defer cancelBlueCtx()

	integration.CreateProject(t, hc, project(projectUUID))
	defer integration.DeleteProject(t, hc, projectUUID)

	integration.CreateNetworkIpam(t, hc, networkIPAM(networkIPAMUUID, projectUUID))
	defer integration.DeleteNetworkIpam(t, hc, networkIPAMUUID)

	integration.CreateVirtualNetwork(t, hc, virtualNetworkRed(vnRedUUID, projectUUID, networkIPAMUUID))
	integration.CreateVirtualNetwork(t, hc, virtualNetworkGreen(vnGreenUUID, projectUUID, networkIPAMUUID))
	integration.CreateVirtualNetwork(t, hc, virtualNetworkBlue(vnBlueUUID, projectUUID, networkIPAMUUID))
	defer deleteVirtualNetworksFromAPIServer(t, hc, vnUUIDs)
	defer ec.DeleteKey(t, integrationetcd.JSONEtcdKey(integrationetcd.VirtualNetworkSchemaID, ""),
		clientv3.WithPrefix()) // delete all VNs

	vnRed := integration.GetVirtualNetwork(t, hc, vnRedUUID)
	vnGreen := integration.GetVirtualNetwork(t, hc, vnGreenUUID)
	vnBlue := integration.GetVirtualNetwork(t, hc, vnBlueUUID)

	integration.SetDefaultSyncConfig()
	sync, err := sync.NewService()
	require.NoError(t, err)

	defer integration.RunNoError(t, sync)(t)

	redEvent := integrationetcd.RetrieveCreateEvent(redCtx, t, vnRedWatch)
	greenEvent := integrationetcd.RetrieveCreateEvent(greenCtx, t, vnGreenWatch)
	blueEvent := integrationetcd.RetrieveCreateEvent(blueCtx, t, vnBlueWatch)

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

func checkNoSuchVirtualNetworksInEtcd(t *testing.T, ec *integrationetcd.EtcdClient, uuids []string) {
	for _, uuid := range uuids {
		ec.CheckKeyDoesNotExist(t, integrationetcd.JSONEtcdKey(integrationetcd.VirtualNetworkSchemaID, uuid))
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
		integration.DeleteVirtualNetwork(t, hc, uuid)
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
