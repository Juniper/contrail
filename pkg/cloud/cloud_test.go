package cloud

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
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
	expectedAZCmdForCreateUpdate     = "./test_data/expected_azure_cmd_for_create_update.yaml"
	expectedAZTopologyCreate         = "./test_data/expected_azure_cloud_topology_create.yaml"
	expectedAZTopologyUpdate         = "./test_data/expected_azure_cloud_topology_update.yaml"
	expectedAZTopologyDeleteVPC      = "./test_data/expected_azure_cloud_delete_vpc.yaml"
	expectedAZSecret                 = "./test_data/expected_azure_cloud_secret.yaml"
	expectedAWSCmdForCreateUpdate    = "./test_data/expected_aws_cmd_for_create_update.yaml"
	expectedAWSTopologyCreate        = "./test_data/expected_aws_cloud_topology_create.yaml"
	expectedAWSTopologyUpdate        = "./test_data/expected_aws_cloud_topology_update.yaml"
	expectedAWSTopologyDeleteVPC     = "./test_data/expected_aws_cloud_delete_vpc.yaml"
	expectedAWSSecret                = "./test_data/expected_aws_cloud_secret.yaml"
	expectedGCPCmdForCreateUpdate    = "./test_data/expected_gcp_cmd_for_create_update.yaml"
	expectedGCPTopologyCreate        = "./test_data/expected_gcp_cloud_topology_create.yaml"
	expectedGCPTopologyUpdate        = "./test_data/expected_gcp_cloud_topology_update.yaml"
	expectedGCPTopologyDeleteVPC     = "./test_data/expected_gcp_cloud_delete_vpc.yaml"
	expectedGCPSecret                = "./test_data/expected_gcp_cloud_secret.yaml"
	expectedOnPremTopology           = "./test_data/expected_onprem_cloud_topology.yaml"
	expectedOnPremSecret             = "./test_data/expected_onprem_cloud_secret.yaml"
	expectedOnPremCmdForCreateUpdate = "./test_data/expected_onprem_cmd_for_create_update.yaml"
	expectedPvtKey                   = "./test_data/cloud_keypair"
	expectedPubKey                   = "./test_data/cloud_keypair.pub"
	cloudID                          = "cloud_uuid"
	cloudTypeKey                     = "CLOUD_TYPE"
	defaultAdminUser                 = "admin"
	defaultAdminPassword             = "contrail123"
	awsAccessKeyFile                 = "/var/tmp/contrail/aws_access.key"
	awsSecretKeyFile                 = "/var/tmp/contrail/aws_secret.key"
)

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
		expectedAZCmdForCreateUpdate,
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
		expectedAWSCmdForCreateUpdate,
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
		expectedGCPCmdForCreateUpdate,
		pongo2.Context{
			"CLOUD_TYPE": gcp,
		},
	)
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
		CloudID:      cloudID,
		Action:       createAction,
		LogLevel:     "info",
		TemplateRoot: "configs/",
		Test:         true,
	}

	// delete previously created
	if _, err = os.Stat(executedCommandsPath()); err == nil {
		// cleanup old executed command file
		err = os.Remove(executedCommandsPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed commands yaml")
		}
	}

	sshKeycleanup := createDummySSHKeyFiles(t)
	defer sshKeycleanup()

	cloud, err := NewCloud(config)
	assert.NoError(t, err, "failed to create cloud struct")

	if context[cloudTypeKey] == AWS {
		createAWSAccessKey(t, awsAccessKeyFile)
		createAWSSecretKey(t, awsSecretKeyFile)
		defer removeAWSCredentials(t, awsAccessKeyFile, awsSecretKeyFile)
	} else if context[cloudTypeKey] == azure {
		createAzureCredentials(t)
		defer removeAzureCredentials(t)
	}
	err = cloud.Manage()
	assert.NoError(t, err, "failed to manage cloud, while creating cloud")

	err = isCloudSecretFilesDeleted()
	require.NoError(t, err, "failed to delete cloud secrets during create")

	assert.True(t, compareGeneratedTopology(t, expectedTopologies),
		"topology file created during cloud create is not as expected")

	if context["CLOUD_TYPE"] != onPrem {

		assert.True(t, verifyNodeType(cloud.ctx, cloud.APIServer, ts),
			"public cloud nodes are not updated as type private")

		assert.True(t, verifyCommandsExecuted(t, expectedCmdFile),
			"Expected list of create commands are not executed")
		// check if ssh keys are created
		assert.True(t, verifyGeneratedSSHKeyFiles(t),
			"Expected ssh key file are not generated")

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
		err = os.Remove(generatedTopoPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete topology.yml file, during update")
		}

		if _, err = os.Stat(executedCommandsPath()); err == nil {
			// cleanup old executed command file
			err = os.Remove(executedCommandsPath())
			if err != nil {
				assert.NoError(t, err, "failed to delete executed cmd yml, during update")
			}
		}

		cloud, err = NewCloud(config)
		assert.NoError(t, err, "failed to create cloud struct for update action")

		if context["CLOUD_TYPE"] == AWS {
			createAWSAccessKey(t, awsAccessKeyFile)
			createAWSSecretKey(t, awsSecretKeyFile)
		} else if context["CLOUD_TYPE"] == azure {
			createAzureCredentials(t)
		}
		err = cloud.Manage()
		assert.NoError(t, err, "failed to manage cloud, while updating onprem cloud")

		err = isCloudSecretFilesDeleted()
		require.NoError(t, err, "failed to delete cloud secrets during update")

		assert.True(t, compareGeneratedTopology(t, expectedTopologies),
			"topology file created during cloud update is not as expected")
		assert.True(t, verifyCommandsExecuted(t, expectedCmdFile),
			"Expected list of update commands are not executed")

		// delete vpc and compare topology
		ts, err = integration.LoadTest(deleteVPCTemplatePath, context)
		require.NoErrorf(t, err, "failed to load cloud test data from file: %s", deleteVPCTemplatePath)
		deleteVPC := integration.RunDirtyTestScenario(t, ts, server)
		deleteVPC()

		// delete previously created files

		// Remove topology file and secret file
		err = os.Remove(generatedTopoPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete topology.yml file, during vpc delete")
		}

		if _, err = os.Stat(executedCommandsPath()); err == nil {
			// cleanup old executed command file
			err = os.Remove(executedCommandsPath())
			if err != nil {
				assert.NoError(t, err, "failed to delete executed cmd yml, during vpc delete")
			}
		}

		cloud, err = NewCloud(config)
		assert.NoError(t, err, "failed to create cloud struct for update action")

		if context["CLOUD_TYPE"] == AWS {
			createAWSAccessKey(t, awsAccessKeyFile)
			createAWSSecretKey(t, awsSecretKeyFile)
		} else if context["CLOUD_TYPE"] == azure {
			createAzureCredentials(t)
		}
		err = cloud.Manage()
		assert.NoError(t, err, "failed to manage cloud, while updating cloud")

		err = isCloudSecretFilesDeleted()
		require.NoError(t, err, "failed to delete cloud secrets during onprem delete")

		assert.True(t, compareGeneratedTopology(t, expectedTopologies),
			"topology file created during cloud delete vpc is not as expected")

	} else {
		config.Action = updateAction
	}

	// delete cloud
	ts, err = integration.LoadTest(allInOneCloudDeleteTemplatePath, context)
	require.NoErrorf(t, err, "failed to load cloud test data from file: %s", allInOneCloudDeleteTemplatePath)
	_ = integration.RunDirtyTestScenario(t, ts, server)

	// delete previously created
	if _, err = os.Stat(executedCommandsPath()); err == nil {
		// cleanup old executed command file
		err = os.Remove(executedCommandsPath())
		if err != nil {
			assert.NoError(t, err, "failed to delete executed cmd yml, during delete")
		}
	}

	cloud, err = NewCloud(config)
	assert.NoError(t, err, "failed to create cloud struct for delete action")

	if context["CLOUD_TYPE"] == AWS {
		createAWSAccessKey(t, awsAccessKeyFile)
		createAWSSecretKey(t, awsSecretKeyFile)
	} else if context["CLOUD_TYPE"] == azure {
		createAzureCredentials(t)
	}
	err = cloud.Manage()
	ok := isCloudSecretFilesDeleted()
	require.NoError(t, ok, "failed to delete cloud secrets during delete")
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

		if context["CLOUD_TYPE"] == AWS {
			createAWSAccessKey(t, awsAccessKeyFile)
			createAWSSecretKey(t, awsSecretKeyFile)
		}
		err = cloud.Manage()
		assert.Error(t, err,
			"delete cloud should fail because cluster p_a is not set to DELETE_CLOUD but p_s is UPDATE_FAILED")

		err = isCloudSecretFilesDeleted()
		require.NoError(t, err, "failed to delete cloud secrets during delete")

		// updates p_a of cluster to DELETE_CLOUD
		// sets p_s of cluster to UPDATED
		ts, err = integration.LoadTest(clusterUpdatedTemplatePath, context)
		require.NoError(t, err, "failed to load updated cluster test data")
		_ = integration.RunDirtyTestScenario(t, ts, server)

		// now delete the cloud again with updated cluster status
		cloud, err = NewCloud(config)
		assert.NoError(t, err, "failed to create cloud struct for delete action")

		if context["CLOUD_TYPE"] == AWS {
			createAWSAccessKey(t, awsAccessKeyFile)
			createAWSSecretKey(t, awsSecretKeyFile)
		}
		err = cloud.Manage()

		ok := isCloudSecretFilesDeleted()
		require.NoError(t, ok, "failed to delete cloud secrets during delete")
	}
	assert.NoError(t, err, "failed to manage cloud, while deleting cloud")

	// make sure cloud is removed
	assert.True(t, verifyCloudDeleted(cloud.ctx, cloud.APIServer),
		"Cloud dir/Cloud object is not deleted during cloud delete")
}

