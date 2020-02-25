package agent

import (
	"fmt"
	"strings"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/deploy"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type eventHandler struct {
	clusterLog *logrus.Entry
	cloudLog   *logrus.Entry
}

func newEventHandler() *eventHandler {
	return &eventHandler{
		clusterLog: logutil.NewLogger("agent-event-handler"),
		cloudLog:   logutil.NewLogger("agent-event-handler"),
	}
}

func (h *eventHandler) handleCluster(e *services.Event, c *Config) error {
	h.clusterLog.Info("Creating cluster config")
	clusterConfig, err := generateClusterConfig(e, c)
	if err != nil {
		return errors.Wrap(err, "creation of cluster config failed")
	}
	h.clusterLog.Info("Created cluster config")

	h.clusterLog.Info("Creating cluster manager")
	manager, err := deploy.NewDeploy(clusterConfig)
	if err != nil {
		return errors.Wrap(err, "could not create cluster manager")
	}
	h.clusterLog.Infof("Created cluster manager")

	h.clusterLog.Info("Running cluster manager")
	operation := strings.ToLower(e.Operation())
	kind := e.Kind()
	logrus.Debug(fmt.Sprintf("AGENT %s %s", operation, kind))
	if err := manager.Manage(); err != nil {
		return errors.Wrapf(err, "AGENT %s %s failed", operation, kind)
	}
	logrus.Debug(fmt.Sprintf("AGENT %s %s complete", operation, kind))
	h.clusterLog.Info("Finished running cluster manager")

	return nil
}

func (h *eventHandler) handleCloud(e *services.Event, c *Config) error {
	h.cloudLog.Info("Creating cloud config")
	cloudConfig := generateCloudConfig(e, c)
	h.cloudLog.Info("Created cloud config")

	h.cloudLog.Info("Creating cloud manager")
	manager, err := cloud.NewCloudManager(cloudConfig)
	if err != nil {
		return errors.Wrap(err, "could not generate cloud manager")
	}
	h.cloudLog.Info("Created cloud manager")

	h.cloudLog.Info("Running cloud manager")
	operation := strings.ToLower(e.Operation())
	kind := e.Kind()
	logrus.Debug(fmt.Sprintf("AGENT %s %s", operation, kind))
	if err := manager.Manage(); err != nil {
		return errors.Wrapf(err, "AGENT %s %s failed", operation, kind)
	}
	logrus.Debug(fmt.Sprintf("AGENT %s %s complete", operation, kind))
	h.cloudLog.Info("Finished running cloud manager")

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
		ResourceType:        strings.Replace(e.Kind(), "-", "_", -1),
		ResourceID:          e.GetUUID(),
		Action:              strings.ToLower(e.Operation()),
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
		Action:       strings.ToLower(e.Operation()),
		LogLevel:     "debug",
		LogFile:      "/var/log/contrail/cloud.log",
		TemplateRoot: "/usr/share/contrail/templates/",
	}
}
