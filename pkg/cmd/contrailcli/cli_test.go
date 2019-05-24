package contrailcli

import (
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/testutil"
)

const (
	vnSchemaID = "virtual_network"

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

func TestCLISchema(t *testing.T) {
	setupClient()
	schema, err := showSchema(vnSchemaID)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworkSchema, schema)
}

func TestCLIHelpMessagesWhenGivenEmptySchemaID(t *testing.T) {
	setupClient()
	o, err := showResource("", "")
	assert.NoError(t, err)
	assert.Contains(t, o, "contrail show virtual_network $UUID")

	o, err = listResources("")
	assert.NoError(t, err)
	assert.Contains(t, o, "contrail list virtual_network")

	o, err = setResourceParameter("", "", "")
	assert.NoError(t, err)
	assert.Contains(t, o, "contrail set virtual_network $UUID $YAML")

	o, err = deleteResource("", "")
	assert.NoError(t, err)
	assert.Contains(t, o, "contrail rm virtual_network $UUID")
}

func TestCLI(t *testing.T) {
	t.Skip("Skipping failing test") // TODO: fix test

	setupClient()
	o, err := deleteResources(virtualNetworks)
	assert.NoError(t, err)
	assert.Equal(t, "", o)

	o, err = syncResources(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	o, err = listResources(vnSchemaID)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksListed, o)

	o, err = showResource(vnSchemaID, "first-uuid")
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworkShowed, o)

	o, err = setResourceParameter(vnSchemaID, "first-uuid", "external_ipam: true")
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksSetOutput, o)

	o, err = listResources(vnSchemaID)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksSetListed, o)

	o, err = syncResources(virtualNetworksUpdate)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksUpdate, o)

	o, err = listResources(vnSchemaID)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksUpdate, o)

	o, err = deleteResources(virtualNetworks)
	assert.NoError(t, err)
	assert.Equal(t, "", o)

	o, err = listResources(vnSchemaID)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksDeletedListed, o)

	o, err = syncResources(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	o, err = deleteResource(vnSchemaID, "second-uuid")
	assert.NoError(t, err)
	assert.Equal(t, "", o)

	o, err = listResources(vnSchemaID)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksRMListed, o)
}

func checkDataEqual(t *testing.T, expectedYAMLFile, actualYAML string) {
	expectedBytes, err := ioutil.ReadFile(expectedYAMLFile)
	require.NoError(t, err, "cannot read expected data file")

	var expected interface{}
	err = yaml.Unmarshal(expectedBytes, &expected)
	require.NoError(t, err, "cannot parse expected data file")

	var actual interface{}
	err = yaml.Unmarshal([]byte(actualYAML), &actual)
	require.NoError(t, err, "cannot parse actual data")

	testutil.AssertEqual(t, expected, actual)
}
