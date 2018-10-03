package cloud

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/common"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
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
	content, err := common.Apply(s.getSecretTemplate(), context)
	if err != nil {
		return err
	}

	err = common.WriteToFile(secretFile, content, defaultRWOnlyPerm)
	if err != nil {
		return err
	}
	s.log.Infof("Created secret file for cloud with uuid: %s ",
		s.cloud.config.CloudID)
	return nil
}

func getCredObject(client *client.HTTP,
	uuid string) (*models.Credential, error) {

	response := new(services.GetCredentialResponse)
	_, err := client.Read("/credential/"+uuid, response)
	if err != nil {
		return nil, err
	}

	return response.GetCredential(), nil
}

func getKeyPairObject(uuid string, c *Cloud) (*models.Keypair, error) {

	response := new(services.GetKeypairResponse)
	_, err := c.APIServer.Read("/keypair/"+uuid, response)
	if err != nil {
		return nil, err
	}

	return response.GetKeypair(), nil

}

func (s *secret) updateFileConfig(d *Data) error {

	// TODO(madhukar) - optimize handling of multiple cloud users
	for _, cred := range d.credentials {
		for _, keyPairRef := range cred.KeypairRefs {
			keypair, err := getKeyPairObject(keyPairRef.UUID, s.cloud)
			if err != nil {
				return err
			}
			s.sfc.keypair = keypair
			break
		}
		break
	}

	if s.sfc.keypair == nil {
		return errors.New("cred ref not found with cloud user Obj")
	}

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

	// create logger for secret
	logger := pkglog.NewLogger("topology")
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	sfc := &secretFileConfig{}

	return &secret{
		cloud:  c,
		log:    logger,
		action: c.config.Action,
		sfc:    sfc,
	}, nil
}
