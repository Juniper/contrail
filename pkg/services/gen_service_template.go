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

//RESTCreateServiceTemplate handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateServiceTemplate(c echo.Context) error {
	requestData := &models.CreateServiceTemplateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_template",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceTemplate(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceTemplate handle a Create API
// nolint
func (service *ContrailService) CreateServiceTemplate(
	ctx context.Context,
	request *models.CreateServiceTemplateRequest) (*models.CreateServiceTemplateResponse, error) {
	model := request.ServiceTemplate
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

	return service.Next().CreateServiceTemplate(ctx, request)
}

//RESTUpdateServiceTemplate handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateServiceTemplate(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceTemplateRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_template",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceTemplate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceTemplate handles a Update request.
// nolint
func (service *ContrailService) UpdateServiceTemplate(
	ctx context.Context,
	request *models.UpdateServiceTemplateRequest) (*models.UpdateServiceTemplateResponse, error) {
	model := request.ServiceTemplate
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateServiceTemplate(ctx, request)
}

//RESTDeleteServiceTemplate delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteServiceTemplate(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceTemplateRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceTemplate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetServiceTemplate a REST Get request.
// nolint
func (service *ContrailService) RESTGetServiceTemplate(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceTemplateRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceTemplate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListServiceTemplate handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListServiceTemplate(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListServiceTemplateRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceTemplate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
