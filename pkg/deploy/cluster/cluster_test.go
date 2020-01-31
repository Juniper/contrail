package cluster_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/ansible"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/deploy/cluster"
	"github.com/Juniper/contrail/pkg/services"
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
	allInOneMCClusterTemplatePath         = "./test_data/test_mc_cluster.tmpl"
	allInOneMCClusterUpdateTemplatePath   = "./test_data/test_mc_update_cluster.tmpl"
	allInOneMCClusterDeleteTemplatePath   = "./test_data/test_mc_delete_cluster.tmpl"
	allInOneMCCloudUpdateTemplatePath     = "./test_data/test_mc_update_pvt_cloud.tmpl"
	allInOneClusterAppformixTemplatePath  = "./test_data/test_all_in_one_with_appformix.tmpl"
	clusterID                             = "test_cluster_uuid"

	generatedInstances   = workRoot + "/" + clusterID + "/" + cluster.DefaultInstanceFile
	generatedInventory   = workRoot + "/" + clusterID + "/" + cluster.DefaultInventoryFile
	generatedVcenterVars = workRoot + "/" + clusterID + "/" + cluster.DefaultVcenterFile
	generatedSecret      = workRoot + "/" + clusterID + "/" + cluster.MCWorkDir + "/" + cluster.DefaultSecretFile
	generatedTopology    = workRoot + "/" + clusterID + "/" + cluster.MCWorkDir + "/" + cluster.DefaultTopologyFile
	executedPlaybooks    = workRoot + "/" + clusterID + "/" + "executed_ansible_playbook.yml"
	executedMCCommand    = workRoot + "/" + clusterID + "/" + "executed_cmd.yml"

	expectedMCClusterTopology   = "./test_data/expected_mc_cluster_topology.yml"
	expectedMCCreateCmdExecuted = "./test_data/expected_mc_create_cmd_executed.yml"
	expectedMCUpdateCmdExecuted = "./test_data/expected_mc_update_cmd_executed.yml"
	expectedMCDeleteCmdExecuted = "./test_data/expected_mc_delete_cmd_executed.yml"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

// TODO(dji): move to testing code and inject as dependency
type mockContainerPlayer struct {
	workingDirectory          string
	t                         *testing.T
	executionParametersFile string
}

func newMockContainerPlayer(workingDirectory string) (*mockContainerPlayer, error) {
	return &mockContainerPlayer{workingDirectory: workingDirectory}, nil
}

func newMCMockContainerPlayer(
	t *testing.T, workingDirectory, executionParametersFile string,
) (*mockContainerPlayer, error) {
	return &mockContainerPlayer{
		t: t,
		workingDirectory: workingDirectory,
		executionParametersFile: executionParametersFile,
	}, nil
}

func (m *mockContainerPlayer) Play(
	ctx context.Context,
	imageRef string,
	imageRefUsername string,
	imageRefPassword string,
	workRoot string,
	ansibleBinaryRepo string,
	ansibleArgs []string,
	keepContainerAlive bool,
) error {
	playBookIndex := len(ansibleArgs) - 1
	content, err := template.Apply("./test_data/test_ansible_playbook.tmpl", pongo2.Context{
		"playBook":    ansibleArgs[playBookIndex],
		"ansibleArgs": strings.Join(ansibleArgs[:playBookIndex], " "),
	})
	if err != nil {
		return err
	}

	return fileutil.AppendToFile(
		filepath.Join(m.workingDirectory, "executed_ansible_playbook.yml"),
		content,
		0600,
	)
}

func (m *mockContainerPlayer) StartExecuteAndRemove(ctx context.Context, cp ansible.ContainerExecParameters) error {
	data, err := yaml.Marshal(cp)
	assert.NoError(m.t, err, "cannot marshal provided container parameters")
	err = fileutil.WriteToFile(m.executionParametersFile, data, 0644)
	assert.NoErrorf(m.t, err, "cannot save provided container parameters to file %s", m.executionParametersFile)
	return nil
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
	_, err := os.Stat(workRoot + "/" + clusterID)
	return os.IsNotExist(err)
}

