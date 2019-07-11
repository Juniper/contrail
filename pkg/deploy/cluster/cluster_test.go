package cluster

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

const (
	defaultAdminUser                      = "admin"
	defaultAdminPassword                  = "contrail123"
	workRoot                              = "/tmp/contrail_cluster"
	allInOneClusterTemplatePath           = "./test_data/test_all_in_one_cluster.tmpl"
	createPlaybooks                       = "./test_data/expected_ansible_create_playbook.yml"
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
	allInOnevcenterClusterTemplatePath    = "./test_data/test_all_in_one_vcenter_server.tmpl"
	upgradePlaybooksvcenter               = "./test_data/expected_ansible_upgrade_playbook_vcenter.yml"
	allInOneMCClusterTemplatePath         = "./test_data/test_mc_cluster.tmpl"
	allInOneMCClusterUpdateTemplatePath   = "./test_data/test_mc_update_cluster.tmpl"
	allInOneMCClusterDeleteTemplatePath   = "./test_data/test_mc_delete_cluster.tmpl"
	allInOneMCCloudUpdateTemplatePath     = "./test_data/test_mc_update_pvt_cloud.tmpl"
	allInOneClusterAppformixTemplatePath  = "./test_data/test_all_in_one_with_appformix.tmpl"
	clusterID                             = "test_cluster_uuid"

	expectedMCClusterTopology = "./test_data/expected_mc_cluster_topology.yml"
	expectedMCClusterSecret   = "./test_data/expected_mc_cluster_secret.yml"
	expectedContrailCommon    = "./test_data/expected_mc_contrail_common.yml"
	expectedGatewayCommon     = "./test_data/expected_mc_gateway_common.yml"

	expectedK8sCreatePlaybook = "./test_data/expected_ansible_create_playbook_kubernetes.yml"
	expectedK8sUpdatePlaybook = "./test_data/expected_ansible_update_playbook_kubernetes.yml"

	expectedVcenterCreatePlaybook = "./test_data/expected_ansible_create_playbook_vcenter.yml"
	expectedVcenterUpdatePlaybook = "./test_data/expected_ansible_update_playbook_vcenter.yml"

	expectedMCCreatePlaybook = "./test_data/expected_ansible_create_mc_playbook.yml"
	expectedMCUpdatePlaybook = "./test_data/expected_ansible_update_mc_playbook.yml"
	expectedMCDeletePlaybook = "./test_data/expected_ansible_delete_mc_playbook.yml"

	executedAnsiblePlaybook = "executed_ansible_playbook.yml"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func verifyEndpoints(t *testing.T, ts *integration.TestScenario, expectedEndpoints map[string]string) error {
	createdEndpoints := map[string]string{}
	url := fmt.Sprintf("/endpoints?parent_uuid=%s", clusterID)
	for _, httpClient := range ts.Clients {
		var response map[string][]interface{}
		_, err := httpClient.Read(context.Background(), url, &response)
		assert.NoError(t, err, "Unable to list endpoints of the cluster")

		for _, endpoint := range response["endpoints"] {
			e := endpoint.(map[string]interface{}) //nolint: errcheck
			// TODO(ijohnson) remove using DisplayName as prefix
			// once UI takes prefix as input.
			prefix := e["display_name"]
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

func verifyClusterDeleted() bool {
	// Make sure working dir is deleted
	_, err := os.Stat(workRoot + "/" + clusterID)
	return os.IsNotExist(err)
}

func verifyMCDeleted(httpClient *client.HTTP) bool {
	// Make sure mc working dir is deleted
	if _, err := os.Stat(workRoot + "/" + clusterID + "/" + mcWorkDir); err == nil {
		// mc working dir not deleted
		return false
	}
	clusterObjResp, err := httpClient.GetContrailCluster(
		context.Background(),
		&services.GetContrailClusterRequest{
			ID: clusterID,
		},
	)
	if err != nil {
		return false
	}

	return clusterObjResp.ContrailCluster.CloudRefs == nil
}

func unmarshalYaml(t *testing.T, path string) map[string]interface{} {
	var yamlMap map[string]interface{}

	yamlBytes, err := ioutil.ReadFile(path)
	assert.NoError(t, err, "Error when reading yaml file %s", path)
	err = yaml.Unmarshal(yamlBytes, &yamlMap)
	assert.NoError(t, err, "Unable to unmarshal yaml file %s", path)

	return yamlMap
}

func assertYamlFileContainsOther(
	t *testing.T, expectedContainedYamlFile, actualYamlFile string, msgAndArgs ...interface{},
) bool {
	expectedMap := unmarshalYaml(t, expectedContainedYamlFile)
	actualMap := unmarshalYaml(t, actualYamlFile)

	for k, v := range expectedMap {
		if !(assert.Contains(t, actualMap, k, msgAndArgs...) && assert.Equal(t, v, actualMap[k], msgAndArgs...)) {
			return false
		}
	}
	return true
}

func assertYamlFilesAreEqual(t *testing.T, expectedYamlFile, actualYamlFile string, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, unmarshalYaml(t, expectedYamlFile), unmarshalYaml(t, actualYamlFile), msgAndArgs...)
}

func compareFiles(t *testing.T, expectedFile, generatedFile string) bool {
	generatedData, err := ioutil.ReadFile(generatedFile)
	assert.NoErrorf(t, err, "Unable to read generated: %s", generatedFile)
	expectedData, err := ioutil.ReadFile(expectedFile)
	assert.NoErrorf(t, err, "Unable to read expected: %s", expectedFile)
	return bytes.Equal(generatedData, expectedData)
}

func pathToClusterFile(file string) string {
	return filepath.Join(workRoot, clusterID, file)
}

func pathToClusterFileMC(file string) string {
	return filepath.Join(workRoot, clusterID, mcWorkDir, file)
}

func removeFile(t *testing.T, file string) {
	if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
		assert.NoError(t, err, "failed to delete "+file)
	}
	return
}

func copyFile(src, dest string) error {
	data, err := fileutil.GetContent(src)
	if err != nil {
		return err
	}
	return fileutil.WriteToFile(dest, data, defaultFilePermRWOnly)
}

func createDummyCloudFiles(t *testing.T) func() {
	files := []struct {
		src  string
		dest string
	}{{
		// public cloud topology.yaml
		src:  "./test_data/public_cloud_topology.yml",
		dest: "/var/tmp/cloud/public_cloud_uuid/topology.yml",
	}, {
		// private cloud topology.yml
		src:  "./test_data/pvt_cloud_topology.yml",
		dest: "/var/tmp/cloud/pvt_cloud_uuid/topology.yml",
	}, {
		// public cloud secret.yml
		src:  "./test_data/public_cloud_secret.yml",
		dest: "/var/tmp/cloud/public_cloud_uuid/secret.yml",
	}}

	for _, f := range files {
		err := copyFile(f.src, f.dest)
		assert.NoErrorf(t, err, "Cannot copy file %s to location %s", f.src, f.dest)
	}

	return func() {
		for _, f := range files {
			err := os.Remove(f.dest)
			assert.NoErrorf(t, err, "Cannot remove file: %s", f.dest)
		}
	}
}

func createDummyAppformixFiles(t *testing.T) func() {

	// create appformix config.yml file
	configFile := workRoot + "/" + "appformix-ansible-deployer/appformix/config.yml"
	configData := []byte("{\n\"appformix_version\": \"3.0.0\"\n}")
	err := fileutil.WriteToFile(configFile, configData, defaultFilePermRWOnly)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to write file: %s", configFile)
	}
	return func() {
		// best effort method of deleting all the files
		// nolint: errcheck
		_ = os.Remove(configFile)
	}
}

// nolint: gocyclo
func runClusterActionTest(
	t *testing.T,
	ts *integration.TestScenario,
	config *Config,
	action, expectedInstance, expectedInventory string,
	expectedPlaybooks string,
	expectedEndpoints map[string]string,
) {
	// set action field in the contrail-cluster resource
	cluster := map[string]string{
		"uuid":                clusterID,
		"provisioning_action": action,
	}

	config.Action = updateAction

	if action == importProvisioningAction ||
		action == addComputeProvisioningAction ||
		action == deleteComputeProvisioningAction ||
		action == addCSNProvisioningAction {
		removeFile(t, pathToClusterFile(defaultInstanceFile))
	}

	switch action {
	case upgradeProvisioningAction:
		cluster["provisioning_state"] = statusNoState
		if expectedInventory != "" {
			expectedPlaybooks = upgradeEncryptPlaybooks
		}
	case addComputeProvisioningAction:
		if expectedInventory != "" {
			expectedPlaybooks = addComputeEncryptPlaybooks
		}
	case importProvisioningAction:
		config.Action = createAction
		cluster["provisioning_action"] = ""
	}
	data := map[string]interface{}{defaultResource: cluster}
	for _, httpClient := range ts.Clients {
		var response interface{}
		url := fmt.Sprintf("/contrail-cluster/%s", clusterID)
		_, err := httpClient.Update(context.Background(), url, &data, &response)
		assert.NoErrorf(t, err, "failed to set %s action in contrail cluster", action)
		break
	}

	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, action)
	if expectedInstance != "" {
		assertYamlFilesAreEqual(t, expectedInstance, pathToClusterFile(defaultInstanceFile),
			"Instance file created during cluster %s is not as expected", action)
	}
	if expectedPlaybooks != "" {
		assert.True(t, compareFiles(t, expectedPlaybooks, pathToClusterFile(executedAnsiblePlaybook)),
			fmt.Sprintf("Expected list of %s playbooks are not executed", action))
	}
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err := verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)
}

