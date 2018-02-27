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

//RESTCreateConfigRoot handle a Create REST service.
func (service *ContrailService) RESTCreateConfigRoot(c echo.Context) error {
	requestData := &models.CreateConfigRootRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_root",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateConfigRoot(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateConfigRoot handle a Create API
func (service *ContrailService) CreateConfigRoot(
	ctx context.Context,
	request *models.CreateConfigRootRequest) (*models.CreateConfigRootResponse, error) {
	model := request.ConfigRoot
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

	return service.Next().CreateConfigRoot(ctx, request)
}

//RESTUpdateConfigRoot handles a REST Update request.
func (service *ContrailService) RESTUpdateConfigRoot(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateConfigRootRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_root",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateConfigRoot(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateConfigRoot handles a Update request.
func (service *ContrailService) UpdateConfigRoot(
	ctx context.Context,
	request *models.UpdateConfigRootRequest) (*models.UpdateConfigRootResponse, error) {
	model := request.ConfigRoot
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateConfigRoot(ctx, request)
}

//RESTDeleteConfigRoot delete a resource using REST service.
func (service *ContrailService) RESTDeleteConfigRoot(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteConfigRootRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteConfigRoot(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetConfigRoot a REST Get request.
func (service *ContrailService) RESTGetConfigRoot(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetConfigRootRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetConfigRoot(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListConfigRoot handles a List REST service Request.
func (service *ContrailService) RESTListConfigRoot(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListConfigRootRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListConfigRoot(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
