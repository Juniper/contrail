package undercloud_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/deploy/rhospd/undercloud"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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

func verifyPorts(t *testing.T, hc *client.HTTP, nodeID string, expectedPorts []string) error {
	resp, err := hc.ListPort(context.Background(), &services.ListPortRequest{
		Spec: &baseservices.ListSpec{
			Fields: []string{"name", "uuid"},
			Filters: []*baseservices.Filter{
				{
					Key:    "parent_uuid",
					Values: []string{nodeID},
				},
			},
		},
	})
	assert.NoError(t, err, "Unable to list ports of the node")

	createdPorts := map[string]string{}
	for _, p := range resp.GetPorts() {
		createdPorts[p.Name] = p.UUID
	}

	for _, p := range expectedPorts {
		if portID, ok := createdPorts[p]; ok {
			url := fmt.Sprintf("/port/%s", portID)
			var response interface{}
			_, err := hc.Delete(context.Background(), url, &response)
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
func runUnderCloudActionTest(
	t *testing.T,
	hc *client.HTTP,
	config *undercloud.Config,
	action,
	expectedSite,
	expectedPlaybooks string,
) {
	// set action field in the rhospd-cloud-manager resource
	var err error
	var data interface{}
	cloudManager := map[string]interface{}{"uuid": cloudManagerID,
		"provisioning_action": action,
	}
	config.Action = undercloud.UpdateAction
	switch action {
	case undercloud.ImportProvisioningAction:
		config.Action = undercloud.CreateAction
		cloudManager["provisioning_action"] = ""
	}
	data = map[string]interface{}{"rhospd-cloud-manager": cloudManager}

	var response map[string]interface{}
	url := fmt.Sprintf("/rhospd-cloud-manager/%s", cloudManagerID)
	_, err = hc.Update(context.Background(), url, &data, &response)
	assert.NoErrorf(t, err, "failed to set %s action in rhospd cloudManager", action)

	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	underCloudManager, err := undercloud.NewUnderCloud(config)
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
	ksPublic := integration.NewKeystoneServerFake(t, keystoneAuthURL, "", "")
	defer ksPublic.Close()
	ksPrivate := integration.NewKeystoneServerFake(t, keystoneAuthURL, "", "")
	defer ksPrivate.Close()

	// Create the cloudManager and related objects
	ts, err := integration.LoadTest(cloudManagerTemplatePath, pContext)
	require.NoError(t, err, "failed to load cloudManager test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	// create cloudManager
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup old executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}

	hc, err := integration.NewAdminHTTPClient(server.URL())
	assert.NoError(t, err)

	config := &undercloud.Config{
		APIServer:    hc,
		ResourceID:   cloudManagerID,
		Action:       undercloud.CreateAction,
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}

	underCloudManager, err := undercloud.NewUnderCloud(config)
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
	err = verifyPorts(t, hc, controlHostOneID, controlHostOnePorts)
	assert.NoError(t, err, "failed to verify and delete ports")
	err = verifyPorts(t, hc, controlHostTwoID, controlHostTwoPorts)
	assert.NoError(t, err, "failed to verify and delete ports")

	// update cloudManager
	config.Action = undercloud.UpdateAction
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
	underCloudManager, err = undercloud.NewUnderCloud(config)
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
	err = verifyPorts(t, hc, controlHostOneID, controlHostOnePorts)
	assert.NoError(t, err, "failed to verify and delete ports")
	err = verifyPorts(t, hc, controlHostTwoID, controlHostTwoPorts)
	assert.NoError(t, err, "failed to verify and delete ports")

	// IMPORT test (expected to create endpoints without triggering playbooks)
	runUnderCloudActionTest(t, hc, config, undercloud.ImportProvisioningAction, "", "")

	// verify ports
	err = verifyPorts(t, hc, controlHostOneID, controlHostOnePorts)
	assert.NoError(t, err, "failed to verify and delete ports")
	err = verifyPorts(t, hc, controlHostTwoID, controlHostTwoPorts)
	assert.NoError(t, err, "failed to verify and delete ports")

	// delete cloudManager
	config.Action = undercloud.DeleteAction
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	underCloudManager, err = undercloud.NewUnderCloud(config)
	assert.NoError(t, err, "failed to create cloudManager manager to delete cloudManager")
	deployer, err = underCloudManager.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cloudManager")
	// make sure cloudManager is removed
	assert.True(t, verifyUnderCloudDeleted(), "Site file is not deleted")
}

func TestUnderCloud(t *testing.T) {
	runUnderCloudTest(t, "./test_data/expected_site.yml", pongo2.Context{})
}
