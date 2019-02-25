package cloud

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	allInOneCloudTemplatePath       = "./test_data/test_all_in_one_public_cloud.tmpl"
	allInOneCloudDeleteTemplatePath = "./test_data/test_all_in_one_public_cloud_delete.tmpl"
	allInOneCloudUpdateTemplatePath = "./test_data/test_all_in_one_public_cloud_update.tmpl"
	expectedAZCmdForCreateUpdate    = "./test_data/expected_azure_cmd_for_create_update.yaml"
	expectedAZTopology              = "./test_data/expected_azure_cloud_topology.yaml"
	expectedAZSecret                = "./test_data/expected_azure_cloud_secret.yaml"
	expectedAWSCmdForCreateUpdate   = "./test_data/expected_aws_cmd_for_create_update.yaml"
	expectedAWSTopology             = "./test_data/expected_aws_cloud_topology.yaml"
	expectedAWSSecret               = "./test_data/expected_aws_cloud_secret.yaml"
	expectedOnPremTopology          = "./test_data/expected_onprem_cloud_topology.yaml"
	expectedPvtKey                  = "./test_data/cloud_keypair"
	cloudID                         = "cloud_uuid"
	defaultAdminUser                = "admin"
	defaultAdminPassword            = "contrail123"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func TestOnPremCloud(t *testing.T) {
	context := pongo2.Context{
		"CLOUD_TYPE": onPrem,
	}
	runCloudTest(t, expectedOnPremTopology, "", "", context)
}
func TestAzureCloud(t *testing.T) {
	context := pongo2.Context{
		"CLOUD_TYPE": azure,
	}
	runCloudTest(t, expectedAZTopology, expectedAZSecret,
		expectedAZCmdForCreateUpdate, context)
}

func TestAWSCloud(t *testing.T) {
	context := pongo2.Context{
		"CLOUD_TYPE": aws,
	}

	runCloudTest(t, expectedAWSTopology, expectedAWSSecret,
		expectedAWSCmdForCreateUpdate, context)
}

