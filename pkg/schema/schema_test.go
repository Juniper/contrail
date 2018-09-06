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
	for k, v := range api.Types {
		t.Logf("Type:    #####> %v => %#v", k, v)
	}
	for _, def := range api.Definitions {
		t.Logf("Def ---------> %#v\n----------->%#v", def, def.JSONSchema)
	}
	enumAsRef := api.Types["EnumTestObject"]
	assert.Nil(t, enumAsRef)
	t.Logf("enumm def schema: %#v", enumAsRef)
	assert.Nil(t, enumAsRef.Enum)
	assert.Equal(t, 3, len(enumAsRef.Enum))
}

func TestReferenceTableName(t *testing.T) {
	assert.Equal(
		t,
		"ref__v_net_i_v_net_i_v_net_i_v_net_i_v_net_i",
		ReferenceTableName("ref_", "virtual_network_interface_virtual_network_interface",
			"virtual_network_interface_virtual_network_interface_virtual_network_interface"))
}
