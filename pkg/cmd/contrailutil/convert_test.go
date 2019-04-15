package contrailutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/convert"
	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/testutil/integration"
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

	projectUUID := "9a76fa43-3c35-4c33-92e9-1133629df0ce"
	vnUUID := "85fa1791-65a3-4797-8732-1d55ba398395"
	riUUID := "088203d7-9b91-400b-9be4-9a513a2088b5"
	rtUUID := "a544fde6-4bc1-4d68-99cf-e20c8e1c0768"

	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: false,
	})
	defer s.CloseT(t)
	hc := integration.NewTestingHTTPClient(t, s.URL())

	// TODO Create keystone project and user.

	defer func() {
		integration.DeleteRoutingInstance(t, hc, riUUID)
		integration.DeleteRouteTarget(t, hc, rtUUID)
		integration.DeleteVirtualNetwork(t, hc, vnUUID)
		integration.DeleteProject(t, hc, projectUUID)
	}()

	vn := integration.GetVirtualNetwork(t, hc, vnUUID)
	assert.Equal(t, projectUUID, vn.GetParentUUID())

	// TODO Check that the children/refs are correct.

	// TODO Clean up.
}
