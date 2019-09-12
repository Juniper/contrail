package cloud

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	allInOneCloudTemplatePath        = "./test_data/test_all_in_one_public_cloud.tmpl"
	allInOneCloudDeleteTemplatePath  = "./test_data/test_all_in_one_public_cloud_delete.tmpl"
	allInOneCloudUpdateTemplatePath  = "./test_data/test_all_in_one_public_cloud_update.tmpl"
	deleteVPCTemplatePath            = "./test_data/test_vpc_delete.tmpl"
	clusterUpdateFailedTemplatePath  = "./test_data/test_update_failed_cluster.tmpl"
	clusterUpdatedTemplatePath       = "./test_data/test_updated_cluster.tmpl"
	expectedAZTopologyCreate         = "./test_data/expected_azure_cloud_topology_create.yaml"
	expectedAZTopologyUpdate         = "./test_data/expected_azure_cloud_topology_update.yaml"
	expectedAZTopologyDeleteVPC      = "./test_data/expected_azure_cloud_delete_vpc.yaml"
	expectedAZSecret                 = "./test_data/expected_azure_cloud_secret.yaml"
	expectedAWSTopologyCreate        = "./test_data/expected_aws_cloud_topology_create.yaml"
	expectedAWSTopologyUpdate        = "./test_data/expected_aws_cloud_topology_update.yaml"
	expectedAWSTopologyDeleteVPC     = "./test_data/expected_aws_cloud_delete_vpc.yaml"
	expectedAWSSecret                = "./test_data/expected_aws_cloud_secret.yaml"
	expectedCmdForCreateUpdate       = "./test_data/expected_cmd_for_create_update.yaml"
	expectedGCPTopologyCreate        = "./test_data/expected_gcp_cloud_topology_create.yaml"
	expectedGCPTopologyUpdate        = "./test_data/expected_gcp_cloud_topology_update.yaml"
	expectedGCPTopologyDeleteVPC     = "./test_data/expected_gcp_cloud_delete_vpc.yaml"
	expectedGCPSecret                = "./test_data/expected_gcp_cloud_secret.yaml"
	expectedOnPremTopology           = "./test_data/expected_onprem_cloud_topology.yaml"
	expectedOnPremSecret             = "./test_data/expected_onprem_cloud_secret.yaml"
	expectedOnPremCmdForCreateUpdate = "./test_data/expected_onprem_cmd_for_create_update.yaml"
	expectedPvtKey                   = "./test_data/cloud_keypair"
	expectedPubKey                   = "./test_data/cloud_keypair.pub"
	defaultAdminUser                 = "admin"
	defaultAdminPassword             = "contrail123"
	awsAccessKeyFile                 = "/var/tmp/contrail/aws_access.key"
	awsSecretKeyFile                 = "/var/tmp/contrail/aws_secret.key"
)

func prepareForTest(
	t *testing.T, dbRequestsFile, cloudUUID string,
) (postActions func(t *testing.T, cloudUUID string)) {
	return prepareForTestWithFiles(t, dbRequestsFile, []string{}, cloudUUID)
}

func prepareForTestWithFiles(
	t *testing.T, dbRequestsFile string, hostFiles []string, cloudUUID string,
) (postActions func(t *testing.T, cloudUUID string)) {
	refreshTestDirectory(t, cloudUUID)

	for _, file := range hostFiles {
		dest := filepath.Join(GetCloudDir(cloudUUID), filepath.Base(file))
		err := fileutil.CopyFile(file, dest, false)
		assert.NoError(t, err, "Failed to copy files for test environment")
	}

	stopServer := authenticate()

	ts, err := integration.LoadTest(dbRequestsFile, nil)
	if err != nil {
		stopServer()
		assert.Fail(t, "failed to load cloud test data from file", "Caused by: %v", err)
	}
	cleanup := integration.RunDirtyTestScenario(t, ts, server)

	return func(t *testing.T, cloudUUID string) {
		stopServer()
		cleanup()
		removeTestDirectory(t, cloudUUID)
	}
}

func authenticate() func() {
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	return func() {
		ksPublic.Close()
		ksPrivate.Close()
	}
}

