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

func (h *eventHandler) handle(e *services.Event, c *Config) error {
	operation := strings.ToLower(e.Operation())

	var kind string

	switch e.Kind() {
	case "contrail-cluster", "rhospd-cloud-manager":
		clusterConfig := generateClusterConfig(e, c)
		manager, err := generateClusterManager(clustrConfig)
		if err != nil {
			return errors.Wrap(err, "could not generate cluster manager")
		}
		h.manager = manager
		kind = "cluster"
	case "cloud":
		cloudConfig := generateCloudConfig(e, c)
		manager, err := generateCloudManager(cloudConfig)
		manager, err := generateClusterManager(clustrConfig)
		if err != nil {
			return errors.Wrap(err, "could not generate cloud manager")
		}
		h.manager = manager
		kind = "cloud"
	}

	logrus.Debug(fmt.Sprintf("AGENT %s %s", operation, kind))
	h.manager = manager
	if err = h.manager.Manage(); err != nil {
		return errors.Wrapf(err, "AGENT %s %s failed", operation, kind)
	}
	logrus.Debug(fmt.Sprintf("AGENT %s %s complete", operation, kind))

	return nil
}

func generateClusterConfig(e *services.Event, c *Config) *deploy.Config {
	provisionerType, ok := e.GetResource().ToMap()["provisioner_type"].(string)
	if !ok {
		return errors.New("provisioner type conversion failed")
	}

	clusterConfig := &deploy.Config{
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
	}

	return clusterConfig
}

func generateClusterManager(c *deploy.Config) (*deploy.Deploy, error) {
	manager, err := deploy.NewDeploy(deployConfig)
	if err != nil {
		return nil, errors.Wrap(err, "cluster manager creation failed")
	}

	return manager, nil
}

func generateClusterConfig(e *services.Event, c *Config) *cloud.Config {

	cloudConfig := &cloud.Config{
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

	return cloudConfig
}

func generateCloudManager(c *cloud.Config) (*cloud.Cloud, error) {
	manager, err := cloud.NewCloudManager(cloudConfig)
	if err != nil {
		return nil, errors.Wrap(err, "cloud manager creation failed")
	}

	return manager, nil
}
