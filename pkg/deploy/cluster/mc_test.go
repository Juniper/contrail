package cluster

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	commonResourcesPath = "./test_data/multicloud/common_resources.yml"
	mcClusterPath       = "./test_data/multicloud/test_mc_cluster.yml"
	mcCreateSuccessPath = "./test_data/multicloud/test_mc_create_success.yml"
	mcClusterUpdatePath = "./test_data/multicloud/test_mc_update_cluster.yml"
	mcClusterDeletePath = "./test_data/multicloud/test_mc_delete_cluster.yml"

	mcClusterID = "test_mc_cluster_uuid"

	generatedSecret         = workRoot + "/" + mcClusterID + "/" + mcWorkDir + "/" + defaultSecretFile
	generatedTopology       = workRoot + "/" + mcClusterID + "/" + mcWorkDir + "/" + defaultTopologyFile
	generatedContrailCommon = workRoot + "/" + mcClusterID + "/" + mcWorkDir + "/" + defaultContrailCommonFile
	generatedGatewayCommon  = workRoot + "/" + mcClusterID + "/" + mcWorkDir + "/" + defaultGatewayCommonFile
	generatedTORCommon      = workRoot + "/" + mcClusterID + "/" + mcWorkDir + "/" + defaultTORCommonFile

	expectedMCClusterTopology   = "./test_data/multicloud/expected_mc_cluster_topology.yml"
	expectedContrailCommon      = "./test_data/multicloud/expected_mc_contrail_common.yml"
	expectedGatewayCommon       = "./test_data/multicloud/expected_mc_gateway_common.yml"
	expectedTORCommon           = "./test_data/multicloud/expected_mc_tor_common.yml"
	expectedMCCreateCmdExecuted = "./test_data/multicloud/expected_mc_create_cmd_executed.yml"
	expectedMCUpdateCmdExecuted = "./test_data/multicloud/expected_mc_update_cmd_executed.yml"
	expectedMCDeleteCmdExecuted = "./test_data/multicloud/expected_mc_delete_cmd_executed.yml"
)

func TestMCHSuccess(t *testing.T) {
	// prerequisites
	prerequisitesCleanup := makeRequests(t, commonResourcesPath)
	defer prerequisitesCleanup()

	cloudFileCleanup := createDummyCloudFiles(t)
	defer cloudFileCleanup()

	c := newTestConfig(t, createAction)

	//create mc
	cleanup := makeRequests(t, mcCreateSuccessPath)
	defer cleanup()

	cleanupMCFiles(t)
	manageWithSuccess(t, c)
	assert.NoErrorf(t, isCloudSecretFilesDeleted(), "failed to delete public cloud secrets during %s", c.Action)

	compareGeneratedFiles(t)
	assert.Truef(t, verifyCommandsExecuted(t, expectedMCCreateCmdExecuted),
		"MC commands executed during cluster %s are not as expected", c.Action)

	//update mc
	_ = makeRequests(t, mcClusterUpdatePath)

	cleanupMCFiles(t)
	c.Action = updateAction
	manageWithSuccess(t, c)
	assert.NoErrorf(t, isCloudSecretFilesDeleted(), "failed to delete public cloud secrets during %s", c.Action)

	compareGeneratedFiles(t)
	assert.Truef(t, verifyCommandsExecuted(t, expectedMCUpdateCmdExecuted),
		"MC commands executed during cluster %s are not as expected", c.Action)

	//remove cloud
	_ = makeRequests(t, mcClusterDeletePath)

	cleanupMCFiles(t)
	manageWithSuccess(t, c)
	assert.NoErrorf(t, isCloudSecretFilesDeleted(), "failed to delete public cloud secrets during %s", c.Action)

	assert.Truef(t, verifyCommandsExecuted(t, expectedMCDeleteCmdExecuted),
		"MC commands executed during cluster %s are not as expected", c.Action)

	assert.True(t, verifyMCDeleted(c.APIServer), "MC folder is not deleted during cluster delete")

	//cleanup
	c.Action = deleteAction
	manageWithSuccess(t, c)
	assert.True(t, verifyClusterDeleted(mcClusterID), "Cluster directory is not deleted during cluster delete")
}

func TestMCCreateFail(t *testing.T) {
	// prerequisites
	prerequisitesCleanup := makeRequests(t, commonResourcesPath)
	defer prerequisitesCleanup()

	cloudFileCleanup := createDummyCloudFiles(t)
	defer cloudFileCleanup()

	c := newTestConfig(t, createAction)

	//create mc
	cleanup := makeRequests(t, mcClusterPath)
	defer cleanup()

	cleanupMCFiles(t)
	manageWithFailure(t, c)
	assert.NoErrorf(t, isCloudSecretFilesDeleted(), "failed to delete public cloud secrets during %s", c.Action)

	//cleanup
	c.Action = deleteAction
	manageWithSuccess(t, c)
	assert.True(t, verifyClusterDeleted(mcClusterID), "Cluster directory is not deleted during cluster delete")
}

