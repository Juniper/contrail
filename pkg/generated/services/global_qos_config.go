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

//RESTCreateGlobalQosConfig handle a Create REST service.
func (service *ContrailService) RESTCreateGlobalQosConfig(c echo.Context) error {
	requestData := &models.CreateGlobalQosConfigRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_qos_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateGlobalQosConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateGlobalQosConfig handle a Create API
func (service *ContrailService) CreateGlobalQosConfig(
	ctx context.Context,
	request *models.CreateGlobalQosConfigRequest) (*models.CreateGlobalQosConfigResponse, error) {
	model := request.GlobalQosConfig
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

	return service.Next().CreateGlobalQosConfig(ctx, request)
}

//RESTUpdateGlobalQosConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateGlobalQosConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateGlobalQosConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_qos_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateGlobalQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateGlobalQosConfig handles a Update request.
func (service *ContrailService) UpdateGlobalQosConfig(
	ctx context.Context,
	request *models.UpdateGlobalQosConfigRequest) (*models.UpdateGlobalQosConfigResponse, error) {
	model := request.GlobalQosConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateGlobalQosConfig(ctx, request)
}

//RESTDeleteGlobalQosConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteGlobalQosConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteGlobalQosConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteGlobalQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetGlobalQosConfig a REST Get request.
func (service *ContrailService) RESTGetGlobalQosConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetGlobalQosConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetGlobalQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListGlobalQosConfig handles a List REST service Request.
func (service *ContrailService) RESTListGlobalQosConfig(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListGlobalQosConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListGlobalQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
