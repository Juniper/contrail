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

//RESTCreateNetworkDeviceConfig handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateNetworkDeviceConfig(c echo.Context) error {
	requestData := &models.CreateNetworkDeviceConfigRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_device_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateNetworkDeviceConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateNetworkDeviceConfig handle a Create API
// nolint
func (service *ContrailService) CreateNetworkDeviceConfig(
	ctx context.Context,
	request *models.CreateNetworkDeviceConfigRequest) (*models.CreateNetworkDeviceConfigResponse, error) {
	model := request.NetworkDeviceConfig
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

	return service.Next().CreateNetworkDeviceConfig(ctx, request)
}

//RESTUpdateNetworkDeviceConfig handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateNetworkDeviceConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateNetworkDeviceConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_device_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateNetworkDeviceConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateNetworkDeviceConfig handles a Update request.
// nolint
func (service *ContrailService) UpdateNetworkDeviceConfig(
	ctx context.Context,
	request *models.UpdateNetworkDeviceConfigRequest) (*models.UpdateNetworkDeviceConfigResponse, error) {
	model := request.NetworkDeviceConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateNetworkDeviceConfig(ctx, request)
}

//RESTDeleteNetworkDeviceConfig delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteNetworkDeviceConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteNetworkDeviceConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteNetworkDeviceConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetNetworkDeviceConfig a REST Get request.
// nolint
func (service *ContrailService) RESTGetNetworkDeviceConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetNetworkDeviceConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetNetworkDeviceConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListNetworkDeviceConfig handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListNetworkDeviceConfig(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListNetworkDeviceConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListNetworkDeviceConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