func deployTestCluster(t *testing.T, config *Config, action string) {
	clusterDeployer, err := NewCluster(config)
	assert.NoErrorf(t, err, "failed to create cluster manager to %s cluster", config.Action)
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	//TODO: is the action parameter necessary?
	assert.NoErrorf(t, err, "failed to manage (%s) cluster", action)
}

func createClusterConfig(t *testing.T) *Config {
	s, err := integration.NewAdminHTTPClient(server.URL())
	assert.NoError(t, err)

	return &Config{
		APIServer:    s,
		ClusterID:    clusterID,
		Action:       createAction,
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}
}

func startMockKeystone() func() {
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357",
		viper.GetString("keystone.authurl"),
		defaultAdminUser,
		defaultAdminPassword,
	)
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000",
		viper.GetString("keystone.authurl"),
		defaultAdminUser,
		defaultAdminPassword,
	)

	return func() {
		defer ksPublic.Close()
		defer ksPrivate.Close()
	}
}

// nolint: gocyclo
func runClusterTest(
	t *testing.T,
	expectedInstance, expectedInventory string,
	pContext map[string]interface{},
	expectedEndpoints map[string]string,
) {
	closeKs := startMockKeystone()
	defer closeKs()

	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneClusterTemplatePath, pContext)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	// create cluster config
	config := createClusterConfig(t)
	// create cluster
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	assertYamlFilesAreEqual(t, expectedInstance, pathToClusterFile(defaultInstanceFile),
		"Instance file created during cluster create is not as expected")
	if expectedInventory != "" {
		assert.True(t, compareFiles(t, expectedInventory, pathToClusterFile(defaultInventoryFile)),
			"Inventory file created during cluster create is not as expected")
		assert.True(t, compareFiles(t, createEncryptPlaybooks, pathToClusterFile(executedAnsiblePlaybook)),
			"Expected list of create playbooks are not executed")
	} else {
		assert.True(t, compareFiles(t, createPlaybooks, pathToClusterFile(executedAnsiblePlaybook)),
			"Expected list of create playbooks are not executed")
	}
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// update cluster
	config.Action = updateAction
	// remove instances.yml to trriger cluster update
	removeFile(t, pathToClusterFile(defaultInstanceFile))
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	assertYamlFilesAreEqual(t, expectedInstance, pathToClusterFile(defaultInstanceFile),
		"Instance file created during cluster update is not as expected")
	if expectedInventory != "" {
		assert.True(t, compareFiles(t, expectedInventory, pathToClusterFile(defaultInventoryFile)),
			"Inventory file created during cluster update is not as expected")
		assert.True(t, compareFiles(t, updateEncryptPlaybooks, pathToClusterFile(executedAnsiblePlaybook)),
			"Expected list of update playbooks are not executed")
	} else {
		assert.True(t, compareFiles(t, updatePlaybooks, pathToClusterFile(executedAnsiblePlaybook)),
			"Expected list of update playbooks are not executed")
	}
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// UPGRADE test
	runClusterActionTest(t, ts, config,
		upgradeProvisioningAction, expectedInstance, expectedInventory,
		upgradePlaybooks, expectedEndpoints)

	// ADD_COMPUTE  test
	runClusterActionTest(t, ts, config,
		addComputeProvisioningAction, expectedInstance, expectedInventory,
		addComputePlaybooks, expectedEndpoints)

	// DELETE_COMPUTE  test
	runClusterActionTest(t, ts, config,
		deleteComputeProvisioningAction, expectedInstance, "",
		deleteComputePlaybooks, expectedEndpoints)

	// ADD_CSN  test
	runClusterActionTest(t, ts, config,
		addCSNProvisioningAction, expectedInstance, expectedInventory,
		addCSNPlaybooks, expectedEndpoints)

	// IMPORT test (expected to create endpoints without triggering playbooks)
	runClusterActionTest(t, ts, config,
		importProvisioningAction, expectedInstance, "",
		"", expectedEndpoints)

	// delete cluster
	config.Action = deleteAction
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}

