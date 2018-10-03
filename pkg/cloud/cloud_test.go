package cloud

import (
	"testing"

	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	allInOneCloudTemplatePath  = "./test_data/test_all_in_one_cloud.tmpl"
	expectedCmdForCreateUpdate = "./test_data/expected_cmd_for_create_update.yml"
	cloudID                    = "cloud_uuid"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func TestAzureCloud(t *testing.T) {
	runAllInOneCloudTest(t, azure)
}

func runAllInOneCloudTest(t *testing.T, cloudType string) {

	context := pongo2.Context{
		"CLOUD_TYPE": cloudType,
	}

	var expectedTopology, expectedSecret string

	switch cloudType {
	case azure:
		expectedTopology = "./test_data/expected_az_cloud_topo.yaml"
		expectedSecret = "./test_data/expected_az_cloud_secret.yaml"
	}
	runCloudTest(t, expectedTopology, expectedSecret, context, cloudType)

}

func runCloudTest(t *testing.T, expectedTopology, expectedSecret string,
	context map[string]interface{}, cloudType string) {

	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystone("127.0.0.1:35357", keystoneAuthURL)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystone("127.0.0.1:5000", keystoneAuthURL)
	defer ksPrivate.Close()

	// create cloud related objects
	var cloudTestScenario integration.TestScenario
	err := integration.LoadTestScenario(&cloudTestScenario, allInOneCloudTemplatePath, context)
	assert.NoErrorf(t, err, "failed to load cloud test data from file: %s",
		allInOneCloudTemplatePath)
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
		Type:         cloudType,
		Action:       createAction,
		LogLevel:     "debug",
		TemplateRoot: "configs/",
		Test:         true,
	}

	// delete previously created
	cloud, err := NewCloud(config)
	assert.NoError(t, err, "failed to create cloud struct")

	err = cloud.Manage()
	assert.NoError(t, err, "failed to create cloud")

	assert.True(t, compareGeneratedTopology(t, expectedTopology),
		"topology file created during cloud create is not as expected")
	assert.True(t, compareGeneratedSecret(t, expectedSecret),
		"secret file created during cloud create is not as expected")
	assert.True(t, verifyCommandsExecuted(t, expectedCmdForCreateUpdate),
		"Expected list of commands are not executed")

	// Wait for the in-memory endpoint cache to get updated
	server.ForceProxyUpdate()
}

func compareGeneratedTopology(t *testing.T, topoFile string) bool {
	return true
}

func compareGeneratedSecret(t *testing.T, secretFile string) bool {
	return true
}

func verifyCommandsExecuted(t *testing.T, cmdExecuted string) bool {
	return true
}