func prepareCloud(t *testing.T, cloudUUID, cloudAction string) *Cloud {
	c := &Config{
		ID:           "alice",
		Password:     "alice_password",
		ProjectID:    "admin",
		AuthURL:      server.URL() + "/keystone/v3",
		Endpoint:     server.URL(),
		InSecure:     true,
		CloudID:      cloudUUID,
		Action:       cloudAction,
		LogLevel:     "info",
		TemplateRoot: "configs/",
		Test:         true,
	}

	cl, err := NewCloud(c)
	assert.NoError(t, err, "failed to create cloud struct")

	return cl
}

func TestOnPremCloud(t *testing.T) {
	runCloudTest(
		t,
		[]string{expectedOnPremTopology},
		expectedOnPremSecret,
		expectedOnPremCmdForCreateUpdate,
		pongo2.Context{
			"CLOUD_TYPE": onPrem,
		},
	)
}

func TestAzureCloud(t *testing.T) {
	runCloudTest(
		t,
		[]string{expectedAZTopologyCreate, expectedAZTopologyUpdate, expectedAZTopologyDeleteVPC},
		expectedAZSecret,
		expectedCmdForCreateUpdate,
		pongo2.Context{
			"CLOUD_TYPE": azure,
		},
	)
}

func TestAWSCloud(t *testing.T) {
	runCloudTest(
		t,
		[]string{expectedAWSTopologyCreate, expectedAWSTopologyUpdate, expectedAWSTopologyDeleteVPC},
		expectedAWSSecret,
		expectedCmdForCreateUpdate,
		pongo2.Context{
			"CLOUD_TYPE": AWS,
		},
	)
}

func TestGCPCloud(t *testing.T) {
	runCloudTest(
		t,
		[]string{expectedGCPTopologyCreate, expectedGCPTopologyUpdate, expectedGCPTopologyDeleteVPC},
		expectedGCPSecret,
		expectedCmdForCreateUpdate,
		pongo2.Context{
			"CLOUD_TYPE": gcp,
		},
	)
}

type providerConfig struct {
	providerName           string
	cloudUUID              string
	filesToCopy            []string
	requestsToStartWith    string
	expectedTopologyFile   string
	expectedSecretFile     string
	expectedCommands       string
	fakeCredentialFiles    []string
	fakeCredentialLocation string
}

func TestCreatingPublicClouds(t *testing.T) {
	for _, providerConfig := range []*providerConfig{
		{
			providerName: "AWS",
			cloudUUID:    "cloud_uuid_aws_create",
			filesToCopy: []string{
				"./test_data/cloud_keypair",
				"./test_data/cloud_keypair.pub",
			},
			requestsToStartWith:  "./test_data/test_aws_create/requests.yml",
			expectedTopologyFile: "./test_data/test_aws_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/test_aws_create/expected_secret.yml",
			expectedCommands:     "./test_data/test_aws_create/expected_commands.yml",
			fakeCredentialFiles: []string{
				"./test_data/test_aws_create/aws_access.key",
				"./test_data/test_aws_create/aws_secret.key",
			},
			fakeCredentialLocation: "/var/tmp/contrail",
		},
		{
			providerName: "GCP",
			cloudUUID:    "cloud_uuid_gcp_create",
			filesToCopy: []string{
				"./test_data/cloud_keypair",
				"./test_data/cloud_keypair.pub",
			},
			requestsToStartWith:  "./test_data/test_gcp_create/requests.yml",
			expectedTopologyFile: "./test_data/test_gcp_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/test_gcp_create/expected_secret.yml",
			expectedCommands:     "./test_data/test_gcp_create/expected_commands.yml",
		},
		{
			providerName: "Azure",
			cloudUUID:    "cloud_uuid_azure_create",
			filesToCopy: []string{
				"./test_data/cloud_keypair",
				"./test_data/cloud_keypair.pub",
			},
			requestsToStartWith:  "./test_data/test_azure_create/requests.yml",
			expectedTopologyFile: "./test_data/test_azure_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/test_azure_create/expected_secret.yml",
			expectedCommands:     "./test_data/test_azure_create/expected_commands.yml",
		},
		{
			providerName: "AWS",
			cloudUUID:    "cloud_uuid_aws_update",
			filesToCopy: []string{
				"./test_data/cloud_keypair",
				"./test_data/cloud_keypair.pub",
				"./test_data/test_aws_update/topology_before_update.yml",
			},
			requestsToStartWith:  "./test_data/test_aws_update/requests.yml",
			expectedTopologyFile: "./test_data/test_aws_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/test_aws_update/expected_secret.yml",
			expectedCommands:     "./test_data/test_aws_update/expected_commands.yml",
			fakeCredentialFiles: []string{
				"./test_data/test_aws_update/aws_access.key",
				"./test_data/test_aws_update/aws_secret.key",
			},
			fakeCredentialLocation: "/var/tmp/contrail",
		},
		{
			providerName: "GCP",
			cloudUUID:    "cloud_uuid_gcp_update",
			filesToCopy: []string{
				"./test_data/cloud_keypair",
				"./test_data/cloud_keypair.pub",
				"./test_data/test_gcp_update/topology_before_update.yml",
			},
			requestsToStartWith:  "./test_data/test_gcp_update/requests.yml",
			expectedTopologyFile: "./test_data/test_gcp_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/test_gcp_update/expected_secret.yml",
			expectedCommands:     "./test_data/test_gcp_update/expected_commands.yml",
		},
		{
			providerName: "Azure",
			cloudUUID:    "cloud_uuid_azure_update",
			filesToCopy: []string{
				"./test_data/cloud_keypair",
				"./test_data/cloud_keypair.pub",
				"./test_data/test_azure_update/topology_before_update.yml",
			},
			requestsToStartWith:  "./test_data/test_azure_update/requests.yml",
			expectedTopologyFile: "./test_data/test_azure_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/test_azure_update/expected_secret.yml",
			expectedCommands:     "./test_data/test_azure_update/expected_commands.yml",
		},
	} {
		testPublicCloudCreation(t, providerConfig)
		testPublicCloudCreationWithoutRemovingSecret(t, providerConfig)
	}
}

