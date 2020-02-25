package agent

import (
	"fmt"
	"strings"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/deploy"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type manager interface {
	Manage() error
}

type eventHandler struct {
	manager manager
}

func newEventHandler() *eventHandler {
	return &eventHandler{}
}

func (h *eventHandler) handleCluster(e *services.Event, c *Config) error {
	clusterConfig, err := generateClusterConfig(e, c)
	if err != nil {
		return errors.Wrap(err, "generate cluster config failed")
	}
	manager, err := deploy.NewDeploy(clusterConfig)
	if err != nil {
		return errors.Wrap(err, "could not generate cluster manager")
	}
	h.manager = manager

	operation := strings.ToLower(e.Operation())
	kind := e.Kind()
	logrus.Debug(fmt.Sprintf("AGENT %s %s", operation, kind))
	if err := h.manager.Manage(); err != nil {
		return errors.Wrapf(err, "AGENT %s %s failed", operation, kind)
	}
	logrus.Debug(fmt.Sprintf("AGENT %s %s complete", operation, kind))

	return nil
}

func (h *eventHandler) handleCloud(e *services.Event, c *Config) error {
	cloudConfig := generateCloudConfig(e, c)
	manager, err := cloud.NewCloudManager(cloudConfig)
	if err != nil {
		return errors.Wrap(err, "could not generate cloud manager")
	}
	h.manager = manager

	operation := strings.ToLower(e.Operation())
	kind := e.Kind()
	logrus.Debug(fmt.Sprintf("AGENT %s %s", operation, kind))
	if err := h.manager.Manage(); err != nil {
		return errors.Wrapf(err, "AGENT %s %s failed", operation, kind)
	}
	logrus.Debug(fmt.Sprintf("AGENT %s %s complete", operation, kind))

	return nil
}

func generateClusterConfig(e *services.Event, c *Config) (*deploy.Config, error) {
	provisionerType, ok := e.GetResource().ToMap()["provisioner_type"].(string)
	if !ok {
		return nil, errors.New("provisioner type conversion failed")
	}

	return &deploy.Config{
		ID:                  c.ID,
		Password:            c.Password,
		DomainID:            c.DomainID,
		ProjectID:           c.ProjectID,
		DomainName:          c.DomainName,
		ProjectName:         c.ProjectName,
		AuthURL:             c.AuthURL,
		Endpoint:            c.Endpoint,
		InSecure:            c.InSecure,
		ResourceType:        basemodels.KindToSchemaID(e.Kind()),
		ResourceID:          e.GetUUID(),
		Action:              e.Operation(),
		ProvisionerType:     provisionerType,
		LogLevel:            "debug",
		LogFile:             "/var/log/contrail/deploy.log",
		TemplateRoot:        "/usr/share/contrail/templates/",
		ServiceUserID:       c.ServiceUserID,
		ServiceUserPassword: c.ServiceUserPassword,
	}, nil
}

func generateCloudConfig(e *services.Event, c *Config) *cloud.Config {
	return &cloud.Config{
		ID:           c.ID,
		Password:     c.Password,
		DomainID:     c.DomainID,
		ProjectID:    c.ProjectID,
		DomainName:   c.DomainName,
		ProjectName:  c.ProjectName,
		AuthURL:      c.AuthURL,
		Endpoint:     c.Endpoint,
		InSecure:     c.InSecure,
		CloudID:      e.GetUUID(),
		Action:       e.Operation(),
		LogLevel:     "debug",
		LogFile:      "/var/log/contrail/cloud.log",
		TemplateRoot: "/usr/share/contrail/templates/",
	}
}
