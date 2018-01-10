//+build integration

package cli

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Juniper/contrail/integration"
	"github.com/Juniper/contrail/pkg/agent"
	"github.com/Juniper/contrail/pkg/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

const (
	virtualNetwork  = "testdata/virtual_network.yaml"
	virtualNetworks = "testdata/virtual_networks.yml"
	//virtualNetworksFiltered = "testdata/virtual_networks_filtered.yaml"
	virtualNetworksSet     = "testdata/virtual_networks_set.yaml"
	virtualNetworksUpdated = "testdata/virtual_networks_updated.yaml"
	virtualNetworksAfterRm = "testdata/virtual_networks_after_rm.yaml"
	virtualNetworksDeleted = "testdata/virtual_networks_deleted.yaml"
)

func TestCLISchema(t *testing.T) {
	s := integration.NewServer(t)
	defer s.Close(t)
	a := givenLoggedInAgent(t, s.URL())

	schema, err := a.SchemaCLI("virtual_network")
	assert.NoError(t, err)
	assert.NotNil(t, schema)
}

func TestCLICreateListAndShowVirtualNetworks(t *testing.T) {
	s := integration.NewServer(t)
	defer s.Close(t)
	common.UseTable(s.Database(), "metadata")
	defer common.ClearTable(s.Database(), "metadata")
	common.UseTable(s.Database(), "virtual_network")
	defer common.ClearTable(s.Database(), "virtual_network")
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	fmt.Println(o)
	checkDataEqual(t, virtualNetworks, o)

	//o, err = a.ListCLI("virtual_network", "", "0", "30")
	//assert.NoError(t, err)
	//checkDataEqual(t, virtualNetworks, o)
	//
	//o, err = a.ListCLI("virtual_network", "id=e176f988-4f44-4fc0-83a2-bead7f288e66", "0", "30")
	//assert.NoError(t, err)
	//
	//checkDataEqual(t, virtualNetworksFiltered, o)
	//
	//o, err = a.ShowCLI("virtual_network", "2484e654-d8c0-4715-9d84-dd46aba23921")
	//assert.NoError(t, err)
	//checkDataEqual(t, virtualNetwork, o)
}

func TestCLISetVirtualNetworks(t *testing.T) {
	s := integration.NewServer(t)
	defer s.Close(t)
	common.UseTable(s.Database(), "metadata")
	defer common.ClearTable(s.Database(), "metadata")
	common.UseTable(s.Database(), "virtual_network")
	defer common.ClearTable(s.Database(), "virtual_network")
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	_, err = a.SetCLI("virtual_network", "2484e654-d8c0-4715-9d84-dd46aba23921", "description: ")
	assert.NoError(t, err)

	o, err = a.ListCLI("virtual_network", "", "0", "30")
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksSet, o)
}

func TestCLIUpdateVirtualNetworks(t *testing.T) {
	s := integration.NewServer(t)
	defer s.Close(t)
	common.UseTable(s.Database(), "metadata")
	defer common.ClearTable(s.Database(), "metadata")
	common.UseTable(s.Database(), "virtual_network")
	defer common.ClearTable(s.Database(), "virtual_network")
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	o, err = a.UpdateCLI(virtualNetworksUpdated)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksUpdated, o)

	o, err = a.ListCLI("virtual_network", "", "0", "30")
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksUpdated, o)
}

func TestCLISyncVirtualNetworks(t *testing.T) {
	s := integration.NewServer(t)
	defer s.Close(t)
	common.UseTable(s.Database(), "metadata")
	defer common.ClearTable(s.Database(), "metadata")
	common.UseTable(s.Database(), "virtual_network")
	defer common.ClearTable(s.Database(), "virtual_network")
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetwork)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	_, err = a.SyncCLI(virtualNetworks)
	assert.NoError(t, err)

	o, err = a.ListCLI("virtual_network", "", "0", "30")
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)
}

func TestCLIRemoveVirtualNetworks(t *testing.T) {
	s := integration.NewServer(t)
	defer s.Close(t)
	common.UseTable(s.Database(), "metadata")
	defer common.ClearTable(s.Database(), "metadata")
	common.UseTable(s.Database(), "virtual_network")
	defer common.ClearTable(s.Database(), "virtual_network")
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	_, err = a.RemoveCLI("virtual_network", "2484e654-d8c0-4715-9d84-dd46aba23921")
	assert.NoError(t, err)

	o, err = a.ListCLI("virtual_network", "", "0", "30")
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksAfterRm, o)
}

func TestCLIDeleteVirtualNetworks(t *testing.T) {
	s := integration.NewServer(t)
	defer s.Close(t)
	common.UseTable(s.Database(), "metadata")
	defer common.ClearTable(s.Database(), "metadata")
	common.UseTable(s.Database(), "virtual_network")
	defer common.ClearTable(s.Database(), "virtual_network")
	a := givenLoggedInAgent(t, s.URL())

	o, err := a.CreateCLI(virtualNetworks)
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworks, o)

	err = a.DeleteCLI(virtualNetworks)
	assert.NoError(t, err)

	o, err = a.ListCLI("virtual_network", "", "0", "30")
	assert.NoError(t, err)
	checkDataEqual(t, virtualNetworksDeleted, o)
}

func givenLoggedInAgent(t *testing.T, serverURL string) *agent.Agent {
	a, err := agent.NewAgent(&agent.Config{
		ID:        "alice",
		Password:  "alice_password",
		ProjectID: "admin",
		AuthURL:   serverURL + "/v3",
		Endpoint:  serverURL,
		Backend:   "file",
		Watcher:   "polling",
	})
	assert.NoError(t, err)

	assert.NoError(t, a.APIServer.Login())
	return a
}

func checkDataEqual(t *testing.T, expectedFile, actualYAML string) {
	expectedBytes, err := ioutil.ReadFile(expectedFile)
	if err != nil {
		assert.NoError(t, err, "cannot read expected data file")
	}

	var expected interface{}
	err = yaml.Unmarshal(expectedBytes, &expected)
	if err != nil {
		assert.NoError(t, err, "cannot parse expected data file")
	}

	var actual interface{}
	err = yaml.Unmarshal([]byte(actualYAML), &actual)
	if err != nil {
		assert.NoError(t, err, "cannot parse actual data file")
	}

	assert.Equal(t, expected, actual)
}