func testPublicCloudCreation(t *testing.T, pc *providerConfig) {
	postActions := prepareForTestWithFiles(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, createAction)

	removeCreds := createFakeCredentials(t, pc.fakeCredentialFiles, pc.fakeCredentialLocation)
	defer removeCreds()

	assert.NoError(t, cl.Manage(), "failed to manage cloud, while creating cloud")

	verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)

	compareTopology(t, pc.expectedTopologyFile, cl.config.CloudID)

	// assert.True(t, verifyNodeType(cloud.ctx, cloud.APIServer, ts),
	// 	"public cloud nodes are not updated as type private")

	verifyCommandsExecuted(t, pc.expectedCommands, cl.config.CloudID)
	verifyGeneratedSSHKeyFiles(t, cl.config.CloudID)
}

func testPublicCloudCreationWithoutRemovingSecret(t *testing.T, pc *providerConfig) {
	postActions := prepareForTestWithFiles(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, createAction)

	removeCreds := createFakeCredentials(t, pc.fakeCredentialFiles, pc.fakeCredentialLocation)
	defer removeCreds()

	assert.NoError(t, cl.manage(), "failed to manage cloud, while creating cloud")

	compareSecret(t, pc.expectedSecretFile, cl.config.CloudID)
}

func compareTopology(t *testing.T, expectedTopologyFile, cloudUUID string) {
	assert.NoError(t, compareFiles(t,
		expectedTopologyFile, GetTopoFile(cloudUUID)),
		"generated topology file is not as expected",
	)
}

func compareSecret(t *testing.T, expectedSecretFile, cloudUUID string) {
	assert.NoError(t, compareFiles(t,
		expectedSecretFile, GetSecretFile(cloudUUID)),
		"generated secret file is not as expected",
	)
}

func TestCloud(t *testing.T) {
	closeServer := authenticate()
	defer closeServer()
}

