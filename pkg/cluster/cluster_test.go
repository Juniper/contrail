package cluster

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/apisrv"
)

const (
	allInOneClusterTemplatePath = "./test_data/test_all_in_one_cluster.tmpl"
	createPlaybooks             = "./test_data/expected_ansible_create_playbook.yml"
	updatePlaybooks             = "./test_data/expected_ansible_update_playbook.yml"
	upgradePlaybooks            = "./test_data/expected_ansible_upgrade_playbook.yml"
	addComputePlaybooks         = "./test_data/expected_ansible_add_compute_playbook.yml"
	addCSNPlaybooks             = "./test_data/expected_ansible_add_csn_playbook.yml"
	createEncryptPlaybooks      = "./test_data/expected_ansible_create_encrypt_playbook.yml"
	updateEncryptPlaybooks      = "./test_data/expected_ansible_update_encrypt_playbook.yml"
	upgradeEncryptPlaybooks     = "./test_data/expected_ansible_upgrade_encrypt_playbook.yml"
	addComputeEncryptPlaybooks  = "./test_data/expected_ansible_add_compute_encrypt_playbook.yml"
	clusterID                   = "test_cluster_uuid"
)

func TestMain(m *testing.M) {
	apisrv.SetupAndRunTest(m)
}

func verifyEndpoints(t *testing.T, testScenario *apisrv.TestScenario,
	expectedEndpoints map[string]string) error {
	createdEndpoints := map[string]string{}
	for _, client := range testScenario.Clients {
		var response map[string][]interface{}
		url := fmt.Sprintf("/endpoints?parent_uuid=%s", clusterID)
		_, err := client.Read(url, &response)
		assert.NoError(t, err, "Unable to list endpoints of the cluster")
		for _, endpoint := range response["endpoints"] {
			e := endpoint.(map[string]interface{})
			// TODO(ijohnson) remove using DisplayName as prefix
			// once UI takes prefix as input.
			var prefix = e["display_name"]
			if v, ok := e["prefix"]; ok {
				prefix = v
			}
			createdEndpoints[prefix.(string)] = e["public_url"].(string)
		}
	}
	for k, e := range expectedEndpoints {
		if v, ok := createdEndpoints[k]; ok {
			if e != v {
				return fmt.Errorf("Endpoint expected: %s, actual: %s for service %s", e, v, k)
			}
		} else {
			return fmt.Errorf("Missing endpoint for service %s", k)
		}
	}
	return nil
}

func verifyClusterDeleted() bool {
	// Make sure working dir is deleted
	if _, err := os.Stat(defaultWorkRoot + "/" + clusterID); err == nil {
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

func compareGeneratedInstances(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedInstancesPath())
}

func compareGeneratedInventory(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedInventoryPath())
}

func verifyPlaybooks(t *testing.T, expected string) bool {
	return compareFiles(t, expected, executedPlaybooksPath())
}

func generatedInstancesPath() string {
	return defaultWorkRoot + "/" + clusterID + "/instances.yml"
}

func generatedInventoryPath() string {
	return defaultWorkRoot + "/" + clusterID + "/inventory.yml"
}

func executedPlaybooksPath() string {
	return defaultWorkRoot + "/" + clusterID + "/executed_ansible_playbook.yml"
}

// nolint: gocyclo
func runClusterActionTest(t *testing.T, testScenario apisrv.TestScenario,
	config *Config, action, expectedInstance, expectedInventory string,
	expectedEndpoints map[string]string) {
	// set action field in the contrail-cluster resource
	var err error
	var expectedPlaybooks string
	var data interface{}
	cluster := map[string]interface{}{"uuid": clusterID,
		"provisioning_action": action,
	}
	config.Action = "update"
	switch action {
	case "UPGRADE":
		cluster["provisioning_state"] = "NOSTATE"
		expectedPlaybooks = upgradePlaybooks
		if expectedInventory != "" {
			expectedPlaybooks = upgradeEncryptPlaybooks
		}
	case "ADD_COMPUTE":
		// remove instances.yml to mock trriger cluster update
		err = os.Remove(generatedInstancesPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete instances.yml")
		}
		expectedPlaybooks = addComputePlaybooks
		if expectedInventory != "" {
			expectedPlaybooks = addComputeEncryptPlaybooks
		}
	case "ADD_CSN":
		// remove instances.yml to mock trriger cluster update
		err = os.Remove(generatedInstancesPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete instances.yml")
		}
		expectedPlaybooks = addCSNPlaybooks
	case "IMPORT":
		config.Action = "create"
		cluster["provisioning_action"] = ""
	}
	data = map[string]interface{}{"contrail-cluster": cluster}
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/contrail-cluster/%s", clusterID)
		_, err = client.Update(url, &data, &response)
		assert.NoErrorf(t, err, "failed to set %s action in contrail cluster", action)
		break
	}
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterManager, err := NewCluster(config)
	assert.NoErrorf(t, err, "failed to create cluster manager to %s cluster", config.Action)
	err = clusterManager.Manage()
	assert.NoErrorf(t, err, "failed to manage(%s) cluster", action)
	if expectedInstance != "" {
		assert.True(t, compareGeneratedInstances(t, expectedInstance),
			fmt.Sprintf("Instance file created during cluster %s is not as expected", action))
	}
	if expectedPlaybooks != "" {
		assert.True(t, verifyPlaybooks(t, expectedPlaybooks),
			fmt.Sprintf("Expected list of %s playbooks are not executed", action))
	}
	// Wait for the in-memory endpoint cache to get updated
	apisrv.APIServer.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}
}