func removeAWSCredentials(t *testing.T, awsAccessKey, awsSecretKey string) {
	// nolint: errcheck
	_ = os.Remove(awsAccessKey)
	// nolint: errcheck
	_ = os.Remove(awsSecretKey)
	_, err := ioutil.ReadFile(awsAccessKey)
	assert.True(t, os.IsNotExist(err), "File %s was not removed!", awsAccessKey)
	_, err = ioutil.ReadFile(awsSecretKey)
	assert.True(t, os.IsNotExist(err), "File %s was not removed!", awsSecretKey)
}

func compareFiles(t *testing.T, expectedFile, generatedFile string) bool {
	generatedData, err := ioutil.ReadFile(generatedFile)
	assert.NoErrorf(t, err, "unable to read generated: %s", generatedFile)
	expectedData, err := ioutil.ReadFile(expectedFile)
	assert.NoErrorf(t, err, "unable to read expected: %s", expectedFile)
	return bytes.Equal(generatedData, expectedData)
}

func compareGeneratedTopology(t *testing.T, expectedTopologies []string) bool {
	for _, topo := range expectedTopologies {
		if compareFiles(t, topo, generatedTopoPath()) {
			return true
		}
	}
	return false
}

func verifyCommandsExecuted(t *testing.T, expectedCmdFile string) bool {
	return compareFiles(t, expectedCmdFile, executedCommandsPath())
}

