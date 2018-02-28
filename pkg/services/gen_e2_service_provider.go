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

//RESTCreateE2ServiceProvider handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateE2ServiceProvider(c echo.Context) error {
	requestData := &models.CreateE2ServiceProviderRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateE2ServiceProvider(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateE2ServiceProvider handle a Create API
// nolint
func (service *ContrailService) CreateE2ServiceProvider(
	ctx context.Context,
	request *models.CreateE2ServiceProviderRequest) (*models.CreateE2ServiceProviderResponse, error) {
	model := request.E2ServiceProvider
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

	return service.Next().CreateE2ServiceProvider(ctx, request)
}

//RESTUpdateE2ServiceProvider handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateE2ServiceProvider(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateE2ServiceProviderRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateE2ServiceProvider(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateE2ServiceProvider handles a Update request.
// nolint
func (service *ContrailService) UpdateE2ServiceProvider(
	ctx context.Context,
	request *models.UpdateE2ServiceProviderRequest) (*models.UpdateE2ServiceProviderResponse, error) {
	model := request.E2ServiceProvider
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateE2ServiceProvider(ctx, request)
}

//RESTDeleteE2ServiceProvider delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteE2ServiceProvider(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteE2ServiceProviderRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteE2ServiceProvider(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetE2ServiceProvider a REST Get request.
// nolint
func (service *ContrailService) RESTGetE2ServiceProvider(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetE2ServiceProviderRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetE2ServiceProvider(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListE2ServiceProvider handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListE2ServiceProvider(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListE2ServiceProviderRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListE2ServiceProvider(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
