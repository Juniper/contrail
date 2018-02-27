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

//RESTCreateServiceObject handle a Create REST service.
func (service *ContrailService) RESTCreateServiceObject(c echo.Context) error {
	requestData := &models.CreateServiceObjectRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceObject(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceObject handle a Create API
func (service *ContrailService) CreateServiceObject(
	ctx context.Context,
	request *models.CreateServiceObjectRequest) (*models.CreateServiceObjectResponse, error) {
	model := request.ServiceObject
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

	return service.Next().CreateServiceObject(ctx, request)
}

//RESTUpdateServiceObject handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceObject(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceObjectRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceObject handles a Update request.
func (service *ContrailService) UpdateServiceObject(
	ctx context.Context,
	request *models.UpdateServiceObjectRequest) (*models.UpdateServiceObjectResponse, error) {
	model := request.ServiceObject
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateServiceObject(ctx, request)
}

//RESTDeleteServiceObject delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceObject(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceObjectRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetServiceObject a REST Get request.
func (service *ContrailService) RESTGetServiceObject(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceObjectRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListServiceObject handles a List REST service Request.
func (service *ContrailService) RESTListServiceObject(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceObjectRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
