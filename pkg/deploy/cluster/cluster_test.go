package cluster

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

const (
	workRoot                              = "/tmp/contrail_cluster"
	allInOneClusterTemplatePath           = "./test_data/test_all_in_one_cluster.tmpl"
	createPlaybooks                       = "./test_data/expected_ansible_create_playbook.yml"
	destroyPlaybooks                      = "./test_data/expected_ansible_destroy_playbook.yml"
	createAppformixPlaybooks              = "./test_data/expected_ansible_create_appformix_playbook.yml"
	updatePlaybooks                       = "./test_data/expected_ansible_update_playbook.yml"
	updateAppformixPlaybooks              = "./test_data/expected_ansible_update_appformix_playbook.yml"
	upgradePlaybooks                      = "./test_data/expected_ansible_upgrade_playbook.yml"
	upgradeAppformixPlaybooks             = "./test_data/expected_ansible_upgrade_appformix_playbook.yml"
	addComputePlaybooks                   = "./test_data/expected_ansible_add_compute_playbook.yml"
	deleteComputePlaybooks                = "./test_data/expected_ansible_delete_compute_playbook.yml"
	addAppformixComputePlaybooks          = "./test_data/expected_ansible_add_appformix_compute_playbook.yml"
	deleteAppformixComputePlaybooks       = "./test_data/expected_ansible_delete_appformix_compute_playbook.yml"
	addCSNPlaybooks                       = "./test_data/expected_ansible_add_csn_playbook.yml"
	addAppformixCSNPlaybooks              = "./test_data/expected_ansible_add_appformix_csn_playbook.yml"
	createEncryptPlaybooks                = "./test_data/expected_ansible_create_encrypt_playbook.yml"
	updateEncryptPlaybooks                = "./test_data/expected_ansible_update_encrypt_playbook.yml"
	upgradeEncryptPlaybooks               = "./test_data/expected_ansible_upgrade_encrypt_playbook.yml"
	addComputeEncryptPlaybooks            = "./test_data/expected_ansible_add_compute_encrypt_playbook.yml"
	allInOneKubernetesClusterTemplatePath = "./test_data/test_all_in_one_kubernetes_cluster.tmpl"
	upgradePlaybooksKubernetes            = "./test_data/expected_ansible_upgrade_playbook_kubernetes.yml"
	addComputePlaybooksKubernetes         = "./test_data/expected_ansible_add_kubernetes_compute.yml"
	deleteComputePlaybooksKubernetes      = "./test_data/expected_ansible_delete_kubernetes_compute.yml"
	allInOneVcenterClusterTemplatePath    = "./test_data/test_all_in_one_vcenter_server.tmpl"
	upgradePlaybooksvcenter               = "./test_data/expected_ansible_upgrade_playbook_vcenter.yml"
	allInOneClusterAppformixTemplatePath  = "./test_data/test_all_in_one_with_appformix.tmpl"
	clusterID                             = "test_cluster_uuid"

	generatedInstances   = workRoot + "/" + clusterID + "/" + defaultInstanceFile
	generatedInventory   = workRoot + "/" + clusterID + "/" + defaultInventoryFile
	generatedVcenterVars = workRoot + "/" + clusterID + "/" + defaultVcenterFile

	executedPlaybooks = workRoot + "/" + clusterID + "/" + "executed_ansible_playbook.yml"
	executedCommands  = workRoot + "/" + clusterID + "/" + "executed_cmd.yml"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func verifyEndpoints(t *testing.T, testScenario *integration.TestScenario,
	expectedEndpoints map[string]string) error {
	createdEndpoints := map[string]string{}
	for _, client := range testScenario.Clients {
		var response map[string][]interface{}
		url := fmt.Sprintf("/endpoints?parent_uuid=%s", clusterID)
		_, err := client.Read(context.Background(), url, &response)
		assert.NoError(t, err, "Unable to list endpoints of the cluster")
		for _, endpoint := range response["endpoints"] {
			e := endpoint.(map[string]interface{}) //nolint: errcheck
			// TODO(ijohnson) remove using DisplayName as prefix
			// once UI takes prefix as input.
			var prefix = e["display_name"]
			if v, ok := e["prefix"]; ok {
				prefix = v
			}
			createdEndpoints[prefix.(string)] = e["public_url"].(string) //nolint: errcheck
		}
	}
	for k, e := range expectedEndpoints {
		if v, ok := createdEndpoints[k]; ok {
			if e != v {
				return fmt.Errorf("endpoint expected: %s, actual: %s for service %s", e, v, k)
			}
		} else {
			return fmt.Errorf("missing endpoint for service %s", k)
		}
	}
	return nil
}

func verifyClusterDeleted(clusterUUID string) bool {
	// Make sure working dir is deleted
	_, err := os.Stat(workRoot + "/" + clusterUUID)
	return os.IsNotExist(err)
}

func unmarshalYaml(t *testing.T, yamlFile string) map[string]interface{} {
	var yamlMap map[string]interface{}
	yamlBytes, err := ioutil.ReadFile(yamlFile)
	assert.NoError(t, err, "Error when reading yaml file %s", yamlFile)
	err = yaml.Unmarshal(yamlBytes, &yamlMap)
	assert.NoError(t, err, "Unable to unmarshal yaml file %s", yamlFile)
	return yamlMap
}

func assertYamlFileContainsOther(
	t *testing.T, expectedContainedYamlFile, actualYamlFile string, msgAndArgs ...interface{},
) bool {
	expectedMap := unmarshalYaml(t, expectedContainedYamlFile)
	actualMap := unmarshalYaml(t, actualYamlFile)
	ok := true
	for k, v := range expectedMap {
		ok = assert.Contains(t, actualMap, k, msgAndArgs...) && assert.Equal(t, v, actualMap[k], msgAndArgs...) && ok
	}
	return ok
}

func assertYamlFilesAreEqual(t *testing.T, expectedYamlFile, actualYamlFile string, msgAndArgs ...interface{}) bool {
	expectedMap := unmarshalYaml(t, expectedYamlFile)
	actualMap := unmarshalYaml(t, actualYamlFile)
	return assert.Equal(t, expectedMap, actualMap, msgAndArgs...)
}

func compareFiles(t *testing.T, expectedFile, generatedFile string) bool {
	generatedData, err := ioutil.ReadFile(generatedFile)
	assert.NoErrorf(t, err, "Unable to read generated: %s", generatedFile)
	expectedData, err := ioutil.ReadFile(expectedFile)
	assert.NoErrorf(t, err, "Unable to read expected: %s", expectedFile)
	return bytes.Equal(generatedData, expectedData)
}

func removeFile(t *testing.T, path string) {
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		assert.NoErrorf(t, err, "failed to delete %s", path)
	}
}

