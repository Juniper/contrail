package cloud

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"

	yaml "gopkg.in/yaml.v2"
)

const (
	bits          = 2048
	sshDirPerm    = 0700
	sshPubKeyPerm = 0644
)

// Keypair holds name and public SSH key value
type Keypair struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// AuthorizedRegistry stores the information about authorized docker registries
type AuthorizedRegistry struct {
	Registry string `yaml:"registry,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Tag      string `yaml:"tag,omitempty"`
}

// SecretFileConfig holds the secret keys of the cloud
type SecretFileConfig struct {
	Keypair              Keypair               `yaml:"public_key"`
	AWSAccessKey         string                `yaml:"aws_access_key,omitempty"`
	AWSSecretKey         string                `yaml:"aws_secret_key,omitempty"`
	AzureSubscriptionID  string                `yaml:"azure_subscription_id,omitempty"`
	AzureClientID        string                `yaml:"azure_client_id,omitempty"`
	AzureClientSecret    string                `yaml:"azure_client_secret,omitempty"`
	AzureTenantID        string                `yaml:"azure_tenant_id,omitempty"`
	GoogleAccount        string                `yaml:"google_credentials,omitempty"`
	AuthorizedRegistries []*AuthorizedRegistry `yaml:"authorized_registries,omitempty"`
}

type secret struct {
	cloud *Cloud
	sfc   *SecretFileConfig
	log   *logrus.Entry
	ctx   context.Context
}

func (s *secret) createSecretFile(clusterUUID string) error {
	if err := s.updateAuthorizedRegistry(clusterUUID); err != nil {
		return err
	}
	marshaled, err := yaml.Marshal(s.sfc)
	if err != nil {
		return errors.Wrapf(err, "couldn't marshal secret to yaml for cloud %s", s.cloud.config.CloudID)
	}
	if err = ioutil.WriteFile(GetSecretFile(s.cloud.config.CloudID), marshaled, defaultRWOnlyPerm); err != nil {
		return errors.Wrapf(err, "couldn't create secret.yml file for cloud %s", s.cloud.config.CloudID)
	}

	s.log.Infof("Created secret file for cloud with uuid: %s ", s.cloud.config.CloudID)
	return nil
}

func (s *secret) updateAuthorizedRegistry(clusterUUID string) error {
	clusterResp, err := s.cloud.APIServer.GetContrailCluster(s.ctx, &services.GetContrailClusterRequest{
		ID: clusterUUID,
	})
	if err != nil {
		return errors.Wrap(err, "cannot resolve Authentication Registries")
	}
	reg, err := NewAuthorizedRegistry(clusterResp.ContrailCluster)
	if err != nil {
		return err
	}
	s.sfc.AuthorizedRegistries = []*AuthorizedRegistry{reg}

	return nil
}

// NewAuthorizedRegistry creates AuthorizedRegistry based on ContrailCluster parameters
func NewAuthorizedRegistry(c *models.ContrailCluster) (*AuthorizedRegistry, error) {
	if c.ContainerRegistry == "" {
		return nil, errors.Errorf("authorized registry isn't specified")
	}
	if c.ContainerRegistryUsername == "" {
		return nil, errors.Errorf("authorized registry username isn't specified")
	}
	if c.ContainerRegistryPassword == "" {
		return nil, errors.Errorf("authorized registry password isn't specified")
	}
	return &AuthorizedRegistry{
		Registry: c.ContainerRegistry,
		Username: c.ContainerRegistryUsername,
		Password: c.ContainerRegistryPassword,
		Tag:      c.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG"),
	}, nil
}

func getCredObject(ctx context.Context, client *client.HTTP, uuid string) (*models.Credential, error) {
	request := new(services.GetCredentialRequest)
	request.ID = uuid

	credResp, err := client.GetCredential(ctx, request)
	if err != nil {
		return nil, err
	}

	return credResp.GetCredential(), nil
}

func getKeyPairObject(ctx context.Context, uuid string, c *Cloud) (*models.Keypair, error) {
	request := new(services.GetKeypairRequest)
	request.ID = uuid

	kpResp, err := c.APIServer.GetKeypair(ctx, request)
	if err != nil {
		return nil, err
	}

	return kpResp.GetKeypair(), nil
}

// Update fills the secret file config
func (sfc *SecretFileConfig) Update(kp *models.Keypair) error {
	sfc.Keypair = Keypair{Name: kp.Name, Value: kp.SSHPublicKey}
	kfd := services.NewKeyFileDefaults()

	if err := sfc.updateAWSCredentials(kfd); err != nil {
		return err
	}
	if err := sfc.updateAzureCredentials(kfd); err != nil {
		return err
	}
	return sfc.updateGCPCredentials(kfd)
}

func (sfc *SecretFileConfig) updateAWSCredentials(kfd *services.KeyFileDefaults) error {
	if awsCredentialsExist(kfd) {
		awsCreds, err := loadAWSCredentials(
			kfd.GetAWSAccessPath(),
			kfd.GetAWSSecretPath(),
		)
		if err != nil {
			return err
		}
		sfc.AWSAccessKey = awsCreds.AccessKey
		sfc.AWSSecretKey = awsCreds.SecretKey
	}
	return nil
}

func awsCredentialsExist(kfd *services.KeyFileDefaults) bool {
	if _, err := os.Stat(kfd.GetAWSAccessPath()); err != nil {
		return false
	}
	if _, err := os.Stat(kfd.GetAWSSecretPath()); err != nil {
		return false
	}
	return true
}

func loadAWSCredentials(accessPath, secretPath string) (*models.AWSCredential, error) {
	accessKey, err := ioutil.ReadFile(accessPath)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve AWS Access Key")
	}

	secretKey, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve AWS Secret Key")
	}

	if len(accessKey) == 0 && len(secretKey) == 0 {
		return &models.AWSCredential{AccessKey: "", SecretKey: ""}, nil
	}

	if len(accessKey) == 0 {
		return nil, errors.New("AWS Access Key not specified")
	}

	if len(secretKey) == 0 {
		return nil, errors.New("AWS Secret Key not specified")
	}

	return &models.AWSCredential{AccessKey: string(accessKey), SecretKey: string(secretKey)}, nil
}

func (sfc *SecretFileConfig) updateAzureCredentials(kfd *services.KeyFileDefaults) error {
	if azureCredentialsExist(kfd) {
		azSfc, err := loadAzureCredentials(kfd)
		if err != nil {
			return err
		}
		sfc.AzureSubscriptionID = azSfc.AzureSubscriptionID
		sfc.AzureClientID = azSfc.AzureClientID
		sfc.AzureClientSecret = azSfc.AzureClientSecret
		sfc.AzureTenantID = azSfc.AzureTenantID
	}
	return nil
}

func azureCredentialsExist(kfd *services.KeyFileDefaults) bool {
	for _, p := range []string{
		kfd.GetAzureSubscriptionIDPath(),
		kfd.GetAzureClientIDPath(),
		kfd.GetAzureClientIDPath(),
		kfd.GetAzureTenantIDPath(),
	} {
		if _, err := os.Stat(p); err != nil {
			return false
		}
	}
	return true
}

func loadAzureCredentials(kfd *services.KeyFileDefaults) (*SecretFileConfig, error) {
	subscriptionID, err := ioutil.ReadFile(kfd.GetAzureSubscriptionIDPath())
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve Azure Subscription ID")
	}
	clientID, err := ioutil.ReadFile(kfd.GetAzureClientIDPath())
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve Azure Client ID")
	}
	clientSecret, err := ioutil.ReadFile(kfd.GetAzureClientSecretPath())
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve Azure Client Secret")
	}
	tenantID, err := ioutil.ReadFile(kfd.GetAzureTenantIDPath())
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve Azure Tenant ID")
	}

	if len(subscriptionID)*len(clientID)*len(clientSecret)*len(tenantID) == 0 &&
		len(subscriptionID)+len(clientID)+len(clientSecret)+len(tenantID) != 0 {
		return nil, errors.New("one of Azure credentials was not specified")
	}

	return &SecretFileConfig{
		AzureSubscriptionID: string(subscriptionID),
		AzureClientID:       string(clientID),
		AzureClientSecret:   string(clientSecret),
		AzureTenantID:       string(tenantID),
	}, nil
}

func (sfc *SecretFileConfig) updateGCPCredentials(kfd *services.KeyFileDefaults) error {
	bytes, err := ioutil.ReadFile(kfd.GetGoogleAccountPath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return errors.Wrap(err, "could not read GCP account file")
	}
	if len(bytes) == 0 {
		return nil
	}
	sfc.GoogleAccount = kfd.GetGoogleAccountPath()
	return nil
}

func newSecret(c *Cloud) *secret {
	return &secret{
		cloud: c,
		log:   logutil.NewFileLogger("secret", c.config.LogFile),
		sfc:   &SecretFileConfig{},
		ctx:   c.ctx,
	}
}

func (s *secret) getKeypair(d *Data) (*models.Keypair, error) {
	// TODO(madhukar) - optimize handling of multiple cloud users
	cloudID := d.info.UUID
	if err := os.MkdirAll(GetCloudSSHKeyDir(cloudID), sshDirPerm); err != nil {
		return nil, err
	}

	keypair, err := s.getCredKeyPairIfExists(d, cloudID)

	if err != nil {
		s.log.Errorf("error while reading cred ref of cloud(%s): %s",
			cloudID, err)
		return nil, err
	}
	return keypair, err
}

func getSSHKeyIfValid(kp *models.Keypair, keyType string) ([]byte, error) {
	var sshKeyFileName string
	if keyType == pubSSHKey {
		sshKeyFileName = filepath.Join(kp.SSHKeyDirPath, kp.DisplayName+".pub")
	} else if keyType == privateSSHKey {
		sshKeyFileName = filepath.Join(kp.SSHKeyDirPath, kp.DisplayName)
	} else {
		return nil, fmt.Errorf("key type: %s is not valid", keyType)
	}

	if _, err := os.Stat(sshKeyFileName); err != nil {
		return nil, fmt.Errorf("ssh key file %s is not located", sshKeyFileName)
	}

	if keyType == privateSSHKey {
		pvtKeyPem, err := fileutil.GetContent("file://" + sshKeyFileName)
		if err != nil {
			return nil, err
		}
		_, err = ssh.ParseRawPrivateKey(pvtKeyPem)
		return pvtKeyPem, err
	}

	pubKey, err := fileutil.GetContent("file://" + sshKeyFileName)
	if err != nil {
		return nil, err
	}

	_, _, _, _, err = ssh.ParseAuthorizedKey(pubKey)
	return pubKey, err
}

// GetCloudSSHKeyDir returns directory of SSH key for given Cloud.
func GetCloudSSHKeyDir(cloudID string) string {
	return filepath.Join(GetCloudDir(cloudID), defaultSSHKeyRepo)
}

// GetCloudSSHKeyPath returns path of SSH key for given Cloud and key name.
func GetCloudSSHKeyPath(cloudID string, name string) string {
	return filepath.Join(GetCloudSSHKeyDir(cloudID), name)
}

//nolint: gocyclo
func (s *secret) getCredKeyPairIfExists(d *Data, cloudID string) (*models.Keypair, error) {
	if d.credentials != nil {
		for _, cred := range d.credentials {
			for _, keyPairRef := range cred.KeypairRefs {
				keypair, err := getKeyPairObject(s.ctx, keyPairRef.UUID, s.cloud)
				if err != nil {
					return nil, err
				}

				// create random ssh key if keypath and pubkey is not given
				if keypair.SSHKeyDirPath == "" {
					return nil, errors.New("ssh_key_dir_path field is empty")
				}

				err = copySHHKeyPairIfValid(keypair, cloudID)
				if err != nil {
					return nil, err
				}
				return keypair, nil
			}

			// if keypair object is not attached to credential object
			// create keypair and attach it cred obj as ref
			kpName := fmt.Sprintf("keypair-%s", cred.UUID)
			keypair := &models.Keypair{
				Name:        kpName,
				DisplayName: kpName,
				FQName:      []string{"default-global-system-config", kpName},
				ParentType:  "global-system-config",
			}

			err := createSSHKey(cloudID, keypair)
			if err != nil {
				s.log.Errorf("error while creating ssh keys: %v", err)
				return nil, err
			}
			keypair.SSHKeyDirPath = GetCloudSSHKeyDir(cloudID)
			kpResp, err := s.cloud.APIServer.CreateKeypair(s.ctx,
				&services.CreateKeypairRequest{
					Keypair: keypair,
				},
			)
			if err != nil {
				return nil, err
			}

			// update cred object with keypair ref
			cred.KeypairRefs = append(cred.KeypairRefs,
				&models.CredentialKeypairRef{
					UUID: kpResp.Keypair.UUID,
				},
			)
			_, err = s.cloud.APIServer.UpdateCredential(s.ctx,
				&services.UpdateCredentialRequest{
					Credential: cred,
				},
			)
			return keypair, err
		}
	}

	return nil, errors.New("credential object is not referred by cloud")
}

func copySHHKeyPairIfValid(keypair *models.Keypair, cloudID string) error {
	// check if pub key is valid
	rawPubkey, err := getSSHKeyIfValid(keypair, pubSSHKey)
	if err != nil {
		return err
	}

	if err = fileutil.WriteToFile(GetCloudSSHKeyPath(cloudID,
		keypair.DisplayName+".pub"), rawPubkey, sshPubKeyPerm); err != nil {
		return err
	}

	keypair.SSHPublicKey = strings.TrimSpace(string(rawPubkey))

	// check if pvt key is valid
	rawPvtKey, err := getSSHKeyIfValid(keypair, privateSSHKey)
	if err != nil {
		return err
	}

	return fileutil.WriteToFile(GetCloudSSHKeyPath(cloudID, keypair.DisplayName),
		rawPvtKey, defaultRWOnlyPerm)
}

func createSSHKey(cloudID string, keypair *models.Keypair) error {
	// logic to handle a ssh key generation if not added as cred ref
	pubKey, pvtKey, err := genKeyPair(bits)
	if err != nil {
		return err
	}
	keypair.SSHPublicKey = strings.TrimSpace(string(pubKey))

	if err = fileutil.WriteToFile(GetCloudSSHKeyPath(cloudID, keypair.DisplayName),
		pvtKey, defaultRWOnlyPerm); err != nil {
		return err
	}
	return fileutil.WriteToFile(GetCloudSSHKeyPath(cloudID, keypair.DisplayName+".pub"),
		pubKey, sshPubKeyPerm)
}
