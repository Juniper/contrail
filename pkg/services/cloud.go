package services

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/labstack/echo"
)

//API Path definitions.
const (
	UploadCloudKeysPath = "upload-cloud-keys"
	keyHomeDir          = "/var/tmp/contrail"
	secretKeyFileName   = "aws_secret.key"
	accessKeyFileName   = "aws_access.key"
	accessTokenFileName = "accessTokens.json"
	profileFileName     = "azureProfile.json"
	accountFileName     = "google-account.json"
	azureDir            = ".azure"
)

// UploadCoudKeysBody defines data format for /upload-cloud-keys endpoint.
type UploadCloudKeysBody struct {
	// CloudProviderUUID is the UUID of the cloud provider who provided keys
	CloudProviderUUID string `json:"cloud_provider_uuid"`
	// AWSSecretKey is the secret key to created on API Server host.
	AWSSecretKey string `json:"aws_secret_key"`
	// AWSAccessKey is the access key to created on API Server host.
	AWSAccessKey string `json:"aws_access_key"`
	// AzureAccessTokenJson is the token file to created on API Server host.
	AzureAccessToken string `json:"azure_access_token_json"`
	// AzureProfileJson is the profile file to created on API Server host.
	AzureProfile string `json:"azure_profile_json"`
	// GoogleAccountJson is the account file to created on API Server host.
	GoogleAccount string `json:"google_account_json"`
}

// KeyFileDefaults defines data format for various cloud secret file
type KeyFileDefaults struct {
	UserHomeDir         string
	KeyHomeDir          string
	SecretKeyFileName   string
	AccessKeyFileName   string
	AccessTokenFileName string
	ProfileFileName     string
	AccountFileName     string
}

// NewKeyFileDefaults returns defaults for various cloud secret files.
func NewKeyFileDefaults() (defaults *KeyFileDefaults, err error) {
	userHomeDir, err := getHomeDir()
	if err != nil {
		return nil, err
	}
	return &KeyFileDefaults{
		userHomeDir,
		keyHomeDir,
		secretKeyFileName,
		accessKeyFileName,
		accessTokenFileName,
		profileFileName,
		accountFileName,
	}, nil
}

// GetAWSSecretPath determines the aws secret key path
func (defaults *KeyFileDefaults) GetAWSSecretPath(cloudProviderUUID string) string {
	return path.Join(defaults.KeyHomeDir, cloudProviderUUID, defaults.SecretKeyFileName)
}

// GetAWSAccessPath determines the aws access key path
func (defaults *KeyFileDefaults) GetAWSAccessPath(cloudProviderUUID string) string {
	return path.Join(defaults.KeyHomeDir, cloudProviderUUID, defaults.AccessKeyFileName)
}

// GetAzureAccessTokenPath determines the azure access token path
func (defaults *KeyFileDefaults) GetAzureAccessTokenPath() string {
	return path.Join(defaults.UserHomeDir, azureDir, defaults.AccessTokenFileName)
}

// GetAzureProfilePath determines the azure profile path
func (defaults *KeyFileDefaults) GetAzureProfilePath() string {
	return path.Join(defaults.UserHomeDir, azureDir, defaults.ProfileFileName)
}

// GetGoogleAccountPath determines the google account path
func (defaults *KeyFileDefaults) GetGoogleAccountPath() string {
	return path.Join(defaults.KeyHomeDir, defaults.AccountFileName)
}

// RESTUploadCloudKeys handles an /upload-cloud-keys REST request.
func (service *ContrailService) RESTUploadCloudKeys(c echo.Context) error {
	var request *UploadCloudKeysBody
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("invalid JSON format: %v", err))
	}
	defaults, err := NewKeyFileDefaults()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			fmt.Sprintf("not able to fetch home dir: %v", err))
	}
	return service.UploadCloudKeys(request, defaults)
}

// UploadCloudKeys stores specified cloud secrets
func (service *ContrailService) UploadCloudKeys(request *UploadCloudKeysBody, keyDefaults *KeyFileDefaults) error {
	if request.CloudProviderUUID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "cloud provider is empty")
	}

	for keyType, secret := range map[string]map[string]string{
		"aws-secret-key": {
			"encoded": request.AWSSecretKey,
			"path":    keyDefaults.GetAWSSecretPath(request.CloudProviderUUID),
		},
		"aws-access-key": {
			"encoded": request.AWSAccessKey,
			"path":    keyDefaults.GetAWSAccessPath(request.CloudProviderUUID),
		},
		"azure-access-token": {
			"encoded": request.AzureAccessToken,
			"path":    keyDefaults.GetAzureAccessTokenPath(),
		},
		"azure-profile": {
			"encoded": request.AzureProfile,
			"path":    keyDefaults.GetAzureProfilePath(),
		},
		"google-account": {
			"encoded": request.GoogleAccount,
			"path":    keyDefaults.GetGoogleAccountPath(),
		},
	} {
		keyPaths := []string{}
		if err := decodeAndStoreCloudKey(keyType, secret["path"], secret["encoded"], keyPaths); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError,
				fmt.Sprintf("failed to store secret key: %v", err))
		} else {
			keyPaths = append(keyPaths, secret["path"])
		}
	}
	return nil
}

func decodeAndStoreCloudKey(keyType, keyPath, encodedSecret string, existingKeyPaths []string) error {
	decodedSecret, err := base64.StdEncoding.DecodeString(encodedSecret)
	if err != nil {
		errstrings := []string{fmt.Sprintf("failed to base64-decode %s: %v", keyType, err)}
		errstrings = append(errstrings, cleanupCloudKeys(existingKeyPaths)...)
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}

	err = os.MkdirAll(path.Dir(keyPath), 0755)
	if err != nil {
		errstrings := []string{fmt.Sprintf("failed to make dir for %s: %v", keyType, err)}
		errstrings = append(errstrings, cleanupCloudKeys(existingKeyPaths)...)
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}

	if err = ioutil.WriteFile(keyPath, decodedSecret, 0644); err != nil {
		errstrings := []string{fmt.Sprintf("failed to store %s: %v", keyType, err)}
		errstrings = append(errstrings, cleanupCloudKeys(existingKeyPaths)...)
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}

	return nil
}

func cleanupCloudKeys(keyPaths []string) (errstrings []string) {
	for _, keyPath := range keyPaths {
		err := os.Remove(keyPath)
		if err != nil {
			errstrings = append(errstrings, fmt.Sprintf("Unable to delete %s: %v", keyPath, err))
		}
	}
	return errstrings
}

func getHomeDir() (homeDir string, err error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}
