package undercloud

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	workRoot                 = "/tmp/rhospd_cloud_manager/"
	cloudManagerTemplatePath = "./test_data/test_undercloud.tmpl"
	createPlaybooks          = "./test_data/expected_ansible_create_command.yml"
	updatePlaybooks          = "./test_data/expected_ansible_update_command.yml"
	cloudManagerID           = "test_rhospd_cloud_manager_uuid"
	controlHostOneID         = "control_host1_node_uuid"
	controlHostTwoID         = "control_host2_node_uuid"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func verifyUnderCloudDeleted() bool {
	// Make sure working dir is deleted
	if _, err := os.Stat(workRoot + "/" + cloudManagerID); err == nil {
		// working dir not deleted
		return false
	}
	return true
}

func compareFiles(t *testing.T, expectedFile, generatedFile string) bool {
	generatedData, err := ioutil.ReadFile(generatedFile)
	assert.NoErrorf(t, err, "Unable to read generated: %s", generatedFile)
	expectedData, err := ioutil.ReadFile(expectedFile)
	assert.NoErrorf(t, err, "Unable to read expected: %s", expectedFile)
	return bytes.Equal(generatedData, expectedData)

}

func compareGeneratedSite(t *testing.T, expected string) bool {
	if !compareFiles(t, expected, generatedSitePath()) {
		return false
	}
	return compareFiles(t, expected, contrailCloudSitePath())
}

func verifyPlaybooks(t *testing.T, expected string) bool {
	return compareFiles(t, expected, executedPlaybooksPath())
}

func generatedSitePath() string {
	return workRoot + "/" + cloudManagerID + "/site.yml"
}

func contrailCloudSitePath() string {
	return workRoot + "/" + cloudManagerID + "/config/site.yml"
}

func executedPlaybooksPath() string {
	return workRoot + "/" + cloudManagerID + "/executed_command.yml"
}

func verifyPorts(t *testing.T, testScenario integration.TestScenario,
	nodeID string, expectedPorts []string) error {
	createdPorts := map[string]string{}
	var httpClient *client.HTTP
	for _, httpClient = range testScenario.Clients {
		var response map[string][]interface{}
		url := fmt.Sprintf("/ports?parent_uuid=%s", nodeID)
		_, err := httpClient.Read(context.Background(), url, &response)
		assert.NoError(t, err, "Unable to list ports of the node")
		for _, port := range response["ports"] {
			p := port.(map[string]interface{})                    //nolint: errcheck
			createdPorts[p["name"].(string)] = p["uuid"].(string) //nolint: errcheck
		}
		break
	}
	for _, p := range expectedPorts {
		if portID, ok := createdPorts[p]; ok {
			url := fmt.Sprintf("/port/%s", portID)
			var response interface{}
			_, err := httpClient.Delete(context.Background(), url, &response)
			assert.NoErrorf(t, err, "Unable to delete port: %s ", portID)
		} else {
			keys := make([]string, len(createdPorts))
			for k := range createdPorts {
				keys = append(keys, k)
			}
			return fmt.Errorf("port expected: %s, actual: %s for node %s",
				expectedPorts, keys, nodeID)
		}
	}
	return nil
}

// nolint: gocyclo
func runUnderCloudActionTest(t *testing.T, testScenario integration.TestScenario,
	config *Config, action, expectedSite, expectedPlaybooks string) {
	// set action field in the rhospd-cloud-manager resource
	var err error
	var data interface{}
	cloudManager := map[string]interface{}{"uuid": cloudManagerID,
		"provisioning_action": action,
	}
	config.Action = updateAction
	switch action {
	case importProvisioningAction:
		config.Action = createAction
		cloudManager["provisioning_action"] = ""
	}
	data = map[string]interface{}{"rhospd-cloud-manager": cloudManager}
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/rhospd-cloud-manager/%s", cloudManagerID)
		_, err = client.Update(context.Background(), url, &data, &response)
		assert.NoErrorf(t, err, "failed to set %s action in rhospd cloudManager", action)
		break
	}
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	underCloudManager, err := NewUnderCloud(config)
	assert.NoErrorf(t, err, "failed to create cloudManager manager to %s cloudManager", config.Action)
	deployer, err := underCloudManager.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoErrorf(t, err, "failed to manage(%s) cloudManager", action)
	if expectedSite != "" {
		assert.True(t, compareGeneratedSite(t, expectedSite),
			fmt.Sprintf("Site file created during action %s is not as expected", action))
	}
	if expectedPlaybooks != "" {
		assert.True(t, verifyPlaybooks(t, expectedPlaybooks),
			fmt.Sprintf("Expected list of %s playbooks are not executed", action))
	}
}