// nolint: gocyclo
func runCloudTest(
	t *testing.T,
	expectedTopologies []string,
	expectedSecret string,
	expectedCmdFile string,
	context map[string]interface{},
) {
	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPrivate.Close()

	// create cloud related objects
	ts, err := integration.LoadTest(allInOneCloudTemplatePath, context)
	require.NoErrorf(t, err, "failed to load cloud test data from file: %s", allInOneCloudTemplatePath)
	cleanup := integration.RunDirtyTestScenario(t, ts, server)

	defer cleanup()

	// creating cloud config
	config := &Config{
		ID:           "alice",
		Password:     "alice_password",
		ProjectID:    "admin",
		AuthURL:      server.URL() + "/keystone/v3",
		Endpoint:     server.URL(),
		InSecure:     true,
		CloudID:      "cloud_uuid",
		Action:       createAction,
		LogLevel:     "info",
		TemplateRoot: "configs/",
		Test:         true,
	}

	// cleanTestEnvironment(t)

	// sshKeycleanup := createDummySSHKeyFiles(t)
	// defer sshKeycleanup()

	cloud, err := NewCloud(config)
	assert.NoError(t, err, "failed to create cloud struct")

	// if context["CLOUD_TYPE"] == AWS {
	// 	createAWSAccessKey(t, awsAccessKeyFile)
	// 	createAWSSecretKey(t, awsSecretKeyFile)
	// 	defer removeAWSCredentials(t, awsAccessKeyFile, awsSecretKeyFile)
	// }

	err = cloud.Manage()
	assert.NoError(t, err, "failed to manage cloud, while creating cloud")

	verifyCloudSecretFilesAreDeleted(t, config.CloudID)

	// assert.True(t, compareGeneratedTopology(t, expectedTopologies),
	// 	"topology file created during cloud create is not as expected")

	if context["CLOUD_TYPE"] != onPrem {

		assert.True(t, verifyNodeType(cloud.ctx, cloud.APIServer, ts),
			"public cloud nodes are not updated as type private")

		verifyCommandsExecuted(t, expectedCmdFile, config.CloudID)
		verifyGeneratedSSHKeyFiles(t, config.CloudID)

		// Wait for the in-memory endpoint cache to get updated
		server.ForceProxyUpdate()

		//update cloud
		config.Action = updateAction

		ts, err = integration.LoadTest(allInOneCloudUpdateTemplatePath, context)
		require.NoErrorf(t, err, "failed to load cloud test data from file: %s", allInOneCloudUpdateTemplatePath)
		updateCleanup := integration.RunDirtyTestScenario(t, ts, server)
		defer updateCleanup()

		// delete previously created files

		// Remove topology file and secret file
		err = os.Remove(generatedTopoPath(config.CloudID))
		if err != nil {
			assert.NoError(t, err, "failed to delete topology.yml file, during update")
		}

		if _, err = os.Stat(executedCommandsPath(config.CloudID)); err == nil {
			// cleanup old executed command file
			err = os.Remove(executedCommandsPath(config.CloudID))
			if err != nil {
				assert.NoError(t, err, "failed to delete executed cmd yml, during update")
			}
		}

		cloud, err = NewCloud(config)
		assert.NoError(t, err, "failed to create cloud struct for update action")

		// if context["CLOUD_TYPE"] == AWS {
		// 	createAWSAccessKey(t, awsAccessKeyFile)
		// 	createAWSSecretKey(t, awsSecretKeyFile)
		// }
		err = cloud.Manage()
		assert.NoError(t, err, "failed to manage cloud, while updating onprem cloud")

		verifyCloudSecretFilesAreDeleted(t, config.CloudID)

		// assert.True(t, compareGeneratedTopology(t, expectedTopologies),
		// 	"topology file created during cloud update is not as expected")
		verifyCommandsExecuted(t, expectedCmdFile, config.CloudID)

		// delete vpc and compare topology
		ts, err = integration.LoadTest(deleteVPCTemplatePath, context)
		require.NoErrorf(t, err, "failed to load cloud test data from file: %s", deleteVPCTemplatePath)
		deleteVPC := integration.RunDirtyTestScenario(t, ts, server)
		deleteVPC()

		// delete previously created files

		// Remove topology file and secret file
		err = os.Remove(generatedTopoPath(config.CloudID))
		if err != nil {
			assert.NoError(t, err, "failed to delete topology.yml file, during vpc delete")
		}

		if _, err = os.Stat(executedCommandsPath(config.CloudID)); err == nil {
			// cleanup old executed command file
			err = os.Remove(executedCommandsPath(config.CloudID))
			if err != nil {
				assert.NoError(t, err, "failed to delete executed cmd yml, during vpc delete")
			}
		}

		cloud, err = NewCloud(config)
		assert.NoError(t, err, "failed to create cloud struct for update action")

		// if context["CLOUD_TYPE"] == AWS {
		// 	createAWSAccessKey(t, awsAccessKeyFile)
		// 	createAWSSecretKey(t, awsSecretKeyFile)
		// }
		err = cloud.Manage()
		assert.NoError(t, err, "failed to manage cloud, while updating cloud")

		verifyCloudSecretFilesAreDeleted(t, config.CloudID)

		// assert.True(t, compareGeneratedTopology(t, expectedTopologies),
		// 	"topology file created during cloud delete vpc is not as expected")

	} else {
		config.Action = updateAction
	}

	// delete cloud
	ts, err = integration.LoadTest(allInOneCloudDeleteTemplatePath, context)
	require.NoErrorf(t, err, "failed to load cloud test data from file: %s", allInOneCloudDeleteTemplatePath)
	_ = integration.RunDirtyTestScenario(t, ts, server)

	// delete previously created
	if _, err = os.Stat(executedCommandsPath(config.CloudID)); err == nil {
		// cleanup old executed command file
		err = os.Remove(executedCommandsPath(config.CloudID))
		if err != nil {
			assert.NoError(t, err, "failed to delete executed cmd yml, during delete")
		}
	}

	cloud, err = NewCloud(config)
	assert.NoError(t, err, "failed to create cloud struct for delete action")

	// if context["CLOUD_TYPE"] == AWS {
	// 	createAWSAccessKey(t, awsAccessKeyFile)
	// 	createAWSSecretKey(t, awsSecretKeyFile)
	// }
	err = cloud.Manage()
	verifyCloudSecretFilesAreDeleted(t, config.CloudID)
	if context["CLOUD_TYPE"] == onPrem {
		assert.Error(t, err,
			"delete cloud should fail because cluster p_action is not set to DELETE_CLOUD")

		// updates p_a of cluster to DELETE_CLOUD
		// sets p_s of cluster to UPDATE_FAILED
		ts, err = integration.LoadTest(clusterUpdateFailedTemplatePath, context)
		require.NoError(t, err, "failed to load cluster update failed test data")
		_ = integration.RunDirtyTestScenario(t, ts, server)

		// now delete the cloud again with update failed cluster status
		cloud, err = NewCloud(config)
		assert.NoError(t, err, "failed to create cloud struct for delete action")

		// if context["CLOUD_TYPE"] == AWS {
		// 	createAWSAccessKey(t, awsAccessKeyFile)
		// 	createAWSSecretKey(t, awsSecretKeyFile)
		// }
		err = cloud.Manage()
		assert.Error(t, err,
			"delete cloud should fail because cluster p_a is not set to DELETE_CLOUD but p_s is UPDATE_FAILED")

		verifyCloudSecretFilesAreDeleted(t, config.CloudID)

		// updates p_a of cluster to DELETE_CLOUD
		// sets p_s of cluster to UPDATED
		ts, err = integration.LoadTest(clusterUpdatedTemplatePath, context)
		require.NoError(t, err, "failed to load updated cluster test data")
		_ = integration.RunDirtyTestScenario(t, ts, server)

		// now delete the cloud again with updated cluster status
		cloud, err = NewCloud(config)
		assert.NoError(t, err, "failed to create cloud struct for delete action")

		// if context["CLOUD_TYPE"] == AWS {
		// 	createAWSAccessKey(t, awsAccessKeyFile)
		// 	createAWSSecretKey(t, awsSecretKeyFile)
		// }
		err = cloud.Manage()

		verifyCloudSecretFilesAreDeleted(t, config.CloudID)
	}
	assert.NoError(t, err, "failed to manage cloud, while deleting cloud")

	// make sure cloud is removed
	assert.True(t, verifyCloudDeleted(cloud.ctx, cloud.APIServer, config.CloudID),
		"Cloud dir/Cloud object is not deleted during cloud delete")
}

