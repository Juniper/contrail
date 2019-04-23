package cluster

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
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
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
	allInOnevcenterClusterTemplatePath    = "./test_data/test_all_in_one_vcenter_server.tmpl"
	upgradePlaybooksvcenter               = "./test_data/expected_ansible_upgrade_playbook_vcenter.yml"
	allInOneMCClusterTemplatePath         = "./test_data/test_mc_cluster.tmpl"
	allInOneMCClusterUpdateTemplatePath   = "./test_data/test_mc_update_cluster.tmpl"
	allInOneMCClusterDeleteTemplatePath   = "./test_data/test_mc_delete_cluster.tmpl"
	allInOneMCCloudUpdateTemplatePath     = "./test_data/test_mc_update_pvt_cloud.tmpl"
	allInOneClusterAppformixTemplatePath  = "./test_data/test_all_in_one_with_appformix.tmpl"
	clusterID                             = "test_cluster_uuid"

	expectedMCClusterTopology   = "./test_data/expected_mc_cluster_topology.yml"
	expectedMCClusterSecret     = "./test_data/expected_mc_cluster_secret.yml"
	expectedContrailCommon      = "./test_data/expected_mc_contrail_common.yml"
	expectedGatewayCommon       = "./test_data/expected_mc_gateway_common.yml"
	expectedTORCommon           = "./test_data/expected_mc_tor_common.yml"
	expectedTORScript           = "./test_data/expected_tor_script.sh"
	expectedMCCreateCmdExecuted = "./test_data/expected_mc_create_cmd_executed.yml"
	expectedMCUpdateCmdExecuted = "./test_data/expected_mc_update_cmd_executed.yml"
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

func compareGeneratedSecret(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedSecretPath())
}

func compareGeneratedContrailCommon(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedContrailCommonPath())
}

func compareGeneratedGatewayCommon(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedGatewayCommonPath())
}

func compareGeneratedTORScript(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedTORScriptPath())
}

func compareGeneratedTORCommon(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedTORCommon())
}

func compareGeneratedVcentervars(t *testing.T, expected string) bool {
	return compareFiles(t, expected, generatedVcenterVarsPath())
}

func verifyCommandsExecuted(t *testing.T, expected string) bool {
	return compareFiles(t, expected, executedMCCommandPath())
}

func verifyPlaybooks(t *testing.T, expected string) bool {
	return compareFiles(t, expected, executedPlaybooksPath())
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

func generatedTORScriptPath() string {
	return workRoot + "/" + clusterID + "/" + mcWorkDir + "/" + defaultMCTORScriptFile
}

func generatedTORCommon() string {
	return workRoot + "/" + clusterID + "/" + mcWorkDir + "/" + defaultTORCommonFile
}

func executedPlaybooksPath() string {
	return workRoot + "/" + clusterID + "/executed_ansible_playbook.yml"
}

func executedMCCommandPath() string {
	return workRoot + "/" + clusterID + "/" + mcWorkDir + "/executed_cmd.yml"
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
	// create public cloud secret.yml
	secretData, err := fileutil.GetContent("./test_data/public_cloud_secret.yml")
	if err != nil {
		assert.NoErrorf(t, err, "Unable to read file: %s", "./test_data/public_cloud_secret.yml")
	}
	err = fileutil.WriteToFile("/var/tmp/cloud/public_cloud_uuid/secret.yml", secretData, defaultFilePermRWOnly)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to write file: %s", "/var/tmp/cloud/config/public_cloud_uuid/secret.yml")
	}

	return func() {
		// best effort method of deleting all the files
		// nolint: errcheck
		_ = os.Remove("/var/tmp/cloud/config/public_cloud_uuid/topology.yml")
		// nolint: errcheck
		_ = os.Remove("/var/tmp/cloud/config/public_cloud_uuid/secret.yml")
		// nolint: errcheck
		_ = os.Remove("/var/tmp/cloud/config/pvt_cloud_uuid/topology.yml")
	}
}

