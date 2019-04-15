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
	apsUUID := "ddc62918-63c1-416b-96a2-e4ad976998fc"
	sgUUID := "c0b52016-498f-4d29-836d-c6629a360f5d"
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

	defer func() {
		integration.DeleteRoutingInstance(t, hc, riUUID)
		integration.DeleteRouteTarget(t, hc, rtUUID)
		integration.DeleteVirtualNetwork(t, hc, vnUUID)
		integration.DeleteProject(t, hc, projectUUID)
	}()

	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())

	project := integration.GetProject(t, hc, projectUUID)
	if assert.Len(t, project.GetApplicationPolicySetRefs(), 1) {
		assert.Equal(t, apsUUID, project.GetApplicationPolicySetRefs()[0].GetUUID())
	}
	if assert.Len(t, project.GetApplicationPolicySets(), 1) {
		assert.Equal(t, apsUUID, project.GetApplicationPolicySets()[0].GetUUID())
	}
	if assert.Len(t, project.GetSecurityGroups(), 1) {
		assert.Equal(t, sgUUID, project.GetSecurityGroups()[0].GetUUID())
	}

	ri := integration.GetRoutingInstance(t, hc, riUUID)
	if assert.Len(t, ri.GetRouteTargetRefs(), 1) {
		assert.Equal(t, rtUUID, ri.GetRouteTargetRefs()[0].GetUUID())
	}

	vn := integration.GetVirtualNetwork(t, hc, vnUUID)
	if assert.Len(t, vn.GetRoutingInstances(), 1) {
		assert.Equal(t, riUUID, vn.GetRoutingInstances()[0].GetUUID())
	}
}
