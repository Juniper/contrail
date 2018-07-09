package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	api, err := MakeAPI("test_data/schema")
	assert.Nil(t, err, "API reading failed")
	assert.Equal(t, 4, len(api.Types))
	assert.Equal(t, 4, len(api.Schemas))

	base := api.SchemaByID("base")
	assert.True(t, base.JSONSchema.Properties["uuid"].Unique, "unique property not set")

	project := api.SchemaByID("project")

	assert.Equal(t, 2, len(project.JSONSchema.Properties))
	assert.Equal(t, 2, len(project.JSONSchema.OrderedProperties))
	assert.Equal(t, 2, len(project.Columns))

	virtualNetwork := api.SchemaByID("virtual_network")

	assert.Equal(t, 3, len(virtualNetwork.JSONSchema.Properties))
	assert.Equal(t, []string{"uuid", "display_name", "virtual_network_network_id"},
		virtualNetwork.JSONSchema.PropertiesOrder)
	assert.Equal(t, 3, len(virtualNetwork.Columns))
	assert.Equal(t, 1004, virtualNetwork.References["network_ipam"].Index)
}

func TestReferenceTableName(t *testing.T) {
	assert.Equal(
		t,
		"ref__v_net_i_v_net_i_v_net_i_v_net_i_v_net_i",
		ReferenceTableName("ref_", "virtual_network_interface_virtual_network_interface",
			"virtual_network_interface_virtual_network_interface_virtual_network_interface"))
}