// func testDeleteOnPrem(t *testing.T) {
// 	ts, err := integration.LoadTest(clusterUpdateFailedTemplatePath, context)
// 	require.NoError(t, err, "failed to load cluster update failed test data")
// 	_ = integration.RunDirtyTestScenario(t, ts, server)
// }

func removeTestDirectory(t *testing.T, cloudUUID string) {
	err := os.RemoveAll(getCloudDir(cloudUUID))
	if err == nil || os.IsNotExist(err) {
		return
	}
	assert.Fail(t, "Cannot clean test environment: %v", err)
}

func refreshTestDirectory(t *testing.T, cloudUUID string) {
	removeTestDirectory(t, cloudUUID)
	assert.NoError(t, os.MkdirAll(getCloudDir(cloudUUID), 0777), "Cannot clean test environment")
}

// func removeAWSCredentials(t *testing.T, awsAccessKey, awsSecretKey string) {
// 	// nolint: errcheck
// 	_ = os.Remove(awsAccessKey)
// 	// nolint: errcheck
// 	_ = os.Remove(awsSecretKey)
// 	_, err := ioutil.ReadFile(awsAccessKey)
// 	assert.True(t, os.IsNotExist(err), "File %s was not removed!", awsAccessKey)
// 	_, err = ioutil.ReadFile(awsSecretKey)
// 	assert.True(t, os.IsNotExist(err), "File %s was not removed!", awsSecretKey)
// }

