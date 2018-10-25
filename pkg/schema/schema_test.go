package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchema(t *testing.T) {
	api, err := MakeAPI([]string{"test_data/schema"}, "")
	assert.Nil(t, err, "API reading failed")
	assert.Equal(t, 4, len(api.Types))
	assert.Equal(t, 4, len(api.Schemas))
	project := api.SchemaByID("project")

	assert.Equal(t, "project", project.Table)
	assert.Equal(t, 3, len(project.JSONSchema.Properties))
	assert.Equal(t, 3, len(project.JSONSchema.OrderedProperties))
	assert.Equal(t, 3, len(project.Columns))

	virtualNetwork := api.SchemaByID("virtual_network")
	assert.Equal(t, "vn", virtualNetwork.Table)
	assert.Equal(t, 4, len(virtualNetwork.JSONSchema.Properties))
	assert.Equal(t, "uint64", virtualNetwork.JSONSchema.Properties["version"].GoType)
	assert.Equal(t, []string{"uuid", "version", "display_name", "virtual_network_network_id"},
		virtualNetwork.JSONSchema.PropertiesOrder)
	assert.Equal(t, 4, len(virtualNetwork.Columns))
	assert.Equal(t, 1005, virtualNetwork.References["network_ipam"].Index)
}

func TestSchemaEnums(t *testing.T) {
	api, err := MakeAPI([]string{"test_data/schema_enums"}, "overrides")
	assert.Nil(t, err, "API reading failed")
	project := api.SchemaByID("project")
	assert.NotNil(t, project, "Project can't be <nil>")
	obj := api.SchemaByID("simple_object")
	assert.NotNil(t, obj, "SimpleObject can't be <nil>")
	// In addition 'uuid' and 'display_name' are added (+2)
	assert.Equal(t, 3+2, len(obj.JSONSchema.Properties))

	assert.NotNil(t, api.Types)
	assert.Equal(t, 4, len(api.Types))
	enumArr, ok := api.Types["ObjectThatReferencesEnumAsArray"]
	assert.True(t, ok)
	checkPropertyRepeated(t, enumArr)
	enumArrOvrd, ok := api.Types["ObjectThatReferencesEnumAsArrayOverriden"]
	assert.True(t, ok)
	checkPropertyRepeated(t, enumArrOvrd)
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

func checkPropertyRepeated(t *testing.T, obj *JSONSchema) {
	assert.NotNil(t, obj)
	assert.Equal(t, 1, len(obj.Properties))
	assert.NotNil(t, obj.Properties[obj.OrderedProperties[0].ID].Items)
}

func TestReferenceTableName(t *testing.T) {
	assert.Equal(
		t,
		"ref__v_net_i_v_net_i_v_net_i_v_net_i_v_net_i",
		ReferenceTableName("ref_", "virtual_network_interface_virtual_network_interface",
			"virtual_network_interface_virtual_network_interface_virtual_network_interface"))
}