func assertGeneratedInstancesContainExpected(t *testing.T, expected string, msgAndArgs ...interface{}) {
	assertYamlFileContainsOther(t, expected, generatedInstances, msgAndArgs...)
}

func assertGeneratedInstancesEqual(t *testing.T, expected string, msgAndArgs ...interface{}) {
	assertYamlFilesAreEqual(t, expected, generatedInstances, msgAndArgs...)
}

func verifyPlaybooks(t *testing.T, expected string) bool {
	return compareFiles(t, expected, executedPlaybooks)
}

func verifyCommandsExecuted(t *testing.T, expected string) bool {
	return compareFiles(t, expected, executedCommands)
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

type clusterActionTestSpec struct {
	action            string
	expectedPlaybooks string
}

func runAllInOneClusterTest(t *testing.T, computeType string) {
	pContext := pongo2.Context{
		"TYPE":         computeType,
		"MGMT_INT_IP":  "127.0.0.1",
		"CLUSTER_NAME": t.Name(),
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
		"appformix": "http://127.0.0.1:9001",
	}
	expectedInstances := "./test_data/expected_all_in_one_instances.yml"
	switch computeType {
	case "dpdk":
		expectedInstances = "./test_data/expected_all_in_one_dpdk_instances.yml"
	case "sriov":
		expectedInstances = "./test_data/expected_all_in_one_sriov_instances.yml"
	}

	runTest(t, expectedInstances, "", pContext, expectedEndpoints, allInOneClusterTemplatePath,
		createPlaybooks,
		updatePlaybooks,
		[]clusterActionTestSpec{{
			action:            upgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            addComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            deleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            addCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: importProvisioningAction,
		}})
}

// nolint: gocyclo
func runTest(t *testing.T, expectedInstance, expectedInventory string,
	pContext map[string]interface{}, expectedEndpoints map[string]string, tsPath string,
	expectedCreatePlaybooks, expectedUpdatePlaybooks string, clusterActionSpecs []clusterActionTestSpec) {

	// Create the cluster and related objects
	ts, err := integration.LoadTest(tsPath, pContext)
	require.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	s, err := integration.NewAdminHTTPClient(server.URL())
	assert.NoError(t, err)

	config := &Config{
		APIServer:    s,
		ClusterID:    clusterID,
		Action:       createAction,
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}

	for _, tt := range []struct {
		action            string
		expectedPlaybooks string
	}{{
		action:            createAction,
		expectedPlaybooks: expectedCreatePlaybooks,
	}, {
		action:            updateAction,
		expectedPlaybooks: expectedUpdatePlaybooks,
	}} {
		runClusterCreateUpdateTest(
			t, ts, config, tt.action, tt.expectedPlaybooks, expectedInstance, expectedInventory, expectedEndpoints,
		)
	}

	for _, spec := range clusterActionSpecs {
		runClusterActionTest(t, ts, config, spec.action, expectedInstance, spec.expectedPlaybooks, expectedEndpoints)
	}

	// delete cluster
	config.Action = deleteAction
	removeFile(t, executedPlaybooks)
	manageCluster(t, config)
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(clusterID), "Instance file is not deleted during cluster delete")
}

