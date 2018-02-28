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

//RESTCreateRouteAggregate handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateRouteAggregate(c echo.Context) error {
	requestData := &models.CreateRouteAggregateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_aggregate",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateRouteAggregate(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateRouteAggregate handle a Create API
// nolint
func (service *ContrailService) CreateRouteAggregate(
	ctx context.Context,
	request *models.CreateRouteAggregateRequest) (*models.CreateRouteAggregateResponse, error) {
	model := request.RouteAggregate
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

	return service.Next().CreateRouteAggregate(ctx, request)
}

//RESTUpdateRouteAggregate handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateRouteAggregate(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateRouteAggregateRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_aggregate",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateRouteAggregate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateRouteAggregate handles a Update request.
// nolint
func (service *ContrailService) UpdateRouteAggregate(
	ctx context.Context,
	request *models.UpdateRouteAggregateRequest) (*models.UpdateRouteAggregateResponse, error) {
	model := request.RouteAggregate
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateRouteAggregate(ctx, request)
}

//RESTDeleteRouteAggregate delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteRouteAggregate(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteRouteAggregateRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteRouteAggregate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetRouteAggregate a REST Get request.
// nolint
func (service *ContrailService) RESTGetRouteAggregate(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetRouteAggregateRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetRouteAggregate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListRouteAggregate handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListRouteAggregate(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListRouteAggregateRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListRouteAggregate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
