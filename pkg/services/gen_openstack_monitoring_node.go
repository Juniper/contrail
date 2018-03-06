package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateOpenstackMonitoringNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateOpenstackMonitoringNode(c echo.Context) error {
	requestData := &models.CreateOpenstackMonitoringNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_monitoring_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateOpenstackMonitoringNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateOpenstackMonitoringNode handle a Create API
// nolint
func (service *ContrailService) CreateOpenstackMonitoringNode(
	ctx context.Context,
	request *models.CreateOpenstackMonitoringNodeRequest) (*models.CreateOpenstackMonitoringNodeResponse, error) {
	model := request.OpenstackMonitoringNode
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}

	if model.FQName == nil {
		if model.DisplayName != "" {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
		} else {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.UUID}
		}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()

	return service.Next().CreateOpenstackMonitoringNode(ctx, request)
}

//RESTUpdateOpenstackMonitoringNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateOpenstackMonitoringNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateOpenstackMonitoringNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_monitoring_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateOpenstackMonitoringNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackMonitoringNode handles a Update request.
// nolint
func (service *ContrailService) UpdateOpenstackMonitoringNode(
	ctx context.Context,
	request *models.UpdateOpenstackMonitoringNodeRequest) (*models.UpdateOpenstackMonitoringNodeResponse, error) {
	model := request.OpenstackMonitoringNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateOpenstackMonitoringNode(ctx, request)
}

//RESTDeleteOpenstackMonitoringNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteOpenstackMonitoringNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteOpenstackMonitoringNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteOpenstackMonitoringNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetOpenstackMonitoringNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetOpenstackMonitoringNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetOpenstackMonitoringNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetOpenstackMonitoringNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListOpenstackMonitoringNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListOpenstackMonitoringNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListOpenstackMonitoringNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListOpenstackMonitoringNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
