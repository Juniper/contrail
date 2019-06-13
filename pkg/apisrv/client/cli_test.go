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
	yaml "gopkg.in/yaml.v2"

	apisrvkeystone "github.com/Juniper/contrail/pkg/apisrv/keystone"

	_ "github.com/go-sql-driver/mysql"
)

const (
	vnSchemaID = "virtual_network"

	project                      = "testdata/project.yml"
	virtualNetworkSchema         = "testdata/virtual_network_schema.yml"
	virtualNetworkShowed         = "testdata/virtual_network_showed.yml"
	virtualNetworks              = "testdata/virtual_networks.yml"
	virtualNetworksListed        = "testdata/virtual_networks_listed.yml"
	virtualNetworksSetOutput     = "testdata/virtual_networks_set_output.yml"
	virtualNetworksSetListed     = "testdata/virtual_networks_set_listed.yml"
	virtualNetworksUpdate        = "testdata/virtual_networks_update.yml"
	virtualNetworksDeletedListed = "testdata/virtual_networks_deleted_listed.yml"
	virtualNetworksRMListed      = "testdata/virtual_networks_rm_listed.yml"
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
		checkDataEqual(t, virtualNetworkSchema, s)
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
		o, err := cli.SyncResources(project)
		require.NoError(t, err)
		checkDataEqual(t, project, o)

		o, err = cli.SyncResources(virtualNetworks)
		require.NoError(t, err, "VNs should be created via sync")
		checkDataEqual(t, virtualNetworks, o)

		o, err = cli.ListResources(vnSchemaID, &client.ListParameters{})
		require.NoError(t, err, "VNs should be listed")
		//checkDataEqual(t, virtualNetworksListed, o)

		t.Run("show", func(t *testing.T) {
			o, err = cli.ShowResource(vnSchemaID, "first-uuid")
			require.NoError(t, err, "VN 'first-uuid' should be get")
			checkDataEqual(t, virtualNetworkShowed, o)
		})

		t.Run("set", func(t *testing.T) {
			o, err = cli.SetResourceParameter(vnSchemaID, "first-uuid", "external_ipam: true")
			require.NoError(t, err, "external_ipam of VN 'first-uuid' should be updated")
			checkDataEqual(t, virtualNetworksSetOutput, o)

			o, err = cli.ListResources(vnSchemaID, &client.ListParameters{})
			require.NoError(t, err, "VNs should be listed")
			//checkDataEqual(t, virtualNetworksSetListed, o)
		})

		t.Run("update", func(t *testing.T) {
			o, err = cli.SyncResources(virtualNetworksUpdate)
			require.NoError(t, err, "VNs should be updated via sync")
			checkDataEqual(t, virtualNetworksUpdate, o)

			o, err = cli.ListResources(vnSchemaID, &client.ListParameters{})
			require.NoError(t, err, "VNs should be listed")
			//checkDataEqual(t, virtualNetworksUpdate, o)

			o, err = cli.DeleteResources(virtualNetworks)
			require.NoError(t, err, "VNs should be deleted")
			require.Equal(t, "", o)

			o, err = cli.ListResources(vnSchemaID, &client.ListParameters{})
			require.NoError(t, err, "VNs should be listed")
			//checkDataEqual(t, virtualNetworksDeletedListed, o)
		})

		t.Run("delete", func(t *testing.T) {
			o, err = cli.DeleteResources(virtualNetworks)
			require.NoError(t, err, "VNs should be deleted")
			require.Equal(t, "", o)

			o, err = cli.ListResources(vnSchemaID, &client.ListParameters{})
			require.NoError(t, err, "VNs should be listed")
			//checkDataEqual(t, virtualNetworksDeletedListed, o)
		})

		t.Run("recreate and delete", func(t *testing.T) {
			o, err = cli.SyncResources(virtualNetworks)
			require.NoError(t, err, "VNs should be recreated via sync")
			checkDataEqual(t, virtualNetworks, o)

			o, err = cli.DeleteResource(vnSchemaID, "second-uuid")
			require.NoError(t, err, "VN 'second-uuid' should deleted")
			require.Equal(t, "", o)

			o, err = cli.ListResources(vnSchemaID, &client.ListParameters{})
			require.NoError(t, err, "VN should be listed")
			//checkDataEqual(t, virtualNetworksRMListed, o)
		})
	}
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
