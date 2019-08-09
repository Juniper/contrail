package cloud

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
	keypair      *models.Keypair	`yaml:"public_key"`
	awsAccessKey string	`yaml:"aws_access_key"`
	awsSecretKey string	`yaml:"aws_secret_key"`
	// TODO: Ensure this is proper tag
	providerType string	`yaml:"provider_type"`
	authorizedRegistries []dockerRegistry `yaml:"authorized_registries"`
}

func (s *secretFileConfig) addRegistries(registries []dockerRegistry) {
	s.authorizedRegistries = append(s.authorizedRegistries, registries...)
	s.removeExistingRegistries()
}

func (s *secretFileConfig) removeExistingRegistries() {
	updatedRegistries := []dockerRegistry{}
	regAlreadyExists := map[[3]string]bool{}

	for _, registry := range s.authorizedRegistries {
		regKey := [3]string{registry.registry, registry.username, registry.password}
		if !regAlreadyExists[regKey] {
			updatedRegistries = append(updatedRegistries, registry)
			regAlreadyExists[regKey] = true
		}
	}
	s.authorizedRegistries = updatedRegistries
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

	context := pongo2.Context{
		"secret": s.sfc,
	}
	content, err := template.Apply(s.getSecretTemplate(), context)
	if err != nil {
		return err
	}

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

func (s *secret) updateFileConfig(d *Data) error {

	keypair, err := s.getKeypair(d)
	if err != nil {
		return err
	}
	s.sfc.keypair = keypair

	if d.hasProviderAWS() {
		awsCreds, err := loadAWSCredentialsFromFile(GetSecretFile(d.cloud.config.CloudID))
		if err != nil {
			return err
		}

		if awsCreds.AccessKey == "" {
			return fmt.Errorf("aws access key not specified")
		}
		s.sfc.awsAccessKey = awsCreds.AccessKey
		if awsCreds.SecretKey == "" {
			return fmt.Errorf("aws secret key not specified")
		}
		s.sfc.awsSecretKey = awsCreds.SecretKey
	}
	if d.hasProviderGCP() {
		s.sfc.providerType = gcp
	}
	return nil
}

func loadAWSCredentialsFromFile(credFile string) (*models.AWSCredential, error) {
	data, err := ioutil.ReadFile(credFile)
    if err != nil {
        return nil, err
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) != 2 {
		return nil, errors.New("Invalid AWS Credential File")
	}

	creds := strings.Split(string(data), ",")
	if len(creds) != 2 {
		return nil, errors.New("Invalid AWS Credential File")
	}

    return &models.AWSCredential{AccessKey: creds[0], SecretKey: creds[1]}, nil
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

//nolint: gocyclo
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
