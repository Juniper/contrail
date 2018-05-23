package contrailutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	inFile = "../../../tools/init_data.yaml"
	configFile = "../../../sample/contrail.yml"
	initConfig()
	resources, err := readYAML()
	assert.NoError(t, err, "read yaml failed")
	err = resources.Sort()
	assert.NoError(t, err, "dependency error")
	err = writeRDBMS(resources)
	assert.NoError(t, err, "write rdbms failed")
}
