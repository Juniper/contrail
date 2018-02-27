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

//RESTCreateAPIAccessList handle a Create REST service.
func (service *ContrailService) RESTCreateAPIAccessList(c echo.Context) error {
	requestData := &models.CreateAPIAccessListRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "api_access_list",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAPIAccessList(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAPIAccessList handle a Create API
func (service *ContrailService) CreateAPIAccessList(
	ctx context.Context,
	request *models.CreateAPIAccessListRequest) (*models.CreateAPIAccessListResponse, error) {
	model := request.APIAccessList
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

	return service.Next().CreateAPIAccessList(ctx, request)
}

//RESTUpdateAPIAccessList handles a REST Update request.
func (service *ContrailService) RESTUpdateAPIAccessList(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAPIAccessListRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "api_access_list",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAPIAccessList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAPIAccessList handles a Update request.
func (service *ContrailService) UpdateAPIAccessList(
	ctx context.Context,
	request *models.UpdateAPIAccessListRequest) (*models.UpdateAPIAccessListResponse, error) {
	model := request.APIAccessList
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateAPIAccessList(ctx, request)
}

//RESTDeleteAPIAccessList delete a resource using REST service.
func (service *ContrailService) RESTDeleteAPIAccessList(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAPIAccessListRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAPIAccessList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetAPIAccessList a REST Get request.
func (service *ContrailService) RESTGetAPIAccessList(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAPIAccessListRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAPIAccessList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListAPIAccessList handles a List REST service Request.
func (service *ContrailService) RESTListAPIAccessList(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAPIAccessListRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAPIAccessList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
