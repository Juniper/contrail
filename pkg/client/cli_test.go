package client_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	yaml "gopkg.in/yaml.v2"

	. "github.com/Juniper/asf/pkg/client"
)

const (
	domainName               = "default-domain"
	projectName              = "project-cli-test"
	secondProjectName        = "project-cli-test-2"
	projectUUID              = "project-cli-test-uuid"
	resourcesPath            = "testdata/resources.yml"
	vmiSchemaID              = "virtual_machine_interface"
	vmiUUID                  = "91611dcc-a7cc-11e9-ad85-27cb7a03275b"
	vnBlueUUID               = "efb6aa60-9d8e-11e9-b056-13df9df3688a"
	vnRedUUID                = "0ce792b6-9d8f-11e9-a76a-5b775b6d8012"
	vnSchemaID               = "virtual_network"
	vnsPath                  = "testdata/vns.yml"
	vnsWithExternalIPAMsPath = "testdata/vns-with-external-ipams.yml"
)

func TestCLI(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		RepoRootPath: "../..",
	})
	defer func() { assert.NoError(t, s.Close()) }()

	cli, err := NewCLI(
		integration.AdminHTTPConfig(s.URL()),
		"/public",
	)
	require.NoError(t, err)

	t.Run("schema is showed", testCLIShowsSchema(cli))
	t.Run("help message is displayed given empty schema ID", testHelpMessageIsDisplayedGivenEmptySchemaID(cli))
	t.Run("CRUD", testCRUD(cli))
}

func testCLIShowsSchema(cli *CLI) func(t *testing.T) {
	return func(t *testing.T) {
		s, err := cli.ShowSchema(vnSchemaID)
		assert.NoError(t, err)
		assertEqual(t, []interface{}{vnSchema(t)}, s)
	}
}

func vnSchema(t *testing.T) map[interface{}]interface{} {
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

func testHelpMessageIsDisplayedGivenEmptySchemaID(cli *CLI) func(t *testing.T) {
	return func(t *testing.T) {
		o, err := cli.ShowResource("", "")
		assert.NoError(t, err)
		assert.Contains(t, o, "contrail show virtual_network $UUID")

		o, err = cli.ListResources("", &ListParameters{})
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

func testCRUD(cli *CLI) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("show", testShow(cli))
		t.Run("list", testList(cli))
		t.Run("set boolean field", testSetBooleanField(cli))
		t.Run("update boolean fields via sync", testUpdateBooleanFieldsViaSync(cli))
		t.Run("delete single (rm)", testDeleteSingle(cli))
		t.Run("delete multiple (delete)", testDeleteMultiple(cli))
	}
}

func testShow(cli *CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)

		o, err := cli.ShowResource(vnSchemaID, vnBlueUUID)

		assert.NoError(t, err, fmt.Sprintf("VN %q should be retrieved", vnBlueUUID))
		assertEqual(t, resources(vnBlue(t)), o)
	}
}

