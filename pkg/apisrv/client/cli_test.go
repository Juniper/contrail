package client_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
)

const (
	projectUUID              = "project-cli-test-uuid"
	resourcesPath            = "testdata/resources.yml"
	vnBlueUUID               = "efb6aa60-9d8e-11e9-b056-13df9df3688a"
	vnRedName                = "vn-red"
	vnSchemaID               = "virtual_network"
	vnsPath                  = "testdata/vns.yml"
	vnsWithExternalIPAMsPath = "testdata/vns-with-external-ipams.yml"
)

func TestCLI(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		RepoRootPath: "../../..",
	})
	defer func() { assert.NoError(t, s.Close()) }()

	cli, err := client.NewCLI(
		integration.AdminHTTPConfig(s.URL()),
		"/public",
	)
	require.NoError(t, err)

	t.Run("schema is showed", testCLIShowsSchema(cli))
	t.Run("help message is displayed given empty schema ID", testHelpMessageIsDisplayedGivenEmptySchemaID(cli))
	t.Run("CRUD", testCRUD(cli))
}

func testCLIShowsSchema(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		s, err := cli.ShowSchema(vnSchemaID)
		assert.NoError(t, err)
		assertEqual(t, []interface{}{vnSchema(t)}, s)
	}
}

func vnSchema(t *testing.T) map[string]interface{} {
	return unmarshalResource(t, vnSchemaYAML())
}

func vnSchemaYAML() string {
	return `
kind: virtual_network
data:
  mac_learning_enabled: False #  (boolean)
  virtual_network_network_id:  #  (integer)
  configuration_version:  #  (integer)
  fq_name:  #  (array)
  ecmp_hashing_include_fields:  #  (object)
  pbb_evpn_enable: False #  (boolean)
  is_shared:  #  (boolean)
  route_target_list:  #  (object)
  flood_unknown_unicast: False #  (boolean)
  import_route_target_list:  #  (object)
  multi_policy_service_chains_enabled:  #  (boolean)
  address_allocation_mode:  #  (string)
  external_ipam:  #  (boolean)
  mac_move_control:  #  (object)
  parent_uuid:  #  (string)
  pbb_etree_enable: False #  (boolean)
  port_security_enabled: True #  (boolean)
  provider_properties:  #  (object)
  display_name:  #  (string)
  layer2_control_word: False #  (boolean)
  perms2:  #  (object)
  uuid:  #  (string)
  parent_type:  #  (string)
  router_external:  #  (boolean)
  export_route_target_list:  #  (object)
  mac_limit_control:  #  (object)
  mac_aging_time: 300 #  (integer)
  virtual_network_properties:  #  (object)
  annotations:  #  (object)
  id_perms:  #  (object) `
}

func testHelpMessageIsDisplayedGivenEmptySchemaID(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		o, err := cli.ShowResource("", "")
		assert.NoError(t, err)
		assert.Contains(t, o, "contrail show virtual_network $UUID")

		o, err = cli.ListResources("", &client.ListParameters{})
		assert.NoError(t, err)
		assert.Contains(t, o, "contrail list virtual_network")

		o, err = cli.SetResourceParameter("", "", "")
		assert.NoError(t, err)
		assert.Contains(t, o, "contrail set virtual_network $UUID $YAML")

		o, err = cli.DeleteResource("", "")
		assert.NoError(t, err)
		assert.Contains(t, o, "contrail rm virtual_network $UUID")
	}
}

func testCRUD(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("show", testShow(cli))
		t.Run("set boolean field", testSetBooleanField(cli))
		t.Run("update boolean fields via sync", testUpdateBooleanFieldsViaSync(cli))
		t.Run("delete single (rm)", testDeleteSingle(cli))
		t.Run("delete multiple (delete)", testDeleteMultiple(cli))
	}
}

func testShow(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.ShowResource(vnSchemaID, vnBlueUUID)
		assert.NoError(t, err, fmt.Sprintf("VN %q should be retrieved", vnBlueUUID))
		assertEqual(
			t,
			resources(vnBlue(t)),
			o,
		)
	}
}

func testSetBooleanField(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.SetResourceParameter(
			vnSchemaID,
			vnBlueUUID,
			"external_ipam: true",
		)

		assert.NoError(t, err)
		assertEqual(
			t,
			resources(withExternalIPAM(t, vnBlue(t), true)),
			o,
		)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(
			t,
			resources(
				vnRed(t),
				withExternalIPAM(t, vnBlue(t), true),
			),
			o,
		)
	}
}

