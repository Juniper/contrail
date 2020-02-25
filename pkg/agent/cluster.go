package agent

import (
	"fmt"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/deploy"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func processCluster(e *services.Event, c *Config) error {
	context := make(map[string]string)
	context["schemaID"] = basemodels.KindToSchemaID(e.Kind())
	context["uuid"] = e.GetUUID()
	provisionerType, ok := e.GetResource().ToMap()["provisioner_type"].(string)
	if !ok {
		return errors.New("provisioner type conversion failed")
	}
	context["provisionerType"] = provisionerType
	context["action"] = e.Operation()
	context["schema_id"] = basemodels.KindToSchemaID(e.Kind())
	context["config_dir"] = fmt.Sprintf("/var/tmp/%s/config/%s", context["schemaID"], context["uuid"])

	switch e.Operation() {
	case services.OperationCreate:
		return clusterCreate(c, context)
	case services.OperationUpdate:
		return clusterUpdate(c, context)
	case services.OperationDelete:
		return clusterDelete(c, context)
	}

	return nil
}

func clusterCreate(c *Config, context map[string]string) error {
	logrus.Debug(fmt.Sprintf("AGENT creating %s", context["config_dir"]))
	if err := directoryHandler("create", context["config_dir"]); err != nil {
		return errors.Wrapf(err, "AGENT create %s failed", context["config_dir"])
	}
	logrus.Debug(fmt.Sprintf("AGENT created %s", context["config_dir"]))

	deployConfig := generateDeployConfig(c, context)

	logrus.Debug(fmt.Sprintf("AGENT creating %s", context["schema_id"]))
	if err := manageCluster(deployConfig); err != nil {
		return errors.Wrapf(err, "AGENT %s create failed", context["schema_id"])
	}
	logrus.Debug(fmt.Sprintf("AGENT %s create complete", context["schema_id"]))

	return nil
}

func clusterUpdate(c *Config, context map[string]string) error {
	deployConfig := generateDeployConfig(c, context)

	logrus.Debug(fmt.Sprintf("AGENT updating %s", context["schema_id"]))
	if err := manageCluster(deployConfig); err != nil {
		return errors.Wrapf(err, "AGENT %s update failed", context["schema_id"])
	}
	logrus.Debug(fmt.Sprintf("AGENT %s update complete", context["schema_id"]))

	return nil
}

func clusterDelete(c *Config, context map[string]string) error {
	deployConfig := generateDeployConfig(c, context)

	logrus.Debug(fmt.Sprintf("AGENT deleting %s", context["schema_id"]))
	if err := manageCluster(deployConfig); err != nil {
		return errors.Wrapf(err, "AGENT %s delete failed", context["schema_id"])
	}
	logrus.Debug(fmt.Sprintf("AGENT deleting %s", context["config_dir"]))

	if err := directoryHandler("delete", context["config_dir"]); err != nil {
		return errors.Wrapf(err, "AGENT delete %s failed", context["config_dir"])
	}
	logrus.Debug(fmt.Sprintf("AGENT deleted %s", context["config_dir"]))

	return nil
}

func generateDeployConfig(c *Config, context map[string]string) *deploy.Config {
	deployConfig := &deploy.Config{
		ID:                  c.ID,
		Password:            c.Password,
		DomainID:            c.DomainID,
		ProjectID:           c.ProjectID,
		DomainName:          c.DomainName,
		ProjectName:         c.ProjectName,
		AuthURL:             c.AuthURL,
		Endpoint:            c.Endpoint,
		InSecure:            c.InSecure,
		ResourceType:        context["schemaID"],
		ResourceID:          context["uuid"],
		Action:              context["operation"],
		ProvisionerType:     context["provisionerType"],
		LogLevel:            "debug",
		LogFile:             "/var/log/contrail/deploy.log",
		TemplateRoot:        "/usr/share/contrail/templates/",
		ServiceUserID:       c.ServiceUserID,
		ServiceUserPassword: c.ServiceUserPassword,
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
