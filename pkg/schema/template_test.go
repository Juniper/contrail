package schema

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	api, err := MakeAPI("test_data/schema")
	assert.Nil(t, err, "no error expected")
	templateConf, err := LoadTemplates("test_data/templates/template_config.yaml")
	assert.Nil(t, err, "no error expected")
	err = ApplyTemplates(api, filepath.Dir("test_data/templates"), templateConf)
	assert.Nil(t, err, "no error expected")
	var schemas []string
	err = LoadFile("test_output/all.yml", &schemas)
	assert.Nil(t, err, "no error expected")
	assert.Equal(t, 4, len(schemas))
	var projectProperty []string
	err = LoadFile("test_output/project_type.yml", &projectProperty)
	assert.Nil(t, err, "no error expected")
	assert.Equal(t, 2, len(projectProperty))
}