// nolint: gocyclo
func runUnderCloudTest(t *testing.T, expectedSite string, pContext map[string]interface{}) {
	// mock keystone to let access server after cloudManager create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystone("127.0.0.1:35357", keystoneAuthURL)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystone("127.0.0.1:5000", keystoneAuthURL)
	defer ksPrivate.Close()

	// Create the cloudManager and related objects
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, cloudManagerTemplatePath, pContext)
	assert.NoError(t, err, "failed to load cloudManager test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()
	// create cloudManager config
	s := &client.HTTP{
		Endpoint: server.URL(),
		InSecure: true,
		AuthURL:  server.URL() + "/keystone/v3",
		ID:       "alice",
		Password: "alice_password",
		Scope: client.GetKeystoneScope(
			"default", "default", "admin", "admin"),
	}
	s.Init()
	_, err = s.Login(context.Background())
	assert.NoError(t, err, "failed to login")
	config := &Config{
		APIServer:    s,
		ResourceID:   cloudManagerID,
		Action:       createAction,
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}
	// create cloudManager
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup old executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	underCloudManager, err := NewUnderCloud(config)
	assert.NoError(t, err, "failed to create cloudManager manager to create cloudManager")
	deployer, err := underCloudManager.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(create) cloudManager")
	assert.True(t, compareGeneratedSite(t, expectedSite),
		"Site file created is not as expected")
	assert.True(t, verifyPlaybooks(t, createPlaybooks),
		"Expected list of create playbooks are not executed")

	// verify ports
	controlHostOnePorts := []string{"ens2f1", "ens2f0"}
	controlHostTwoPorts := []string{"ens2f1"}
	err = verifyPorts(t, testScenario, controlHostOneID, controlHostOnePorts)
	assert.NoError(t, err, "failed to verify and delete ports")
	err = verifyPorts(t, testScenario, controlHostTwoID, controlHostTwoPorts)
	assert.NoError(t, err, "failed to verify and delete ports")

	// update cloudManager
	config.Action = updateAction
	// remove site.yml to trriger cloudManager update
	err = os.Remove(generatedSitePath())
	if err != nil {
		assert.NoError(t, err, "failed to delete site.yml")
	}
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	underCloudManager, err = NewUnderCloud(config)
	assert.NoError(t, err, "failed to create cloudManager manager to update cloudManager")
	deployer, err = underCloudManager.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(update) cloudManager")
	assert.True(t, compareGeneratedSite(t, expectedSite),
		"Site file updated is not as expected")
	assert.True(t, verifyPlaybooks(t, updatePlaybooks),
		"Expected list of update playbooks are not executed")

	// verify ports
	err = verifyPorts(t, testScenario, controlHostOneID, controlHostOnePorts)
	assert.NoError(t, err, "failed to verify and delete ports")
	err = verifyPorts(t, testScenario, controlHostTwoID, controlHostTwoPorts)
	assert.NoError(t, err, "failed to verify and delete ports")

	// IMPORT test (expected to create endpoints without triggering playbooks)
	runUnderCloudActionTest(t, testScenario, config, importProvisioningAction, "", "")

	// verify ports
	err = verifyPorts(t, testScenario, controlHostOneID, controlHostOnePorts)
	assert.NoError(t, err, "failed to verify and delete ports")
	err = verifyPorts(t, testScenario, controlHostTwoID, controlHostTwoPorts)
	assert.NoError(t, err, "failed to verify and delete ports")

	// delete cloudManager
	config.Action = deleteAction
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	underCloudManager, err = NewUnderCloud(config)
	assert.NoError(t, err, "failed to create cloudManager manager to delete cloudManager")
	deployer, err = underCloudManager.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cloudManager")
	// make sure cloudManager is removed
	assert.True(t, verifyUnderCloudDeleted(), "Site file is not deleted")
}

func TestUnderCloud(t *testing.T) {
	pContext := pongo2.Context{}
	expectedSites := "./test_data/expected_site.yml"

	runUnderCloudTest(t, expectedSites, pContext)
}
