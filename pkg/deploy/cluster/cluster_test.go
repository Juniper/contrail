package cluster

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	yaml "gopkg.in/yaml.v2"
)

const (
	defaultAdminUser                      = "admin"
	defaultAdminPassword                  = "contrail123"
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
	allInOnevcenterClusterTemplatePath    = "./test_data/test_all_in_one_vcenter_server.tmpl"
	upgradePlaybooksvcenter               = "./test_data/expected_ansible_upgrade_playbook_vcenter.yml"
	allInOneMCClusterTemplatePath         = "./test_data/test_mc_cluster.tmpl"
	allInOneMCClusterUpdateTemplatePath   = "./test_data/test_mc_update_cluster.tmpl"
	allInOneMCClusterDeleteTemplatePath   = "./test_data/test_mc_delete_cluster.tmpl"
	allInOneMCCloudUpdateTemplatePath     = "./test_data/test_mc_update_pvt_cloud.tmpl"
	allInOneClusterAppformixTemplatePath  = "./test_data/test_all_in_one_with_appformix.tmpl"
	clusterID                             = "test_cluster_uuid"

	expectedMCClusterTopology   = "./test_data/expected_mc_cluster_topology.yml"
	expectedContrailCommon      = "./test_data/expected_mc_contrail_common.yml"
	expectedGatewayCommon       = "./test_data/expected_mc_gateway_common.yml"
	expectedMCCreateCmdExecuted = "./test_data/expected_mc_create_cmd_executed.yml"
	expectedMCUpdateCmdExecuted = "./test_data/expected_mc_update_cmd_executed.yml"
	expectedMCDeleteCmdExecuted = "./test_data/expected_mc_delete_cmd_executed.yml"
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

func verifyClusterDeleted() bool {
	// Make sure working dir is deleted
	if _, err := os.Stat(workRoot + "/" + clusterID); err == nil {
		// working dir not deleted
		return false
	}
	return true
}

func verifyMCDeleted(httpClient *client.HTTP) bool {
	// Make sure mc working dir is deleted
	if _, err := os.Stat(workRoot + "/" + clusterID + "/" + mcWorkDir); err == nil {
		// mc working dir not deleted
		return false
	}
	clusterObjResp, err := httpClient.GetContrailCluster(context.Background(),
		&services.GetContrailClusterRequest{
			ID: clusterID,
		},
	)
	if err != nil {
		return false
	}

	if clusterObjResp.ContrailCluster.CloudRefs == nil {
		return true
	}
	return false
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
	t *testing.T,
	expectedContainedYamlFile, actualYamlFile string,
	msgAndArgs ...interface{}) bool {

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

func assertGeneratedInstancesContainExpected(t *testing.T, expected string, msgAndArgs ...interface{}) {
	assertYamlFileContainsOther(t, expected, generatedInstancesPath(), msgAndArgs...)
}

func assertGeneratedInstancesAreExpected(t *testing.T, expected string, msgAndArgs ...interface{}) {
	assertYamlFilesAreEqual(t, expected, generatedInstancesPath(), msgAndArgs...)
}

func compareGeneratedInventory(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedInventoryPath())
}

func compareGeneratedTopology(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedTopologyPath())
}

func compareGeneratedContrailCommon(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedContrailCommonPath())
}

func compareGeneratedGatewayCommon(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedGatewayCommonPath())
}

func compareGeneratedVcentervars(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedVcenterVarsPath())
}

func verifyPlaybooks(t *testing.T, expected string) bool {
	return compareFiles(t, expected, executedPlaybooksPath())
}

func verifyCommandsExecuted(t *testing.T, expected string) bool {
	return compareFiles(t, expected, executedMCCommandPath())
}

func generatedInstancesPath() string {
	return workRoot + "/" + clusterID + "/instances.yml"
}

func generatedInventoryPath() string {
	return workRoot + "/" + clusterID + "/inventory.yml"
}

func generatedVcenterVarsPath() string {
	return workRoot + "/" + clusterID + "/vcenter_vars.yml"
}

func generatedSecretPath() string {
	return workRoot + "/" + clusterID + "/" + mcWorkDir + "/" + defaultSecretFile
}

