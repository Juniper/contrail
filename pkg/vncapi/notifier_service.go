package vncapi

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
)

// NotifierService is a service that proxies events to VNC API.
type NotifierService struct {
	services.BaseService
	client *client.HTTP
	log    *log.Entry
}

// Config is NotifierService config.
type Config struct {
	Endpoint string
}

// NewNotifierService makes a NotifierService service.
func NewNotifierService(c *Config) *NotifierService {
	client := client.NewHTTP(
		c.Endpoint,
		"",
		"",
		"",
		"",
		true,
		nil,
	)
	client.Init()

	return &NotifierService{
		client: client,
		log:    pkglog.NewLogger("vnc-api-notifier"),
	}
}

// CreateProject creates Project in VNC API.
func (ns *NotifierService) CreateProject(
	ctx context.Context, request *services.CreateProjectRequest,
) (*services.CreateProjectResponse, error) {
	response, err := ns.BaseService.CreateProject(ctx, request)
	if err != nil {
		return nil, err
	}

	vncResponse, err := ns.client.CreateProject(ctx, &services.CreateProjectRequest{
		Project: response.GetProject(),
	})
	if err != nil {
		ns.log.WithError(err).WithFields(log.Fields{
			"uuid":        response.GetProject().UUID,
			"vncResponse": vncResponse,
		}).Error("Failed to create Project in VNC API")
	} else {
		ns.log.WithFields(log.Fields{
			"uuid": response.GetProject().UUID,
		}).Debug("Project created in VNC API")
	}

	return response, nil
}

// TODO: generate for all resources

// UpdateProject updates Project in VNC API.
func (ns *NotifierService) UpdateProject(
	ctx context.Context, request *services.UpdateProjectRequest,
) (*services.UpdateProjectResponse, error) {
	response, err := ns.BaseService.UpdateProject(ctx, request)
	if err != nil {
		return nil, err
	}

	// TODO: apply field mask
	vncResponse, err := ns.client.UpdateProject(ctx, &services.UpdateProjectRequest{
		Project: response.GetProject(),
	})
	if err != nil {
		ns.log.WithError(err).WithFields(log.Fields{
			"uuid":        response.GetProject().UUID,
			"vncResponse": vncResponse,
		}).Error("Failed to update Project in VNC API")
	} else {
		ns.log.WithFields(log.Fields{
			"uuid": response.GetProject().UUID,
		}).Debug("Project updated in VNC API")
	}

	return response, nil
}

// DeleteProject deletes Project in VNC API.
func (ns *NotifierService) DeleteProject(
	ctx context.Context, request *services.DeleteProjectRequest,
) (*services.DeleteProjectResponse, error) {
	response, err := ns.BaseService.DeleteProject(ctx, request)
	if err != nil {
		return nil, err
	}

	vncResponse, err := ns.client.DeleteProject(ctx, &services.DeleteProjectRequest{
		ID: response.ID,
	})
	if err != nil {
		ns.log.WithError(err).WithFields(log.Fields{
			"uuid":        response.ID,
			"vncResponse": vncResponse,
		}).Error("Failed to delete Project in VNC API")
	} else {
		ns.log.WithFields(log.Fields{
			"uuid": response.ID,
		}).Debug("Project deleted in VNC API")
	}

	return response, nil
}
