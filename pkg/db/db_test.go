package db

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDBScanRow(t *testing.T) {
	tests := []struct {
		name     string
		schemaID string
		row      map[string]interface{}
		fails    bool
		expected proto.Message
	}{
		{name: "empty", fails: true},
		{name: "empty with valid schemaID", schemaID: "logical_interface", expected: models.MakeLogicalInterface()},
		{name: "valid logical_interface_row", schemaID: "logical_interface",
			row: map[string]interface{}{
				"configuration_version":      1,
				"created":                    "test created",
				"creator":                    "test creator",
				"description":                "test description",
				"display_name":               "test display name",
				"enable":                     true,
				"fq_name":                    []byte(`["first", "second"]`),
				"global_access":              2,
				"group":                      "test group",
				"group_access":               3,
				"key_value_pair":             []byte(`[{"key": "some key", "value": "some value"}]`),
				"last_modified":              "test last modified",
				"logical_interface_type":     "test type",
				"logical_interface_vlan_tag": 4,
				"other_access":               5,
				"owner":                      "test owner",
				"owner_access":               6,
				"parent_type":                "test parent type",
				"parent_uuid":                "test parent uuid",
				"permissions_owner":          "test perms owner",
				"permissions_owner_access":   7,
				"share":                      []byte(`[{"tenant_access": 1337, "tenant": "leet"}]`),
				"user_visible":               true,
				"uuid":                       "test uuid",
				"uuid_lslong":                8,
				"uuid_mslong":                9,
			},
			expected: &models.LogicalInterface{
				UUID:       "test uuid",
				ParentUUID: "test parent uuid",
				ParentType: "test parent type",
				FQName:     []string{"first", "second"},
				IDPerms: &models.IdPermsType{
					Enable:       true,
					Description:  "test description",
					Created:      "test created",
					Creator:      "test creator",
					UserVisible:  true,
					LastModified: "test last modified",
					Permissions: &models.PermType{
						Owner:       "test perms owner",
						OwnerAccess: 7,
						OtherAccess: 5,
						Group:       "test group",
						GroupAccess: 3,
					},
					UUID: &models.UuidType{
						UUIDMslong: 9,
						UUIDLslong: 8,
					},
				},
				DisplayName: "test display name",
				Annotations: &models.KeyValuePairs{
					KeyValuePair: []*models.KeyValuePair{{Value: "some value", Key: "some key"}},
				},
				Perms2: &models.PermType2{
					Owner:        "test owner",
					OwnerAccess:  6,
					GlobalAccess: 2,
					Share:        []*models.ShareType{{TenantAccess: 1337, Tenant: "leet"}},
				},
				ConfigurationVersion:        1,
				LogicalInterfaceVlanTag:     4,
				LogicalInterfaceType:        "test type",
				VirtualMachineInterfaceRefs: nil,
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := db.ScanRow(tt.schemaID, tt.row)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, result, tt.expected)
			}
		})
	}
}

var exampleVN = &models.VirtualNetwork{
	UUID:       "vn_uuid",
	ParentType: "project",
	ParentUUID: "beefbeef-beef-beef-beef-beefbeef0003",
	FQName:     []string{"default-domain", "default-project", "vn-db-create-ref"},
}

var exampleRI = &models.RoutingInstance{
	UUID:       "ri_uuid",
	ParentType: "virtual-network",
	ParentUUID: "vn_uuid",
	FQName:     []string{"default-domain", "default-project", "vn-db-create-ref", "ri-db-create-ref"},
}

var exampleRT = &models.RouteTarget{
	UUID: "rt_uuid",
}

