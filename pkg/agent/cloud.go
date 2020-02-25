package agent

import (
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (h *eventHandler) processCloud(e *services.Event, c *Config) error {
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
		return createCloud(v)
	case services.OperationUpdate:
		return updateCloud(v)
	case services.OperationDelete:
		return deleteCloud(v)
	}

	return nil
}

func createCloud(v *variables) error {
	cloudConfig := generateCloudConfig(v)

	logrus.Debug("AGENT creating cloud")
	if err := manageCloud(cloudConfig); err != nil {
		return errors.Wrap(err, "AGENT cloud creation failed")
	}
	logrus.Debug("AGENT cloud creation complete")

	return nil
}

func updateCloud(v *variables) error {
	cloudConfig := generateCloudConfig(v)

	logrus.Debug("AGENT updating cloud")
	if err := manageCloud(cloudConfig); err != nil {
		return errors.Wrap(err, "AGENT cloud update failed")
	}
	logrus.Debug("AGENT cloud update complete")

	return nil
}

func deleteCloud(v *variables) error {
	cloudConfig := generateCloudConfig(v)

	logrus.Debug("AGENT deleting cloud")
	if err := manageCloud(cloudConfig); err != nil {
		return errors.Wrap(err, "AGENT cloud delete failed")
	}
	logrus.Debug("AGENT cloud deletion complete")

	return nil
}

func generateCloudConfig(v *variables) *cloud.Config {
	deployConfig := &cloud.Config{
		ID:           v.config.ID,
		Password:     v.config.Password,
		DomainID:     v.config.DomainID,
		ProjectID:    v.config.ProjectID,
		DomainName:   v.config.DomainName,
		ProjectName:  v.config.ProjectName,
		AuthURL:      v.config.AuthURL,
		Endpoint:     v.config.Endpoint,
		InSecure:     v.config.InSecure,
		CloudID:      v.uuid,
		Action:       v.action,
		LogLevel:     "debug",
		LogFile:      "/var/log/contrail/cloud.log",
		TemplateRoot: "/usr/share/contrail/templates/",
	}

	return deployConfig
}

func manageCloud(c *cloud.Config) error {
	manager, err := cloud.NewCloudManager(c)
	if err != nil {
		return errors.Wrap(err, "cloud manager creation failed")
	}
	if err = manager.Manage(); err != nil {
		return errors.Wrap(err, "cloud management failed")
	}

	return nil
}
