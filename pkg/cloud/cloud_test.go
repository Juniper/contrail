package cloud

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type providerConfig struct {
	cloudUUID            string
	filesToCopy          []*fileToCopy
	requestsToStartWith  []string
	expectedTopologyFile string
	expectedSecretFile   string
	expectedCommandsFile string
	cloudAction          string
	manageFails          bool
}

type fileToCopy struct {
	source      string
	destination string
}

func TestCreatingUpdatingPublicClouds(t *testing.T) {
	for _, providerConfig := range []*providerConfig{
		{
			cloudUUID:   "cloud_uuid_aws",
			cloudAction: createAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_aws"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_aws"),
				},
				{
					source:      "./test_data/aws/aws_access.key",
					destination: "/var/tmp/contrail/aws_access.key",
				},
				{
					source:      "./test_data/aws/aws_secret.key",
					destination: "/var/tmp/contrail/aws_secret.key",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/aws/sample_create_requests.yml",
			},
			expectedTopologyFile: "./test_data/aws/test_aws_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/aws/expected_secret.yml",
			expectedCommandsFile: "./test_data/aws/expected_commands.yml",
		},
		{
			cloudUUID:   "cloud_uuid_gcp",
			cloudAction: createAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_gcp"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_gcp"),
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/gcp/sample_create_requests.yml",
			},
			expectedTopologyFile: "./test_data/gcp/test_gcp_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/gcp/expected_secret.yml",
			expectedCommandsFile: "./test_data/gcp/expected_commands.yml",
		},
		{
			cloudUUID:   "cloud_uuid_azure",
			cloudAction: createAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_azure"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_azure"),
				},
				{
					source:      "./test_data/azure/subscription_id",
					destination: "/var/tmp/contrail/subscription_id",
				},
				{
					source:      "./test_data/azure/client_id",
					destination: "/var/tmp/contrail/client_id",
				},
				{
					source:      "./test_data/azure/client_secret",
					destination: "/var/tmp/contrail/client_secret",
				},
				{
					source:      "./test_data/azure/tenant_id",
					destination: "/var/tmp/contrail/tenant_id",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/azure/sample_create_requests.yml",
			},
			expectedTopologyFile: "./test_data/azure/test_azure_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/azure/expected_secret.yml",
			expectedCommandsFile: "./test_data/azure/expected_commands.yml",
		},
		{
			cloudUUID:   "cloud_uuid_aws",
			cloudAction: updateAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_aws"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_aws"),
				},
				{
					source:      "./test_data/aws/test_aws_update/topology_before_update.yml",
					destination: "/var/tmp/cloud/cloud_uuid_aws/topology.yml",
				},
				{
					source:      "./test_data/aws/aws_access.key",
					destination: "/var/tmp/contrail/aws_access.key",
				},
				{
					source:      "./test_data/aws/aws_secret.key",
					destination: "/var/tmp/contrail/aws_secret.key",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/aws/sample_create_requests.yml",
				"./test_data/aws/test_aws_update/requests.yml",
			},
			expectedTopologyFile: "./test_data/aws/test_aws_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/aws/expected_secret.yml",
			expectedCommandsFile: "./test_data/aws/expected_commands.yml",
		},
		{
			cloudUUID:   "cloud_uuid_gcp",
			cloudAction: updateAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_gcp"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_gcp"),
				},
				{
					source:      "./test_data/gcp/test_gcp_update/topology_before_update.yml",
					destination: "/var/tmp/cloud/cloud_uuid_gcp/topology.yml",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/gcp/sample_create_requests.yml",
				"./test_data/gcp/test_gcp_update/requests.yml",
			},
			expectedTopologyFile: "./test_data/gcp/test_gcp_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/gcp/expected_secret.yml",
			expectedCommandsFile: "./test_data/gcp/expected_commands.yml",
		},
		{
			cloudUUID:   "cloud_uuid_azure",
			cloudAction: updateAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_azure"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_azure"),
				},
				{
					source:      "./test_data/azure/subscription_id",
					destination: "/var/tmp/contrail/subscription_id",
				},
				{
					source:      "./test_data/azure/client_id",
					destination: "/var/tmp/contrail/client_id",
				},
				{
					source:      "./test_data/azure/client_secret",
					destination: "/var/tmp/contrail/client_secret",
				},
				{
					source:      "./test_data/azure/tenant_id",
					destination: "/var/tmp/contrail/tenant_id",
				},
				{
					source:      "./test_data/azure/test_azure_update/topology_before_update.yml",
					destination: "/var/tmp/cloud/cloud_uuid_azure/topology.yml",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/azure/sample_create_requests.yml",
				"./test_data/azure/test_azure_update/requests.yml",
			},
			expectedTopologyFile: "./test_data/azure/test_azure_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/azure/expected_secret.yml",
			expectedCommandsFile: "./test_data/azure/expected_commands.yml",
		},
	} {
		testPublicCloudUpdate(t, providerConfig)
		testPublicCloudUpdateWithoutRemovingSecret(t, providerConfig)
	}
}