// nolint: gocyclo
func runAppformixClusterTest(
	t *testing.T,
	expectedInstance, expectedInventory string,
	pContext map[string]interface{},
	expectedEndpoints map[string]string,
) {
	closeKs := startMockKeystone()
	defer closeKs()

	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneClusterAppformixTemplatePath, pContext)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()
	// create cluster config
	config := createClusterConfig(t)
	// create cluster
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	assertYamlFilesAreEqual(t, expectedInstance, pathToClusterFile(defaultInstanceFile),
		"Instance file created during cluster create is not as expected")
	if expectedInventory != "" {
		assert.True(t, compareFiles(t, expectedInventory, pathToClusterFile(defaultInventoryFile)),
			"Inventory file created during cluster create is not as expected")
		assert.True(t, compareFiles(t, createEncryptPlaybooks, pathToClusterFile(executedAnsiblePlaybook)),
			"Expected list of create playbooks are not executed")
	} else {
		assert.True(t, compareFiles(t, createAppformixPlaybooks, pathToClusterFile(executedAnsiblePlaybook)),
			"Expected list of create playbooks are not executed")
	}
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// update cluster
	config.Action = updateAction
	// remove instances.yml to trigger cluster update
	removeFile(t, pathToClusterFile(defaultInstanceFile))
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	assertYamlFilesAreEqual(t, expectedInstance, pathToClusterFile(defaultInstanceFile),
		"Instance file created during cluster update is not as expected")
	if expectedInventory != "" {
		assert.True(t, compareFiles(t, expectedInventory, pathToClusterFile(defaultInventoryFile)),
			"Inventory file created during cluster update is not as expected")
		assert.True(t, compareFiles(t, updateEncryptPlaybooks, pathToClusterFile(executedAnsiblePlaybook)),
			"Expected list of update playbooks are not executed")
	} else {
		assert.True(t, compareFiles(t, updateAppformixPlaybooks, pathToClusterFile(executedAnsiblePlaybook)),
			"Expected list of update playbooks are not executed")
	}
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// UPGRADE test
	runClusterActionTest(t, ts, config,
		upgradeProvisioningAction, expectedInstance, expectedInventory,
		upgradeAppformixPlaybooks, expectedEndpoints)

	// ADD_COMPUTE  test
	runClusterActionTest(t, ts, config,
		addComputeProvisioningAction, expectedInstance, expectedInventory,
		addAppformixComputePlaybooks, expectedEndpoints)

	// DELETE_COMPUTE  test
	runClusterActionTest(t, ts, config,
		deleteComputeProvisioningAction, expectedInstance, "",
		deleteAppformixComputePlaybooks, expectedEndpoints)

	// ADD_CSN  test
	runClusterActionTest(t, ts, config,
		addCSNProvisioningAction, expectedInstance, expectedInventory,
		addAppformixCSNPlaybooks, expectedEndpoints)

	// IMPORT test (expected to create endpoints without triggering playbooks)
	runClusterActionTest(t, ts, config,
		importProvisioningAction, expectedInstance,
		"", "", expectedEndpoints)

	// delete cluster
	config.Action = deleteAction
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}