func verifyMCDeleted(httpClient *client.HTTP) bool {
	// Make sure mc working dir is deleted
	if _, err := os.Stat(workRoot + "/" + clusterID + "/" + cluster.MCWorkDir); err == nil || !os.IsNotExist(err) {
		// mc working dir not deleted
		return false
	}
	clusterObjResp, err := httpClient.GetContrailCluster(context.Background(),
		&services.GetContrailClusterRequest{
			ID: clusterID,
		},
	)
	return err == nil && clusterObjResp.ContrailCluster.CloudRefs == nil
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

func verifyCommandsExecuted(t *testing.T, expected, msg string) {
	assertYAMLFileEqual(t, expected, executedMCCommand, msg)
}

func assertYAMLFileEqual(t *testing.T, expectedFilePath, actualFilePath string, msg string) {
	var actualYAML interface{}
	require.NoErrorf(t, fileutil.LoadFile(actualFilePath, &actualYAML), "Failed read yaml from %s", actualFilePath)

	var expectedYAML interface{}
	require.NoErrorf(t, fileutil.LoadFile(expectedFilePath, &expectedYAML), "Failed read yaml from %s", expectedFilePath)

	testutil.AssertEqual(t, expectedYAML, actualYAML,
		fmt.Sprintf("YAML files %s and %s are not equal", expectedFilePath, actualFilePath), msg)
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
			action:            cluster.UpgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            cluster.AddComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            cluster.DeleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            cluster.AddCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: cluster.ImportProvisioningAction,
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

	config := &cluster.Config{
		APIServer:           s,
		ClusterID:           clusterID,
		Action:              cluster.CreateAction,
		LogLevel:            "debug",
		TemplateRoot:        "templates/",
		WorkRoot:            workRoot,
		Test:                true,
		LogFile:             workRoot + "/deploy.log",
		ServiceUserID:       integration.ServiceUserName,
		ServiceUserPassword: integration.ServiceUserPassword,
	}

	for _, tt := range []struct {
		action            string
		expectedPlaybooks string
	}{{
		action:            cluster.CreateAction,
		expectedPlaybooks: expectedCreatePlaybooks,
	}, {
		action:            cluster.UpdateAction,
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
	config.Action = cluster.DeleteAction
	removeFile(t, executedPlaybooks)
	player, err := newMockContainerPlayer(filepath.Join(config.WorkRoot, config.ClusterID))
	assert.NoError(t, err, "failed to create mock container player")
	manageCluster(t, config, player)
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}

func runClusterCreateUpdateTest(
	t *testing.T, ts *integration.TestScenario, config *cluster.Config, action string,
	expectedPlaybooks, expectedInstance, expectedInventory string, expectedEndpoints map[string]string) {

	//cleanup old files
	removeFile(t, generatedInstances)
	removeFile(t, generatedInventory)
	removeFile(t, executedPlaybooks)

	config.Action = action
	player, err := newMockContainerPlayer(filepath.Join(config.WorkRoot, config.ClusterID))
	assert.NoError(t, err, "failed to create mock container player")
	manageCluster(t, config, player)

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

func manageCluster(t *testing.T, c *Config, player Player) {
	clusterDeployer, err := cluster.NewCluster(c, testutil.NewFileWritingExecutor(executedMCCommand), player)
	assert.NoErrorf(t, err, "failed to create cluster manager to %s cluster", c.Action)
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoErrorf(t, err, "failed to manage(%s) cluster", c.Action)
}

// nolint: gocyclo
func runClusterActionTest(
	t *testing.T, ts *integration.TestScenario, config *cluster.Config, action, expectedInstance string,
	expectedPlaybooks string, expectedEndpoints map[string]string,
) {
	// set action field in the contrail-cluster resource
	cl := map[string]interface{}{"uuid": clusterID, "provisioning_action": action}
	config.Action = cluster.UpdateAction
	switch action {
	case cluster.UpgradeProvisioningAction:
		cl["provisioning_state"] = "NOSTATE"
	case cluster.ImportProvisioningAction:
		config.Action = cluster.CreateAction
		cl["provisioning_action"] = ""
		err := os.Remove(generatedInstances)
		assert.NoError(t, err, "failed to delete instances.yml")
	case cluster.AddComputeProvisioningAction, cluster.AddCVFMProvisioningAction,
		cluster.DeleteComputeProvisioningAction, cluster.AddCSNProvisioningAction, cluster.DestroyAction:
		err := os.Remove(generatedInstances)
		assert.NoError(t, err, "failed to delete instances.yml")
	}

	data := map[string]interface{}{"contrail-cluster": cl}
	for _, client := range ts.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/contrail-cluster/%s", clusterID)
		_, err := client.Update(context.Background(), url, &data, &response)
		assert.NoErrorf(t, err, "failed to set %s action in contrail cluster", action)
		break
	}
	removeFile(t, executedPlaybooks)
	player, err := newMockContainerPlayer(filepath.Join(config.WorkRoot, config.ClusterID))
	assert.NoError(t, err, "failed to create mock container player")
	manageCluster(t, config, player)
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
			action:            cluster.UpgradeProvisioningAction,
			expectedPlaybooks: upgradeAppformixPlaybooks,
		}, {
			action:            cluster.AddComputeProvisioningAction,
			expectedPlaybooks: addAppformixComputePlaybooks,
		}, {
			action:            cluster.DeleteComputeProvisioningAction,
			expectedPlaybooks: deleteAppformixComputePlaybooks,
		}, {
			action:            cluster.AddCSNProvisioningAction,
			expectedPlaybooks: addAppformixCSNPlaybooks,
		}, {
			action: cluster.ImportProvisioningAction,
		}, {
			action:            cluster.DestroyAction,
			expectedPlaybooks: destroyPlaybooks,
		}},
	)

}

