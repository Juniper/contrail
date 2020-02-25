package agent

import (
	"fmt"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/deploy"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type eventHandler struct{}

type variables struct {
	uuid            string
	schemaID        string
	action          string
	provisionerType string
	config          *Config
}

func newEventHandler() *eventHandler {
	return &eventHandler{}
}

func (h *eventHandler) processCluster(e *services.Event, c *Config) error {
	provisionerType, ok := e.GetResource().ToMap()["provisioner_type"].(string)
	if !ok {
		return errors.New("provisioner type conversion failed")
	}
	v := &variables{
		uuid:            e.GetUUID(),
		schemaID:        basemodels.KindToSchemaID(e.Kind()),
		action:          e.Operation(),
		provisionerType: provisionerType,
		config:          c,
	}

	switch e.Operation() {
	case services.OperationCreate:
		return createCluster(v)
	case services.OperationUpdate:
		return updateCluster(v)
	case services.OperationDelete:
		return deleteCluster(v)
	}

	return nil
}

func createCluster(v *variables) error {
	deployConfig := generateDeployConfig(v)

	logrus.Debug(fmt.Sprintf("AGENT creating %s", v.schemaID))
	if err := manageCluster(deployConfig); err != nil {
		return errors.Wrapf(err, "AGENT %s create failed", v.schemaID)
	}
	logrus.Debug(fmt.Sprintf("AGENT %s create complete", v.schemaID))

	return nil
}

func updateCluster(v *variables) error {
	deployConfig := generateDeployConfig(v)

	logrus.Debug(fmt.Sprintf("AGENT updating %s", v.schemaID))
	if err := manageCluster(deployConfig); err != nil {
		return errors.Wrapf(err, "AGENT %s update failed", v.schemaID)
	}
	logrus.Debug(fmt.Sprintf("AGENT %s update complete", v.schemaID))

	return nil
}

func deleteCluster(v *variables) error {
	deployConfig := generateDeployConfig(v)

	logrus.Debug(fmt.Sprintf("AGENT deleting %s", v.schemaID))
	if err := manageCluster(deployConfig); err != nil {
		return errors.Wrapf(err, "AGENT %s delete failed", v.schemaID)
	}
	logrus.Debug(fmt.Sprintf("AGENT deleting %s", v.schemaID))

	return nil
}

func generateDeployConfig(v *variables) *deploy.Config {
	deployConfig := &deploy.Config{
		ID:                  v.config.ID,
		Password:            v.config.Password,
		DomainID:            v.config.DomainID,
		ProjectID:           v.config.ProjectID,
		DomainName:          v.config.DomainName,
		ProjectName:         v.config.ProjectName,
		AuthURL:             v.config.AuthURL,
		Endpoint:            v.config.Endpoint,
		InSecure:            v.config.InSecure,
		ResourceType:        v.schemaID,
		ResourceID:          v.uuid,
		Action:              v.action,
		ProvisionerType:     v.provisionerType,
		LogLevel:            "debug",
		LogFile:             "/var/log/contrail/deploy.log",
		TemplateRoot:        "/usr/share/contrail/templates/",
		ServiceUserID:       v.config.ServiceUserID,
		ServiceUserPassword: v.config.ServiceUserPassword,
	}

	return deployConfig
}

func manageCluster(c *deploy.Config) error {
	manager, err := deploy.NewDeploy(c)
	if err != nil {
		return errors.Wrap(err, "creation of cluster manager failed")
	}

	if err = manager.Manage(); err != nil {
		return errors.Wrap(err, "management of cluster failed")
	}

	return nil
}