func runAllInOneClusterTest(t *testing.T, computeType string) {
	expectedInstances := "./test_data/expected_all_in_one_instances.yml"
	switch computeType {
	case "dpdk":
		expectedInstances = "./test_data/expected_all_in_one_dpdk_instances.yml"
	case "sriov":
		expectedInstances = "./test_data/expected_all_in_one_sriov_instances.yml"
	}

	runClusterTest(
		t,
		expectedInstances,
		"",
		pongo2.Context{
			"TYPE":            computeType,
			"MGMT_INT_IP":     "127.0.0.1",
			"CONTROL_NODES":   "",
			"OPENSTACK_NODES": "",
		},
		map[string]string{
			"config":                "http://127.0.0.1:8082",
			"nodejs":                "https://127.0.0.1:8143",
			"telemetry":             "http://127.0.0.1:8081",
			"baremetal":             "http://127.0.0.1:6385",
			"swift":                 "http://127.0.0.1:8080",
			"glance":                "http://127.0.0.1:9292",
			"compute":               "http://127.0.0.1:8774",
			"keystone":              "http://127.0.0.1:5000",
			"endpoint_user_created": "http://127.0.0.1:8082",
		},
	)
}

func runAllInOneAppformixTest(t *testing.T, computeType string) {
	expectedInstances := "./test_data/expected_all_in_one_with_appformix.yml"
	switch computeType {
	case "dpdk":
		expectedInstances = "./test_data/expected_all_in_one_dpdk_instances.yml"
	case "sriov":
		expectedInstances = "./test_data/expected_all_in_one_sriov_instances.yml"
	}
	appformixFilesCleanup := createDummyAppformixFiles(t)
	defer appformixFilesCleanup()
	runAppformixClusterTest(
		t,
		expectedInstances,
		"",
		pongo2.Context{
			"TYPE":            computeType,
			"MGMT_INT_IP":     "127.0.0.1",
			"CONTROL_NODES":   "",
			"OPENSTACK_NODES": "",
		},
		map[string]string{
			"config":    "http://127.0.0.1:9100",
			"nodejs":    "https://127.0.0.1:8144",
			"telemetry": "http://127.0.0.1:9101",
			"baremetal": "http://127.0.0.1:6386",
			"swift":     "http://127.0.0.1:8081",
			"glance":    "http://127.0.0.1:9293",
			"compute":   "http://127.0.0.1:8775",
			"keystone":  "http://127.0.0.1:5000",
			"appformix": "http://127.0.0.1:9001",
		},
	)
}