func generatedTopologyPath() string {
	return workRoot + "/" + clusterID + "/" + mcWorkDir + "/" + defaultTopologyFile
}

func generatedContrailCommonPath() string {
	return workRoot + "/" + clusterID + "/" + mcWorkDir + "/" + defaultContrailCommonFile
}

func generatedGatewayCommonPath() string {
	return workRoot + "/" + clusterID + "/" + mcWorkDir + "/" + defaultGatewayCommonFile
}

func executedPlaybooksPath() string {
	return workRoot + "/" + clusterID + "/executed_ansible_playbook.yml"
}

func executedMCCommandPath() string {
	return workRoot + "/" + clusterID + "/executed_cmd.yml"
}

func createDummyCloudFiles(t *testing.T) func() {

	// create public cloud topology.yaml
	publicTopoData, err := fileutil.GetContent("file://./test_data/public_cloud_topology.yml")
	if err != nil {
		assert.NoErrorf(t, err, "Unable to read file: %s", "./test_data/public_cloud_topology.yml")
	}
	err = fileutil.WriteToFile("/var/tmp/cloud/public_cloud_uuid/topology.yml", publicTopoData, defaultFilePermRWOnly)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to write file: %s", "/var/tmp/cloud/config/public_cloud_uuid/topology.yml")
	}
	// create pvt cloud topology.yml
	pvtTopoData, err := fileutil.GetContent("./test_data/pvt_cloud_topology.yml")
	if err != nil {
		assert.NoErrorf(t, err, "Unable to read file: %s", "./test_data/pvt_cloud_topology.yml")
	}
	err = fileutil.WriteToFile("/var/tmp/cloud/pvt_cloud_uuid/topology.yml", pvtTopoData, defaultFilePermRWOnly)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to write file: %s", "/var/tmp/cloud/config/pvt_cloud_uuid/topology.yml")
	}

	return func() {
		// best effort method of deleting all the files
		// nolint: errcheck
		_ = os.Remove("/var/tmp/cloud/config/public_cloud_uuid/topology.yml")
		// nolint: errcheck
		_ = os.Remove("/var/tmp/cloud/config/pvt_cloud_uuid/topology.yml")
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
func runClusterActionTest(t *testing.T, ts *integration.TestScenario,
	config *Config, action, expectedInstance, expectedInventory string,
	expectedPlaybooks string, expectedEndpoints map[string]string) {
	// set action field in the contrail-cluster resource
	var err error
	var data interface{}
	cluster := map[string]interface{}{"uuid": clusterID,
		"provisioning_action": action,
	}
	config.Action = updateAction
	switch action {
	case upgradeProvisioningAction:
		cluster["provisioning_state"] = "NOSTATE"
		if expectedInventory != "" {
			expectedPlaybooks = upgradeEncryptPlaybooks
		}
	case addComputeProvisioningAction:
		// remove instances.yml to mock trigger cluster update
		err = os.Remove(generatedInstancesPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete instances.yml")
		}
		if expectedInventory != "" {
			expectedPlaybooks = addComputeEncryptPlaybooks
		}
	case deleteComputeProvisioningAction, addCSNProvisioningAction, addCVFMProvisioningAction, destroyAction:
		// remove instances.yml to mock trigger cluster update
		err = os.Remove(generatedInstancesPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete instances.yml")
		}
	case importProvisioningAction:
		// cleanup instances.yml
		err = os.Remove(generatedInstancesPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete instances.yml")
		}
		config.Action = createAction
		cluster["provisioning_action"] = ""
	}
	data = map[string]interface{}{"contrail-cluster": cluster}
	for _, client := range ts.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/contrail-cluster/%s", clusterID)
		_, err = client.Update(context.Background(), url, &data, &response)
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
	clusterDeployer, err := NewCluster(config)
	assert.NoErrorf(t, err, "failed to create cluster manager to %s cluster", config.Action)
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoErrorf(t, err, "failed to manage(%s) cluster", action)
	if expectedInstance != "" {
		assertGeneratedInstancesAreExpected(t, expectedInstance,
			"Instance file created during cluster %s is not as expected", action)
	}
	if expectedPlaybooks != "" {
		assert.True(t, verifyPlaybooks(t, expectedPlaybooks),
			fmt.Sprintf("Expected list of %s playbooks are not executed", action))
	}
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}
}