func verifyGeneratedSSHKeyFiles(t *testing.T) bool {
	pvtKeyPath := getCloudSSHKeyPath(cloudID, "cloud_keypair")
	pubKeyPath := getCloudSSHKeyPath(cloudID, "cloud_keypair.pub")
	return compareFiles(t, expectedPvtKey,
		pvtKeyPath) && compareFiles(t, expectedPubKey, pubKeyPath)
}

func createAWSAccessKey(t *testing.T, path string) {
	err := fileutil.WriteToFile(path, []byte("access_key"), sshPubKeyPerm)
	assert.NoErrorf(t, err, "Unable to write file: %s", path)
}

func createAWSSecretKey(t *testing.T, path string) {
	err := fileutil.WriteToFile(path, []byte("secret_key"), sshPubKeyPerm)
	assert.NoErrorf(t, err, "Unable to write file: %s", path)
}

func createAzureCredentials(t *testing.T) {
	kfd := services.NewKeyFileDefaults()

	err := fileutil.WriteToFile(kfd.GetAzureSubscriptionIDPath(), []byte("subscription_id"), defaultRWOnlyPerm)
	assert.NoErrorf(t, err, "Unable to write file: %s", kfd.GetAzureSubscriptionIDPath())

	err = fileutil.WriteToFile(kfd.GetAzureClientIDPath(), []byte("client_id"), defaultRWOnlyPerm)
	assert.NoErrorf(t, err, "Unable to write file: %s", kfd.GetAzureClientIDPath())

	err = fileutil.WriteToFile(kfd.GetAzureClientSecretPath(), []byte("client_secret"), defaultRWOnlyPerm)
	assert.NoErrorf(t, err, "Unable to write file: %s", kfd.GetAzureClientSecretPath())

	err = fileutil.WriteToFile(kfd.GetAzureTenantIDPath(), []byte("tenant_id"), defaultRWOnlyPerm)
	assert.NoErrorf(t, err, "Unable to write file: %s", kfd.GetAzureTenantIDPath())
}

