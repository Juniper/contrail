package client_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	yaml "gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
)

const (
	projectName              = "project-cli-test"
	projectUUID              = "project-cli-test-uuid"
	resourcesPath            = "testdata/resources.yml"
	vmiSchemaID              = "virtual_machine_interface"
	vmiUUID                  = "91611dcc-a7cc-11e9-ad85-27cb7a03275b"
	vnBlueName               = "vn-blue"
	vnBlueUUID               = "efb6aa60-9d8e-11e9-b056-13df9df3688a"
	vnRedName                = "vn-red"
	vnRedUUID                = "0ce792b6-9d8f-11e9-a76a-5b775b6d8012"
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
		t.Run("list", testList(cli))
		t.Run("set boolean field", testSetBooleanField(cli))
		t.Run("update boolean fields via sync", testUpdateBooleanFieldsViaSync(cli))
		t.Run("delete single (rm)", testDeleteSingle(cli))
		t.Run("delete multiple (delete)", testDeleteMultiple(cli))
	}
}

func testShow(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)

		o, err := cli.ShowResource(vnSchemaID, vnBlueUUID)

		assert.NoError(t, err, fmt.Sprintf("VN %q should be retrieved", vnBlueUUID))
		assertEqual(t, resources(vnBlue(t)), o)
	}
}

