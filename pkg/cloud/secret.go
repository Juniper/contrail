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
	keypair      *models.Keypair
	awsAccessKey string
	awsSecretKey string
}

type secret struct {
	cloud  *Cloud
	sfc    *secretFileConfig
	log    *logrus.Entry
	action string
}

func (s *secret) getSecretTemplate() string {
	return filepath.Join(s.cloud.getTemplateRoot(), defaultSecretTemplate)
}

func (s *secret) createSecretFile() error {
	secretFile := GetSecretFile(s.cloud.config.CloudID)

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

	cloud, err := GetCloud(s.cloud.APIServer, s.cloud.config.CloudID)
	if err != nil {
		return err
	}

	for _, credRef := range cloud.CredentialRefs {
		credObject, err := getCredObject(credRef.UUID, s.cloud)
		if err != nil {
			return err
		}

		for _, keyPairRef := range credObject.KeypairRefs {
			s.sfc.keypair, err = getKeyPairObject(keyPairRef.UUID, s.cloud)
			if err != nil {
				return err
			}
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
			return fmt.Errorf("aws access key not specified")
		}
		s.sfc.awsAccessKey = userObject.AwsCredential.AccessKey

		if userObject.AwsCredential.SecretKey == "" {
			return fmt.Errorf("aws secret key not specified")
		}
		s.sfc.awsSecretKey = userObject.AwsCredential.SecretKey
	}

	return nil
}

func (c *Cloud) newSecret() (*secret, error) {

	// create logger for secret
	logger := pkglog.NewFileLogger("topology", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	sfc := &secretFileConfig{}

	return &secret{
		cloud:  c,
		log:    logger,
		action: c.config.Action,
		sfc:    sfc,
	}, nil
}
