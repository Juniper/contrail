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

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type providerConfig struct {
	name                 string
	cloudUUID            string
	filesToCopy          []*fileToCopy
	requestsToStartWith  []string
	expectedTopologyFile string
	expectedSecretFile   string
	expectedCommandsFile string
	cloudAction          string
	manageFails          bool
}

type terraformStateReaderStub struct{}

func (r terraformStateReaderStub) Read() (terraformState, error) {
	return terraformStateStub{}, nil
}

type terraformStateStub struct{}

func (s terraformStateStub) GetPublicIP(hostname string) (string, error) {
	return "1.1.1.1", nil
}

func (s terraformStateStub) GetPrivateIP(hostname string) (string, error) {
	return "2.2.2.2", nil
}

type fileToCopy struct {
	source      string
	destination string
}

// TODO: Use relative path in every test that uses absolute paths.
func TestCreatingUpdatingPublicClouds(t *testing.T) {
	for _, providerConfig := range []*providerConfig{
		{
			name:        "Test Create AWS Cloud",
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
				"./test_data/aws/test_aws_create/requests.yml",
			},
			expectedTopologyFile: "./test_data/aws/test_aws_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/aws/expected_secret.yml",
			expectedCommandsFile: "./test_data/aws/expected_commands.yml",
		}, {
			name:        "Test Create AWS Cloud Failure",
			cloudUUID:   "cloud_uuid_aws",
			cloudAction: createAction,
			manageFails: true,
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/aws/test_aws_create/requests.yml",
			},
			expectedTopologyFile: "./test_data/aws/test_aws_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/aws/expected_secret.yml",
			expectedCommandsFile: "./test_data/aws/expected_commands.yml",
		}, {
			name:        "Test Create GCP Cloud",
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
				{
					source:      "./test_data/gcp/google-account.json",
					destination: "/var/tmp/contrail/google-account.json",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/gcp/test_gcp_create/requests.yml",
			},
			expectedTopologyFile: "./test_data/gcp/test_gcp_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/gcp/expected_secret.yml",
			expectedCommandsFile: "./test_data/gcp/expected_commands.yml",
		}, {
			name:        "Test Create Azure Cloud",
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
				"./test_data/azure/test_azure_create/requests.yml",
			},
			expectedTopologyFile: "./test_data/azure/test_azure_create/expected_topology.yml",
			expectedSecretFile:   "./test_data/azure/expected_secret.yml",
			expectedCommandsFile: "./test_data/azure/expected_commands.yml",
		}, {
			name:        "Test Update AWS Cloud",
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
				"./test_data/aws/test_aws_update/prerequisites.yml",
				"./test_data/aws/test_aws_update/requests.yml",
			},
			expectedTopologyFile: "./test_data/aws/test_aws_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/aws/expected_secret.yml",
			expectedCommandsFile: "./test_data/aws/expected_commands.yml",
		}, {
			name:        "Test Update AWS Cloud Fail",
			cloudUUID:   "cloud_uuid_aws",
			cloudAction: updateAction,
			manageFails: true,
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/aws/test_aws_update/prerequisites.yml",
				"./test_data/aws/test_aws_update/requests.yml",
			},
			expectedTopologyFile: "./test_data/aws/test_aws_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/aws/expected_secret.yml",
			expectedCommandsFile: "./test_data/aws/expected_commands.yml",
		}, {
			name:        "Test Update GCP Cloud",
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
				{
					source:      "./test_data/gcp/google-account.json",
					destination: "/var/tmp/contrail/google-account.json",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/gcp/test_gcp_update/prerequisites.yml",
				"./test_data/gcp/test_gcp_update/requests.yml",
			},
			expectedTopologyFile: "./test_data/gcp/test_gcp_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/gcp/expected_secret.yml",
			expectedCommandsFile: "./test_data/gcp/expected_commands.yml",
		}, {
			name:        "Test Update Azure Cloud",
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
				"./test_data/azure/test_azure_update/prerequisites.yml",
				"./test_data/azure/test_azure_update/requests.yml",
			},
			expectedTopologyFile: "./test_data/azure/test_azure_update/expected_topology.yml",
			expectedSecretFile:   "./test_data/azure/expected_secret.yml",
			expectedCommandsFile: "./test_data/azure/expected_commands.yml",
		},
	} {
		t.Run(providerConfig.name, func(t *testing.T) {
			testPublicCloudUpdate(t, providerConfig)
		})
		t.Run(providerConfig.name+" without removing secrets", func(t *testing.T) {
			testPublicCloudUpdateWithoutRemovingSecret(t, providerConfig)
		})
	}
}

