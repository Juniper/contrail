package db

import (
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
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

func TestDBCreateRef(t *testing.T) {
	tests := []struct {
		name string
		//request  services.CreateAccessControlListRequest
		fails    bool
		expected proto.Message
	}{
		{name: "empty", fails: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO(Michał)
		})
	}
}