func TestMCUpdateFail(t *testing.T) {
	// prerequisites
	prerequisitesCleanup := makeRequests(t, commonResourcesPath)
	defer prerequisitesCleanup()

	cloudFileCleanup := createDummyCloudFiles(t)
	defer cloudFileCleanup()

	c := newTestConfig(t, updateAction)

	//create mc
	cleanup := makeRequests(t, mcClusterPath)
	defer cleanup()

	//update mc
	_ = makeRequests(t, mcClusterUpdatePath)

	cleanupMCFiles(t)
	manageWithFailure(t, c)
	assert.NoErrorf(t, isCloudSecretFilesDeleted(), "failed to delete public cloud secrets during %s", c.Action)

	//cleanup
	c.Action = deleteAction
	manageWithSuccess(t, c)
	assert.True(t, verifyClusterDeleted(mcClusterID), "Cluster directory is not deleted during cluster delete")
}

func TestMCDeleteFail(t *testing.T) {
	// prerequisites
	prerequisitesCleanup := makeRequests(t, commonResourcesPath)
	defer prerequisitesCleanup()

	c := newTestConfig(t, updateAction)

	//create mc
	cleanup := makeRequests(t, mcClusterPath)
	defer cleanup()

	//delete mc
	_ = makeRequests(t, mcClusterDeletePath)

	cleanupMCFiles(t)
	manageWithFailure(t, c)
	assert.NoErrorf(t, isCloudSecretFilesDeleted(), "failed to delete public cloud secrets during %s", c.Action)

	//cleanup
	c.Action = deleteAction
	manageWithSuccess(t, c)
	assert.True(t, verifyClusterDeleted(mcClusterID), "Cluster directory is not deleted during cluster delete")
}

func makeRequests(t *testing.T, yamlPath string) func() {
	ts, err := integration.LoadTest(yamlPath, nil)
	require.NoError(t, err, "failed to load test data file %s", yamlPath)
	return integration.RunDirtyTestScenario(t, ts, server)
}

func newTestConfig(t *testing.T, action string) *Config {
	s, err := integration.NewAdminHTTPClient(server.URL())
	require.NoError(t, err)

	return &Config{
		APIServer:                 s,
		ClusterID:                 mcClusterID,
		Action:                    action,
		LogLevel:                  "debug",
		TemplateRoot:              "templates/",
		WorkRoot:                  workRoot,
		Test:                      true,
		LogFile:                   workRoot + "/deploy.log",
		AnsibleFetchURL:           "ansibleFetchURL",
		AnsibleCherryPickRevision: "ansibleCherryPickRevision",
		AnsibleRevision:           "ansibleRevision",
	}
}

func manageWithSuccess(t *testing.T, c *Config) {
	clusterDeployer, err := NewCluster(c, testutil.NewFileWritingExecutor(executedCommands))
	assert.NoErrorf(t, err, "failed to create cluster manager to %s cluster", c.Action)
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoErrorf(t, err, "failed to manage(%s) cluster", c.Action)
}

func manageWithFailure(t *testing.T, c *Config) {
	clusterDeployer, err := NewCluster(c, testutil.NewFileWritingExecutor(executedCommands))
	assert.NoErrorf(t, err, "failed to create cluster manager to %s cluster", c.Action)
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.Errorf(t, err, "cluster managed(%s) successfully, expected failure", c.Action)
}

func compareGeneratedFiles(t *testing.T) {
	assert.Truef(t, compareFiles(t, expectedMCClusterTopology, generatedTopology),
		"Generated topolgy file is not as expected")
	assert.Truef(t, compareFiles(t, expectedContrailCommon, generatedContrailCommon),
		"Generated contrail common file is not as expected")
	assert.Truef(t, compareFiles(t, expectedGatewayCommon, generatedGatewayCommon),
		"Generated gateway common file is not as expected")
	assert.Truef(t, compareFiles(t, expectedTORCommon, generatedTORCommon),
		"Generated TOR common file is not as expected")
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

func verifyMCDeleted(httpClient *client.HTTP) bool {
	// Make sure mc working dir is deleted
	if _, err := os.Stat(workRoot + "/" + mcClusterID + "/" + mcWorkDir); err == nil || !os.IsNotExist(err) {
		// mc working dir not deleted
		return false
	}
	clusterObjResp, err := httpClient.GetContrailCluster(context.Background(),
		&services.GetContrailClusterRequest{
			ID: mcClusterID,
		},
	)
	return err == nil && clusterObjResp.ContrailCluster.CloudRefs == nil
}

func cleanupMCFiles(t *testing.T) {
	removeFile(t, executedCommands)
	removeFile(t, generatedTopology)
	removeFile(t, generatedSecret)
	removeFile(t, generatedContrailCommon)
	removeFile(t, generatedGatewayCommon)
	removeFile(t, generatedTORCommon)
}