func expectedSSHPrivKeyPath(cloudUUID string) string {
	return path.Join(GetCloudDir(cloudUUID), "cloud_keypair")
}

func expectedSSHPubKeyPath(cloudUUID string) string {
	return path.Join(GetCloudDir(cloudUUID), "cloud_keypair.pub")
}

func testPublicCloudUpdate(t *testing.T, pc *providerConfig) {
	postActions, usedScenarios := prepareForTest(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, pc.cloudAction)

	err := cl.Manage()

	if pc.manageFails {
		assert.Errorf(t, err, "manage cloud should fail, while cloud %s", pc.cloudAction)
		assert.False(t, assertModifiedStatusRemoval(cl.ctx, t, cl.APIServer, cl.config.CloudID),
			"modified status is removed")
		verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)
		return
	}

	assert.NoErrorf(t, err, "failed to manage cloud, while cloud %s", pc.cloudAction)
	assert.True(t, assertModifiedStatusRemoval(cl.ctx, t, cl.APIServer, cl.config.CloudID),
		"modified status is not removed")

	verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)

	assert.True(t, verifyNodeType(cl.ctx, cl.APIServer, usedScenarios),
		"public cloud nodes are not updated as type private")

	compareTopology(t, pc.expectedTopologyFile, cl.config.CloudID)

	verifyCommandsExecuted(t, pc.expectedCommandsFile, cl.config.CloudID)
	verifyGeneratedSSHKeyFiles(t, cl.config.CloudID)
}

func assertModifiedStatusRemoval(ctx context.Context, t *testing.T, APIServer *client.HTTP, cloudUUID string) bool {
	c, err := APIServer.GetCloud(ctx, &services.GetCloudRequest{
		ID: cloudUUID,
	})
	assert.NoError(t, err)
	return !(c.Cloud.AwsModified || c.Cloud.AzureModified || c.Cloud.GCPModified)
}