func runClusterCreateUpdateTest(
	t *testing.T, ts *integration.TestScenario, config *Config, action string,
	expectedPlaybooks, expectedInstance, expectedInventory string, expectedEndpoints map[string]string) {

	//cleanup old files
	removeFile(t, generatedInstances)
	removeFile(t, generatedInventory)
	removeFile(t, executedPlaybooks)

	config.Action = action
	manageCluster(t, config)

	assertGeneratedInstancesEqual(t, expectedInstance,
		fmt.Sprintf("Instance file created during cluster %s is not as expected", action))
	if expectedInventory != "" {
		assert.Truef(t, compareFiles(t, expectedInventory, generatedInventory),
			"Inventory file created during cluster %s is not as expected", action)
	}
	assert.Truef(t, verifyPlaybooks(t, expectedPlaybooks),
		"Expected list of %s playbooks are not executed", action)

	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	assert.NoError(t, verifyEndpoints(t, ts, expectedEndpoints))
}

func manageCluster(t *testing.T, c *Config) {
	clusterDeployer, err := NewCluster(c, testutil.NewFileWritingExecutor(executedCommands))
	assert.NoErrorf(t, err, "failed to create cluster manager to %s cluster", c.Action)
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoErrorf(t, err, "failed to manage(%s) cluster", c.Action)
}

// nolint: gocyclo
func runClusterActionTest(
	t *testing.T, ts *integration.TestScenario, config *Config, action, expectedInstance string,
	expectedPlaybooks string, expectedEndpoints map[string]string,
) {
	// set action field in the contrail-cluster resource
	cluster := map[string]interface{}{"uuid": clusterID, "provisioning_action": action}
	config.Action = updateAction
	switch action {
	case upgradeProvisioningAction:
		cluster["provisioning_state"] = "NOSTATE"
	case importProvisioningAction:
		config.Action = createAction
		cluster["provisioning_action"] = ""
		err := os.Remove(generatedInstances)
		assert.NoError(t, err, "failed to delete instances.yml")
	case addComputeProvisioningAction, addCVFMProvisioningAction,
		deleteComputeProvisioningAction, addCSNProvisioningAction, destroyAction:
		err := os.Remove(generatedInstances)
		assert.NoError(t, err, "failed to delete instances.yml")
	}

	data := map[string]interface{}{"contrail-cluster": cluster}
	for _, client := range ts.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/contrail-cluster/%s", clusterID)
		_, err := client.Update(context.Background(), url, &data, &response)
		assert.NoErrorf(t, err, "failed to set %s action in contrail cluster", action)
		break
	}
	removeFile(t, executedPlaybooks)
	manageCluster(t, config)
	if expectedInstance != "" {
		assertGeneratedInstancesEqual(t, expectedInstance,
			"Instance file created during cluster %s is not as expected", action)
	}
	if expectedPlaybooks != "" {
		assert.True(t, verifyPlaybooks(t, expectedPlaybooks),
			fmt.Sprintf("Expected list of %s playbooks are not executed", action))
	}
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	assert.NoError(t, verifyEndpoints(t, ts, expectedEndpoints))
}

