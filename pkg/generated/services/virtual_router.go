package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateVirtualRouter handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualRouter(c echo.Context) error {
	requestData := &models.CreateVirtualRouterRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_router",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualRouter(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualRouter handle a Create API
func (service *ContrailService) CreateVirtualRouter(
	ctx context.Context,
	request *models.CreateVirtualRouterRequest) (*models.CreateVirtualRouterResponse, error) {
	model := request.VirtualRouter
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

	return service.Next().CreateVirtualRouter(ctx, request)
}

//RESTUpdateVirtualRouter handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualRouter(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualRouterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_router",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualRouter handles a Update request.
func (service *ContrailService) UpdateVirtualRouter(
	ctx context.Context,
	request *models.UpdateVirtualRouterRequest) (*models.UpdateVirtualRouterResponse, error) {
	model := request.VirtualRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateVirtualRouter(ctx, request)
}

//RESTDeleteVirtualRouter delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetVirtualRouter a REST Get request.
func (service *ContrailService) RESTGetVirtualRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListVirtualRouter handles a List REST service Request.
func (service *ContrailService) RESTListVirtualRouter(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVirtualRouterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