// nolint: gocyclo
func runClusterActionTest(t *testing.T, testScenario integration.TestScenario,
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
		// remove instances.yml to mock trriger cluster update
		err = os.Remove(generatedInstancesPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete instances.yml")
		}
		if expectedInventory != "" {
			expectedPlaybooks = addComputeEncryptPlaybooks
		}
	case deleteComputeProvisioningAction:
		// remove instances.yml to mock trriger cluster update
		err = os.Remove(generatedInstancesPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete instances.yml")
		}
	case addCSNProvisioningAction:
		// remove instances.yml to mock trriger cluster update
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
	for _, client := range testScenario.Clients {
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
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
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
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, allInOneClusterTemplatePath, pContext)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()
	// create cluster config
	s := &client.HTTP{
		Endpoint: server.URL(),
		InSecure: true,
		AuthURL:  server.URL() + "/keystone/v3",
		ID:       "alice",
		Password: "alice_password",
		Scope: keystone.NewScope(
			"default", "default", "admin", "admin"),
	}
	s.Init()
	_, err = s.Login(context.Background())
	assert.NoError(t, err, "failed to login")
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
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// update cluster
	config.Action = updateAction
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
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// UPGRADE test
	runClusterActionTest(t, testScenario, config,
		upgradeProvisioningAction, expectedInstance, expectedInventory,
		upgradePlaybooks, expectedEndpoints)

	// ADD_COMPUTE  test
	runClusterActionTest(t, testScenario, config,
		addComputeProvisioningAction, expectedInstance, expectedInventory,
		addComputePlaybooks, expectedEndpoints)

	// DELETE_COMPUTE  test
	runClusterActionTest(t, testScenario, config,
		deleteComputeProvisioningAction, expectedInstance, "",
		deleteComputePlaybooks, expectedEndpoints)

	// ADD_CSN  test
	runClusterActionTest(t, testScenario, config,
		addCSNProvisioningAction, expectedInstance, expectedInventory,
		addCSNPlaybooks, expectedEndpoints)

	// IMPORT test (expected to create endpoints without triggering playbooks)
	runClusterActionTest(t, testScenario, config,
		importProvisioningAction, expectedInstance, "", "", expectedEndpoints)

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

func getConfig(apiserver *client.HTTP, klusterID string) *Config {
	if klusterID == "" {
		klusterID = clusterID
	}

	return &Config{
		APIServer:    apiserver,
		ClusterID:    klusterID,
		Action:       "create",
		LogLevel:     "debug",
		TemplateRoot: "templates/",
		WorkRoot:     workRoot,
		Test:         true,
		LogFile:      workRoot + "/deploy.log",
	}
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
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, allInOneClusterAppformixTemplatePath, pContext)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()
	// create cluster config
	s := &client.HTTP{
		Endpoint: server.URL(),
		InSecure: true,
		AuthURL:  server.URL() + "/keystone/v3",
		ID:       "alice",
		Password: "alice_password",
		Scope: keystone.NewScope(
			"default", "default", "admin", "admin"),
	}
	s.Init()
	_, err = s.Login(context.Background())
	assert.NoError(t, err, "failed to login")
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
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
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
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// UPGRADE test
	runClusterActionTest(t, testScenario, config,
		upgradeProvisioningAction, expectedInstance, expectedInventory,
		upgradeAppformixPlaybooks, expectedEndpoints)

	// ADD_COMPUTE  test
	runClusterActionTest(t, testScenario, config,
		addComputeProvisioningAction, expectedInstance, expectedInventory,
		addAppformixComputePlaybooks, expectedEndpoints)

	// DELETE_COMPUTE  test
	runClusterActionTest(t, testScenario, config,
		deleteComputeProvisioningAction, expectedInstance, "",
		deleteAppformixComputePlaybooks, expectedEndpoints)

	// ADD_CSN  test
	runClusterActionTest(t, testScenario, config,
		addCSNProvisioningAction, expectedInstance, expectedInventory,
		addAppformixCSNPlaybooks, expectedEndpoints)

	// IMPORT test (expected to create endpoints without triggering playbooks)
	runClusterActionTest(t, testScenario, config,
		importProvisioningAction, expectedInstance, "", "", expectedEndpoints)

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

	runAppformixClusterTest(t, expectedInstances, "", context, expectedEndpoints)
}

