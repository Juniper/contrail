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

//RESTCreateServiceHealthCheck handle a Create REST service.
func (service *ContrailService) RESTCreateServiceHealthCheck(c echo.Context) error {
	requestData := &models.CreateServiceHealthCheckRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_health_check",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceHealthCheck(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceHealthCheck handle a Create API
func (service *ContrailService) CreateServiceHealthCheck(
	ctx context.Context,
	request *models.CreateServiceHealthCheckRequest) (*models.CreateServiceHealthCheckResponse, error) {
	model := request.ServiceHealthCheck
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

	return service.Next().CreateServiceHealthCheck(ctx, request)
}

//RESTUpdateServiceHealthCheck handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceHealthCheck(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceHealthCheckRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_health_check",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceHealthCheck(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceHealthCheck handles a Update request.
func (service *ContrailService) UpdateServiceHealthCheck(
	ctx context.Context,
	request *models.UpdateServiceHealthCheckRequest) (*models.UpdateServiceHealthCheckResponse, error) {
	model := request.ServiceHealthCheck
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateServiceHealthCheck(ctx, request)
}

//RESTDeleteServiceHealthCheck delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceHealthCheck(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceHealthCheckRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceHealthCheck(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetServiceHealthCheck a REST Get request.
func (service *ContrailService) RESTGetServiceHealthCheck(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceHealthCheckRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceHealthCheck(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListServiceHealthCheck handles a List REST service Request.
func (service *ContrailService) RESTListServiceHealthCheck(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceHealthCheckRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceHealthCheck(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
