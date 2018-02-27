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

//RESTCreateServiceConnectionModule handle a Create REST service.
func (service *ContrailService) RESTCreateServiceConnectionModule(c echo.Context) error {
	requestData := &models.CreateServiceConnectionModuleRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceConnectionModule(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceConnectionModule handle a Create API
func (service *ContrailService) CreateServiceConnectionModule(
	ctx context.Context,
	request *models.CreateServiceConnectionModuleRequest) (*models.CreateServiceConnectionModuleResponse, error) {
	model := request.ServiceConnectionModule
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

	return service.Next().CreateServiceConnectionModule(ctx, request)
}

//RESTUpdateServiceConnectionModule handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceConnectionModule(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceConnectionModuleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceConnectionModule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceConnectionModule handles a Update request.
func (service *ContrailService) UpdateServiceConnectionModule(
	ctx context.Context,
	request *models.UpdateServiceConnectionModuleRequest) (*models.UpdateServiceConnectionModuleResponse, error) {
	model := request.ServiceConnectionModule
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateServiceConnectionModule(ctx, request)
}

//RESTDeleteServiceConnectionModule delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceConnectionModule(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceConnectionModuleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceConnectionModule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetServiceConnectionModule a REST Get request.
func (service *ContrailService) RESTGetServiceConnectionModule(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceConnectionModuleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceConnectionModule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListServiceConnectionModule handles a List REST service Request.
func (service *ContrailService) RESTListServiceConnectionModule(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceConnectionModuleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceConnectionModule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