func compareFiles(t *testing.T, expectedFile, generatedFile string) error {
	generatedData, err := ioutil.ReadFile(generatedFile)
	assert.NoErrorf(t, err, "unable to read generated: %s", generatedFile)
	expectedData, err := ioutil.ReadFile(expectedFile)
	assert.NoErrorf(t, err, "unable to read expected: %s", expectedFile)
	if !bytes.Equal(generatedData, expectedData) {
		return fmt.Errorf("Expected:\n%s\n\nActual:\n%s", expectedData, generatedData)
	}
	return nil
}

// func compareGeneratedTopology(t *testing.T, expectedTopologies []string) bool {
// 	for _, topo := range expectedTopologies {
// 		if compareFiles(t, topo, generatedTopoPath()) {
// 			return true
// 		}
// 		assert.NoError(t, )
// 	}
// 	return false
// }

// Change this const to pass variable
func verifyCommandsExecuted(t *testing.T, expectedCmdFile string, cloudUUID string) {
	assert.NoError(t, compareFiles(t, expectedCmdFile, executedCommandsPath(cloudUUID)),
		"Expected list of create commands are not executed")
}

func verifyGeneratedSSHKeyFiles(t *testing.T, cloudUUID string) {
	pvtKeyPath := getCloudSSHKeyPath(cloudUUID, "cloud_keypair")
	pubKeyPath := getCloudSSHKeyPath(cloudUUID, "cloud_keypair.pub")
	assert.NoError(t, compareFiles(t, expectedPvtKey, pvtKeyPath), "Private SSH keys doesn't match!")
	assert.NoError(t, compareFiles(t, expectedPubKey, pubKeyPath), "Public SSH keys doesn't match!")
}

// func createAWSAccessKey(t *testing.T, path string) {
// 	data := []byte("access_key")

// 	err := fileutil.WriteToFile(path, data, sshPubKeyPerm)
// 	assert.NoErrorf(t, err, "Unable to write file: %s", path)
// }

// func createAWSSecretKey(t *testing.T, path string) {
// 	data := []byte("secret_key")

// 	err := fileutil.WriteToFile(path, data, sshPubKeyPerm)
// 	assert.NoErrorf(t, err, "Unable to write file: %s", path)
// }

func createFakeCredentials(
	t *testing.T, files []string, location string,
) (cleanCredentials func()) {
	filesToRemove := []string{}
	for _, file := range files {
		dest := filepath.Join(location, filepath.Base(file))
		filesToRemove = append(filesToRemove, dest)
		err := fileutil.CopyFile(file, dest, true)
		assert.NoErrorf(t, err, "Failed to copy file %s", file)
	}
	return func() {
		for _, file := range filesToRemove {
			if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
				assert.Fail(t, "Unable to remove credential file", err.Error())
			}
		}
	}
}