func createDummyAppformixFiles(t *testing.T) func() {
	// create appformix config.yml file
	configFile := workRoot + "/" + "appformix-ansible-deployer/appformix/config.yml"
	configData := []byte(`{"appformix_version": "3.0.0"}`)
	err := fileutil.WriteToFile(configFile, configData, cluster.DefaultFilePermRWOnly)
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

	d, ok := getClusterDeployer(t, &cluster.Config{
		APIServer:           ts.Clients["default"],
		ClusterID:           clusterID,
		Action:              "create",
		LogLevel:            "debug",
		TemplateRoot:        "templates/",
		WorkRoot:            workRoot,
		Test:                true,
		LogFile:             workRoot + "/deploy.log",
		ServiceUserID:       integration.ServiceUserName,
		ServiceUserPassword: integration.ServiceUserPassword,
	}).(*cluster.ContrailAnsibleDeployer)
	require.True(t, ok, "unable to cast deployer to ContrailAnsibleDeployer")

	err = d.CreateInventory()
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

	d, ok := getClusterDeployer(t, &cluster.Config{
		APIServer:           ts.Clients["default"],
		ClusterID:           clusterID,
		Action:              "create",
		LogLevel:            "debug",
		TemplateRoot:        "templates/",
		WorkRoot:            workRoot,
		Test:                true,
		LogFile:             workRoot + "/deploy.log",
		ServiceUserID:       integration.ServiceUserName,
		ServiceUserPassword: integration.ServiceUserPassword,
	}).(*cluster.ContrailAnsibleDeployer)
	require.True(t, ok, "unable to cast deployer to ContrailAnsibleDeployer")

	err = d.CreateInventory()
	assert.NoError(t, err, "unable to create inventory")

	expectedInstance := "test_data/expected_xflow_inband_instances.yaml"

	assertGeneratedInstancesContainExpected(t, expectedInstance,
		"Instance file created during cluster create is not as expected")
}