func prepareForTest(
	t *testing.T, requestsToStartWith []string, hostFiles []*fileToCopy, cloudUUID string,
) (postActions func(t *testing.T, cloudUUID string), usedScenarios []*integration.TestScenario) {
	clearTestDirectory(t, cloudUUID)

	removeFiles := copyFiles(t, hostFiles)

	cleanDB, usedScenarios, err := doAPIRequests(t, requestsToStartWith)

	cleanTestEnvironment := func(t *testing.T, cloudUUID string) {
		cleanDB()
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

func doAPIRequests(
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
		cleanupFunctions = append([]func(){cleanDB}, cleanupFunctions...)
	}
	return func() {
		for _, fun := range cleanupFunctions {
			fun()
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

	cl, err := NewCloud(c, terraformStateReaderStub{})
	assert.NoError(t, err, "failed to create cloud struct")

	return cl
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
		if err == nil {
			assert.Failf(t, "File %s is not deleted", secret)
		} else if !os.IsNotExist(err) {
			assert.Failf(t, "Could not verify if %s file is deleted: %v", secret, err)
		}
	}
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

func compareTopology(t *testing.T, expectedTopologyFile, cloudUUID string) {
	assertYAMLFileEqual(t, expectedTopologyFile, GetTopoFile(cloudUUID),
		"generated topology file is not as expected")
}

func assertYAMLFileEqual(t *testing.T, expectedFilePath, actualFilePath string, msg string) {
	var actualYAML interface{}
	require.NoErrorf(t, fileutil.LoadFile(actualFilePath, &actualYAML), "Failed read yaml from %s", actualFilePath)

	var expectedYAML interface{}
	require.NoErrorf(t, fileutil.LoadFile(expectedFilePath, &expectedYAML), "Failed read yaml from %s", expectedFilePath)

	testutil.AssertEqual(t, expectedYAML, actualYAML,
		fmt.Sprintf("YAML files %s and %s are not equal", expectedFilePath, actualFilePath), msg)
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
	assertYAMLFileEqual(t, expectedCmdFile, executedCommandsPath(cloudUUID),
		"Expected list of create commands are not executed")
}

func executedCommandsPath(cloudUUID string) string {
	return path.Join(defaultWorkRoot, cloudUUID, executedCmdTestFile)
}

func verifyGeneratedSSHKeyFiles(t *testing.T, cloudUUID string) {
	pvtKeyPath := getCloudSSHKeyPath(cloudUUID, "cloud_keypair")
	pubKeyPath := getCloudSSHKeyPath(cloudUUID, "cloud_keypair.pub")
	expectedPvtKeyPath := expectedSSHPrivKeyPath(cloudUUID)
	expectedPubKeyPath := expectedSSHPubKeyPath(cloudUUID)
	assert.NoError(t, compareFiles(t, expectedPvtKeyPath, pvtKeyPath), "Private SSH keys doesn't match!")
	assert.NoError(t, compareFiles(t, expectedPubKeyPath, pubKeyPath), "Public SSH keys doesn't match!")
}

func testPublicCloudUpdateWithoutRemovingSecret(t *testing.T, pc *providerConfig) {
	postActions, _ := prepareForTest(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, pc.cloudAction)
	err := cl.manage()

	if pc.manageFails {
		assert.Errorf(t, err, "manage cloud should fail, while cloud %s", pc.cloudAction)
		assert.False(t, assertModifiedStatusRemoval(cl.ctx, t, cl.APIServer, cl.config.CloudID),
			"modified status is removed")
		return
	}

	assert.NoErrorf(t, cl.manage(), "failed to manage cloud, while cloud %s", pc.cloudAction)
	assert.True(t, assertModifiedStatusRemoval(cl.ctx, t, cl.APIServer, cl.config.CloudID),
		"modified status is not removed")
	compareSecret(t, pc.expectedSecretFile, cl.config.CloudID)
}

func compareSecret(t *testing.T, expectedSecretFile, cloudUUID string) {
	assertYAMLFileEqual(t, expectedSecretFile, GetSecretFile(cloudUUID),
		"generated secret file is not as expected")
}

func TestDeletingPublicClouds(t *testing.T) {
	for _, providerConfig := range []*providerConfig{
		{
			name:      "Delete AWS Cloud",
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
				"./test_data/aws/test_aws_delete/prerequisites.yml",
				"./test_data/aws/test_aws_delete/requests.yml",
			},
			expectedTopologyFile: "./test_data/aws/test_aws_delete/expected_topology.yml",
			expectedSecretFile:   "./test_data/aws/expected_secret.yml",
			expectedCommandsFile: "./test_data/aws/expected_commands.yml",
		}, {
			name:        "Delete AWS Cloud Fail",
			cloudUUID:   "cloud_uuid_aws",
			manageFails: true,
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/aws/test_aws_delete/prerequisites.yml",
				"./test_data/aws/test_aws_delete/requests.yml",
			},
			expectedTopologyFile: "./test_data/aws/test_aws_delete/expected_topology.yml",
			expectedSecretFile:   "./test_data/aws/expected_secret.yml",
			expectedCommandsFile: "./test_data/aws/expected_commands.yml",
		}, {
			name:      "Delete GCP Cloud",
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
				{
					source:      "./test_data/gcp/google-account.json",
					destination: "/var/tmp/contrail/google-account.json",
				},
			},
			requestsToStartWith: []string{
				"./test_data/cluster_with_credentials_request.yml",
				"./test_data/gcp/test_gcp_delete/prerequisites.yml",
				"./test_data/gcp/test_gcp_delete/requests.yml",
			},
			expectedTopologyFile: "./test_data/gcp/test_gcp_delete/expected_topology.yml",
			expectedSecretFile:   "./test_data/gcp/expected_secret.yml",
			expectedCommandsFile: "./test_data/gcp/expected_commands.yml",
		}, {
			name:      "Delete Azure Cloud",
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
				"./test_data/azure/test_azure_delete/prerequisites.yml",
				"./test_data/azure/test_azure_delete/requests.yml",
			},
			expectedTopologyFile: "./test_data/azure/test_azure_delete/expected_topology.yml",
			expectedSecretFile:   "./test_data/azure/expected_secret.yml",
			expectedCommandsFile: "./test_data/azure/expected_commands.yml",
		},
	} {
		t.Run(providerConfig.name, func(t *testing.T) {
			testPublicCloudDeletion(t, providerConfig)
		})
	}
}