func TestXflow(t *testing.T) {
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(
		&testScenario,
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

	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	config := getConfig(testScenario.Clients["default"], "")

	contrailDeployer, ok := getClusterDeployer(t, config).(*contrailAnsibleDeployer)
	assert.True(t, ok, "unable to cast deployer to contrailAnsibleDeployer")

	err = contrailDeployer.createInventory()
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
		"telemetry": "http://10.1.1.100:8081",
		"baremetal": "http://127.0.0.1:6385",
		"swift":     "http://127.0.0.1:8080",
		"glance":    "http://127.0.0.1:9292",
		"compute":   "http://127.0.0.1:8774",
		"keystone":  "http://127.0.0.1:5000",
	}

	runClusterTest(t, "./test_data/expected_multi_interface_instances.yml", "", pContext, expectedEndpoints)
}

func TestCredAllInOneClusterTest(t *testing.T) {
	pContext := pongo2.Context{
		"CUSTOMIZE":       true,
		"CREDS":           true,
		"TYPE":            "",
		"MGMT_INT_IP":     "127.0.0.1",
		"CONTROL_NODES":   "",
		"OPENSTACK_NODES": "",
		"ENABLE_ZTP":      true,
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
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, allInOneKubernetesClusterTemplatePath, pContext)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()
	// create cluster config
	s := &client.HTTP{
		Endpoint: server.URL(),
		InSecure: true,
		AuthURL:  server.URL() + "/keystone/v3",
		ID:       "alice",
		Password: "alice_password",
		Scope: keystone.NewScope(
			"default", "default", "admin", "admin"),
	}
	s.Init()
	_, err = s.Login(context.Background())
	assert.NoError(t, err, "failed to login")
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
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// update cluster
	config.Action = updateAction
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
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// UPGRADE test
	runClusterActionTest(t, testScenario, config,
		upgradeProvisioningAction, expectedOutput, "",
		upgradePlaybooksKubernetes, expectedEndpoints)

	// IMPORT test (expected to create endpoints withtout triggering playbooks)
	runClusterActionTest(t, testScenario, config,
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
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, allInOnevcenterClusterTemplatePath, pContext)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()
	// create cluster config
	s := &client.HTTP{
		Endpoint: server.URL(),
		InSecure: true,
		AuthURL:  server.URL() + "/keystone/v3",
		ID:       "alice",
		Password: "alice_password",
		Scope: keystone.NewScope(
			"default", "default", "admin", "admin"),
	}
	s.Init()
	_, err = s.Login(context.Background())
	assert.NoError(t, err, "failed to login")
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
	err = verifyEndpoints(t, &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	// UPGRADE test
	runClusterActionTest(t, testScenario, config,
		"UPGRADE", expectedOutput, "",
		upgradePlaybooksvcenter, expectedEndpoints)

	// IMPORT test (expected to create endpoints withtout triggering playbooks)
	runClusterActionTest(t, testScenario, config,
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
	testScenario, err := integration.LoadTest("./test_data/test_windows_compute.yml", nil)
	assert.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, testScenario, server)
}

// nolint: gocyclo
func runMCClusterTest(t *testing.T, pContext map[string]interface{},
	expectedEndpoints map[string]string) {

	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPrivate.Close()
	// Create the cluster and related objects
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, allInOneMCClusterTemplatePath, pContext)
	assert.NoError(t, err, "failed to load mc cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()
	// create cluster config
	s := &client.HTTP{
		Endpoint: server.URL(),
		InSecure: true,
		AuthURL:  server.URL() + "/keystone/v3",
		ID:       "alice",
		Password: "alice_password",
		Scope: keystone.NewScope(
			"default", "default", "admin", "admin"),
	}
	s.Init()
	_, err = s.Login(context.Background())
	assert.NoError(t, err, "failed to login")
	// create cluster config
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

	var updateCloudTestScenario integration.TestScenario
	err = integration.LoadTestScenario(&updateCloudTestScenario, allInOneMCCloudUpdateTemplatePath, pContext)
	assert.NoError(t, err, "failed to load mc pvt cloud update test data")
	_ = integration.RunDirtyTestScenario(t, &updateCloudTestScenario, server)

	// now get cluster data again
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(create) cluster")

	assert.True(t, compareGeneratedTopology(t, expectedMCClusterTopology),
		"Topolgy file created during cluster create is not as expected")
	assert.True(t, compareGeneratedSecret(t, expectedMCClusterSecret),
		"Secret file created during cluster create is not as expected")
	assert.True(t, compareGeneratedContrailCommon(t, expectedContrailCommon),
		"Contrail common file created during cluster create is not as expected")
	assert.True(t, compareGeneratedGatewayCommon(t, expectedGatewayCommon),
		"Gateway common file created during cluster create is not as expected")
	assert.True(t, compareGeneratedTORCommon(t, expectedTORCommon),
		"TOR common file created during cluster create is not as expected")
	assert.True(t, compareGeneratedTORScript(t, expectedTORScript),
		"Gateway common file created during cluster create is not as expected")
	assert.True(t, verifyCommandsExecuted(t, expectedMCCreateCmdExecuted),
		"commands executed during cluster create is not as expected")
	assert.True(t, verifyPlaybooks(t, "./test_data/expected_ansible_create_mc_playbook.yml"),
		"Expected list of playbooks are not executed during create")

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
	if _, err = os.Stat(generatedTORScriptPath()); err == nil {
		// cleanup executed playbook file
		err = os.Remove(generatedTORScriptPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete generated TOR script file")
		}
	}

	var updateTestScenario integration.TestScenario
	err = integration.LoadTestScenario(&updateTestScenario, allInOneMCClusterUpdateTemplatePath, pContext)
	assert.NoError(t, err, "failed to load mc cluster test data")
	_ = integration.RunDirtyTestScenario(t, &updateTestScenario, server)
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to update cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(update) cluster")

	assert.True(t, compareGeneratedTopology(t, expectedMCClusterTopology),
		"Topolgy file created during cluster update is not as expected")
	assert.True(t, compareGeneratedSecret(t, expectedMCClusterSecret),
		"Secret file created during cluster update is not as expected")
	assert.True(t, compareGeneratedContrailCommon(t, expectedContrailCommon),
		"Contrail common file created during cluster update is not as expected")
	assert.True(t, compareGeneratedGatewayCommon(t, expectedGatewayCommon),
		"Gateway common file created during cluster update is not as expected")
	assert.True(t, compareGeneratedTORScript(t, expectedTORScript),
		"Gateway common file created during cluster create is not as expected")
	assert.True(t, compareGeneratedTORCommon(t, expectedTORCommon),
		"TOR common file created during cluster create is not as expected")
	assert.True(t, verifyCommandsExecuted(t, expectedMCUpdateCmdExecuted),
		"commands executed during cluster update is not as expected")
	assert.True(t, verifyPlaybooks(t, "./test_data/expected_ansible_update_mc_playbook.yml"),
		"Expected list of playbooks are not executed during update")

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
	var deleteTestScenario integration.TestScenario
	err = integration.LoadTestScenario(&deleteTestScenario, allInOneMCClusterDeleteTemplatePath, pContext)
	assert.NoError(t, err, "failed to load mc cluster test data")
	_ = integration.RunDirtyTestScenario(t, &deleteTestScenario, server)
	clusterDeployer, err = NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to delete cloud")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cloud")
	assert.True(t, verifyPlaybooks(t, "./test_data/expected_ansible_delete_mc_playbook.yml"),
		"Expected list of delete playbooks are not executed")
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

func TestMCCluster(t *testing.T) {
	context := pongo2.Context{
		"CONTROL_NODES": "",
	}
	expectedEndpoints := map[string]string{
		"config":    "http://1.1.1.1:8082",
		"nodejs":    "https://1.1.1.1:8143",
		"telemetry": "http://1.1.1.1:8081",
	}
	runMCClusterTest(t, context, expectedEndpoints)
}