func getClusterDeployer(t *testing.T, config *Config) base.Deployer {
	player, err := newMockContainerPlayer(filepath.Join(config.WorkRoot, config.ClusterID))
	assert.NoError(t, err, "failed to create mock container player")
	cluster, err := cluster.NewCluster(config, testutil.NewFileWritingExecutor(executedMCCommand), player)
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
			action:            cluster.UpgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            cluster.AddComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            cluster.DeleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            cluster.AddCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: cluster.ImportProvisioningAction,
		}, {
			action:            cluster.AddCVFMProvisioningAction,
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
			action:            cluster.UpgradeProvisioningAction,
			expectedPlaybooks: upgradeEncryptPlaybooks,
		}, {
			action:            cluster.AddComputeProvisioningAction,
			expectedPlaybooks: addComputeEncryptPlaybooks,
		}, {
			action:            cluster.DeleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            cluster.AddCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: cluster.ImportProvisioningAction,
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
			action:            cluster.UpgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            cluster.AddComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            cluster.DeleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            cluster.AddCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: cluster.ImportProvisioningAction,
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
			action:            cluster.UpgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            cluster.AddComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            cluster.DeleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            cluster.AddCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: cluster.ImportProvisioningAction,
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
			action:            cluster.UpgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooks,
		}, {
			action:            cluster.AddComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooks,
		}, {
			action:            cluster.DeleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooks,
		}, {
			action:            cluster.AddCSNProvisioningAction,
			expectedPlaybooks: addCSNPlaybooks,
		}, {
			action: cluster.ImportProvisioningAction,
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
			action:            cluster.UpgradeProvisioningAction,
			expectedPlaybooks: upgradePlaybooksKubernetes,
		}, {
			action:            cluster.AddComputeProvisioningAction,
			expectedPlaybooks: addComputePlaybooksKubernetes,
		}, {
			action:            cluster.DeleteComputeProvisioningAction,
			expectedPlaybooks: deleteComputePlaybooksKubernetes,
		}, {
			action: cluster.ImportProvisioningAction,
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

	config := &cluster.Config{
		APIServer:           s,
		ClusterID:           clusterID,
		Action:              cluster.CreateAction,
		LogLevel:            "debug",
		TemplateRoot:        "templates/",
		WorkRoot:            workRoot,
		Test:                true,
		LogFile:             workRoot + "/deploy.log",
		ServiceUserID:       integration.ServiceUserName,
		ServiceUserPassword: integration.ServiceUserPassword,
	}
	// create cluster
	removeFile(t, executedPlaybooks)
	player, err := newMockContainerPlayer(filepath.Join(config.WorkRoot, config.ClusterID))
	assert.NoError(t, err, "failed to create mock container player")
	manageCluster(t, config, player)
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
	config.Action = cluster.UpdateAction
	// remove instances.yml to trigger cluster update
	removeFile(t, generatedInstances)
	removeFile(t, executedPlaybooks)

	manageCluster(t, config, player)
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
		cluster.UpgradeProvisioningAction, expectedInstance,
		upgradePlaybooksvcenter, expectedEndpoints)

	// IMPORT test (expected to create endpoints withtout triggering playbooks)
	runClusterActionTest(t, ts, config,
		cluster.ImportProvisioningAction, expectedInstance, "", expectedEndpoints)

	// delete cluster
	config.Action = cluster.DeleteAction
	removeFile(t, executedPlaybooks)

	manageCluster(t, config, player)
	// make sure cluster is removed
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}
func TestWindowsCompute(t *testing.T) {
	ts, err := integration.LoadTest("./test_data/test_windows_compute.yml", nil)
	require.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, ts, server)
}

func TestMCCluster(t *testing.T) {
	runMCClusterTest(t, pongo2.Context{
		"CONTROL_NODES": "",
	})
}

