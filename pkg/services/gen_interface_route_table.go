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

//RESTCreateInterfaceRouteTable handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateInterfaceRouteTable(c echo.Context) error {
	requestData := &models.CreateInterfaceRouteTableRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateInterfaceRouteTable(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateInterfaceRouteTable handle a Create API
// nolint
func (service *ContrailService) CreateInterfaceRouteTable(
	ctx context.Context,
	request *models.CreateInterfaceRouteTableRequest) (*models.CreateInterfaceRouteTableResponse, error) {
	model := request.InterfaceRouteTable
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

	return service.Next().CreateInterfaceRouteTable(ctx, request)
}

//RESTUpdateInterfaceRouteTable handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateInterfaceRouteTable(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateInterfaceRouteTableRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateInterfaceRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateInterfaceRouteTable handles a Update request.
// nolint
func (service *ContrailService) UpdateInterfaceRouteTable(
	ctx context.Context,
	request *models.UpdateInterfaceRouteTableRequest) (*models.UpdateInterfaceRouteTableResponse, error) {
	model := request.InterfaceRouteTable
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateInterfaceRouteTable(ctx, request)
}

//RESTDeleteInterfaceRouteTable delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteInterfaceRouteTable(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteInterfaceRouteTableRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteInterfaceRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetInterfaceRouteTable a REST Get request.
// nolint
func (service *ContrailService) RESTGetInterfaceRouteTable(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetInterfaceRouteTableRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetInterfaceRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListInterfaceRouteTable handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListInterfaceRouteTable(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListInterfaceRouteTableRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListInterfaceRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
