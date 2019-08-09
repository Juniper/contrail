package services_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	uuid "github.com/satori/go.uuid"
)

func TestUploadCloudKeys(t *testing.T) {
	outputDirectory := "/tmp/test_upload_cloud_keys"

	for _, tt := range []struct {
		name              string
		cloudProviderUUID string
		awsAccessKey      string
		awsSecretKey      string
		azureAccessToken  string
		azureProfile      string
		googleAccount     string
		statusCode        int
	}{
		{
			name:              "upload aws/azure/gcp keys",
			cloudProviderUUID: uuid.NewV4().String(),
			awsAccessKey:      "test access key",
			awsSecretKey:      "test secret key",
			azureAccessToken:  "{\"test\": \"access token\"}",
			azureProfile:      "{\"test\": \"profile\"}",
			googleAccount:     "{\"test\": \"account\"}",
		},
		{
			name:              "upload aws keys",
			cloudProviderUUID: uuid.NewV4().String(),
			awsAccessKey:      "test access key",
			awsSecretKey:      "test secret key",
		},
		{
			name:              "upload goolgle keys",
			cloudProviderUUID: uuid.NewV4().String(),
			googleAccount:     "{\"test\": \"account\"}",
		},
		{
			name:              "upload azure keys",
			cloudProviderUUID: uuid.NewV4().String(),
			azureAccessToken:  "{\"test\": \"access token\"}",
			azureProfile:      "{\"test\": \"profile\"}",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := os.RemoveAll(outputDirectory)
			require.NoError(t, err)

			err = os.MkdirAll(outputDirectory, 0755)
			require.NoError(t, err)

			defer func() {
				if err = os.RemoveAll(outputDirectory); err != nil {
					fmt.Println("Failed to clean up", outputDirectory, ":", err)
				}
			}()

			defaults, err := services.NewKeyFileDefaults()
			require.NoError(t, err)
			defaults.UserHomeDir = outputDirectory
			defaults.KeyHomeDir = outputDirectory

			request := &services.UploadCloudKeysBody{
				CloudProviderUUID: tt.cloudProviderUUID,
				AWSAccessKey:      base64.StdEncoding.EncodeToString([]byte(tt.awsAccessKey)),
				AWSSecretKey:      base64.StdEncoding.EncodeToString([]byte(tt.awsSecretKey)),
				AzureAccessTokens: base64.StdEncoding.EncodeToString([]byte(tt.azureAccessToken)),
				AzureProfile:      base64.StdEncoding.EncodeToString([]byte(tt.azureProfile)),
				GoogleAccount:     base64.StdEncoding.EncodeToString([]byte(tt.googleAccount)),
			}

			cs := services.ContrailService{}
			err = cs.UploadCloudKeys(request, defaults)
			require.NoError(t, err)

			for keyPath, content := range map[string]string{
				defaults.GetAWSSecretPath(tt.cloudProviderUUID): tt.awsSecretKey,
				defaults.GetAWSAccessPath(tt.cloudProviderUUID): tt.awsAccessKey,
				defaults.GetAzureAccessTokenPath():              tt.azureAccessToken,
				defaults.GetAzureProfilePath():                  tt.azureProfile,
				defaults.GetGoogleAccountPath():                 tt.googleAccount,
			} {
				b, err := ioutil.ReadFile(keyPath)
				assert.NoError(t, err)
				assert.Equal(t, content, string(b))
			}
		})
	}
}
