package cloud

import (
	"context"
	"bytes"
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
	keypair      *models.Keypair
	awsAccessKey string
	awsSecretKey string
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
		user, err := d.getDefaultCloudUser()
		if err != nil {
			return err
		}

		if user.AwsCredential.AccessKey == "" {
			return fmt.Errorf("aws access key not specified")
		}
		s.sfc.awsAccessKey = user.AwsCredential.AccessKey
		if user.AwsCredential.SecretKey == "" {
			return fmt.Errorf("aws secret key not specified")
		}
		s.sfc.awsSecretKey = user.AwsCredential.SecretKey
	}
	return nil
}

func (c *Cloud) newSecret() (*secret, error) {
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

func getPvtKeyIfValid(kp *models.Keypair) ([]byte, error) {

	if _, err := os.Stat(filepath.Join(kp.SSHKeyPath, kp.Name)); err != nil {
		return nil, errors.New("ssh private key path given in keypair is not valid")
	}
	pvtKeyPem, err := fileutil.GetContent("file://" + kp.SSHKeyPath + kp.Name)
	if err != nil {
		return nil, err
	}
	_, err = ssh.ParseRawPrivateKey(pvtKeyPem)
	return pvtKeyPem, err
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
				if keypair.SSHKeyPath == "" && keypair.SSHPublicKey == "" {
					err = createSSHKey(cloudID, keypair.Name)
					if err != nil {
						s.log.Errorf("error while creating ssh keys: %v", err)
						return nil, err
					}
					// get pub key content
					pubKey, err := fileutil.GetContent("file://" + getCloudSSHKeyPath(cloudID, keypair.Name+".pub")) //nolint: govet
					if err != nil {
						return nil, err
					}
					// update keypair
					keypair.SSHPublicKey = bytes.NewBuffer(pubKey).String()
					keypair.SSHKeyPath = getCloudSSHKeyDir(cloudID)

					kpUpdateReq := new(services.UpdateKeypairRequest)
					kpUpdateReq.Keypair = keypair

					_, err = s.cloud.APIServer.UpdateKeypair(s.ctx,
						&services.UpdateKeypairRequest{
							Keypair: keypair,
						},
					)
					return keypair, err

				}

				if keypair.SSHKeyPath == "" || keypair.SSHPublicKey == "" {
					return nil, errors.New(`ssh private key path and public key
						both needs to be given in keypair object
						attached to cloud: ` + cloudID)
				}

				if ifKeyFileAlreadyExists(keypair, cloudID) {
					return keypair, nil
				}

				err = copySHHKeyPairIfValid(keypair, cloudID)
				if err != nil {
					return nil, err
				}
				return keypair, nil
			}
		}
	}

	return nil, errors.New("credential object is not referred by cloud")
}

func ifKeyFileAlreadyExists(keypair *models.Keypair, cloudID string) bool {

	if keypair.SSHKeyPath == "" || keypair.SSHPublicKey == "" {
		return false
	}
	objKeyByte := bytes.NewBufferString(keypair.SSHPublicKey).Bytes()
	pubKey, err := fileutil.GetContent("file://" + getCloudSSHKeyPath(cloudID, keypair.Name+".pub"))
	if err != nil {
		return false
	}
	return bytes.Equal(objKeyByte, pubKey)

}

func copySHHKeyPairIfValid(keypair *models.Keypair, cloudID string) error {

	// check if pub key is valid
	rawPubkey := []byte(keypair.SSHPublicKey)

	if err := fileutil.WriteToFile(getCloudSSHKeyPath(cloudID,
		keypair.Name+".pub"), rawPubkey, sshPubKeyPerm); err != nil {
		return err
	}

	// check if pvt key is valid
	rawPvtKey, err := getPvtKeyIfValid(keypair)
	if err != nil {
		return err
	}

	return fileutil.WriteToFile(getCloudSSHKeyPath(cloudID, keypair.Name),
		rawPvtKey, defaultRWOnlyPerm)
}

func createSSHKey(cloudID string, name string) error {
	// logic to handle a ssh key generation if not added as cred ref
	pubKey, pvtKey, err := genKeyPair(bits)
	if err != nil {
		return err
	}

	if err = fileutil.WriteToFile(getCloudSSHKeyPath(cloudID, name),
		pvtKey, defaultRWOnlyPerm); err != nil {
		return err
	}
	return fileutil.WriteToFile(getCloudSSHKeyPath(cloudID, name+".pub"),
		pubKey, sshPubKeyPerm)
}