func runCloudTest(t *testing.T, expectedTopology, expectedSecret string,
	expectedCmdForCreateUpdate string, context map[string]interface{}) {

	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:35357", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"127.0.0.1:5000", keystoneAuthURL, defaultAdminUser, defaultAdminPassword)
	defer ksPrivate.Close()

	// create cloud related objects
	var cloudTestScenario integration.TestScenario
	err := integration.LoadTestScenario(&cloudTestScenario, allInOneCloudTemplatePath, context)
	assert.NoErrorf(t, err, "failed to load cloud test data from file: %s", allInOneCloudTemplatePath)
	cleanup := integration.RunDirtyTestScenario(t, &cloudTestScenario, server)

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
		LogLevel:     "debug",
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

	pvtKeycleanup := createDummyPvtKeyFile(t)
	defer pvtKeycleanup()

	cloud, err := NewCloud(config)
	assert.NoError(t, err, "failed to create cloud struct")

	err = cloud.Manage()
	assert.NoError(t, err, "failed to manage cloud, while creating cloud")

	assert.True(t, compareGeneratedTopology(t, expectedTopology),
		"topology file created during cloud create is not as expected")
	if context["CLOUD_TYPE"] == onPrem {
		return
	}
	assert.True(t, compareGeneratedSecret(t, expectedSecret),
		"secret file created during cloud create is not as expected")
	assert.True(t, verifyCommandsExecuted(t, expectedCmdForCreateUpdate),
		"Expected list of create commands are not executed")
	// check if ssh keys are created
	assert.True(t, verifyGeneratedSSHKeyFile(t),
		"Expected ssh key file are not generated")

	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()

	//update cloud
	config.Action = updateAction

	var cloudUpdateTestScenario integration.TestScenario
	err = integration.LoadTestScenario(&cloudUpdateTestScenario, allInOneCloudUpdateTemplatePath, context)
	assert.NoErrorf(t, err, "failed to load cloud test data from file: %s", allInOneCloudUpdateTemplatePath)
	_ = integration.RunDirtyTestScenario(t, &cloudUpdateTestScenario, server)

	// delete previously created files

	// Remove topology file and secret file
	err = os.Remove(generatedTopoPath())
	if err != nil {
		assert.NoError(t, err, "failed to delete topology.yml file, during update")
	}
	err = os.Remove(generatedSecretPath())
	if err != nil {
		assert.NoError(t, err, "failed to delete secret.yml file, during update")
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

	err = cloud.Manage()
	assert.NoError(t, err, "failed to manage cloud, while updating cloud")

	assert.True(t, compareGeneratedTopology(t, expectedTopology),
		"topology file created during cloud update is not as expected")
	assert.True(t, compareGeneratedSecret(t, expectedSecret),
		"secret file created during cloud update is not as expected")
	assert.True(t, verifyCommandsExecuted(t, expectedCmdForCreateUpdate),
		"Expected list of update commands are not executed")

	// delete cloud
	var cloudDeleteTestScenario integration.TestScenario
	err = integration.LoadTestScenario(&cloudDeleteTestScenario, allInOneCloudDeleteTemplatePath, context)
	assert.NoErrorf(t, err, "failed to load cloud test data from file: %s", allInOneCloudDeleteTemplatePath)
	_ = integration.RunDirtyTestScenario(t, &cloudDeleteTestScenario, server)

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

	err = cloud.Manage()
	assert.NoError(t, err, "failed to manage cloud, while deleting cloud")

	// make sure cloud is removed
	assert.True(t, verifyCloudDeleted(), "Topo/Secret file is not deleted during cloud delete")

}

func compareFiles(t *testing.T, expectedFile, generatedFile string) bool {
	generatedData, err := ioutil.ReadFile(generatedFile)
	assert.NoErrorf(t, err, "unable to read generated: %s", generatedFile)
	expectedData, err := ioutil.ReadFile(expectedFile)
	assert.NoErrorf(t, err, "unable to read expected: %s", expectedFile)
	return bytes.Equal(generatedData, expectedData)

}

func compareGeneratedTopology(t *testing.T, expectedTopology string) bool {
	return compareFiles(t, expectedTopology, generatedTopoPath())
}

func compareGeneratedSecret(t *testing.T, expectedSecretFile string) bool {
	return compareFiles(t, expectedSecretFile, generatedSecretPath())
}

func verifyCommandsExecuted(t *testing.T, expectedCmdForCreateUpdate string) bool {
	return compareFiles(t, expectedCmdForCreateUpdate, executedCommandsPath())
}

func verifyGeneratedSSHKeyFile(t *testing.T) bool {
	pvtKeyPath := getCloudSSHKeyPath(cloudID, "cloud_keypair")
	return compareFiles(t, expectedPvtKey, pvtKeyPath)
}

func createDummyPvtKeyFile(t *testing.T) func() {
	// create public cloud topology.yaml
	publicTopoData, err := fileutil.GetContent("file://" + expectedPvtKey)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to read file: %s", expectedPvtKey)
	}
	err = fileutil.WriteToFile("/tmp/cloud_keypair", publicTopoData, defaultRWOnlyPerm)
	if err != nil {
		assert.NoErrorf(t, err, "Unable to write file: %s", "/tmp/cloud_keypair")
	}

	return func() {
		// best effort method of deleting all the files
		// nolint: errcheck
		_ = os.Remove("/tmp/cloud_keypair")
	}
}

func generatedTopoPath() string {
	return defaultWorkRoot + "/" + cloudID + "/topology.yml"
}

//new-comment

func generatedSecretPath() string {
	return defaultWorkRoot + "/" + cloudID + "/secret.yml"
}

func executedCommandsPath() string {
	return defaultWorkRoot + "/" + cloudID + "/" + executedCmdTestFile
}

func verifyCloudDeleted() bool {

	if _, err := os.Stat(defaultWorkRoot + "/" + cloudID); err == nil {
		// working dir not deleted
		return false
	}
	return true

}