func testPublicCloudDeletion(t *testing.T, pc *providerConfig) {
	postActions, _ := prepareForTest(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, updateAction)

	err := cl.Manage()
	if pc.manageFails {
		assert.Errorf(t, err, "manage cloud should fail, while deleting cloud", pc.cloudAction)
		assert.False(t, assertModifiedStatusRemoval(cl.ctx, t, cl.APIServer, cl.config.CloudID),
			"modified status is removed")
		verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)
		return
	}

	assert.NoErrorf(t, err, "failed to manage cloud, while deleting", pc.cloudAction)

	verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)

	verifyCloudDeleted(cl.ctx, t, cl.APIServer, pc.cloudUUID)
}

func verifyCloudDeleted(ctx context.Context, t *testing.T, httpClient *client.HTTP, cloudUUID string) {
	_, err := os.Stat(defaultWorkRoot + "/" + cloudUUID)
	assert.Error(t, err, "The cloud directory exists on host and it shouldn't")
	httpResp, err := httpClient.Read(ctx, "/cloud/"+cloudUUID, &services.GetCloudResponse{})
	assert.Error(t, err, "HTTP get request should return an error due to non existing cloud")
	assert.NotNil(t, httpResp.StatusCode, http.StatusNotFound, "Couldn't verify if Cloud doesn't exist in DB")
	assert.Equal(t, httpResp.StatusCode, http.StatusNotFound, "The cloud still exists in DB and it shouldn't")
}

func TestOnPremUpdate(t *testing.T) {
	for _, providerConfig := range []*providerConfig{
		{
			name:        "Create OnPrem Cloud",
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
				"./test_data/onprem/create_cloud_resources.yml",
			},
			expectedTopologyFile: "./test_data/onprem/test_onprem_create/expected_topology.yml",
		},
		{
			name:        "Update OnPrem Cloud",
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
				"./test_data/onprem/create_cloud_resources.yml",
				"./test_data/onprem/test_onprem_update/requests.yml",
			},
			expectedTopologyFile: "./test_data/onprem/test_onprem_update/expected_topology.yml",
		},
	} {
		t.Run(providerConfig.name, func(t *testing.T) {
			testOnPremUpdate(t, providerConfig)
		})
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
			name:        "Delete OnPrem Cloud (unsuccessfully)",
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
				"./test_data/onprem/create_cloud_resources.yml",
				"./test_data/onprem/test_onprem_delete/invalid_requests.yml",
			},
		},
		{
			name:        "Delete OnPrem Cloud",
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
				"./test_data/onprem/create_cloud_resources.yml",
				"./test_data/onprem/test_onprem_delete/requests.yml",
			},
		},
	} {
		t.Run(providerConfig.name, func(t *testing.T) {
			testOnPremDelete(t, providerConfig)
		})
	}
}

func testOnPremDelete(t *testing.T, pc *providerConfig) {
	postActions, _ := prepareForTest(t, pc.requestsToStartWith, pc.filesToCopy, pc.cloudUUID)
	defer postActions(t, pc.cloudUUID)

	cl := prepareCloud(t, pc.cloudUUID, pc.cloudAction)

	if pc.manageFails {
		assert.Error(t, cl.Manage(), "manage cloud succeeded but it shouldn't")
	} else {
		assert.NoError(t, cl.Manage(), "failed to manage cloud, while deleting cloud")
		verifyCloudDeleted(cl.ctx, t, cl.APIServer, pc.cloudUUID)
	}

	verifyCloudSecretFilesAreDeleted(t, cl.config.CloudID)
}
