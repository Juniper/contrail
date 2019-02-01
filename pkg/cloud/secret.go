package cloud

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/log"
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
		log:    log.NewFileLogger("topology", c.config.LogFile),
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

	keypair, err := s.getCredKeyPairIfExists(d)

	if err == nil {
		return keypair, nil
	}
	s.log.Debugf("Error while reading cred ref of cloud(%s): %s",
		cloudID, err)

	// logic to handle a ssh key generation if not added as cred ref
	pubKey, pvtKey, err := genKeyPair(bits)
	if err != nil {
		return nil, err
	}

	if err = fileutil.WriteToFile(getCloudSSHKeyPath(cloudID, defaultSSHPvtKey),
		pvtKey, defaultRWOnlyPerm); err != nil {
		return nil, err
	}
	if err = fileutil.WriteToFile(getCloudSSHKeyPath(cloudID, defaultSSHPubKey),
		pubKey, sshPubKeyPerm); err != nil {
		return nil, err
	}
	return &models.Keypair{
		Name:         defaultSSHPvtKey,
		SSHPublicKey: string(pubKey),
	}, nil
}

func getPvtKeyIfValid(kp *models.Keypair) ([]byte, error) {

	if _, err := os.Stat(filepath.Join(kp.SSHKeyPath, kp.Name)); err != nil {
		return nil, errors.New("ssh private key path give in keypair is not valid")
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

func (s *secret) getCredKeyPairIfExists(d *Data) (*models.Keypair, error) {

	if d.credentials != nil {
		for _, cred := range d.credentials {
			for _, keyPairRef := range cred.KeypairRefs {
				keypair, err := getKeyPairObject(s.ctx, keyPairRef.UUID, s.cloud)
				if err != nil {
					return nil, err
				}
				if s.cloud.config.Test {
					return keypair, nil
				}
				rawPvtKey, err := getPvtKeyIfValid(keypair)
				if err != nil {
					return nil, err
				}
				if err = fileutil.WriteToFile(getCloudSSHKeyPath(d.info.UUID, keypair.Name),
					rawPvtKey, defaultRWOnlyPerm); err != nil {
					return nil, err
				}
				return keypair, nil
			}
			break
		}
	}

	return nil, errors.New("credential object is not referred by cloud")
}