func TestDBCreateRef(t *testing.T) {
	vnUUID, riUUID, rtUUID := exampleVN.UUID, exampleRI.UUID, exampleRT.UUID

	tests := []struct {
		name     string
		request  services.CreateRoutingInstanceRouteTargetRefRequest
		fails    bool
		expected *services.CreateRoutingInstanceRouteTargetRefResponse
	}{
		{name: "empty", fails: true},
		{
			name: "objects missing",
			request: services.CreateRoutingInstanceRouteTargetRefRequest{
				ID:                            "foo",
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{UUID: "bar"},
			},
			fails: true,
		},
		{
			name: "valid ID invalid ref UUID",
			request: services.CreateRoutingInstanceRouteTargetRefRequest{
				ID:                            riUUID,
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{UUID: "bar"},
			},
			fails: true,
		},
		{
			name: "valid ID valid ref UUID",
			request: services.CreateRoutingInstanceRouteTargetRefRequest{
				ID:                            riUUID,
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{UUID: rtUUID},
			},
			expected: &services.CreateRoutingInstanceRouteTargetRefResponse{
				ID: riUUID,
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
					UUID: rtUUID,
					Attr: &models.InstanceTargetType{},
				},
			},
		},
		{
			name: "valid ID valid ref UUID with attrs",
			request: services.CreateRoutingInstanceRouteTargetRefRequest{
				ID: riUUID,
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
					UUID: rtUUID,
					Attr: &models.InstanceTargetType{ImportExport: "import:export"},
				},
			},
			expected: &services.CreateRoutingInstanceRouteTargetRefResponse{
				ID: riUUID,
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
					UUID: rtUUID,
					Attr: &models.InstanceTargetType{ImportExport: "import:export"},
				},
			},
		},
	}

	setup := func(t *testing.T) {
		ctx := context.Background()
		_, err := db.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{VirtualNetwork: exampleVN})
		require.NoError(t, err)
		_, err = db.CreateRouteTarget(ctx, &services.CreateRouteTargetRequest{RouteTarget: exampleRT})
		require.NoError(t, err)
		_, err = db.CreateRoutingInstance(ctx, &services.CreateRoutingInstanceRequest{RoutingInstance: exampleRI})
		require.NoError(t, err)
	}
	teardown := func(t *testing.T) {
		ctx := context.Background()
		_, err := db.DeleteRoutingInstance(ctx, &services.DeleteRoutingInstanceRequest{ID: riUUID})
		assert.NoError(t, err)
		_, err = db.DeleteRouteTarget(ctx, &services.DeleteRouteTargetRequest{ID: rtUUID})
		assert.NoError(t, err)
		_, err = db.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: vnUUID})
		assert.NoError(t, err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer teardown(t)
			setup(t)

			response, err := db.CreateRoutingInstanceRouteTargetRef(context.Background(), &tt.request)
			assert.Equal(t, tt.expected, response)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				getResp, err := db.GetRoutingInstance(context.Background(), &services.GetRoutingInstanceRequest{ID: riUUID})
				assert.NoError(t, err)

				assert.Len(t, getResp.RoutingInstance.RouteTargetRefs, 1)
			}
		})
	}
}

func TestDBDeleteRef(t *testing.T) {
	vnUUID, riUUID, rtUUID := exampleVN.UUID, exampleRI.UUID, exampleRT.UUID

	tests := []struct {
		name           string
		request        services.DeleteRoutingInstanceRouteTargetRefRequest
		fails          bool
		expected       *services.DeleteRoutingInstanceRouteTargetRefResponse
		shouldRefExist bool
	}{
		{name: "empty", fails: true},
		{
			name: "objects missing",
			request: services.DeleteRoutingInstanceRouteTargetRefRequest{
				ID:                            "foo",
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{UUID: "bar"},
			},
			expected: &services.DeleteRoutingInstanceRouteTargetRefResponse{
				ID:                            "foo",
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{UUID: "bar"},
			},
			shouldRefExist: true,
		},
		{
			name: "valid ID invalid ref UUID",
			request: services.DeleteRoutingInstanceRouteTargetRefRequest{
				ID:                            riUUID,
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{UUID: "bar"},
			},
			expected: &services.DeleteRoutingInstanceRouteTargetRefResponse{
				ID:                            riUUID,
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{UUID: "bar"},
			},
			shouldRefExist: true,
		},
		{
			name: "valid ID valid ref UUID",
			request: services.DeleteRoutingInstanceRouteTargetRefRequest{
				ID:                            riUUID,
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{UUID: rtUUID},
			},
			expected: &services.DeleteRoutingInstanceRouteTargetRefResponse{
				ID:                            riUUID,
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{UUID: rtUUID},
			},
			shouldRefExist: false,
		},
	}

	setup := func(t *testing.T) {
		ctx := context.Background()
		_, err := db.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{VirtualNetwork: exampleVN})
		require.NoError(t, err)
		_, err = db.CreateRouteTarget(ctx, &services.CreateRouteTargetRequest{RouteTarget: exampleRT})
		require.NoError(t, err)

		// create routing instance with ref to rout target
		ri := *exampleRI
		ri.RouteTargetRefs = []*models.RoutingInstanceRouteTargetRef{{UUID: rtUUID}}
		_, err = db.CreateRoutingInstance(ctx, &services.CreateRoutingInstanceRequest{RoutingInstance: &ri})
		require.NoError(t, err)
	}
	teardown := func(t *testing.T) {
		ctx := context.Background()
		_, err := db.DeleteRoutingInstance(ctx, &services.DeleteRoutingInstanceRequest{ID: riUUID})
		assert.NoError(t, err)
		_, err = db.DeleteRouteTarget(ctx, &services.DeleteRouteTargetRequest{ID: rtUUID})
		assert.NoError(t, err)
		_, err = db.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: vnUUID})
		assert.NoError(t, err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer teardown(t)
			setup(t)

			response, err := db.DeleteRoutingInstanceRouteTargetRef(context.Background(), &tt.request)
			assert.Equal(t, tt.expected, response)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				getResp, err := db.GetRoutingInstance(context.Background(), &services.GetRoutingInstanceRequest{ID: riUUID})
				assert.NoError(t, err)

				if tt.shouldRefExist {
					assert.Len(t, getResp.RoutingInstance.RouteTargetRefs, 1)
				} else {
					assert.Len(t, getResp.RoutingInstance.RouteTargetRefs, 0)
				}
			}
		})
	}
}
