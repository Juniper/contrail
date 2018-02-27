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

//RESTCreateLogicalRouter handle a Create REST service.
func (service *ContrailService) RESTCreateLogicalRouter(c echo.Context) error {
	requestData := &models.CreateLogicalRouterRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_router",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLogicalRouter(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLogicalRouter handle a Create API
func (service *ContrailService) CreateLogicalRouter(
	ctx context.Context,
	request *models.CreateLogicalRouterRequest) (*models.CreateLogicalRouterResponse, error) {
	model := request.LogicalRouter
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

	return service.Next().CreateLogicalRouter(ctx, request)
}

//RESTUpdateLogicalRouter handles a REST Update request.
func (service *ContrailService) RESTUpdateLogicalRouter(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLogicalRouterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_router",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLogicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLogicalRouter handles a Update request.
func (service *ContrailService) UpdateLogicalRouter(
	ctx context.Context,
	request *models.UpdateLogicalRouterRequest) (*models.UpdateLogicalRouterResponse, error) {
	model := request.LogicalRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateLogicalRouter(ctx, request)
}

//RESTDeleteLogicalRouter delete a resource using REST service.
func (service *ContrailService) RESTDeleteLogicalRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLogicalRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLogicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetLogicalRouter a REST Get request.
func (service *ContrailService) RESTGetLogicalRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLogicalRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLogicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListLogicalRouter handles a List REST service Request.
func (service *ContrailService) RESTListLogicalRouter(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLogicalRouterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLogicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
