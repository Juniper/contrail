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

//RESTCreateLoadbalancerHealthmonitor handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateLoadbalancerHealthmonitor(c echo.Context) error {
	requestData := &models.CreateLoadbalancerHealthmonitorRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLoadbalancerHealthmonitor(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerHealthmonitor handle a Create API
// nolint
func (service *ContrailService) CreateLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.CreateLoadbalancerHealthmonitorRequest) (*models.CreateLoadbalancerHealthmonitorResponse, error) {
	model := request.LoadbalancerHealthmonitor
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

	return service.Next().CreateLoadbalancerHealthmonitor(ctx, request)
}

//RESTUpdateLoadbalancerHealthmonitor handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateLoadbalancerHealthmonitor(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLoadbalancerHealthmonitorRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLoadbalancerHealthmonitor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerHealthmonitor handles a Update request.
// nolint
func (service *ContrailService) UpdateLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.UpdateLoadbalancerHealthmonitorRequest) (*models.UpdateLoadbalancerHealthmonitorResponse, error) {
	model := request.LoadbalancerHealthmonitor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateLoadbalancerHealthmonitor(ctx, request)
}

//RESTDeleteLoadbalancerHealthmonitor delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteLoadbalancerHealthmonitor(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLoadbalancerHealthmonitorRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLoadbalancerHealthmonitor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetLoadbalancerHealthmonitor a REST Get request.
// nolint
func (service *ContrailService) RESTGetLoadbalancerHealthmonitor(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLoadbalancerHealthmonitorRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLoadbalancerHealthmonitor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListLoadbalancerHealthmonitor handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListLoadbalancerHealthmonitor(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListLoadbalancerHealthmonitorRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLoadbalancerHealthmonitor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
