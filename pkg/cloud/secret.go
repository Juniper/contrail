package cloud

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	bits          = 2048
	sshDirPerm    = 0700
	sshPubKeyPerm = 0644
)

type secretFileConfig struct {
	keypair             *models.Keypair
	providerType        string
	awsAccessKey        string
	awsSecretKey        string
	azureSubscriptionId string
	azureClientId       string
	azureClientSecret   string
	azureTenantId       string
}

type secret struct {
	cloud  *Cloud
	sfc    *secretFileConfig
	log    *logrus.Entry
	action string
	ctx    context.Context
}

func (s *secret) getSecretTemplate() string {
	return filepath.Join(s.cloud.getTemplateRoot(), defaultSecretTemplate)
}

func (s *secret) createSecretFile() error {
	secretFile := GetSecretFile(s.cloud.config.CloudID)

	templateContext := pongo2.Context{
		"secret":            s.sfc,
		"gcpCredentialFile": defaultGCPCredentialFile,
	}
	content, err := template.Apply(s.getSecretTemplate(), templateContext)
	if err != nil {
		return err
	}

	_ = fileutil.WriteToFile("/var/tmp/cloud/aws_secret.yaml", content, defaultRWOnlyPerm)
	err = fileutil.WriteToFile(secretFile, content, defaultRWOnlyPerm)
	if err != nil {
		return err
	}
	s.log.Infof("Created secret file for cloud with uuid: %s ",
		s.cloud.config.CloudID)
	return nil
}

func getCredObject(ctx context.Context, client *client.HTTP,
	uuid string) (*models.Credential, error) {

	request := new(services.GetCredentialRequest)
	request.ID = uuid

	credResp, err := client.GetCredential(ctx, request)
	if err != nil {
		return nil, err
	}

	return credResp.GetCredential(), nil
}

func getKeyPairObject(ctx context.Context, uuid string,
	c *Cloud) (*models.Keypair, error) {

	request := new(services.GetKeypairRequest)
	request.ID = uuid

	kpResp, err := c.APIServer.GetKeypair(ctx, request)
	if err != nil {
		return nil, err
	}

	return kpResp.GetKeypair(), nil
}

func (s *secret) updateCloudData(d *Data) error {
	keypair, err := s.getKeypair(d)
	if err != nil {
		return err
	}
	s.sfc.keypair = keypair

	user, err := d.getDefaultCloudUser()
	if err != nil {
		return err
	}

	if d.hasProviderAWS() {
		s.sfc.providerType = aws
		if user.AwsCredential.AccessKey == "" {
			return fmt.Errorf("aws access key not specified")
		}
		s.sfc.awsAccessKey = user.AwsCredential.AccessKey
		if user.AwsCredential.SecretKey == "" {
			return fmt.Errorf("aws secret key not specified")
		}
		s.sfc.awsSecretKey = user.AwsCredential.SecretKey
	}
	if d.hasProviderGCP() {
		s.sfc.providerType = gcp
		if user.GCPCredential == "" {
			return fmt.Errorf("gcp credentials not specified")
		}
	}

	if d.hasProviderAzure() {
		s.sfc.providerType = azure
		if user.AzureCredential.AzureTenantID == "" {
			return fmt.Errorf("azure tenant id not specified")
		}
		s.sfc.azureTenantId = user.AzureCredential.AzureTenantID

		if user.AzureCredential.AzureSubscriptionID == "" {
			return fmt.Errorf("azure subscription id not specified")
		}
		s.sfc.azureSubscriptionId = user.AzureCredential.AzureSubscriptionID

		if user.AzureCredential.AzureClientID == "" {
			return fmt.Errorf("azure client id not specified")
		}
		s.sfc.azureClientId = user.AzureCredential.AzureClientID

		if user.AzureCredential.AzureClientSecret == "" {
			return fmt.Errorf("azure client secret not specified")
		}
		s.sfc.azureClientSecret = user.AzureCredential.AzureClientSecret
	}
	return nil
}

func (s *secret) saveCloudCredentials(d *Data) error {
	if d.hasProviderGCP() {
		if err := d.saveGCPCredentialsToDisk(); err != nil {
			return err
		}
	}
	return nil
}

func newSecret(c *Cloud) (*secret, error) {
	return &secret{
		cloud:  c,
		log:    logutil.NewFileLogger("topology", c.config.LogFile),
		action: c.config.Action,
		sfc:    &secretFileConfig{},
		ctx:    c.ctx,
	}, nil
}

func (s *secret) getKeypair(d *Data) (*models.Keypair, error) {
	// TODO(madhukar) - optimize handling of multiple cloud users
	cloudID := d.info.UUID
	if err := os.MkdirAll(getCloudSSHKeyDir(cloudID), sshDirPerm); err != nil {
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

func getCloudSSHKeyDir(cloudID string) string {
	return filepath.Join(GetCloudDir(cloudID), defaultSSHKeyRepo)
}

func getCloudSSHKeyPath(cloudID string, name string) string {
	return filepath.Join(getCloudSSHKeyDir(cloudID), name)
}

// nolint: gocyclo
func (s *secret) getCredKeyPairIfExists(d *Data,
	cloudID string) (*models.Keypair, error) {

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
			keypair.SSHKeyDirPath = getCloudSSHKeyDir(cloudID)
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

	if err = fileutil.WriteToFile(getCloudSSHKeyPath(cloudID,
		keypair.DisplayName+".pub"), rawPubkey, sshPubKeyPerm); err != nil {
		return err
	}

	keypair.SSHPublicKey = string(rawPubkey)

	// check if pvt key is valid
	rawPvtKey, err := getSSHKeyIfValid(keypair, privateSSHKey)
	if err != nil {
		return err
	}

	return fileutil.WriteToFile(getCloudSSHKeyPath(cloudID, keypair.DisplayName),
		rawPvtKey, defaultRWOnlyPerm)
}

func createSSHKey(cloudID string, keypair *models.Keypair) error {
	// logic to handle a ssh key generation if not added as cred ref
	pubKey, pvtKey, err := genKeyPair(bits)
	if err != nil {
		return err
	}
	keypair.SSHPublicKey = string(pubKey)

	if err = fileutil.WriteToFile(getCloudSSHKeyPath(cloudID, keypair.DisplayName),
		pvtKey, defaultRWOnlyPerm); err != nil {
		return err
	}
	return fileutil.WriteToFile(getCloudSSHKeyPath(cloudID, keypair.DisplayName+".pub"),
		pubKey, sshPubKeyPerm)
}