func TestXflow(t *testing.T) {
	ts, err := integration.LoadTest(
		"test_data/test_xflow_cluster.tmpl",
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
		t.Error("Unable to load test scenario", err)
	}
	appformixFilesCleanup := createDummyAppformixFiles(t)
	defer appformixFilesCleanup()

	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	config := createClusterConfig(t)

	cluster, err := NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := cluster.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")

	d, ok := deployer.(*contrailAnsibleDeployer)
	assert.True(t, ok, "unable to cast deployer to contrailAnsibleDeployer")

	err = d.createInventory()
	assert.NoError(t, err, "unable to create inventory")

	expectedInstance := "test_data/expected_xflow_instances.yaml"

	assertYamlFileContainsOther(t, expectedInstance, pathToClusterFile(defaultInstanceFile),
		"Instance file created during cluster create is not as expected")
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
func TestAllInOneAppformix(t *testing.T) {
	runAllInOneAppformixTest(t, "kernel")
}

func TestAllInOneClusterWithDatapathEncryption(t *testing.T) {
	runClusterTest(
		t,
		"./test_data/expected_all_in_one_instances.yml",
		"./test_data/expected_all_in_one_inventory.yml",
		pongo2.Context{
			"DATAPATH_ENCRYPT": true,
			"MGMT_INT_IP":      "127.0.0.1",
			"CONTROL_NODES":    "",
			"OPENSTACK_NODES":  "",
		},
		map[string]string{
			"config":    "http://127.0.0.1:8082",
			"nodejs":    "https://127.0.0.1:8143",
			"telemetry": "http://127.0.0.1:8081",
			"baremetal": "http://127.0.0.1:6385",
			"swift":     "http://127.0.0.1:8080",
			"glance":    "http://127.0.0.1:9292",
			"compute":   "http://127.0.0.1:8774",
			"keystone":  "http://127.0.0.1:5000",
		},
	)
}

func TestClusterWithDeploymentNetworkAsControlDataNet(t *testing.T) {
	runClusterTest(
		t,
		"./test_data/expected_same_mgmt_ctrldata_net_instances.yml",
		"",
		pongo2.Context{
			"MGMT_INT_IP":     "127.0.0.1",
			"CONTROL_NODES":   "127.0.0.1",
			"OPENSTACK_NODES": "127.0.0.1",
		},
		map[string]string{
			"config":    "http://127.0.0.1:8082",
			"nodejs":    "https://127.0.0.1:8143",
			"telemetry": "http://127.0.0.1:8081",
			"baremetal": "http://127.0.0.1:6385",
			"swift":     "http://127.0.0.1:8080",
			"glance":    "http://127.0.0.1:9292",
			"compute":   "http://127.0.0.1:8774",
			"keystone":  "http://127.0.0.1:5000",
		},
	)
}

func TestClusterWithSeperateDeploymentAndControlDataNet(t *testing.T) {
	runClusterTest(
		t,
		"./test_data/expected_multi_interface_instances.yml",
		"",
		pongo2.Context{
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
		},
		map[string]string{
			"config":    "https://10.1.1.100:8082",
			"nodejs":    "https://10.1.1.100:8143",
			"telemetry": "http://10.1.1.100:8081",
			"baremetal": "http://127.0.0.1:6385",
			"swift":     "http://127.0.0.1:8080",
			"glance":    "http://127.0.0.1:9292",
			"compute":   "http://127.0.0.1:8774",
			"keystone":  "http://127.0.0.1:5000",
		},
	)
}

func TestCredAllInOneClusterTest(t *testing.T) {
	runClusterTest(
		t,
		"./test_data/expected_creds_all_in_one_instances.yml",
		"",
		pongo2.Context{
			"CUSTOMIZE":       true,
			"CREDS":           true,
			"TYPE":            "",
			"MGMT_INT_IP":     "127.0.0.1",
			"CONTROL_NODES":   "127.0.0.1",
			"OPENSTACK_NODES": "",
			"ENABLE_ZTP":      true,
			"WEBUI_NODES":     "10.1.1.35",
		},
		map[string]string{
			"config":    "http://127.0.0.1:8082",
			"nodejs":    "https://10.1.1.35:8143",
			"telemetry": "http://127.0.0.1:8081",
			"baremetal": "http://127.0.0.1:6385",
			"swift":     "http://127.0.0.1:8080",
			"glance":    "http://127.0.0.1:9292",
			"compute":   "http://127.0.0.1:8774",
			"keystone":  "http://127.0.0.1:5000",
		},
	)
}

// nolint: gocyclo
func runKubernetesClusterTest(
	t *testing.T,
	expectedOutput string,
	pContext map[string]interface{},
	expectedEndpoints map[string]string,
) {
	closeKs := startMockKeystone()
	defer closeKs()

	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneKubernetesClusterTemplatePath, pContext)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()
	// create cluster config
	config := createClusterConfig(t)
	// create cluster
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	assertYamlFilesAreEqual(t, expectedOutput, pathToClusterFile(defaultInstanceFile),
		"Instance file created during cluster create is not as expected")
	assert.True(t, compareFiles(t, expectedK8sCreatePlaybook, pathToClusterFile(executedAnsiblePlaybook)),
		"Expected list of create playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// update cluster
	config.Action = updateAction
	// remove instances.yml to trriger cluster update

	removeFile(t, pathToClusterFile(defaultInstanceFile))
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	assertYamlFilesAreEqual(t, expectedOutput, pathToClusterFile(defaultInstanceFile),
		"Instance file created during cluster update is not as expected")
	assert.True(t, compareFiles(t, expectedK8sUpdatePlaybook, pathToClusterFile(executedAnsiblePlaybook)),
		"Expected list of update playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// UPGRADE test
	runClusterActionTest(t, ts, config,
		upgradeProvisioningAction, expectedOutput, "",
		upgradePlaybooksKubernetes, expectedEndpoints)

	// ADD_COMPUTE  test
	runClusterActionTest(t, ts, config,
		addComputeProvisioningAction, expectedOutput, "",
		addComputePlaybooksKubernetes, expectedEndpoints)

	// DELETE_COMPUTE  test
	runClusterActionTest(t, ts, config,
		deleteComputeProvisioningAction, expectedOutput, "",
		deleteComputePlaybooksKubernetes, expectedEndpoints)

	// IMPORT test (expected to create endpoints withtout triggering playbooks)
	runClusterActionTest(t, ts, config,
		importProvisioningAction, expectedOutput, "",
		"", expectedEndpoints)

	// delete cluster
	config.Action = deleteAction
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}

func TestKubernetesCluster(t *testing.T) {
	runKubernetesClusterTest(
		t,
		"./test_data/expected_all_in_one_kubernetes_instances.yml",
		pongo2.Context{
			"TYPE":          "kernel",
			"MGMT_INT_IP":   "127.0.0.1",
			"CONTROL_NODES": "",
		},
		map[string]string{
			"config":    "http://127.0.0.1:8082",
			"nodejs":    "https://127.0.0.1:8143",
			"telemetry": "http://127.0.0.1:8081",
		},
	)
}

//vcenter
// nolint: gocyclo
func runvcenterClusterTest(
	t *testing.T,
	expectedOutput, expectedVcenterVars string,
	pContext map[string]interface{},
	expectedEndpoints map[string]string,
) {
	closeKs := startMockKeystone()
	defer closeKs()

	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOnevcenterClusterTemplatePath, pContext)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()
	// create cluster config
	config := createClusterConfig(t)
	// create cluster
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	assertYamlFilesAreEqual(t, expectedOutput, pathToClusterFile(defaultInstanceFile),
		"Instance file created during cluster create is not as expected")
	assert.True(t, compareFiles(t, expectedVcenterVars, pathToClusterFile(defaultVcenterFile)),
		"Vcenter_vars file created during cluster create is not as expected")
	assert.True(t, compareFiles(t, expectedVcenterCreatePlaybook, pathToClusterFile(executedAnsiblePlaybook)),
		"Expected list of create playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// update cluster
	config.Action = updateAction
	// remove instances.yml to trriger cluster update
	removeFile(t, pathToClusterFile(defaultInstanceFile))
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	assertYamlFilesAreEqual(t, expectedOutput, pathToClusterFile(defaultInstanceFile),
		"Instance file created during cluster update is not as expected")
	assert.True(t, compareFiles(t, expectedVcenterUpdatePlaybook, pathToClusterFile(executedAnsiblePlaybook)),
		"Expected list of update playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	assert.NoError(t, err)

	// UPGRADE test
	runClusterActionTest(t, ts, config,
		upgradeProvisioningAction, expectedOutput, "",
		upgradePlaybooksvcenter, expectedEndpoints)

	// IMPORT test (expected to create endpoints withtout triggering playbooks)
	runClusterActionTest(t, ts, config,
		importProvisioningAction, expectedOutput, "",
		"", expectedEndpoints)

	// delete cluster
	config.Action = deleteAction
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	deployTestCluster(t, config, config.Action)
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}

func TestVcenterCluster(t *testing.T) {
	runvcenterClusterTest(
		t,
		"./test_data/expected_all_in_one_vcenter_instances.yml",
		"./test_data/expected_all_in_one_vcenter_vars.yml",
		pongo2.Context{
			"TYPE":             "ESXI",
			"ESXI":             "10.84.16.11",
			"MGMT_INT_IP":      "127.0.0.1",
			"CONTROLLER_NODES": "127.0.0.1",
		},
		map[string]string{
			"config":    "http://127.0.0.1:8082",
			"nodejs":    "https://127.0.0.1:8143",
			"telemetry": "http://127.0.0.1:8081",
		},
	)
}

func TestWindowsCompute(t *testing.T) {
	ts, err := integration.LoadTest("./test_data/test_windows_compute.yml", nil)
	assert.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, ts, server)
}

// nolint: gocyclo
func runMCClusterTest(t *testing.T, pContext map[string]interface{}) {
	closeKs := startMockKeystone()
	defer closeKs()

	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneMCClusterTemplatePath, pContext)
	assert.NoError(t, err, "failed to load mc cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	// create cluster config
	config := createClusterConfig(t)

	cloudFileCleanup := createDummyCloudFiles(t)
	defer cloudFileCleanup()
	// create cluster
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))

	clusterDeployer, err := NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.Error(t, err,
		"mc deployment should fail because cloud provisioning has failed")

	ts, err = integration.LoadTest(allInOneMCCloudUpdateTemplatePath, pContext)
	assert.NoError(t, err, "failed to load mc pvt cloud update test data")
	_ = integration.RunDirtyTestScenario(t, ts, server)

	// now get cluster data again
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(create) cluster")

	assertMCFiles(t, config.Action)

	// update cluster
	config.Action = updateAction
	//cleanup all the files
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	removeFile(t, pathToClusterFileMC(defaultTopologyFile))
	removeFile(t, pathToClusterFileMC(defaultSecretFile))
	removeFile(t, pathToClusterFileMC(defaultContrailCommonFile))
	removeFile(t, pathToClusterFileMC(defaultGatewayCommonFile))

	ts, err = integration.LoadTest(allInOneMCClusterUpdateTemplatePath, pContext)
	assert.NoError(t, err, "failed to load mc cluster test data")
	_ = integration.RunDirtyTestScenario(t, ts, server)
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to update cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(update) cluster")

	assertMCFiles(t, config.Action)

	// delete cloud secanrio
	//cleanup all the files
	removeFile(t, pathToClusterFile(executedAnsiblePlaybook))
	//var deleteTestScenario integration.TestScenario
	ts, err = integration.LoadTest(allInOneMCClusterDeleteTemplatePath, pContext)
	//err = integration.LoadTestScenario(&deleteTestScenario, allInOneMCClusterDeleteTemplatePath, pContext)
	assert.NoError(t, err, "failed to load mc cluster test data")
	_ = integration.RunDirtyTestScenario(t, ts, server)
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to delete cloud")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cloud")
	assert.True(t, compareFiles(t, expectedMCDeletePlaybook, pathToClusterFile(executedAnsiblePlaybook)),
		"Expected list of delete playbooks are not executed",
	)
	// make sure cluster is removed
	assert.True(t, verifyMCDeleted(clusterDeployer.APIServer), "MC folder is not deleted during cluster delete")

	// delete cluster itself
	config.Action = deleteAction
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to delete cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cluster")
}

