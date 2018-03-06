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

//RESTCreatePhysicalRouter handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreatePhysicalRouter(c echo.Context) error {
	requestData := &models.CreatePhysicalRouterRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_router",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreatePhysicalRouter(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreatePhysicalRouter handle a Create API
// nolint
func (service *ContrailService) CreatePhysicalRouter(
	ctx context.Context,
	request *models.CreatePhysicalRouterRequest) (*models.CreatePhysicalRouterResponse, error) {
	model := request.PhysicalRouter
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

	return service.Next().CreatePhysicalRouter(ctx, request)
}

//RESTUpdatePhysicalRouter handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdatePhysicalRouter(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdatePhysicalRouterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_router",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdatePhysicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdatePhysicalRouter handles a Update request.
// nolint
func (service *ContrailService) UpdatePhysicalRouter(
	ctx context.Context,
	request *models.UpdatePhysicalRouterRequest) (*models.UpdatePhysicalRouterResponse, error) {
	model := request.PhysicalRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdatePhysicalRouter(ctx, request)
}

//RESTDeletePhysicalRouter delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeletePhysicalRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeletePhysicalRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeletePhysicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetPhysicalRouter a REST Get request.
// nolint
func (service *ContrailService) RESTGetPhysicalRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetPhysicalRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetPhysicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListPhysicalRouter handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListPhysicalRouter(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListPhysicalRouterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListPhysicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