func testList(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			skip     bool // TODO(Daniel): fix implementation and remove skipping possibility
			name     string
			lp       *client.ListParameters
			expected interface{}
			assert   func(t *testing.T, response string)
		}{
			{
				name: "with filters",
				lp: &client.ListParameters{
					Filters: fmt.Sprintf("uuid==%s", vnBlueUUID),
				},
				expected: resources(vnBlue(t)),
			},
			{
				name: "with page limit",
				lp: &client.ListParameters{
					PageLimit: 1,
				},
				expected: resources(vnRed(t)),
			},
			{
				name: "with parent UUID and page marker",
				lp: &client.ListParameters{
					ParentUUIDs: projectUUID,
					PageMarker:  vnRedUUID,
				},
				expected: resources(vnBlue(t)),
			},
			{
				name: "with parent UUID and detail",
				lp: &client.ListParameters{
					ParentUUIDs: projectUUID,
					Detail:      true,
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				skip: true, // TODO(Daniel): fix implementation and remove
				name: "with parent UUID and count",
				lp: &client.ListParameters{
					ParentUUIDs: projectUUID,
					Count:       true,
				},
				expected: resources(vnRed(t), vnBlue(t)),
				assert: func(t *testing.T, response string) {
					// TODO(Daniel): check response body while fixing implementation
				},
			},
			{
				// TODO(Daniel): improve this test
				name: "with parent UUID and shared",
				lp: &client.ListParameters{
					ParentUUIDs: projectUUID,
					Shared:      true,
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				skip: true, // TODO(Daniel): fix implementation and remove
				name: "with parent UUID and exclude hrefs",
				lp: &client.ListParameters{
					ParentUUIDs:  projectUUID,
					ExcludeHRefs: true,
				},
				expected: resources(vnRed(t), vnBlue(t)),
				assert: func(t *testing.T, response string) {
					for _, e := range unmarshalEventList(t, response).Events {
						assert.Equal(t, "", e.GetCreateVirtualNetworkRequest().GetVirtualNetwork().GetHref())
					}
				},
			},
			{
				name: "with parent UUID and parent type",
				lp: &client.ListParameters{
					ParentUUIDs: projectUUID,
					ParentType:  "project",
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				skip: true, // TODO(Daniel): fix implementation and remove
				name: "with parent FQ Name",
				lp: &client.ListParameters{
					ParentFQName: strings.Join([]string{integration.DefaultDomainUUID, projectName}, ":"),
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with parent UUID",
				lp: &client.ListParameters{
					ParentUUIDs: projectUUID,
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with backref UUIDs",
				lp: &client.ListParameters{
					BackrefUUIDs: vmiUUID,
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with object UUIDs",
				lp: &client.ListParameters{
					ObjectUUIDs: strings.Join([]string{vnRedUUID, vnBlueUUID}, ","),
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				skip: true, // TODO(Daniel): fix implementation and remove
				name: "with parent UUID and fields",
				lp: &client.ListParameters{
					ParentUUIDs: projectUUID,
					Fields:      "name,uuid",
				},
				expected: resources(vnRedFiltered(t), vnBlueFiltered(t)),
				assert: func(t *testing.T, response string) {
					el := unmarshalEventList(t, response).Events
					assert.Equal(t, vnRedUUID, el[0].GetCreateVirtualNetworkRequest().GetVirtualNetwork().GetUUID())
					assert.Equal(t, vnRedName, el[0].GetCreateVirtualNetworkRequest().GetVirtualNetwork().GetName())

					assert.Equal(t, vnBlueUUID, el[1].GetCreateVirtualNetworkRequest().GetVirtualNetwork().GetUUID())
					assert.Equal(t, vnBlueName, el[1].GetCreateVirtualNetworkRequest().GetVirtualNetwork().GetName())

					for _, e := range el {
						assert.Equal(t, "", e.GetCreateVirtualNetworkRequest().GetVirtualNetwork().GetParentUUID())
						assert.Equal(t, 0, len(e.GetCreateVirtualNetworkRequest().GetVirtualNetwork().GetFQName()))
						assert.Equal(t, false, e.GetCreateVirtualNetworkRequest().GetVirtualNetwork().GetIsShared())
						// FIXME: this field has non zero value (it should not)
						assert.Equal(t, "", e.GetCreateVirtualNetworkRequest().GetVirtualNetwork().GetHref())
					}
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// TODO(Daniel): fix implementation and remove skipping possibility
				if tt.skip {
					t.Skip("Broken implementation")
				}

				createTestResources(t, cli)

				r, err := cli.ListResources(vnSchemaID, tt.lp)

				require.NoError(t, err)
				assertEqual(t, tt.expected, r)
				if tt.assert != nil {
					tt.assert(t, r)
				}
			})
		}
	}
}

func testSetBooleanField(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)

		o, err := cli.SetResourceParameter(
			vnSchemaID,
			vnBlueUUID,
			"external_ipam: true",
		)

		assert.NoError(t, err)
		assertEqual(t, resources(withExternalIPAM(t, vnBlue(t), true)), o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(t, resources(vnRed(t), withExternalIPAM(t, vnBlue(t), true)), o)
	}
}

func testUpdateBooleanFieldsViaSync(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)

		o, err := cli.SyncResources(vnsWithExternalIPAMsPath)

		assert.NoError(t, err)
		assertEqual(t, resources(withExternalIPAM(t, vnRed(t), true), withExternalIPAM(t, vnBlue(t), true)), o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(t, resources(withExternalIPAM(t, vnRed(t), true), withExternalIPAM(t, vnBlue(t), true)), o)
	}
}

func testDeleteSingle(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)
		deleteVMI(t, cli) // avoid DB constraint violation on VN delete

		o, err := cli.DeleteResource(vnSchemaID, vnRedUUID)

		assert.NoError(t, err)
		assert.Equal(t, "", o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(t, resources(vnBlue(t)), o)
	}
}

func testDeleteMultiple(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)
		deleteVMI(t, cli) // avoid DB constraint violation on VNs delete

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

func createTestResources(t *testing.T, cli *client.CLI) {
	o, err := cli.SyncResources(resourcesPath)

	require.NoError(t, err)
	assertEqualByFile(t, resourcesPath, o)
}

func deleteVMI(t *testing.T, cli *client.CLI) {
	o, err := cli.DeleteResource(vmiSchemaID, vmiUUID)
	require.NoError(t, err)
	require.Equal(t, "", o)
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

func vnBlueFiltered(t *testing.T) map[string]interface{} {
	return unmarshalResource(t, vnBlueFilteredYAML())
}

func vnRed(t *testing.T) map[string]interface{} {
	return unmarshalResource(t, vnRedYAML())
}

func vnRedFiltered(t *testing.T) map[string]interface{} {
	return unmarshalResource(t, vnRedFilteredYAML())
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

func vnBlueFilteredYAML() string {
	return `
kind: virtual_network
data:
  name: vn-blue
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

func vnRedFilteredYAML() string {
	return `
kind: virtual_network
data:
  name: vn-red
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
		unmarshalData(t, actualYAML),
	)
}

func assertEqualByFile(t *testing.T, expectedYAMLFile, actualYAML string) {
	testutil.AssertEqual(
		t,
		unmarshalDataFromFile(t, expectedYAMLFile),
		unmarshalData(t, actualYAML),
	)
}

func unmarshalDataFromFile(t *testing.T, expectedYAMLFile string) interface{} {
	expectedBytes, err := ioutil.ReadFile(expectedYAMLFile)
	require.NoError(t, err, "cannot read expected data file")

	return unmarshalData(t, string(expectedBytes))
}

func unmarshalData(t *testing.T, yamlData string) interface{} {
	var d interface{}
	err := yaml.Unmarshal([]byte(yamlData), &d)
	require.NoError(t, err, "cannot parse data")
	return d
}

func unmarshalEventList(t *testing.T, yamlEL string) *services.EventList {
	var el services.EventList
	err := yaml.Unmarshal([]byte(yamlEL), &el)
	require.NoError(t, err, "cannot parse actual data")
	return &el
}
