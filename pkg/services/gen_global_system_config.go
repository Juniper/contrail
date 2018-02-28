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

//RESTCreateGlobalSystemConfig handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateGlobalSystemConfig(c echo.Context) error {
	requestData := &models.CreateGlobalSystemConfigRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_system_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateGlobalSystemConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateGlobalSystemConfig handle a Create API
// nolint
func (service *ContrailService) CreateGlobalSystemConfig(
	ctx context.Context,
	request *models.CreateGlobalSystemConfigRequest) (*models.CreateGlobalSystemConfigResponse, error) {
	model := request.GlobalSystemConfig
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

	return service.Next().CreateGlobalSystemConfig(ctx, request)
}

//RESTUpdateGlobalSystemConfig handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateGlobalSystemConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateGlobalSystemConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_system_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateGlobalSystemConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateGlobalSystemConfig handles a Update request.
// nolint
func (service *ContrailService) UpdateGlobalSystemConfig(
	ctx context.Context,
	request *models.UpdateGlobalSystemConfigRequest) (*models.UpdateGlobalSystemConfigResponse, error) {
	model := request.GlobalSystemConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateGlobalSystemConfig(ctx, request)
}

//RESTDeleteGlobalSystemConfig delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteGlobalSystemConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteGlobalSystemConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteGlobalSystemConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetGlobalSystemConfig a REST Get request.
// nolint
func (service *ContrailService) RESTGetGlobalSystemConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetGlobalSystemConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetGlobalSystemConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListGlobalSystemConfig handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListGlobalSystemConfig(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListGlobalSystemConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListGlobalSystemConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