func testUpdateBooleanFieldsViaSync(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.SyncResources(vnsWithExternalIPAMsPath)

		assert.NoError(t, err)
		assertEqual(
			t,
			resources(
				withExternalIPAM(t, vnRed(t), true),
				withExternalIPAM(t, vnBlue(t), true),
			),
			o,
		)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(
			t,
			resources(
				withExternalIPAM(t, vnRed(t), true),
				withExternalIPAM(t, vnBlue(t), true),
			),
			o,
		)
	}
}

func testDeleteSingle(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.DeleteResource(vnSchemaID, vnRedName)

		assert.NoError(t, err)
		require.Equal(t, "", o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(
			t,
			resources(
				vnRed(t),
				vnBlue(t),
			),
			o,
		) // TODO(Daniel): only vn-blue should be left, fix implementation
	}
}

func testDeleteMultiple(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.DeleteResources(vnsPath)

		assert.NoError(t, err)
		require.Equal(t, "", o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(t, resources(), o)
	}
}

func createTestVirtualNetworks(t *testing.T, cli *client.CLI) {
	o, err := cli.SyncResources(resourcesPath)

	require.NoError(t, err)
	assertEqualByFile(t, resourcesPath, o)

	o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
		ParentUUIDs: projectUUID,
	})
	require.NoError(t, err)
	assertEqual(
		t,
		resources(
			vnRed(t),
			vnBlue(t),
		),
		o,
	)
}

func withExternalIPAM(t *testing.T, resource map[string]interface{}, ei bool) map[string]interface{} {
	data, ok := resource["data"].(map[interface{}]interface{})
	require.True(t, ok)

	data["external_ipam"] = ei
	return resource
}

func vnBlue(t *testing.T) map[string]interface{} {
	return unmarshalResource(t, vnBlueYAML())
}

func vnRed(t *testing.T) map[string]interface{} {
	return unmarshalResource(t, vnRedYAML())
}

func vnBlueYAML() string {
	return `
kind: virtual_network
data:
  fq_name:
  - default-domain
  - project-cli-test
  - vn-blue
  parent_type: project
  parent_uuid: project-cli-test-uuid
  perms2:
    owner: TestCLI
  uuid: efb6aa60-9d8e-11e9-b056-13df9df3688a`
}

func vnRedYAML() string {
	return `
kind: virtual_network
data:
  flood_unknown_unicast: true
  fq_name:
  - default-domain
  - project-cli-test
  - vn-red
  is_shared: true
  layer2_control_word: true
  mac_learning_enabled: true
  multi_policy_service_chains_enabled: true
  parent_type: project
  parent_uuid: project-cli-test-uuid
  pbb_etree_enable: true
  pbb_evpn_enable: true
  perms2:
    owner: TestCLI
  port_security_enabled: true
  router_external: true
  uuid: 0ce792b6-9d8f-11e9-a76a-5b775b6d8012`
}

func resources(resources ...interface{}) map[string]interface{} {
	return map[string]interface{}{
		"resources": append([]interface{}{}, resources...),
	}
}

func unmarshalResource(t *testing.T, yamlData string) map[string]interface{} {
	var r map[string]interface{}
	err := yaml.Unmarshal([]byte(yamlData), &r)
	require.NoError(t, err)
	return r
}

func assertEqual(t *testing.T, expected interface{}, actualYAML string) {
	testutil.AssertEqual(
		t,
		expected,
		actualData(t, actualYAML),
	)
}

func assertEqualByFile(t *testing.T, expectedYAMLFile, actualYAML string) {
	testutil.AssertEqual(
		t,
		expectedData(t, expectedYAMLFile),
		actualData(t, actualYAML),
	)
}

func expectedData(t *testing.T, expectedYAMLFile string) interface{} {
	expectedBytes, err := ioutil.ReadFile(expectedYAMLFile)
	require.NoError(t, err, "cannot read expected data file")

	var expected interface{}
	err = yaml.Unmarshal(expectedBytes, &expected)
	require.NoError(t, err, "cannot parse expected data file")

	return expected
}

func actualData(t *testing.T, actualYAML string) interface{} {
	var actual interface{}
	err := yaml.Unmarshal([]byte(actualYAML), &actual)
	require.NoError(t, err, "cannot parse actual data")
	return actual
}