func prepareForTest(
	t *testing.T, dbRequestsFiles []string, hostFiles []*fileToCopy, cloudUUID string,
) (postActions func(t *testing.T, cloudUUID string), usedScenarios []*integration.TestScenario) {
	clearTestDirectory(t, cloudUUID)

	removeFiles := copyFiles(t, hostFiles)

	stopServer := authenticate()

	cleanDB, usedScenarios, err := applyDBRequests(t, dbRequestsFiles)

	cleanTestEnvironment := func(t *testing.T, cloudUUID string) {
		cleanDB()
		stopServer()
		removeFiles()
		removeTestDirectory(t, cloudUUID)
	}

	if err != nil {
		cleanTestEnvironment(t, cloudUUID)
		assert.Fail(t, "Test failed due to error with loading requests", "Error: %v", err)
	}
	return cleanTestEnvironment, usedScenarios
}

func clearTestDirectory(t *testing.T, cloudUUID string) {
	removeTestDirectory(t, cloudUUID)
	assert.NoError(t, os.MkdirAll(GetCloudDir(cloudUUID), 0777), "Cannot clean test environment")
}

func removeTestDirectory(t *testing.T, cloudUUID string) {
	err := os.RemoveAll(GetCloudDir(cloudUUID))
	if err == nil || os.IsNotExist(err) {
		return
	}
	assert.Fail(t, "Cannot clean test environment: %v", err)
}

func copyFiles(t *testing.T, files []*fileToCopy) (cleanup func()) {
	filesToRemove := []string{}
	for _, file := range files {
		filesToRemove = append(filesToRemove, file.destination)
		err := fileutil.CopyFile(file.source, file.destination, true)
		assert.NoErrorf(t, err, "Failed to copy file %s", file.source)
	}
	return func() {
		for _, file := range filesToRemove {
			if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
				assert.Fail(t, "Unable to remove file", err.Error())
			}
		}
	}
}

func authenticate() func() {
	defaultAdminUser := "admin"
	defaultAdminPassword := "contrail123"
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

func applyDBRequests(
	t *testing.T, files []string,
) (cleanup func(), usedScenarios []*integration.TestScenario, err error) {
	cleanupFunctions := []func(){}
	for _, file := range files {
		var ts *integration.TestScenario
		ts, err = integration.LoadTest(file, nil)
		if err != nil {
			return cleanup, nil, err
		}
		usedScenarios = append(usedScenarios, ts)
		cleanDB := integration.RunDirtyTestScenario(t, ts, server)
		cleanupFunctions = append(cleanupFunctions, cleanDB)
	}
	return func() {
		for id := range cleanupFunctions {
			cleanupFunctions[len(cleanupFunctions)-id-1]()
		}
	}, usedScenarios, nil
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

func testPublicCloudUpdate(t *testing.T, pc *providerConfig) {
	postActions, usedScenarios := prepareForTest(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, pc.cloudAction)

	assert.NoErrorf(t, cl.Manage(), "failed to manage cloud, while cloud %s", pc.cloudAction)

	verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)

	assert.True(t, verifyNodeType(cl.ctx, cl.APIServer, usedScenarios),
		"public cloud nodes are not updated as type private")

	compareTopology(t, pc.expectedTopologyFile, cl.config.CloudID)

	verifyCommandsExecuted(t, pc.expectedCommandsFile, cl.config.CloudID)
	verifyGeneratedSSHKeyFiles(t, cl.config.CloudID)
}

func testPublicCloudUpdateWithoutRemovingSecret(t *testing.T, pc *providerConfig) {
	postActions, _ := prepareForTest(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, pc.cloudAction)

	assert.NoErrorf(t, cl.manage(), "failed to manage cloud, while cloud %s", pc.cloudAction)

	compareSecret(t, pc.expectedSecretFile, cl.config.CloudID)
}

func TestDeletingPublicClouds(t *testing.T) {
	for _, providerConfig := range []*providerConfig{
		{
			cloudUUID: "cloud_uuid_aws",
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_aws"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_aws"),
				},
				{
					source:      "./test_data/aws/test_aws_delete/topology_before_delete.yml",
					destination: "/var/tmp/cloud/cloud_uuid_aws/topology.yml",
				},
				{
					source:      "./test_data/aws/aws_access.key",
					destination: "/var/tmp/contrail/aws_access.key",
				},
				{
					source:      "./test_data/aws/aws_secret.key",
					destination: "/var/tmp/contrail/aws_secret.key",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/aws/sample_create_requests.yml",
				"./test_data/aws/test_aws_delete/requests.yml",
			},
			expectedTopologyFile: "./test_data/aws/test_aws_delete/expected_topology.yml",
			expectedSecretFile:   "./test_data/aws/expected_secret.yml",
			expectedCommandsFile: "./test_data/aws/expected_commands.yml",
		},
		{
			cloudUUID: "cloud_uuid_gcp",
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_gcp"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_gcp"),
				},
				{
					source:      "./test_data/gcp/test_gcp_delete/topology_before_delete.yml",
					destination: "/var/tmp/cloud/cloud_uuid_gcp/topology.yml",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/gcp/sample_create_requests.yml",
				"./test_data/gcp/test_gcp_delete/requests.yml",
			},
			expectedTopologyFile: "./test_data/gcp/test_gcp_delete/expected_topology.yml",
			expectedSecretFile:   "./test_data/gcp/expected_secret.yml",
			expectedCommandsFile: "./test_data/gcp/expected_commands.yml",
		},
		{
			cloudUUID: "cloud_uuid_azure",
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_azure"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_azure"),
				},
				{
					source:      "./test_data/azure/subscription_id",
					destination: "/var/tmp/contrail/subscription_id",
				},
				{
					source:      "./test_data/azure/client_id",
					destination: "/var/tmp/contrail/client_id",
				},
				{
					source:      "./test_data/azure/client_secret",
					destination: "/var/tmp/contrail/client_secret",
				},
				{
					source:      "./test_data/azure/tenant_id",
					destination: "/var/tmp/contrail/tenant_id",
				},
				{
					source:      "./test_data/azure/test_azure_delete/topology_before_delete.yml",
					destination: "/var/tmp/cloud/cloud_uuid_azure/topology.yml",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/azure/sample_create_requests.yml",
				"./test_data/azure/test_azure_delete/requests.yml",
			},
			expectedTopologyFile: "./test_data/azure/test_azure_delete/expected_topology.yml",
			expectedSecretFile:   "./test_data/azure/expected_secret.yml",
			expectedCommandsFile: "./test_data/azure/expected_commands.yml",
		},
	} {
		testPublicCloudDeletion(t, providerConfig)
	}
}