func removeAzureCredentials(t *testing.T) {
	kfd := services.NewKeyFileDefaults()
	for _, f := range []string{
		kfd.GetAzureSubscriptionIDPath(),
		kfd.GetAzureClientIDPath(),
		kfd.GetAzureClientSecretPath(),
		kfd.GetAzureTenantIDPath(),
	} {
		// nolint: errcheck
		_ = os.Remove(f)
		_, err := ioutil.ReadFile(f)
		assert.True(t, os.IsNotExist(err), "File %s was not removed!", f)
	}
}

func createDummySSHKeyFiles(t *testing.T) func() {
	// create ssh pub key
	pubKeyData, err := fileutil.GetContent("file://" + expectedPubKey)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to read file: %s", expectedPubKey)
	}

	err = fileutil.WriteToFile("/tmp/cloud_keypair.pub", pubKeyData, sshPubKeyPerm)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to write file: %s", "/tmp/cloud_keypair.pub")
	}

	pvtKeyData, err := fileutil.GetContent("file://" + expectedPvtKey)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to read file: %s", expectedPvtKey)
	}
	err = fileutil.WriteToFile("/tmp/cloud_keypair", pvtKeyData, defaultRWOnlyPerm)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to write file: %s", "/tmp/cloud_keypair")
	}

	return func() {
		// best effort method of deleting all the files
		// nolint: errcheck
		_ = os.Remove("/tmp/cloud_keypair")
		// nolint: errcheck
		_ = os.Remove("/tmp/cloud_keypair.pub")
	}
}

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

func generatedTopoPath() string {
	return defaultWorkRoot + "/" + cloudID + "/topology.yml"
}

func executedCommandsPath() string {
	return defaultWorkRoot + "/" + cloudID + "/" + executedCmdTestFile
}

func verifyCloudDeleted(ctx context.Context, httpClient *client.HTTP) bool {
	if _, err := os.Stat(defaultWorkRoot + "/" + cloudID); err == nil {
		// working dir not deleted
		return false
	}

	_, err := httpClient.GetCloud(ctx,
		&services.GetCloudRequest{
			ID: cloudID,
		},
	)
	if err == nil {
		return false
	}
	return true
}

func isCloudSecretFilesDeleted() error {
	keyDefaults := services.NewKeyFileDefaults()
	errstrings := []string{}
	for _, secret := range []string{
		GetTerraformAWSPlanFile(cloudID),
		GetTerraformAzurePlanFile(cloudID),
		GetTerraformGCPPlanFile(cloudID),
		keyDefaults.GetAWSAccessPath(),
		keyDefaults.GetAWSSecretPath(),
		keyDefaults.GetAzureSubscriptionIDPath(),
		keyDefaults.GetAzureClientIDPath(),
		keyDefaults.GetAzureClientSecretPath(),
		keyDefaults.GetAzureTenantIDPath(),
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
	if len(errstrings) != 0 {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}