// nolint: gocyclo
func runClusterTest(t *testing.T, expectedInstance, expectedInventory string,
	context map[string]interface{}, expectedEndpoints map[string]string) {
	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := apisrv.MockServerWithKeystone("127.0.0.1:35357", keystoneAuthURL)
	defer ksPublic.Close()
	ksPrivate := apisrv.MockServerWithKeystone("127.0.0.1:5000", keystoneAuthURL)
	defer ksPrivate.Close()

	// Create the cluster and related objects
	var testScenario apisrv.TestScenario
	err := apisrv.LoadTestScenario(&testScenario, allInOneClusterTemplatePath, context)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := apisrv.RunDirtyTestScenario(t, &testScenario)
	defer cleanup()
	// create cluster config
	config := &Config{
		ID:           "alice",
		Password:     "alice_password",
		ProjectID:    "admin",
		AuthURL:      apisrv.TestServer.URL + "/keystone/v3",
		Endpoint:     apisrv.TestServer.URL,
		InSecure:     true,
		ClusterID:    clusterID,
		Action:       "create",
		LogLevel:     "debug",
		TemplateRoot: "configs/",
		Test:         true,
	}
	// create cluster
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup old executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterManager, err := NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	err = clusterManager.Manage()
	assert.NoError(t, err, "failed to manage(create) cluster")
	assert.True(t, compareGeneratedInstances(t, expectedInstance),
		"Instance file created during cluster create is not as expected")
	if expectedInventory != "" {
		assert.True(t, compareGeneratedInventory(t, expectedInventory),
			"Inventory file created during cluster create is not as expected")
		assert.True(t, verifyPlaybooks(t, createEncryptPlaybooks),
			"Expected list of create playbooks are not executed")
	} else {
		assert.True(t, verifyPlaybooks(t, createPlaybooks),
			"Expected list of create playbooks are not executed")
	}
	// Wait for the in-memory endpoint cache to get updated
	apisrv.APIServer.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// update cluster
	config.Action = "update"
	// remove instances.yml to trriger cluster update
	err = os.Remove(generatedInstancesPath())
	if err != nil {
		assert.NoError(t, err, "failed to delete instances.yml")
	}
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterManager, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to update cluster")
	err = clusterManager.Manage()
	assert.NoError(t, err, "failed to manage(update) cluster")
	assert.True(t, compareGeneratedInstances(t, expectedInstance),
		"Instance file created during cluster update is not as expected")
	if expectedInventory != "" {
		assert.True(t, compareGeneratedInventory(t, expectedInventory),
			"Inventory file created during cluster update is not as expected")
		assert.True(t, verifyPlaybooks(t, updateEncryptPlaybooks),
			"Expected list of update playbooks are not executed")
	} else {
		assert.True(t, verifyPlaybooks(t, updatePlaybooks),
			"Expected list of update playbooks are not executed")
	}
	// Wait for the in-memory endpoint cache to get updated
	apisrv.APIServer.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// UPGRADE test
	runClusterActionTest(t, testScenario, config,
		"UPGRADE", expectedInstance, expectedInventory,
		expectedEndpoints)

	// ADD_COMPUTE  test
	runClusterActionTest(t, testScenario, config,
		"ADD_COMPUTE", expectedInstance, expectedInventory,
		expectedEndpoints)

	// ADD_CSN  test
	runClusterActionTest(t, testScenario, config,
		"ADD_CSN", expectedInstance, expectedInventory,
		expectedEndpoints)

	// IMPORT test (expected to create endpoints without triggering playbooks)
	runClusterActionTest(t, testScenario, config,
		"IMPORT", "", "", expectedEndpoints)

	// delete cluster
	config.Action = "delete"
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterManager, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to delete cluster")
	err = clusterManager.Manage()
	assert.NoError(t, err, "failed to manage(delete) cluster")
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}

