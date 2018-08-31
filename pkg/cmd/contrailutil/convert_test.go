package contrailutil

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/convert"
)

const (
	contrailConfigFile = "../../../sample/contrail.yml"
	initDataFile       = "../../../tools/init_data.yaml"

	initDataToSortFile = "../../../pkg/apisrv/test_data/init_data_to_sort.yaml"
	sortedInitDataFile = "../../../pkg/apisrv/test_data/init_data_sorted.yaml"
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

func TestConvertSort(t *testing.T) {
	configFile = contrailConfigFile
	initConfig()

	err := convert.Convert(&convert.Config{
		InType:  convert.YAMLType,
		InFile:  initDataToSortFile,
		OutType: convert.YAMLType,
		OutFile: sortedInitDataFile,
	})

	assert.NoError(t, err)
}