func TestAllInOneAppformix(t *testing.T) {
	runAllInOneAppformixTest(t, "kernel")
}

func runAllInOneAppformixTest(t *testing.T, computeType string) {
	context := pongo2.Context{
		"TYPE":        computeType,
		"MGMT_INT_IP": "127.0.0.1",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:9100",
		"nodejs":    "https://127.0.0.1:8144",
		"telemetry": "http://127.0.0.1:9101",
		"baremetal": "http://127.0.0.1:6386",
		"swift":     "http://127.0.0.1:8081",
		"glance":    "http://127.0.0.1:9293",
		"compute":   "http://127.0.0.1:8775",
		"keystone":  "http://127.0.0.1:5000",
		"appformix": "http://127.0.0.1:9001",
	}
	expectedInstances := "./test_data/expected_all_in_one_with_appformix.yml"
	switch computeType {
	case "dpdk":
		expectedInstances = "./test_data/expected_all_in_one_dpdk_instances.yml"
	case "sriov":
		expectedInstances = "./test_data/expected_all_in_one_sriov_instances.yml"
	}
	appformixFilesCleanup := createDummyAppformixFiles(t)
	defer appformixFilesCleanup()
	runTest(t, expectedInstances, "", context, expectedEndpoints, allInOneClusterAppformixTemplatePath,
		createAppformixPlaybooks,
		updateAppformixPlaybooks,
		[]clusterActionTestSpec{{
			action:            upgradeProvisioningAction,
			expectedPlaybooks: upgradeAppformixPlaybooks,
		}, {
			action:            addComputeProvisioningAction,
			expectedPlaybooks: addAppformixComputePlaybooks,
		}, {
			action:            deleteComputeProvisioningAction,
			expectedPlaybooks: deleteAppformixComputePlaybooks,
		}, {
			action:            addCSNProvisioningAction,
			expectedPlaybooks: addAppformixCSNPlaybooks,
		}, {
			action: importProvisioningAction,
		}, {
			action:            destroyAction,
			expectedPlaybooks: destroyPlaybooks,
		}},
	)

}

func createDummyAppformixFiles(t *testing.T) func() {
	// create appformix config.yml file
	configFile := workRoot + "/" + "appformix-ansible-deployer/appformix/config.yml"
	configData := []byte(`{"appformix_version": "3.0.0"}`)
	err := fileutil.WriteToFile(configFile, configData, defaultFilePermRWOnly)
	assert.NoErrorf(t, err, "Unable to write file: %s", configFile)

	return func() {
		if err = os.Remove(configFile); err != nil && !os.IsNotExist(err) {
			t.Logf("Couldn't cleanup file: %s, err: %v", configFile, err)
		}
	}
}

func TestXflowOutOfBand(t *testing.T) {
	ts, err := integration.LoadTest(
		"test_data/test_xflow_outofband_cluster.tmpl",
		map[string]interface{}{
			"appformixFlowsUUID":     "f0151fa2-2db7-476f-b4a9-58fcda2130b7",
			"nodeUUID":               "d63c420a-d4af-49d2-a7de-dd0665f2134a",
			"contrailClusterUUID":    clusterID,
			"appformixClusterUUID":   "appformix-cluster-uuid",
			"openstackClusterUUID":   "openstack-cluster-uuid",
			"appformixFlowsNodeUUID": "appformix-flows-node-uuid",
		},
	)
	if err != nil {
		t.Fatal("Unable to load test scenario", err)
	}
	appformixFilesCleanup := createDummyAppformixFiles(t)
	defer appformixFilesCleanup()

	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	d, ok := getClusterDeployer(t, &Config{
		APIServer:    ts.Clients["default"],
		ClusterID:    clusterID,
		Action:       "create",
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}).(*contrailAnsibleDeployer)
	require.True(t, ok, "unable to cast deployer to contrailAnsibleDeployer")

	err = d.createInventory()
	assert.NoError(t, err, "unable to create inventory")

	expectedInstance := "test_data/expected_xflow_outofband_instances.yaml"

	assertGeneratedInstancesContainExpected(t, expectedInstance,
		"Instance file created during cluster create is not as expected")
}