// nolint: gocyclo
func runClusterTest(t *testing.T, expectedInstance, expectedInventory string,
	pContext map[string]interface{}, expectedEndpoints map[string]string) {
	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPrivate.Close()

	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneClusterTemplatePath, pContext)
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
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup old executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterDeployer, err := NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(create) cluster")
	assertGeneratedInstancesAreExpected(t, expectedInstance,
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
	server.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, ts, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// update cluster
	config.Action = updateAction
	// remove instances.yml to trigger cluster update
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
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to update cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(update) cluster")
	assertGeneratedInstancesAreExpected(t, expectedInstance,
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
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

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
		importProvisioningAction, expectedInstance, "", "", expectedEndpoints)

	// ADD_CVFM test
	if pContext["VFABRIC_MANAGER"] == true {
		runClusterActionTest(t, ts, config,
			addCVFMProvisioningAction, expectedInstance, expectedInventory,
			updatePlaybooks, expectedEndpoints)
	}

	// delete cluster
	config.Action = deleteAction
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to delete cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cluster")
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}

func getClusterDeployer(t *testing.T, config *Config) base.Deployer {
	cluster, err := NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := cluster.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	return deployer
}

// nolint: gocyclo
func runAppformixClusterTest(t *testing.T, expectedInstance, expectedInventory string,
	pContext map[string]interface{}, expectedEndpoints map[string]string) {
	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPrivate.Close()

	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneClusterAppformixTemplatePath, pContext)
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
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup old executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterDeployer, err := NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(create) cluster")
	assertGeneratedInstancesAreExpected(t, expectedInstance,
		"Instance file created during cluster create is not as expected")
	if expectedInventory != "" {
		assert.True(t, compareGeneratedInventory(t, expectedInventory),
			"Inventory file created during cluster create is not as expected")
		assert.True(t, verifyPlaybooks(t, createEncryptPlaybooks),
			"Expected list of create playbooks are not executed")
	} else {
		assert.True(t, verifyPlaybooks(t, createAppformixPlaybooks),
			"Expected list of create playbooks are not executed")
	}
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, ts, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// update cluster
	config.Action = updateAction
	// remove instances.yml to trigger cluster update
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
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to update cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(update) cluster")
	assertGeneratedInstancesAreExpected(t, expectedInstance,
		"Instance file created during cluster update is not as expected")
	if expectedInventory != "" {
		assert.True(t, compareGeneratedInventory(t, expectedInventory),
			"Inventory file created during cluster update is not as expected")
		assert.True(t, verifyPlaybooks(t, updateEncryptPlaybooks),
			"Expected list of update playbooks are not executed")
	} else {
		assert.True(t, verifyPlaybooks(t, updateAppformixPlaybooks),
			"Expected list of update playbooks are not executed")
	}
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

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
		importProvisioningAction, expectedInstance, "", "", expectedEndpoints)

	// DESTROY test
	runClusterActionTest(t, ts, config,
		destroyAction, expectedInstance, "", destroyPlaybooks, expectedEndpoints)

	// delete cluster
	config.Action = deleteAction
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to delete cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cluster")
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}

func runAllInOneClusterTest(t *testing.T, computeType string) {
	pContext := pongo2.Context{
		"TYPE":            computeType,
		"MGMT_INT_IP":     "127.0.0.1",
		"CONTROL_NODES":   "",
		"OPENSTACK_NODES": "",
	}
	expectedEndpoints := map[string]string{
		"config":                "http://127.0.0.1:8082",
		"nodejs":                "https://127.0.0.1:8143",
		"telemetry":             "http://127.0.0.1:8081",
		"baremetal":             "http://127.0.0.1:6385",
		"swift":                 "http://127.0.0.1:8080",
		"glance":                "http://127.0.0.1:9292",
		"compute":               "http://127.0.0.1:8774",
		"keystone":              "http://127.0.0.1:5000",
		"endpoint_user_created": "http://127.0.0.1:8082",
	}
	expectedInstances := "./test_data/expected_all_in_one_instances.yml"
	switch computeType {
	case "dpdk":
		expectedInstances = "./test_data/expected_all_in_one_dpdk_instances.yml"
	case "sriov":
		expectedInstances = "./test_data/expected_all_in_one_sriov_instances.yml"
	}

	runClusterTest(t, expectedInstances, "", pContext, expectedEndpoints)
}

