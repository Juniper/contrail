package cloud

import (
	//"fmt"
	//"path/filepath"

	//"github.com/flosch/pongo2"
	"fmt"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	defaultSecretFile     = "secret.yml"
	defaultSecretTemplate = "secret.tmpl"
)

type secretFileConfig struct {
	pubKey       map[string]string
	awsAccessKey string
	awsSecretKey string
}

type secret struct {
	cloud  *Cloud
	sfc    *secretFileConfig
	log    *logrus.Entry
	action string
}

func (s *secret) getSecretFile() string {
	return filepath.Join(s.cloud.getWorkingDir(), defaultSecretFile)
}

func (s *secret) getSecretTemplate() string {
	return filepath.Join(s.cloud.getTemplateRoot(), defaultSecretTemplate)
}

func (s *secret) createSecretFile() error {
	secretFile := s.getSecretFile()

	context := pongo2.Context{
		"secret": s.sfc,
	}
	content, err := applyTemplate(s.getSecretTemplate(), context)
	if err != nil {
		return err
	}

	err = writeToFile(secretFile, content)
	if err != nil {
		return err
	}
	s.log.Infof("Created secret file for cloud with uuid: %s ", s.cloud.config.CloudID)
	return nil
}

func getCredObject(uuid string, c *Cloud) (*models.Credential, error) {

	ctx := returnContext()
	request := new(services.GetCredentialRequest)
	request.ID = uuid

	credResp, err := c.APIServer.GetCredential(ctx, request)
	if err != nil {
		return nil, err
	}

	return credResp.GetCredential(), nil
}

func getKeyPairObject(uuid string, c *Cloud) (*models.Keypair, error) {

	ctx := returnContext()
	request := new(services.GetKeypairRequest)
	request.ID = uuid

	kpResp, err := c.APIServer.GetKeypair(ctx, request)
	if err != nil {
		return nil, err
	}

	return kpResp.GetKeypair(), nil

}

func (s *secret) updateFileConfig(d *Data) error {

	pk := make(map[string]string)

	cloud, err := s.cloud.getCloudObject()
	if err != nil {
		return err
	}

	sfc := &secretFileConfig{}

	for _, credRef := range cloud.CredentialRefs {
		credObject, err := getCredObject(credRef.UUID, s.cloud)
		if err != nil {
			return err
		}

		for _, keyPairRef := range credObject.KeypairRefs {
			keyPairObject, err := getKeyPairObject(keyPairRef.UUID, s.cloud)
			if err != nil {
				return err
			}
			pk[keyPairObject.Name] = keyPairObject.SSHPublicKey
			sfc.pubKey = pk
			break
		}
		break
	}

	if cloud.Type == "aws" {
		// TODO(madhukar) - Needs to handle multiple projects
		userObject, err := s.cloud.getCloudUser(d)
		if err != nil {
			return err
		}

		if userObject.AwsCredential.AccessKey == "" {
			return fmt.Errorf("AWS access key not specified")
		}
		sfc.awsAccessKey = userObject.AwsCredential.AccessKey

		if userObject.AwsCredential.SecretKey == "" {
			return fmt.Errorf("AWS secret key not specified")
		}
		sfc.awsSecretKey = userObject.AwsCredential.SecretKey
	}

	return nil
}

func (c *Cloud) newSecret(d *Data) (*secret, error) {

	// create logger for secret
	logger := pkglog.NewFileLogger("topology", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	s := &secret{
		cloud:  c,
		log:    logger,
		action: c.config.Action,
	}

	err := s.updateFileConfig(d)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *Cloud) getSecret(d *Data) (*secret, error) {

	s, err := c.newSecret(d)

	if err != nil {
		return nil, err
	}

	err = s.createSecretFile()

	if err != nil {
		return nil, err
	}

	return s, nil

}
