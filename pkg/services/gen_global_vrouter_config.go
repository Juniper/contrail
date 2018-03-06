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

//RESTCreateGlobalVrouterConfig handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateGlobalVrouterConfig(c echo.Context) error {
	requestData := &models.CreateGlobalVrouterConfigRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_vrouter_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateGlobalVrouterConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateGlobalVrouterConfig handle a Create API
// nolint
func (service *ContrailService) CreateGlobalVrouterConfig(
	ctx context.Context,
	request *models.CreateGlobalVrouterConfigRequest) (*models.CreateGlobalVrouterConfigResponse, error) {
	model := request.GlobalVrouterConfig
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

	return service.Next().CreateGlobalVrouterConfig(ctx, request)
}

//RESTUpdateGlobalVrouterConfig handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateGlobalVrouterConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateGlobalVrouterConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_vrouter_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateGlobalVrouterConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateGlobalVrouterConfig handles a Update request.
// nolint
func (service *ContrailService) UpdateGlobalVrouterConfig(
	ctx context.Context,
	request *models.UpdateGlobalVrouterConfigRequest) (*models.UpdateGlobalVrouterConfigResponse, error) {
	model := request.GlobalVrouterConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateGlobalVrouterConfig(ctx, request)
}

//RESTDeleteGlobalVrouterConfig delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteGlobalVrouterConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteGlobalVrouterConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteGlobalVrouterConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetGlobalVrouterConfig a REST Get request.
// nolint
func (service *ContrailService) RESTGetGlobalVrouterConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetGlobalVrouterConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetGlobalVrouterConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListGlobalVrouterConfig handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListGlobalVrouterConfig(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListGlobalVrouterConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListGlobalVrouterConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