func runAllInOneAppformixTest(t *testing.T, computeType string) {
	context := pongo2.Context{
		"TYPE":            computeType,
		"MGMT_INT_IP":     "127.0.0.1",
		"CONTROL_NODES":   "",
		"OPENSTACK_NODES": "",
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
	runAppformixClusterTest(t, expectedInstances, "", context, expectedEndpoints)
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

	expectedInstance := "test_data/expected_xflow_instances.yaml"

	assertGeneratedInstancesContainExpected(t, expectedInstance,
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
func TestAllInOneVfabricManager(t *testing.T) {
	pContext := pongo2.Context{
		"TYPE":            "kernel",
		"MGMT_INT_IP":     "127.0.0.1",
		"CONTROL_NODES":   "",
		"OPENSTACK_NODES": "",
		"VFABRIC_MANAGER": true,
	}
	expectedEndpoints := map[string]string{
		"config":                "http://127.0.0.1:8082",
		"nodejs":                "https://127.0.0.1:8143",
		"telemetry":             "http://127.0.0.1:8081",
		"baremetal":             "http://127.0.0.1:6385",
		"swift":                 "http://127.0.0.1:8080",
		"glance":                "http://127.0.0.1:9292",
		"compute":               "http://127.0.0.1:8774",
		"keystone":              "http://127.0.0.1:5000",
		"endpoint_user_created": "http://127.0.0.1:8082",
	}
	expectedInstances := "./test_data/expected_all_in_one_vfabric_manager_instances.yml"

	runClusterTest(t, expectedInstances, "", pContext, expectedEndpoints)
}

func TestAllInOneClusterWithDatapathEncryption(t *testing.T) {
	pContext := pongo2.Context{
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
		"./test_data/expected_all_in_one_inventory.yml", pContext, expectedEndpoints)
}

func TestClusterWithDeploymentNetworkAsControlDataNet(t *testing.T) {
	pContext := pongo2.Context{
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
	runClusterTest(t, "./test_data/expected_same_mgmt_ctrldata_net_instances.yml", "", pContext, expectedEndpoints)
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

	runClusterTest(t, "./test_data/expected_multi_interface_instances.yml", "", pContext, expectedEndpoints)
}

func TestCredAllInOneClusterTest(t *testing.T) {
	pContext := pongo2.Context{
		"CUSTOMIZE":       true,
		"CREDS":           true,
		"TYPE":            "",
		"MGMT_INT_IP":     "127.0.0.1",
		"CONTROL_NODES":   "127.0.0.1",
		"OPENSTACK_NODES": "",
		"ENABLE_ZTP":      true,
		"WEBUI_NODES":     "10.1.1.35",
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

	runClusterTest(t, expectedInstances, "", pContext, expectedEndpoints)
}

// nolint: gocyclo
func runKubernetesClusterTest(t *testing.T, expectedOutput string,
	pContext map[string]interface{}, expectedEndpoints map[string]string) {
	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPrivate.Close()
	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneKubernetesClusterTemplatePath, pContext)
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
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup old executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterDeployer, err := NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(create) cluster")
	assertGeneratedInstancesAreExpected(t, expectedOutput,
		"Instance file created during cluster create is not as expected")
	assert.True(t, verifyPlaybooks(t, "./test_data/expected_ansible_create_playbook_kubernetes.yml"),
		"Expected list of create playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, ts, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// update cluster
	config.Action = updateAction
	// remove instances.yml to trigger cluster update
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
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to update cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(update) cluster")
	assertGeneratedInstancesAreExpected(t, expectedOutput,
		"Instance file created during cluster update is not as expected")
	assert.True(t, verifyPlaybooks(t, "./test_data/expected_ansible_update_playbook_kubernetes.yml"),
		"Expected list of update playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

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
		importProvisioningAction, expectedOutput, "", "", expectedEndpoints)

	// delete cluster
	config.Action = deleteAction
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to delete cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cluster")
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}
func TestKubernetesCluster(t *testing.T) {
	pContext := pongo2.Context{
		"TYPE":          "kernel",
		"MGMT_INT_IP":   "127.0.0.1",
		"CONTROL_NODES": "",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://127.0.0.1:8082",
		"nodejs":    "https://127.0.0.1:8143",
		"telemetry": "http://127.0.0.1:8081",
	}
	runKubernetesClusterTest(t, "./test_data/expected_all_in_one_kubernetes_instances.yml", pContext, expectedEndpoints)
}

//vcenter
// nolint: gocyclo
func runvcenterClusterTest(t *testing.T, expectedOutput, expectedVcentervars string,
	pContext map[string]interface{}, expectedEndpoints map[string]string) {
	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPrivate.Close()
	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOnevcenterClusterTemplatePath, pContext)
	require.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	s, err := integration.NewAdminHTTPClient(server.URL())
	assert.NoError(t, err)

	config := &Config{
		APIServer:    s,
		ClusterID:    clusterID,
		Action:       "create",
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}
	// create cluster
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup old executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterDeployer, err := NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(create) cluster")
	assertGeneratedInstancesAreExpected(t, expectedOutput,
		"Instance file created during cluster create is not as expected")
	assert.True(t, compareGeneratedVcentervars(t, expectedVcentervars),
		"Vcenter_vars file created during cluster create is not as expected")
	assert.True(t, verifyPlaybooks(t, "./test_data/expected_ansible_create_playbook_vcenter.yml"),
		"Expected list of create playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created
	err = verifyEndpoints(t, ts, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// update cluster
	config.Action = "update"
	// remove instances.yml to trigger cluster update
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
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to update cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(update) cluster")
	assertGeneratedInstancesAreExpected(t, expectedOutput,
		"Instance file created during cluster update is not as expected")
	assert.True(t, verifyPlaybooks(t, "./test_data/expected_ansible_update_playbook_vcenter.yml"),
		"Expected list of update playbooks are not executed")
	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are recreated as part of update
	err = verifyEndpoints(t, ts, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// UPGRADE test
	runClusterActionTest(t, ts, config,
		"UPGRADE", expectedOutput, "",
		upgradePlaybooksvcenter, expectedEndpoints)

	// IMPORT test (expected to create endpoints withtout triggering playbooks)
	runClusterActionTest(t, ts, config,
		"IMPORT", expectedOutput, "", "", expectedEndpoints)

	// delete cluster
	config.Action = "delete"
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to delete cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cluster")
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
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
	runvcenterClusterTest(t, "./test_data/expected_all_in_one_vcenter_instances.yml",
		"./test_data/expected_all_in_one_vcenter_vars.yml", pContext, expectedEndpoints)
}

func TestWindowsCompute(t *testing.T) {
	ts, err := integration.LoadTest("./test_data/test_windows_compute.yml", nil)
	require.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, ts, server)
}

// nolint: gocyclo
func runMCClusterTest(t *testing.T, pContext map[string]interface{}) {
	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPrivate.Close()
	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneMCClusterTemplatePath, pContext)
	require.NoError(t, err, "failed to load mc cluster test data")
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

	cloudFileCleanup := createDummyCloudFiles(t)
	defer cloudFileCleanup()
	// create cluster
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	if _, err = os.Stat(executedMCCommandPath()); err == nil {
		// cleanup old executed playbook file
		err = os.Remove(executedMCCommandPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}

	clusterDeployer, err := NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.Error(t, err,
		"mc deployment should fail because cloud provisioning has failed")

	err = isCloudSecretFilesDeleted()
	require.NoError(t, err, "failed to delete public cloud secrets during create")

	ts, err = integration.LoadTest(allInOneMCCloudUpdateTemplatePath, pContext)
	require.NoError(t, err, "failed to load mc pvt cloud update test data")
	_ = integration.RunDirtyTestScenario(t, ts, server)

	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(create) cluster")

	err = isCloudSecretFilesDeleted()
	require.NoError(t, err, "failed to delete public cloud secrets during create")

	assert.True(t, compareGeneratedTopology(t, expectedMCClusterTopology),
		"Topolgy file created during cluster create is not as expected")
	assert.True(t, compareGeneratedContrailCommon(t, expectedContrailCommon),
		"Contrail common file created during cluster create is not as expected")
	assert.True(t, compareGeneratedGatewayCommon(t, expectedGatewayCommon),
		"Gateway common file created during cluster create is not as expected")
	assert.True(t, verifyCommandsExecuted(t, expectedMCCreateCmdExecuted),
		"commands executed during cluster create are not as expected")

	// update cluster
	config.Action = updateAction
	//cleanup all the files
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	if _, err = os.Stat(executedMCCommandPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedMCCommandPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed mc command file")
		}
	}
	if _, err = os.Stat(generatedTopologyPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(generatedTopologyPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete generated topology file")
		}
	}
	if _, err = os.Stat(generatedSecretPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(generatedSecretPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete generated secret file")
		}
	}
	if _, err = os.Stat(generatedContrailCommonPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(generatedContrailCommonPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete generated contrail common file")
		}
	}
	if _, err = os.Stat(generatedGatewayCommonPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(generatedGatewayCommonPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete generated gateway common file")
		}
	}

	ts, err = integration.LoadTest(allInOneMCClusterUpdateTemplatePath, pContext)
	require.NoError(t, err, "failed to load mc cluster test data")
	_ = integration.RunDirtyTestScenario(t, ts, server)
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to update cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(update) cluster")

	err = isCloudSecretFilesDeleted()
	require.NoError(t, err, "failed to delete public cloud secrets during update")

	assert.True(t, compareGeneratedTopology(t, expectedMCClusterTopology),
		"Topolgy file created during cluster update is not as expected")
	assert.True(t, compareGeneratedContrailCommon(t, expectedContrailCommon),
		"Contrail common file created during cluster update is not as expected")
	assert.True(t, compareGeneratedGatewayCommon(t, expectedGatewayCommon),
		"Gateway common file created during cluster update is not as expected")
	assert.True(t, verifyCommandsExecuted(t, expectedMCUpdateCmdExecuted),
		"commands executed during cluster update are not as expected")

	// delete cloud secanrio
	//cleanup all the files
	if _, err = os.Stat(executedMCCommandPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedMCCommandPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed mc command file")
		}
	}
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}

	ts, err = integration.LoadTest(allInOneMCClusterDeleteTemplatePath, pContext)
	require.NoError(t, err, "failed to load mc cluster test data")
	_ = integration.RunDirtyTestScenario(t, ts, server)
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to delete cloud")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cloud")
	err = isCloudSecretFilesDeleted()
	require.NoError(t, err, "failed to delete public cloud secrets during delete")
	assert.True(t, verifyCommandsExecuted(t, expectedMCDeleteCmdExecuted),
		"commands executed during cluster delete are not as expected")
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
	err = isCloudSecretFilesDeleted()
	require.NoError(t, err, "failed to delete cloud secrets during delete")

}

func TestMCCluster(t *testing.T) {
	runMCClusterTest(t, pongo2.Context{
		"CONTROL_NODES": "",
	})
}

func isCloudSecretFilesDeleted() error {
	errstrings := []string{}
	for _, secret := range []string{
		"/var/tmp/cloud/public_cloud_uuid/secret.yml",
		generatedSecretPath(),
	} {
		_, err := os.Stat(secret)
		if err != nil {
			if !os.IsNotExist(err) {
				errstrings = append(errstrings, fmt.Sprintf(
					"Unable to verify the non existence of secret file: %s is not deleted: %v", secret, err))
			}
		} else {
			errstrings = append(errstrings, fmt.Sprintf("secret file: %s is not deleted: %v", secret, err))
		}
	}
	if len(errstrings) != 0 {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}

func TestTripleoClusterImport(t *testing.T) {
	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPrivate.Close()

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
	if _, err = os.Stat(executedPlaybooksPath()); err == nil {
		// cleanup old executed playbook file
		err = os.Remove(executedPlaybooksPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed ansible playbooks yaml")
		}
	}
	clusterDeployer, err := NewCluster(config)
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
	if err != nil {
		assert.NoError(t, err, err.Error())
	}
}