func TestXflowInBand(t *testing.T) {
	ts, err := integration.LoadTest(
		"test_data/test_xflow_inband_cluster.tmpl",
		map[string]interface{}{
			"appformixFlowsUUID":     "f0151fa2-2db7-476f-b4a9-58fcda2130b7",
			"nodeUUID":               "d63c420a-d4af-49d2-a7de-dd0665f2134a",
			"contrailClusterUUID":    clusterID,
			"appformixClusterUUID":   "appformix-cluster-uuid",
			"openstackClusterUUID":   "openstack-cluster-uuid",
			"appformixFlowsNodeUUID": "appformix-flows-node-uuid",
		},
	)
	if err != nil {
		t.Fatal("Unable to load test scenario", err)
	}
	appformixFilesCleanup := createDummyAppformixFiles(t)
	defer appformixFilesCleanup()

	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	d, ok := getClusterDeployer(t, &Config{
		APIServer:    ts.Clients["default"],
		ClusterID:    clusterID,
		Action:       "create",
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}).(*contrailAnsibleDeployer)
	require.True(t, ok, "unable to cast deployer to contrailAnsibleDeployer")

	err = d.createInventory()
	assert.NoError(t, err, "unable to create inventory")

	expectedInstance := "test_data/expected_xflow_inband_instances.yaml"

	assertGeneratedInstancesContainExpected(t, expectedInstance,
		"Instance file created during cluster create is not as expected")
}

func getClusterDeployer(t *testing.T, config *Config) base.Deployer {
	cluster, err := NewCluster(config, testutil.NewFileWritingExecutor(executedCommands))
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := cluster.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	return deployer
}

func TestAllInOneVfabricManager(t *testing.T) {
	pContext := pongo2.Context{
		"TYPE":            "kernel",
		"MGMT_INT_IP":     "127.0.0.1",
		"VFABRIC_MANAGER": true,
		"CLUSTER_NAME":    t.Name(),
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
		"appformix": "http://127.0.0.1:9001",
	}
	expectedInstances := "./test_data/expected_all_in_one_vfabric_manager_instances.yml"

	runTest(t, expectedInstances, "", pContext, expectedEndpoints, allInOneClusterTemplatePath,
		createPlaybooks,
		updatePlaybooks,
		[]clusterActionTestSpec{{
			action:            upgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            addComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            deleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            addCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: importProvisioningAction,
		}, {
			action:            addCVFMProvisioningAction,
			expectedPlaybooks: updatePlaybooks,
		}})
}