// nolint: gocyclo
func runMCClusterTest(t *testing.T, pContext map[string]interface{}) {
	// Create the cluster and related objects
	ts, err := integration.LoadTest(allInOneMCClusterTemplatePath, pContext)
	require.NoError(t, err, "failed to load mc cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	s, err := integration.NewAdminHTTPClient(server.URL())
	assert.NoError(t, err)

	config := &cluster.Config{
		APIServer:                 s,
		ClusterID:                 clusterID,
		Action:                    cluster.CreateAction,
		LogLevel:                  "debug",
		TemplateRoot:              "templates/",
		WorkRoot:                  workRoot,
		Test:                      true,
		LogFile:                   workRoot + "/deploy.log",
		AnsibleFetchURL:           "ansibleFetchURL",
		AnsibleCherryPickRevision: "ansibleCherryPickRevision",
		AnsibleRevision:           "ansibleRevision",
		ServiceUserID:             integration.ServiceUserName,
		ServiceUserPassword:       integration.ServiceUserPassword,
	}

	cloudFileCleanup := createDummyCloudFiles(t)
	defer cloudFileCleanup()
	// create cluster
	removeFile(t, executedPlaybooks)
	removeFile(t, executedMCCommand)

	player, err := newMCMockContainerPlayer(t, filepath.Join(config.WorkRoot, config.ClusterID), executedMCCommand)
	assert.NoError(t, err, "failed to create mock container player")

	clusterDeployer, err := cluster.NewCluster(config, testutil.NewFileWritingExecutor(executedMCCommand), player)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.Error(t, err,
		"mc deployment should fail because cloud provisioning has failed")

	verifyClusterProvisioningStatus(
		context.Background(), t, config.APIServer, config.ClusterID, cluster.StatusCreateFailed,
	)
	err = isCloudSecretFilesDeleted()
	require.NoError(t, err, "failed to delete public cloud secrets during create")

	for _, tt := range []struct {
		tsPath           string
		action           string
		expectedCommands string
		expectedStatus   string
	}{{
		tsPath:           allInOneMCCloudUpdateTemplatePath,
		action:           cluster.CreateAction,
		expectedCommands: expectedMCCreateCmdExecuted,
		expectedStatus:   cluster.StatusCreated,
	}, {
		tsPath:           allInOneMCClusterUpdateTemplatePath,
		action:           cluster.UpdateAction,
		expectedCommands: expectedMCUpdateCmdExecuted,
		expectedStatus:   cluster.StatusUpdated,
	}} {
		//cleanup all the files
		removeFile(t, executedMCCommand)
		removeFile(t, generatedTopology)
		removeFile(t, generatedSecret)

		config.Action = tt.action

		ts, err = integration.LoadTest(tt.tsPath, pContext)
		require.NoErrorf(t, err, "failed to load MC test data for cluster %s", tt.action)
		_ = integration.RunDirtyTestScenario(t, ts, server)

		manageCluster(t, config, player)
		verifyClusterProvisioningStatus(context.Background(), t, config.APIServer, config.ClusterID, tt.expectedStatus)
		err = isCloudSecretFilesDeleted()
		require.NoErrorf(t, err, "failed to delete public cloud secrets during %s", tt.action)

		assert.Truef(t, compareFiles(t, expectedMCClusterTopology, generatedTopology),
			"Topolgy file created during cluster %s is not as expected", tt.action)
		verifyCommandsExecuted(t, tt.expectedCommands,
			fmt.Sprintf("MC commands executed during cluster %s are not as expected", tt.action))
	}

	// delete cloud secanrio
	//cleanup all the files
	removeFile(t, executedPlaybooks)
	removeFile(t, executedMCCommand)

	ts, err = integration.LoadTest(allInOneMCClusterDeleteTemplatePath, pContext)
	require.NoError(t, err, "failed to load mc cluster test data")
	_ = integration.RunDirtyTestScenario(t, ts, server)
	clusterDeployer, err = cluster.NewCluster(config, testutil.NewFileWritingExecutor(executedMCCommand), player)
	assert.NoError(t, err, "failed to create cluster manager to delete cloud")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(delete) cloud")
	err = isCloudSecretFilesDeleted()
	require.NoError(t, err, "failed to delete public cloud secrets during delete")
	verifyClusterProvisioningStatus(context.Background(), t, config.APIServer, config.ClusterID, cluster.StatusUpdated)
	verifyCommandsExecuted(t, expectedMCDeleteCmdExecuted,
		"commands executed during cluster delete are not as expected")
	// make sure cluster is removed
	assert.True(t, verifyMCDeleted(clusterDeployer.APIServer), "MC folder is not deleted during cluster delete")

	// delete cluster itself
	config.Action = cluster.DeleteAction
	manageCluster(t, config, player)
	err = isCloudSecretFilesDeleted()
	require.NoError(t, err, "failed to delete cloud secrets during delete")
	assert.True(t, verifyClusterDeleted(), "Instance file is not deleted during cluster delete")
}

func verifyClusterProvisioningStatus(
	ctx context.Context, t *testing.T, apiServer *client.HTTP, clusterUUID, expectedStatus string,
) {
	c, err := apiServer.GetContrailCluster(ctx, &services.GetContrailClusterRequest{
		ID: clusterUUID,
	})
	require.NoError(t, err)
	assert.Equal(t, expectedStatus, c.ContrailCluster.ProvisioningState)
}

func createDummyCloudFiles(t *testing.T) func() {
	files := []struct {
		src  string
		dest string
	}{{
		src:  "./test_data/public_cloud_topology.yml",
		dest: "/var/tmp/cloud/public_cloud_uuid/topology.yml",
	}, {
		src:  "./test_data/pvt_cloud_topology.yml",
		dest: "/var/tmp/cloud/pvt_cloud_uuid/topology.yml",
	}}

	for _, f := range files {
		err := fileutil.CopyFile(f.src, f.dest, true)
		assert.NoErrorf(t, err, "Couldn't copy file %s to %s", f.src, f.dest)
	}

	return func() {
		for _, f := range files {
			assert.NoError(t, os.Remove(f.dest), "Couldn't cleanup file %s", f.dest)
		}
	}
}

func isCloudSecretFilesDeleted() error {
	errstrings := []string{}
	for _, secret := range []string{
		"/var/tmp/cloud/public_cloud_uuid/secret.yml",
		generatedSecret,
	} {
		if _, err := os.Stat(secret); err != nil {
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

	config := &cluster.Config{
		APIServer:           s,
		ClusterID:           clusterID,
		Action:              cluster.CreateAction,
		LogLevel:            "debug",
		TemplateRoot:        "templates/",
		WorkRoot:            workRoot,
		Test:                true,
		LogFile:             workRoot + "/deploy.log",
		ServiceUserID:       integration.ServiceUserName,
		ServiceUserPassword: integration.ServiceUserPassword,
	}
	// create cluster
	removeFile(t, executedPlaybooks)

	player, err := newMockContainerPlayer(filepath.Join(config.WorkRoot, config.ClusterID))
	assert.NoError(t, err, "failed to create mock container player")
	clusterDeployer, err := cluster.NewCluster(config, testutil.NewFileWritingExecutor(executedMCCommand), player)
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
	config.Action = cluster.DeleteAction
	clusterDeployer, err = cluster.NewCluster(config, testutil.NewFileWritingExecutor(executedMCCommand), player)
	assert.NoError(t, err, "failed to create cluster manager to delete cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to delete tripleo cluster")
}

func TestJUJUClusterImport(t *testing.T) {
	// Create the cluster and related objects
	ts, err := integration.LoadTest(
		allInOneClusterTemplatePath,
		pongo2.Context{
			"TYPE":                  "kernel",
			"MGMT_INT_IP":           "127.0.0.1",
			"KEYSTONE_VIP":          "10.10.10.100",
			"NOVA_VIP":              "20.20.20.100",
			"SWIFT_VIP":             "30.30.30.100",
			"IRONIC_VIP":            "40.40.40.100",
			"GLANCE_VIP":            "50.50.50.100",
			"CONTRAIL_EXTERNAL_VIP": "overcloud.localdomain",
			"SSL_ENABLE":            "yes",
			"PROVISIONER_TYPE":      "juju",
			"PROVISIONING_STATE":    "CREATED",
			"PROVISIONING_ACTION":   "",
			"CLUSTER_NAME":          t.Name(),
		},
	)
	require.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	s, err := integration.NewAdminHTTPClient(server.URL())
	assert.NoError(t, err)

	config := &cluster.Config{
		APIServer:           s,
		ClusterID:           clusterID,
		Action:              cluster.CreateAction,
		LogLevel:            "debug",
		TemplateRoot:        "templates/",
		WorkRoot:            workRoot,
		Test:                true,
		LogFile:             workRoot + "/deploy.log",
		ServiceUserID:       integration.ServiceUserName,
		ServiceUserPassword: integration.ServiceUserPassword,
	}
	// create cluster
	removeFile(t, executedPlaybooks)

	player, err := newMockContainerPlayer(filepath.Join(config.WorkRoot, config.ClusterID))
	assert.NoError(t, err, "failed to create mock container player")
	clusterDeployer, err := cluster.NewCluster(config, testutil.NewFileWritingExecutor(executedMCCommand), player)
	assert.NoError(t, err, "failed to create cluster manager to import juju cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(import) juju cluster")

	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
	// make sure all endpoints are created as part of import
	err = verifyEndpoints(
		t, ts,
		map[string]string{
			"config":    "https://overcloud.localdomain:8082",
			"nodejs":    "https://overcloud.localdomain:8143",
			"telemetry": "https://overcloud.localdomain:8081",
			"baremetal": "https://40.40.40.100:6385",
			"swift":     "https://30.30.30.100:8080",
			"glance":    "https://50.50.50.100:9292",
			"compute":   "https://20.20.20.100:8774",
			"keystone":  "https://10.10.10.100:5000",
		},
	)
	assert.NoError(t, err)
	// delete cluster
	config.Action = cluster.DeleteAction
	clusterDeployer, err = cluster.NewCluster(config, testutil.NewFileWritingExecutor(executedMCCommand), player)
	assert.NoError(t, err, "failed to create cluster manager to delete cluster")
	deployer, err = clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to delete juju cluster")
}
