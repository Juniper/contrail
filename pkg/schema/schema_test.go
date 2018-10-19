package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchema(t *testing.T) {
	api, err := MakeAPI("test_data/schema")
	assert.Nil(t, err, "API reading failed")
	assert.Equal(t, 4, len(api.Types))
	assert.Equal(t, 4, len(api.Schemas))
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

func TestReferencesExtendBase(t *testing.T) {
	api, err := MakeAPI([]string{"test_data/schema_extend"}, "")
	require.Nil(t, err, "API reading failed")
	assert.Equal(t, 5, len(api.Schemas))

	base := api.SchemaByID("base")
	require.NotNil(t, base, "Base object can't be <nil>")
	assert.Equal(t, 1, len(base.ReferencesSlice))
	assert.Equal(t, 0, len(base.References)) // References in base schemas are not processed

	zeroRefObj := api.SchemaByID("derived_object")
	require.NotNil(t, zeroRefObj, "derived_object schema shouldn't be <nil>")
	assert.Equal(t, 1, len(zeroRefObj.ReferencesSlice))
	assert.Equal(t, 1, len(zeroRefObj.References))

	ownRefObj := api.SchemaByID("derived_own_refs_object")
	require.NotNil(t, ownRefObj, "derived_own_refs_object schema shouldn't be <nil>")
	assert.Equal(t, 2, len(ownRefObj.ReferencesSlice))
	assert.Equal(t, 2, len(ownRefObj.References))
}

func TestReferenceTableName(t *testing.T) {
	assert.Equal(
		t,
		"ref__v_net_i_v_net_i_v_net_i_v_net_i_v_net_i",
		ReferenceTableName("ref_", "virtual_network_interface_virtual_network_interface",
			"virtual_network_interface_virtual_network_interface_virtual_network_interface"))
}
