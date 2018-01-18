package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	api, err := MakeAPI("test_data/schema")
	assert.Nil(t, err, "API reading failed")
	assert.Equal(t, 4, len(api.Types))
	assert.Equal(t, 4, len(api.Schemas))
	project := api.schemaByID("project")

	assert.Equal(t, 2, len(project.JSONSchema.Properties))
	assert.Equal(t, 2, len(project.Columns))

	virtualNetwork := api.schemaByID("virtual_network")

	assert.Equal(t, 3, len(virtualNetwork.JSONSchema.Properties))
	assert.Equal(t, 3, len(virtualNetwork.Columns))
}
