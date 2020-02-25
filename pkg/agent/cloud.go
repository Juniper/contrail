package agent

import (
	"fmt"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func cloudHandler(e *services.Event, c *Config) error {
	context := make(map[string]string)
	context["schemaID"] = basemodels.KindToSchemaID(e.Kind())
	context["uuid"] = e.GetUUID()
	context["provisionerType"] = e.GetResource().ToMap()["provisioner_type"].(string)
	context["action"] = e.Operation()
	context["schema_id"] = basemodels.KindToSchemaID(e.Kind())
	context["config_dir"] = fmt.Sprintf("/var/tmp/%s/config/%s", context["schemaID"], context["uuid"])

	switch e.Operation() {
	case services.OperationCreate:
		return cloudCreate(c, context)
	case services.OperationUpdate:
		return cloudUpdate(c, context)
	case services.OperationDelete:
		return cloudDelete(c, context)
	}

	return nil
}

func cloudCreate(c *Config, context map[string]string) error {
	logrus.Debug(fmt.Sprintf("AGENT creating %s", context["config_dir"]))
	if err := directoryHandler("create", context["config_dir"]); err != nil {
		return errors.Wrapf(err, "AGENT create %s failed", context["config_dir"])
	}
	logrus.Debug(fmt.Sprintf("AGENT created %s", context["config_dir"]))

	cloudConfig := generateCloudConfig(c, context)

	logrus.Debug("AGENT creating cloud (contrailgo cloud -c %s/contrail-cloud-config.yml)")
	if err := manageCloud(cloudConfig); err != nil {
		return errors.Wrap(err, "AGENT cloud create failed")
	}
	logrus.Debug("AGENT cloud create complete")

	return nil
}

func cloudUpdate(c *Config, context map[string]string) error {
	cloudConfig := generateCloudConfig(c, context)

	logrus.Debug("AGENT updating cloud")
	if err := manageCloud(cloudConfig); err != nil {
		return errors.Wrap(err, "AGENT cloud update failed")
	}
	logrus.Debug("AGENT cloud update complete")

	return nil
}

func cloudDelete(c *Config, context map[string]string) error {
	cloudConfig := generateCloudConfig(c, context)

	logrus.Debug("AGENT deleting cloud")
	if err := manageCloud(cloudConfig); err != nil {
		return errors.Wrap(err, "AGENT cloud delete failed")
	}
	logrus.Debug("AGENT cloud delete complete")

	logrus.Debug(fmt.Sprintf("AGENT deleting %s", context["config_dir"]))
	if err := directoryHandler("delete", context["config_dir"]); err != nil {
		return errors.Wrapf(err, "AGENT delete %s FAILED", context["config_dir"])
	}
	logrus.Debug(fmt.Sprintf("AGENT deleted %s", context["config_dir"]))

	return nil
}

func generateCloudConfig(c *Config, context map[string]string) *cloud.Config {
	deployConfig := &cloud.Config{
		ID:           c.ID,
		Password:     c.Password,
		DomainID:     c.DomainID,
		ProjectID:    c.ProjectID,
		DomainName:   c.DomainName,
		ProjectName:  c.ProjectName,
		AuthURL:      c.AuthURL,
		Endpoint:     c.Endpoint,
		InSecure:     c.InSecure,
		CloudID:      context["uuid"],
		Action:       context["operation"],
		LogLevel:     "debug",
		LogFile:      "/var/log/contrail/cloud.log",
		TemplateRoot: "/usr/share/contrail/templates/",
	}

	return deployConfig
}

func manageCloud(c *cloud.Config) error {
	manager, err := cloud.NewCloudConfigAndCommandExecutor(c, &osCommandExecutor{})
	if err != nil {
		return errors.Wrap(err, "cloud manager creation failed")
	}
	if err = manager.Manage(); err != nil {
		return errors.Wrap(err, "cloud management failed")
	}

	return nil
}
