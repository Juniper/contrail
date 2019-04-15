package contrailutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/convert"
)

const (
	contrailConfigFile = "../../../sample/contrail.yml"
	initDataFile       = "../../../tools/init_data.yaml"
)

func TestConvertYAMLToRDBMS(t *testing.T) {
	configFile = contrailConfigFile
	initConfig()

	err := convert.Convert(&convert.Config{
		InType:  convert.YAMLType,
		InFile:  initDataFile,
		OutType: convert.RDBMSType,
	})
	assert.NoError(t, err)
}

func TestConvertYAMLToRDBMSWithRefs(t *testing.T) {
	configFile = contrailConfigFile
	initConfig()

	err := convert.Convert(&convert.Config{
		InType:  convert.YAMLType,
		InFile:  initDataFile,
		OutType: convert.RDBMSType,
	})
	require.NoError(t, err)

	err = convert.Convert(&convert.Config{
		InType:  convert.YAMLType,
		InFile:  "test_data/test_with_refs.yaml",
		OutType: convert.RDBMSType,
	})
	assert.NoError(t, err)

	// TODO Test that the expected data is in the database.
}