func testPublicCloudDeletion(t *testing.T, pc *providerConfig) {
	postActions, _ := prepareForTest(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, updateAction)

	assert.NoError(t, cl.Manage(), "failed to manage cloud, while deleting cloud")

	verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)

	verifyCloudDeleted(cl.ctx, t, cl.APIServer, pc.cloudUUID)
}

func TestOnPremUpdate(t *testing.T) {
	for _, providerConfig := range []*providerConfig{
		{
			cloudUUID:   "cloud_uuid_onprem",
			cloudAction: createAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_onprem"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_onprem"),
				},
			},
			requestsToStartWith: []string{
				"./test_data/onprem/sample_create_requests.yml",
			},
			expectedTopologyFile: "./test_data/onprem/test_onprem_create/expected_topology.yml",
		},
		{
			cloudUUID:   "cloud_uuid_onprem",
			cloudAction: updateAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_onprem"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_onprem"),
				},
			},
			requestsToStartWith: []string{
				"./test_data/onprem/sample_create_requests.yml",
				"./test_data/onprem/test_onprem_update/requests.yml",
			},
			expectedTopologyFile: "./test_data/onprem/test_onprem_update/expected_topology.yml",
		},
	} {
		testOnPremUpdate(t, providerConfig)
	}
}

func testOnPremUpdate(t *testing.T, pc *providerConfig) {
	postActions, _ := prepareForTest(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, pc.cloudAction)
	assert.NoErrorf(t, cl.Manage(), "failed to manage cloud, while cloud %s", pc.cloudAction)

	verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)
	compareTopology(t, pc.expectedTopologyFile, cl.config.CloudID)
}

func TestOnPremDelete(t *testing.T) {
	for _, providerConfig := range []*providerConfig{
		{
			cloudUUID:   "cloud_uuid_onprem",
			manageFails: true,
			cloudAction: updateAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_onprem"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_onprem"),
				},
			},
			requestsToStartWith: []string{
				"./test_data/onprem/sample_create_requests.yml",
				"./test_data/onprem/test_onprem_delete/invalid_requests.yml",
			},
		},
		{
			cloudUUID:   "cloud_uuid_onprem",
			cloudAction: updateAction,
			filesToCopy: []*fileToCopy{
				{
					source:      "./test_data/cloud_keypair",
					destination: expectedSSHPrivKeyPath("cloud_uuid_onprem"),
				},
				{
					source:      "./test_data/cloud_keypair.pub",
					destination: expectedSSHPubKeyPath("cloud_uuid_onprem"),
				},
			},
			requestsToStartWith: []string{
				"./test_data/onprem/sample_create_requests.yml",
				"./test_data/onprem/test_onprem_delete/requests.yml",
			},
		},
	} {
		testOnPremDelete(t, providerConfig)
	}
}