func runAllInOneClusterTest(t *testing.T, computeType string) {
	context := pongo2.Context{
		"TYPE":            computeType,
		"MGMT_INT_IP":     "127.0.0.1",
		"CONTROL_NODES":   "",
		"OPENSTACK_NODES": "",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:8082",
		"nodejs":    "https://127.0.0.1:8143",
		"telemetry": "http://127.0.0.1:8081",
		"baremetal": "http://127.0.0.1:6385",
		"swift":     "http://127.0.0.1:8080",
		"glance":    "http://127.0.0.1:9292",
		"compute":   "http://127.0.0.1:8774",
		"keystone":  "http://127.0.0.1:5000",
	}
	expectedInstances := "./test_data/expected_all_in_one_instances.yml"
	switch computeType {
	case "dpdk":
		expectedInstances = "./test_data/expected_all_in_one_dpdk_instances.yml"
	case "sriov":
		expectedInstances = "./test_data/expected_all_in_one_sriov_instances.yml"
	}

	runClusterTest(t, expectedInstances, "", context, expectedEndpoints)
}

func TestAllInOneCluster(t *testing.T) {
	runAllInOneClusterTest(t, "kernel")
}
func TestAllInOneDpdkCluster(t *testing.T) {
	runAllInOneClusterTest(t, "dpdk")
}

func TestAllInOneSriovCluster(t *testing.T) {
	runAllInOneClusterTest(t, "sriov")
}

func TestAllInOneClusterWithDatapathEncryption(t *testing.T) {
	context := pongo2.Context{
		"DATAPATH_ENCRYPT": true,
		"MGMT_INT_IP":      "127.0.0.1",
		"CONTROL_NODES":    "",
		"OPENSTACK_NODES":  "",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:8082",
		"nodejs":    "https://127.0.0.1:8143",
		"telemetry": "http://127.0.0.1:8081",
		"baremetal": "http://127.0.0.1:6385",
		"swift":     "http://127.0.0.1:8080",
		"glance":    "http://127.0.0.1:9292",
		"compute":   "http://127.0.0.1:8774",
		"keystone":  "http://127.0.0.1:5000",
	}
	runClusterTest(t, "./test_data/expected_all_in_one_instances.yml",
		"./test_data/expected_all_in_one_inventory.yml", context, expectedEndpoints)
}

func TestClusterWithManagementNetworkAsControlDataNet(t *testing.T) {
	context := pongo2.Context{
		"MGMT_INT_IP":     "127.0.0.1",
		"CONTROL_NODES":   "127.0.0.1",
		"OPENSTACK_NODES": "127.0.0.1",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:8082",
		"nodejs":    "https://127.0.0.1:8143",
		"telemetry": "http://127.0.0.1:8081",
		"baremetal": "http://127.0.0.1:6385",
		"swift":     "http://127.0.0.1:8080",
		"glance":    "http://127.0.0.1:9292",
		"compute":   "http://127.0.0.1:8774",
		"keystone":  "http://127.0.0.1:5000",
	}
	runClusterTest(t, "./test_data/expected_same_mgmt_ctrldata_net_instances.yml", "", context, expectedEndpoints)
}

func TestClusterWithSeperateManagementAndControlDataNet(t *testing.T) {
	context := pongo2.Context{
		"MGMT_INT_IP":            "10.1.1.1",
		"CONTROL_NODES":          "127.0.0.1",
		"CONTROLLER_NODES":       "127.0.0.1",
		"OPENSTACK_NODES":        "127.0.0.1",
		"OPENSTACK_INTERNAL_VIP": "127.0.0.1",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:8082",
		"nodejs":    "https://127.0.0.1:8143",
		"telemetry": "http://127.0.0.1:8081",
		"baremetal": "http://127.0.0.1:6385",
		"swift":     "http://127.0.0.1:8080",
		"glance":    "http://127.0.0.1:9292",
		"compute":   "http://127.0.0.1:8774",
		"keystone":  "http://127.0.0.1:5000",
	}

	runClusterTest(t, "./test_data/expected_multi_interface_instances.yml", "", context, expectedEndpoints)
}

func TestCredAllInOneClusterTest(t *testing.T) {
	context := pongo2.Context{
		"CREDS":           true,
		"TYPE":            "",
		"MGMT_INT_IP":     "127.0.0.1",
		"CONTROL_NODES":   "",
		"OPENSTACK_NODES": "",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:8082",
		"nodejs":    "https://127.0.0.1:8143",
		"telemetry": "http://127.0.0.1:8081",
		"baremetal": "http://127.0.0.1:6385",
		"swift":     "http://127.0.0.1:8080",
		"glance":    "http://127.0.0.1:9292",
		"compute":   "http://127.0.0.1:8774",
		"keystone":  "http://127.0.0.1:5000",
	}
	expectedInstances := "./test_data/expected_creds_all_in_one_instances.yml"

	runClusterTest(t, expectedInstances, context, expectedEndpoints)
}