func TestAllInOneClusterWithDatapathEncryption(t *testing.T) {
	pContext := pongo2.Context{
		"DATAPATH_ENCRYPT": true,
		"MGMT_INT_IP":      "127.0.0.1",
		"CLUSTER_NAME":     t.Name(),
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
	runTest(t, "./test_data/expected_all_in_one_instances.yml",
		"./test_data/expected_all_in_one_inventory.yml", pContext, expectedEndpoints, allInOneClusterTemplatePath,
		createEncryptPlaybooks,
		updateEncryptPlaybooks,
		[]clusterActionTestSpec{{
			action:            upgradeProvisioningAction,
			expectedPlaybooks: upgradeEncryptPlaybooks,
		}, {
			action:            addComputeProvisioningAction,
			expectedPlaybooks: addComputeEncryptPlaybooks,
		}, {
			action:            deleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            addCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: importProvisioningAction,
		}})
}

func TestClusterWithDeploymentNetworkAsControlDataNet(t *testing.T) {
	pContext := pongo2.Context{
		"MGMT_INT_IP":     "127.0.0.1",
		"CONTROL_NODES":   "127.0.0.1",
		"OPENSTACK_NODES": "127.0.0.1",
		"CLUSTER_NAME":    t.Name(),
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
	runTest(t, "./test_data/expected_same_mgmt_ctrldata_net_instances.yml", "",
		pContext, expectedEndpoints, allInOneClusterTemplatePath,
		createPlaybooks,
		updatePlaybooks,
		[]clusterActionTestSpec{{
			action:            upgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            addComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            deleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            addCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: importProvisioningAction,
		}})
}

func TestClusterWithSeperateDeploymentAndControlDataNet(t *testing.T) {
	pContext := pongo2.Context{
		"MGMT_INT_IP":                 "10.1.1.1",
		"CONTROL_NODES":               "127.0.0.1",
		"CONTROLLER_NODES":            "127.0.0.1",
		"OPENSTACK_NODES":             "127.0.0.1",
		"OPENSTACK_INTERNAL_VIP":      "127.0.0.1",
		"CONTRAIL_EXTERNAL_VIP":       "10.1.1.100",
		"CONTAINER_REGISTRY_USERNAME": "testRegistry",
		"CONTAINER_REGISTRY_PASSWORD": "testRegistry123",
		"ZTP_ROLE":                    true,
		"SSL_ENABLE":                  "yes",
		"CLUSTER_NAME":                t.Name(),
	}
	expectedEndpoints := map[string]string{
		"config":    "https://10.1.1.100:8082",
		"nodejs":    "https://10.1.1.100:8143",
		"telemetry": "https://10.1.1.100:8081",
		"baremetal": "https://127.0.0.1:6385",
		"swift":     "https://127.0.0.1:8080",
		"glance":    "https://127.0.0.1:9292",
		"compute":   "https://127.0.0.1:8774",
		"keystone":  "https://127.0.0.1:5000",
	}

	runTest(t, "./test_data/expected_multi_interface_instances.yml", "",
		pContext, expectedEndpoints, allInOneClusterTemplatePath,
		createPlaybooks,
		updatePlaybooks,
		[]clusterActionTestSpec{{
			action:            upgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            addComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            deleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            addCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: importProvisioningAction,
		}})
}

func TestCredAllInOneClusterTest(t *testing.T) {
	pContext := pongo2.Context{
		"CUSTOMIZE":     true,
		"CREDS":         true,
		"MGMT_INT_IP":   "127.0.0.1",
		"CONTROL_NODES": "127.0.0.1",
		"ENABLE_ZTP":    true,
		"WEBUI_NODES":   "10.1.1.35",
		"CLUSTER_NAME":  t.Name(),
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:8082",
		"nodejs":    "https://10.1.1.35:8143",
		"telemetry": "http://127.0.0.1:8081",
		"baremetal": "http://127.0.0.1:6385",
		"swift":     "http://127.0.0.1:8080",
		"glance":    "http://127.0.0.1:9292",
		"compute":   "http://127.0.0.1:8774",
		"keystone":  "http://127.0.0.1:5000",
	}
	expectedInstances := "./test_data/expected_creds_all_in_one_instances.yml"

	runTest(t, expectedInstances, "", pContext, expectedEndpoints, allInOneClusterTemplatePath,
		createPlaybooks,
		updatePlaybooks,
		[]clusterActionTestSpec{{
			action:            upgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            addComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            deleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            addCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: importProvisioningAction,
		}})
}

func TestKubernetesCluster(t *testing.T) {
	pContext := pongo2.Context{
		"TYPE":        "kernel",
		"MGMT_INT_IP": "127.0.0.1",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:8082",
		"nodejs":    "https://127.0.0.1:8143",
		"telemetry": "http://127.0.0.1:8081",
	}
	runTest(t, "./test_data/expected_all_in_one_kubernetes_instances.yml", "",
		pContext, expectedEndpoints, allInOneKubernetesClusterTemplatePath,
		"./test_data/expected_ansible_create_playbook_kubernetes.yml",
		"./test_data/expected_ansible_update_playbook_kubernetes.yml",
		[]clusterActionTestSpec{{
			action:            upgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooksKubernetes,
		}, {
			action:            addComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooksKubernetes,
		}, {
			action:            deleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooksKubernetes,
		}, {
			action: importProvisioningAction,
		}})
}

func TestVcenterCluster(t *testing.T) {
	pContext := pongo2.Context{
		"TYPE":             "ESXI",
		"ESXI":             "10.84.16.11",
		"MGMT_INT_IP":      "127.0.0.1",
		"CONTROLLER_NODES": "127.0.0.1",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:8082",
		"nodejs":    "https://127.0.0.1:8143",
		"telemetry": "http://127.0.0.1:8081",
	}
	runVcenterClusterTest(t, "./test_data/expected_all_in_one_vcenter_instances.yml",
		"./test_data/expected_all_in_one_vcenter_vars.yml", pContext, expectedEndpoints)
}

// nolint: gocyclo
func runVcenterClusterTest(t *testing.T, expectedInstance, expectedVcenterVars string,
	pContext map[string]interface{}, expectedEndpoints map[string]string) {
	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneVcenterClusterTemplatePath, pContext)
	require.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	s, err := integration.NewAdminHTTPClient(server.URL())
	assert.NoError(t, err)

	config := &Config{
		APIServer:    s,
		ClusterID:    clusterID,
		Action:       createAction,
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}
	// create cluster
	removeFile(t, executedPlaybooks)
	manageCluster(t, config)
	assertGeneratedInstancesEqual(t, expectedInstance,
		"Instance file created during cluster create is not as expected")
	assert.True(t, compareFiles(t, expectedVcenterVars, generatedVcenterVars),
		"Vcenter_vars file created during cluster create is not as expected")
	assert.True(t, verifyPlaybooks(t, "./test_data/expected_ansible_create_playbook_vcenter.yml"),
		"Expected list of create playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// update cluster
	config.Action = updateAction
	// remove instances.yml to trigger cluster update
	removeFile(t, generatedInstances)
	removeFile(t, executedPlaybooks)

	manageCluster(t, config)
	assertGeneratedInstancesEqual(t, expectedInstance,
		"Instance file created during cluster update is not as expected")
	assert.True(t, verifyPlaybooks(t, "./test_data/expected_ansible_update_playbook_vcenter.yml"),
		"Expected list of update playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// UPGRADE test
	runClusterActionTest(t, ts, config,
		upgradeProvisioningAction, expectedInstance,
		upgradePlaybooksvcenter, expectedEndpoints)

	// IMPORT test (expected to create endpoints withtout triggering playbooks)
	runClusterActionTest(t, ts, config,
		importProvisioningAction, expectedInstance, "", expectedEndpoints)

	// delete cluster
	config.Action = deleteAction
	removeFile(t, executedPlaybooks)

	manageCluster(t, config)
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(clusterID), "Instance file is not deleted during cluster delete")
}
func TestWindowsCompute(t *testing.T) {
	ts, err := integration.LoadTest("./test_data/test_windows_compute.yml", nil)
	require.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, ts, server)
}

func TestTripleoClusterImport(t *testing.T) {
	// Create the cluster and related objects
	ts, err := integration.LoadTest(
		allInOneClusterTemplatePath,
		pongo2.Context{
			"TYPE":                   "kernel",
			"MGMT_INT_IP":            "127.0.0.1",
			"OPENSTACK_INTERNAL_VIP": "overcloud.localdomain",
			"CONTRAIL_EXTERNAL_VIP":  "overcloud.localdomain",
			"SSL_ENABLE":             "yes",
			"PROVISIONER_TYPE":       "tripleo",
			"PROVISIONING_STATE":     "CREATED",
			"PROVISIONING_ACTION":    "",
			"CLUSTER_NAME":           t.Name(),
		},
	)
	require.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	s, err := integration.NewAdminHTTPClient(server.URL())
	assert.NoError(t, err)

	config := &Config{
		APIServer:    s,
		ClusterID:    clusterID,
		Action:       createAction,
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}
	// create cluster
	removeFile(t, executedPlaybooks)

	clusterDeployer, err := NewCluster(config, testutil.NewFileWritingExecutor(executedCommands))
	assert.NoError(t, err, "failed to create cluster manager to import tripleo cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(import) tripleo cluster")

	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created as part of import
	err = verifyEndpoints(
		t, ts,
		map[string]string{
			"config":    "https://overcloud.localdomain:9100",
			"nodejs":    "https://overcloud.localdomain:8144",
			"telemetry": "https://overcloud.localdomain:9101",
			"baremetal": "https://overcloud.localdomain:6386",
			"swift":     "https://overcloud.localdomain:8081",
			"glance":    "https://overcloud.localdomain:9293",
			"compute":   "https://overcloud.localdomain:8775",
			"keystone":  "https://overcloud.localdomain:5000",
		},
	)
	assert.NoError(t, err)
	// delete cluster
	config.Action = deleteAction
	clusterDeployer, err = NewCluster(config, testutil.NewFileWritingExecutor(executedCommands))
	assert.NoError(t, err, "failed to create cluster manager to delete cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to delete triple0 cluster")
}