func assertMCFiles(t *testing.T, action string) {
	assert.Truef(t, compareFiles(t, expectedMCClusterTopology, pathToClusterFileMC(defaultTopologyFile)),
		"Topology file created during cluster %s is not as expected", action)
	assert.Truef(t, compareFiles(t, expectedMCClusterSecret, pathToClusterFileMC(defaultSecretFile)),
		"Secret file created during cluster %s is not as expected", action)
	assert.Truef(t, compareFiles(t, expectedContrailCommon, pathToClusterFileMC(defaultContrailCommonFile)),
		"Contrail common file created during cluster %s is not as expected", action)
	assert.Truef(t, compareFiles(t, expectedGatewayCommon, pathToClusterFileMC(defaultGatewayCommonFile)),
		"Gateway common file created during cluster %s is not as expected", action)

	var playbook string
	switch action {
	case createAction:
		playbook = expectedMCCreatePlaybook
	case updateAction:
		playbook = expectedMCUpdatePlaybook
	default:
		assert.Failf(t, "Failed to assert MC test files", "Unexpected action %s", action)
	}

	assert.Truef(t, compareFiles(t, playbook, pathToClusterFile(executedAnsiblePlaybook)),
		"Expected list of playbooks are not executed during %s", action)
}

func TestMCCluster(t *testing.T) {
	runMCClusterTest(t, pongo2.Context{
		"CONTROL_NODES": "",
	})
}
