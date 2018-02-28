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

//RESTCreateServiceEndpoint handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateServiceEndpoint(c echo.Context) error {
	requestData := &models.CreateServiceEndpointRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_endpoint",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceEndpoint(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceEndpoint handle a Create API
// nolint
func (service *ContrailService) CreateServiceEndpoint(
	ctx context.Context,
	request *models.CreateServiceEndpointRequest) (*models.CreateServiceEndpointResponse, error) {
	model := request.ServiceEndpoint
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

	return service.Next().CreateServiceEndpoint(ctx, request)
}

//RESTUpdateServiceEndpoint handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateServiceEndpoint(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceEndpointRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_endpoint",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceEndpoint(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceEndpoint handles a Update request.
// nolint
func (service *ContrailService) UpdateServiceEndpoint(
	ctx context.Context,
	request *models.UpdateServiceEndpointRequest) (*models.UpdateServiceEndpointResponse, error) {
	model := request.ServiceEndpoint
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateServiceEndpoint(ctx, request)
}

//RESTDeleteServiceEndpoint delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteServiceEndpoint(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceEndpointRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceEndpoint(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetServiceEndpoint a REST Get request.
// nolint
func (service *ContrailService) RESTGetServiceEndpoint(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceEndpointRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceEndpoint(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListServiceEndpoint handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListServiceEndpoint(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListServiceEndpointRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceEndpoint(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