// func copyFile(t *testing.T, src, dest string, permission int) func() {
// 	data, err := fileutil.GetContent("file://" + src)
// 	if err != nil {
// 		assert.NoErrorf(t, err, "Unable to read file: %s", src)
// 	}
// 	err = fileutil.WriteToFile(dest, data, sshPubKeyPerm)
// 	if err != nil {
// 		assert.NoErrorf(t, err, "Unable to write file: %s", dest)
// 	}
// }

// func createDummySSHKeyFiles(t *testing.T) func() {
// 	// create ssh pub key
// 	pubKeyData, err := fileutil.GetContent("file://" + expectedPubKey)
// 	if err != nil {
// 		assert.NoErrorf(t, err, "Unable to read file: %s", expectedPubKey)
// 	}

// 	err = fileutil.WriteToFile("/tmp/cloud_keypair.pub", pubKeyData, sshPubKeyPerm)
// 	if err != nil {
// 		assert.NoErrorf(t, err, "Unable to write file: %s", "/tmp/cloud_keypair.pub")
// 	}

// 	pvtKeyData, err := fileutil.GetContent("file://" + expectedPvtKey)
// 	if err != nil {
// 		assert.NoErrorf(t, err, "Unable to read file: %s", expectedPvtKey)
// 	}
// 	err = fileutil.WriteToFile("/tmp/cloud_keypair", pvtKeyData, defaultRWOnlyPerm)
// 	if err != nil {
// 		assert.NoErrorf(t, err, "Unable to write file: %s", "/tmp/cloud_keypair")
// 	}

// 	return func() {
// 		// nolint: errcheck
// 		_ = os.Remove("/tmp/cloud_keypair")
// 		// nolint: errcheck
// 		_ = os.Remove("/tmp/cloud_keypair.pub")
// 	}
// }

func verifyNodeType(ctx context.Context, httpClient *client.HTTP, testScenario *integration.TestScenario) bool {
	for _, task := range testScenario.Workflow {
		if task.Request.Path == "/nodes" {
			//nolint: errcheck
			expectMap, _ := task.Expect.(map[string]interface{})
			//nolint: errcheck
			nodeData, _ := expectMap["node"].(map[string]interface{})
			//nolint: errcheck
			nodeUUID, _ := nodeData["uuid"].(string)

			nodeResp, err := httpClient.GetNode(ctx,
				&services.GetNodeRequest{
					ID: nodeUUID,
				},
			)
			if err != nil {
				return false
			}
			if nodeResp.Node.Type != "private" {
				return false
			}
		}
	}
	return true
}

func getCloudDir(cloudUUID string) string {
	return defaultWorkRoot + "/" + cloudUUID
}

func generatedTopoPath(cloudUUID string) string {
	return defaultWorkRoot + "/" + cloudUUID + "/topology.yml"
}

// Maybe take care of this
func executedCommandsPath(cloudUUID string) string {
	return defaultWorkRoot + "/" + cloudUUID + "/" + executedCmdTestFile
}

func verifyCloudDeleted(ctx context.Context, httpClient *client.HTTP, cloudUUID string) bool {
	if _, err := os.Stat(defaultWorkRoot + "/" + cloudUUID); err == nil {
		// working dir not deleted
		return false
	}

	_, err := httpClient.GetCloud(ctx,
		&services.GetCloudRequest{
			ID: cloudUUID,
		},
	)
	if err == nil {
		return false
	}
	return true
}

func verifyCloudSecretFilesAreDeleted(t *testing.T, cloudUUID string) {
	keyDefaults, err := services.NewKeyFileDefaults()
	if err != nil {
		assert.NoError(t, err, "Cannot verify if cloud secret files are deleted")
	}
	errstrings := []string{}
	for _, secret := range []string{
		GetTerraformAWSPlanFile(cloudUUID),
		GetTerraformAzurePlanFile(cloudUUID),
		GetTerraformGCPPlanFile(cloudUUID),
		awsAccessKeyFile,
		awsSecretKeyFile,
		keyDefaults.GetAzureAccessTokenPath(),
		keyDefaults.GetAzureProfilePath(),
		keyDefaults.GetGoogleAccountPath(),
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
	if len(errstrings) == 0 {
		return
	}
	assert.Fail(t, strings.Join(errstrings, "\n"))
}