func testOnPremDelete(t *testing.T, pc *providerConfig) {
	postActions, _ := prepareForTest(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, pc.cloudAction)

	if pc.manageFails {
		assert.Error(t, cl.Manage(), "manage cloud succeded but it shouldn't")
	} else {
		assert.NoError(t, cl.Manage(), "failed to manage cloud, while deleting cloud")
		verifyCloudDeleted(cl.ctx, t, cl.APIServer, pc.cloudUUID)
	}

	verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)
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

func compareFiles(t *testing.T, expectedFile, generatedFile string) error {
	generatedData, err := ioutil.ReadFile(generatedFile)
	assert.NoErrorf(t, err, "unable to read generated: %s", generatedFile)
	expectedData, err := ioutil.ReadFile(expectedFile)
	assert.NoErrorf(t, err, "unable to read expected: %s", expectedFile)
	if !bytes.Equal(generatedData, expectedData) {
		return fmt.Errorf("expected:\n%s\n\nActual:\n%s", expectedData, generatedData)
	}
	return nil
}

func verifyCommandsExecuted(t *testing.T, expectedCmdFile string, cloudUUID string) {
	assert.NoError(t, compareFiles(t, expectedCmdFile, executedCommandsPath(cloudUUID)),
		"Expected list of create commands are not executed")
}

func verifyGeneratedSSHKeyFiles(t *testing.T, cloudUUID string) {
	pvtKeyPath := getCloudSSHKeyPath(cloudUUID, "cloud_keypair")
	pubKeyPath := getCloudSSHKeyPath(cloudUUID, "cloud_keypair.pub")
	expectedPvtKeyPath := expectedSSHPrivKeyPath(cloudUUID)
	expectedPubKeyPath := expectedSSHPubKeyPath(cloudUUID)
	assert.NoError(t, compareFiles(t, expectedPvtKeyPath, pvtKeyPath), "Private SSH keys doesn't match!")
	assert.NoError(t, compareFiles(t, expectedPubKeyPath, pubKeyPath), "Public SSH keys doesn't match!")
}

func executedCommandsPath(cloudUUID string) string {
	return path.Join(defaultWorkRoot, cloudUUID, executedCmdTestFile)
}

func verifyCloudDeleted(ctx context.Context, t *testing.T, httpClient *client.HTTP, cloudUUID string) {
	_, err := os.Stat(defaultWorkRoot + "/" + cloudUUID)
	assert.Error(t, err, "The cloud directory exists on host and it shouldn't")
	httpResp, err := httpClient.Read(ctx, "/cloud/"+cloudUUID, &services.GetCloudResponse{})
	assert.Error(t, err, "HTTP get request should return an error due to non existing cloud")
	assert.NotNil(t, httpResp.StatusCode, http.StatusNotFound, "Couldn't verify if Cloud doesn't exist in DB")
	assert.Equal(t, httpResp.StatusCode, http.StatusNotFound, "The cloud still exists in DB and it shouldn't")
}

func verifyNodeType(ctx context.Context, httpClient *client.HTTP, testScenarios []*integration.TestScenario) bool {
	for _, ts := range testScenarios {
		for _, task := range ts.Workflow {
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
	}
	return true
}

func verifyCloudSecretFilesAreDeleted(t *testing.T, cloudUUID string) {
	keyDefaults := services.NewKeyFileDefaults()
	for _, secret := range []string{
		GetTerraformAWSPlanFile(cloudUUID),
		GetTerraformAzurePlanFile(cloudUUID),
		GetTerraformGCPPlanFile(cloudUUID),
		keyDefaults.GetAWSAccessPath(),
		keyDefaults.GetAWSSecretPath(),
		keyDefaults.GetAzureSubscriptionIDPath(),
		keyDefaults.GetAzureClientIDPath(),
		keyDefaults.GetAzureClientSecretPath(),
		keyDefaults.GetAzureTenantIDPath(),
		keyDefaults.GetGoogleAccountPath(),
	} {
		_, err := os.Stat(secret)
		assert.True(t, err != nil, "File %s is not deleted", secret)
		assert.True(t, os.IsNotExist(err), "Could not verify if %s file is deleted", secret)
	}
}

func expectedSSHPrivKeyPath(cloudUUID string) string {
	return path.Join(GetCloudDir(cloudUUID), "cloud_keypair")
}

func expectedSSHPubKeyPath(cloudUUID string) string {
	return path.Join(GetCloudDir(cloudUUID), "cloud_keypair.pub")
}
