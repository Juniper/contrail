package client_test

import (
	"io/ioutil"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	apisrvkeystone "github.com/Juniper/contrail/pkg/apisrv/keystone"
	yaml "gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
)

const (
	vnSchemaID = "virtual_network"

	noResources                 = "testdata/no-resources.yml"
	resources                   = "testdata/resources.yml"
	vns                         = "testdata/vns.yml"
	vnBlue                      = "testdata/vn-blue.yml"
	vnBlueWithExternalIPAM      = "testdata/vn-blue-with-external-ipam.yml"
	vnSchema                    = "testdata/vn-schema.yml"
	vnsWithBlueWithExternalIPAM = "testdata/vns-with-blue-with-external-ipam.yml"
	vnsWithExternalIPAMs        = "testdata/vns-with-external-ipams.yml"
)

func TestCLI(t *testing.T) {
	server := integration.NewRunningAPIServer(t,
		&integration.APIServerConfig{
			RepoRootPath: "../../..",
		})
	defer func() { assert.NoError(t, server.Close()) }()

	cli, err := client.NewCLI(
		&client.HTTPConfig{
			ID:       integration.AdminUserID,
			Password: integration.AdminUserPassword,
			Endpoint: server.URL(),
			AuthURL:  server.URL() + apisrvkeystone.AuthEndpointSuffix,
			Scope: keystone.NewScope(
				integration.DefaultDomainID,
				integration.DefaultDomainName,
				integration.AdminProjectID,
				integration.AdminProjectName,
			),
			Insecure: true,
		},
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
		checkDataEqual(t, vnSchema, s)
	}
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
		t.Run("set", testSet(cli))
		t.Run("update via sync", testUpdateViaSync(cli))
		t.Run("delete single (rm)", testDeleteSingle(cli))
		t.Run("delete multiple (delete)", testDeleteMultiple(cli))
	}
}

func testShow(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.ShowResource(vnSchemaID, "efb6aa60-9d8e-11e9-b056-13df9df3688a")
		assert.NoError(t, err, "VN 'efb6aa60-9d8e-11e9-b056-13df9df3688a' should be get")
		checkDataEqual(t, vnBlue, o)
	}
}

func testSet(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.SetResourceParameter(
			vnSchemaID,
			"efb6aa60-9d8e-11e9-b056-13df9df3688a",
			"external_ipam: true",
		)

		assert.NoError(t, err)
		checkDataEqual(t, vnBlueWithExternalIPAM, o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: "project-cli-test-uuid",
		})
		assert.NoError(t, err)
		checkDataEqual(t, vnsWithBlueWithExternalIPAM, o)
	}
}

func testUpdateViaSync(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.SyncResources(vnsWithExternalIPAMs)

		assert.NoError(t, err)
		checkDataEqual(t, vnsWithExternalIPAMs, o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: "project-cli-test-uuid",
		})
		assert.NoError(t, err)
		checkDataEqual(t, vnsWithExternalIPAMs, o)
	}
}

func testDeleteSingle(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.DeleteResource(vnSchemaID, "vn-red")

		assert.NoError(t, err)
		require.Equal(t, "", o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: "project-cli-test-uuid",
		})
		assert.NoError(t, err)
		checkDataEqual(t, vns, o) // TODO(Daniel): only vn-blue should be left, fix implementation
	}
}

func testDeleteMultiple(cli *client.CLI) func(t *testing.T) {
	return func(t *testing.T) {
		createTestVirtualNetworks(t, cli)

		o, err := cli.DeleteResources(vns)

		assert.NoError(t, err)
		require.Equal(t, "", o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
			ParentUUIDs: "project-cli-test-uuid",
		})
		assert.NoError(t, err)
		checkDataEqual(t, noResources, o)
	}
}

func createTestVirtualNetworks(t *testing.T, cli *client.CLI) {
	o, err := cli.SyncResources(resources)

	require.NoError(t, err)
	checkDataEqual(t, resources, o)

	o, err = cli.ListResources(vnSchemaID, &client.ListParameters{
		ParentUUIDs: "project-cli-test-uuid",
	})
	require.NoError(t, err)
	checkDataEqual(t, vns, o)
}

func checkDataEqual(t *testing.T, expectedYAMLFile, actualYAML string) {
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
