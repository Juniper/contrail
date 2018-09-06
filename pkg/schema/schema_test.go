package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findDefinitionByID(id string, api *API) *Schema {
	for _, def := range api.Definitions {
		if def.ID == id {
			return def
		}
	}
	return nil
}

func TestSchema(t *testing.T) {
	api, err := MakeAPI([]string{"test_data/schema"}, "")
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
	assert.Equal(t, true, ok)
	checkPropertyRepeated(t, enumArr)
	enumArrOvrd, ok := api.Types["ObjectThatReferencesEnumAsArrayOverriden"]
	assert.Equal(t, true, ok)
	checkPropertyRepeated(t, enumArrOvrd)
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
