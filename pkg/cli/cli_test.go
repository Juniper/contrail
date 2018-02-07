// Package cli contains integration tests of CLI for API Server.
// Keep command example inputs and outputs in doc/cli.md up to date.
package cli

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/Juniper/contrail/pkg/agent"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/testutil"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

const (
	vnSchemaID       = "virtual_network"
	metadataSchemaID = "metadata"

	virtualNetworkSchema         = "testdata/virtual_network_schema.yml"
	virtualNetwork               = "testdata/virtual_network.yml"
	virtualNetworkListed         = "testdata/virtual_network_listed.yml"
	virtualNetworkShowed         = "testdata/virtual_network_showed.yml"
	virtualNetworks              = "testdata/virtual_networks.yml"
	virtualNetworksListed        = "testdata/virtual_networks_listed.yml"
	virtualNetworksSetOutput     = "testdata/virtual_networks_set_output.yml"
	virtualNetworksSetListed     = "testdata/virtual_networks_set_listed.yml"
	virtualNetworksUpdate        = "testdata/virtual_networks_update.yml"
	virtualNetworksUpdateOutput  = "testdata/virtual_networks_update_output.yml"
	virtualNetworksUpdatedListed = "testdata/virtual_networks_updated_listed.yml"
	virtualNetworksDeletedListed = "testdata/virtual_networks_deleted_listed.yml"
)

func TestCLISchema(t *testing.T) {
	s := givenAPIServer(t)
	defer s.Close(t)
	a := givenLoggedInAgent(t, s.URL())

	schema, err := a.SchemaCLI(vnSchemaID)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworkSchema, schema)
}

func TestCLIHelpMessagesWhenGivenEmptySchemaID(t *testing.T) {
	s := givenAPIServer(t)
	defer s.Close(t)
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.ShowCLI("", "")
	assert.NoError(t, err)
	assert.Contains(t, o, "contrail show user $UUID")
	assert.Contains(t, o, "contrail show virtual_network $UUID")

	o, err = a.ListCLI("", nil)
	assert.NoError(t, err)
	assert.Contains(t, o, "contrail list user")
	assert.Contains(t, o, "contrail list virtual_network")

	o, err = a.SetCLI("", "", "")
	assert.NoError(t, err)
	assert.Contains(t, o, "contrail set user $UUID $YAML")
	assert.Contains(t, o, "contrail set virtual_network $UUID $YAML")

	o, err = a.RemoveCLI("", "")
	assert.NoError(t, err)
	assert.Contains(t, o, "contrail rm user $UUID")
	assert.Contains(t, o, "contrail rm virtual_network $UUID")
}

func TestCLICreateListAndShowVirtualNetworks(t *testing.T) {
	s := givenAPIServer(t)
	defer s.Close(t)
	testutil.LockAndClearTables(s.Database(), metadataSchemaID, vnSchemaID)
	defer testutil.ClearAndUnlockTables(s.Database(), metadataSchemaID, vnSchemaID)
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	o, err = a.ListCLI(vnSchemaID, nil)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksListed, o)

	o, err = a.ListCLI(vnSchemaID, url.Values{
		common.FiltersKey: []string{fmt.Sprintf("uuid==%s", "first-uuid")},
	})
	assert.NoError(t, err)
	fmt.Println(o)
	checkDataEqual(t, virtualNetworkListed, o)

	o, err = a.ShowCLI(vnSchemaID, "first-uuid")
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworkShowed, o)
}

//TODO(nati) Skip until update implemented
func TestCLISetVirtualNetworks(t *testing.T) {
	t.Skip("Skipping until update implemented")

	s := givenAPIServer(t)
	defer s.Close(t)
	testutil.LockAndClearTables(s.Database(), metadataSchemaID, vnSchemaID)
	defer testutil.ClearAndUnlockTables(s.Database(), metadataSchemaID, vnSchemaID)
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	o, err = a.SetCLI(vnSchemaID, "first-uuid", "external_ipam: true")
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksSetOutput, o)

	o, err = a.ListCLI(vnSchemaID, nil)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksSetListed, o)
}

//TODO(nati) Skip until update implemented
func TestCLIUpdateVirtualNetworks(t *testing.T) {
	t.Skip("Skipping until update implemented")

	s := givenAPIServer(t)
	defer s.Close(t)
	testutil.LockAndClearTables(s.Database(), metadataSchemaID, vnSchemaID)
	defer testutil.ClearAndUnlockTables(s.Database(), metadataSchemaID, vnSchemaID)
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	o, err = a.UpdateCLI(virtualNetworksUpdate)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksUpdateOutput, o)

	o, err = a.ListCLI(vnSchemaID, nil)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksUpdatedListed, o)
}

func TestCLISyncVirtualNetworks(t *testing.T) {
	// TODO(daniel): Enable when API Server behavior is fixed: https://github.com/Juniper/contrail/issues/69
	t.Skip("Skipping till API Server Show() behavior is fixed")

	s := givenAPIServer(t)
	defer s.Close(t)
	testutil.LockAndClearTables(s.Database(), metadataSchemaID, vnSchemaID)
	defer testutil.ClearAndUnlockTables(s.Database(), metadataSchemaID, vnSchemaID)
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.SyncCLI(virtualNetwork)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetwork, o)

	o, err = a.ListCLI(vnSchemaID, nil)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworkListed, o)

	o, err = a.SyncCLI(virtualNetworksUpdate)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksUpdateOutput, o)

	o, err = a.ListCLI(vnSchemaID, nil)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksUpdatedListed, o)
}

func TestCLIRemoveVirtualNetworks(t *testing.T) {
	s := givenAPIServer(t)
	defer s.Close(t)
	testutil.LockAndClearTables(s.Database(), metadataSchemaID, vnSchemaID)
	defer testutil.ClearAndUnlockTables(s.Database(), metadataSchemaID, vnSchemaID)
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	o, err = a.RemoveCLI(vnSchemaID, "second-uuid")
	assert.NoError(t, err)
	assert.Equal(t, "", o)

	o, err = a.ListCLI(vnSchemaID, nil)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworkListed, o)
}

func TestCLIDeleteVirtualNetworks(t *testing.T) {
	s := givenAPIServer(t)
	defer s.Close(t)
	testutil.LockAndClearTables(s.Database(), metadataSchemaID, vnSchemaID)
	defer testutil.ClearAndUnlockTables(s.Database(), metadataSchemaID, vnSchemaID)
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	err = a.DeleteCLI(virtualNetworks)
	assert.NoError(t, err)

	o, err = a.ListCLI(vnSchemaID, nil)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksDeletedListed, o)
}

func givenAPIServer(t *testing.T) *testutil.APIServer {
	return testutil.NewAPIServer(t, "../..")
}

func givenLoggedInAgent(t *testing.T, apiServerURL string) *agent.Agent {
	a, err := agent.NewAgent(&agent.Config{
		ID:        testutil.AdminUserID,
		Password:  testutil.AdminPassword,
		ProjectID: testutil.AdminProjectID,
		AuthURL:   apiServerURL + testutil.AuthEndpointSuffix,
		Endpoint:  apiServerURL,
		Backend:   agent.FileBackend,
		Watcher:   agent.PollingWatcher,
	})
	assert.NoError(t, err)

	assert.NoError(t, a.APIServer.Login())
	return a
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

	assert.Equal(t, expected, actual)
}
