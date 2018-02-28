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

//RESTCreateRouteTable handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateRouteTable(c echo.Context) error {
	requestData := &models.CreateRouteTableRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_table",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateRouteTable(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateRouteTable handle a Create API
// nolint
func (service *ContrailService) CreateRouteTable(
	ctx context.Context,
	request *models.CreateRouteTableRequest) (*models.CreateRouteTableResponse, error) {
	model := request.RouteTable
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

	return service.Next().CreateRouteTable(ctx, request)
}

//RESTUpdateRouteTable handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateRouteTable(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateRouteTableRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_table",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateRouteTable handles a Update request.
// nolint
func (service *ContrailService) UpdateRouteTable(
	ctx context.Context,
	request *models.UpdateRouteTableRequest) (*models.UpdateRouteTableResponse, error) {
	model := request.RouteTable
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateRouteTable(ctx, request)
}

//RESTDeleteRouteTable delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteRouteTable(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteRouteTableRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetRouteTable a REST Get request.
// nolint
func (service *ContrailService) RESTGetRouteTable(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetRouteTableRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListRouteTable handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListRouteTable(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListRouteTableRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