func testList(cli *CLI) func(t *testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			name     string
			lp       *ListParameters
			expected interface{}
			assert   func(t *testing.T, response string)
		}{
			{
				name: "with filters",
				lp: &ListParameters{
					Filters: fmt.Sprintf("uuid==%s", vnBlueUUID),
				},
				expected: resources(vnBlue(t)),
			},
			{
				name: "with parent UUID and page limit",
				lp: &ListParameters{
					ParentUUIDs: projectUUID,
					PageLimit:   1,
				},
				expected: resources(vnRed(t)),
			},
			{
				name: "with parent UUID and page marker",
				lp: &ListParameters{
					ParentUUIDs: projectUUID,
					PageMarker:  vnRedUUID,
				},
				expected: resources(vnBlue(t)),
			},
			{
				name: "with parent UUID and detail",
				lp: &ListParameters{
					ParentUUIDs: projectUUID,
					Detail:      true,
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with parent UUID and count",
				lp: &ListParameters{
					ParentUUIDs: projectUUID,
					Count:       true,
				},
				expected: map[string]interface{}{
					"virtual-networks": map[string]interface{}{
						"count": 2,
					},
				},
			},
			{
				// TODO(Daniel): improve this test
				name: "with parent UUID and shared",
				lp: &ListParameters{
					ParentUUIDs: projectUUID,
					Shared:      true,
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with parent UUID and exclude hrefs",
				lp: &ListParameters{
					ParentUUIDs:  projectUUID,
					ExcludeHRefs: true,
				},
				expected: resources(vnRed(t), vnBlue(t)),
				assert: func(t *testing.T, response string) {
					for _, r := range unmarshalResources(t, response)[ResourcesKey] {
						data, ok := r[DataKey].(map[interface{}]interface{})
						assert.True(t, ok)

						_, ok = data["href"]
						assert.False(t, ok, "There should be no Href field in data, but there is")
					}
				},
			},
			{
				name: "with parent UUID and parent type",
				lp: &ListParameters{
					ParentUUIDs: projectUUID,
					ParentType:  "project",
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with parent FQ Name",
				lp: &ListParameters{
					ParentFQName: strings.Join([]string{domainName, projectName}, ":"),
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with parent's parent FQ Name",
				lp: &ListParameters{
					ParentFQName: strings.Join([]string{domainName}, ":"),
				},
				expected: resources(),
			},
			{
				name: "with different parent FQ Name",
				lp: &ListParameters{
					ParentFQName: strings.Join([]string{domainName, secondProjectName}, ":"),
				},
				expected: resources(vnGreen(t)),
			},
			{
				name: "with parent UUID",
				lp: &ListParameters{
					ParentUUIDs: projectUUID,
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with backref UUIDs",
				lp: &ListParameters{
					BackrefUUIDs: vmiUUID,
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with object UUIDs",
				lp: &ListParameters{
					ObjectUUIDs: strings.Join([]string{vnRedUUID, vnBlueUUID}, ","),
				},
				expected: resources(vnRed(t), vnBlue(t)),
			},
			{
				name: "with object UUID and fields",
				lp: &ListParameters{
					ObjectUUIDs: vnRedUUID,
					Fields:      "name,uuid",
				},
				expected: resources(vnRedFiltered(t)),
			},
			{
				name: "with object UUID and invalid field",
				lp: &ListParameters{
					ObjectUUIDs: vnRedUUID,
					Fields:      "name,uuid,invalid_field123",
				},
				expected: resources(vnRedFiltered(t)),
			},
			{
				name: "with object UUID and no valid field",
				lp: &ListParameters{
					ObjectUUIDs: vnRedUUID,
					Fields:      "invalid_field123",
				},
				expected: vnEmptyResource(),
			},
			{
				name: "with object UUID, fields and detail",
				lp: &ListParameters{
					ObjectUUIDs: vnRedUUID,
					Fields:      "name,uuid",
					Detail:      true,
				},
				expected: resources(vnRed(t)),
			},
			{
				name: "with object UUID, invalid fields and detail",
				lp: &ListParameters{
					ObjectUUIDs: vnRedUUID,
					Fields:      "name,uuid,invalidfield123",
					Detail:      true,
				},
				expected: resources(vnRed(t)),
			},
			{
				name: "with object UUID, no valid field and detail",
				lp: &ListParameters{
					ObjectUUIDs: vnRedUUID,
					Fields:      "invalidfield123",
					Detail:      true,
				},
				expected: resources(vnRed(t)),
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
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

func testSetBooleanField(cli *CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)

		o, err := cli.SetResourceParameter(
			vnSchemaID,
			vnBlueUUID,
			"external_ipam: true",
		)

		assert.NoError(t, err)
		assertEqual(t, resources(withExternalIPAM(t, vnBlue(t), true)), o)

		o, err = cli.ListResources(vnSchemaID, &ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(t, resources(vnRed(t), withExternalIPAM(t, vnBlue(t), true)), o)
	}
}

func testUpdateBooleanFieldsViaSync(cli *CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)

		o, err := cli.SyncResources(vnsWithExternalIPAMsPath)

		assert.NoError(t, err)
		assertEqual(t, resources(withExternalIPAM(t, vnRed(t), true), withExternalIPAM(t, vnBlue(t), true)), o)

		o, err = cli.ListResources(vnSchemaID, &ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(t, resources(withExternalIPAM(t, vnRed(t), true), withExternalIPAM(t, vnBlue(t), true)), o)
	}
}

func testDeleteSingle(cli *CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)
		deleteVMI(t, cli) // avoid DB constraint violation on VN delete

		o, err := cli.DeleteResource(vnSchemaID, vnRedUUID)

		assert.NoError(t, err)
		assert.Equal(t, "", o)

		o, err = cli.ListResources(vnSchemaID, &ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(t, resources(vnBlue(t)), o)
	}
}

func testDeleteMultiple(cli *CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestResources(t, cli)
		deleteVMI(t, cli) // avoid DB constraint violation on VNs delete

		o, err := cli.DeleteResources(vnsPath)

		assert.NoError(t, err)
		require.Equal(t, "", o)

		o, err = cli.ListResources(vnSchemaID, &ListParameters{
			ParentUUIDs: projectUUID,
		})
		assert.NoError(t, err)
		assertEqual(t, nil, o)
	}
}

func createTestResources(t *testing.T, cli *CLI) {
	o, err := cli.SyncResources(resourcesPath)

	require.NoError(t, err)
	assertEqualByFile(t, resourcesPath, o)
}

func deleteVMI(t *testing.T, cli *CLI) {
	o, err := cli.DeleteResource(vmiSchemaID, vmiUUID)
	require.NoError(t, err)
	require.Equal(t, "", o)
}

func withExternalIPAM(t *testing.T, resource map[interface{}]interface{}, ei bool) map[interface{}]interface{} {
	data, ok := resource["data"].(map[interface{}]interface{})
	require.True(t, ok)

	data["external_ipam"] = ei
	return resource
}

func vnBlue(t *testing.T) map[interface{}]interface{} {
	return unmarshalResource(t, vnBlueYAML())
}

func vnRed(t *testing.T) map[interface{}]interface{} {
	return unmarshalResource(t, vnRedYAML())
}

func vnRedFiltered(t *testing.T) map[interface{}]interface{} {
	return unmarshalResource(t, vnRedFilteredYAML())
}

func vnGreen(t *testing.T) map[interface{}]interface{} {
	return unmarshalResource(t, vnGreenYAML())
}

func vnEmptyResource() map[interface{}]interface{} {
	return map[interface{}]interface{}{
		"resources": []interface{}{
			map[interface{}]interface{}{
				"data": map[interface{}]interface{}{},
				"kind": "virtual_network"}}}
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

func vnRedFilteredYAML() string {
	return `
kind: virtual_network
data:
  name: vn-red
  uuid: 0ce792b6-9d8f-11e9-a76a-5b775b6d8012`
}

func vnGreenYAML() string {
	return `
kind: virtual_network
data:
  fq_name:
  - default-domain
  - project-cli-test-2
  - vn-green
  parent_type: project
  parent_uuid: bf4d34df-3807-4573-929a-415224af0fc0
  perms2:
    owner: TestCLI
  uuid: 84a182ea-9c0a-4f8e-b570-6183b4697c40`
}

func resources(resources ...interface{}) map[interface{}]interface{} {
	if len(resources) == 0 {
		return map[interface{}]interface{}{
			ResourcesKey: nil,
		}
	}
	return map[interface{}]interface{}{
		ResourcesKey: append([]interface{}{}, resources...),
	}
}

func unmarshalResources(t *testing.T, yamlData string) Resources {
	var r Resources
	err := yaml.Unmarshal([]byte(yamlData), &r)
	require.NoError(t, err)
	return r
}

func unmarshalResource(t *testing.T, yamlData string) map[interface{}]interface{} {
	var r map[interface{}]interface{}
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
