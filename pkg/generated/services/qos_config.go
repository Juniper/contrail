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

//RESTCreateQosConfig handle a Create REST service.
func (service *ContrailService) RESTCreateQosConfig(c echo.Context) error {
	requestData := &models.CreateQosConfigRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateQosConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateQosConfig handle a Create API
func (service *ContrailService) CreateQosConfig(
	ctx context.Context,
	request *models.CreateQosConfigRequest) (*models.CreateQosConfigResponse, error) {
	model := request.QosConfig
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

	return service.Next().CreateQosConfig(ctx, request)
}

//RESTUpdateQosConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateQosConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateQosConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateQosConfig handles a Update request.
func (service *ContrailService) UpdateQosConfig(
	ctx context.Context,
	request *models.UpdateQosConfigRequest) (*models.UpdateQosConfigResponse, error) {
	model := request.QosConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateQosConfig(ctx, request)
}

//RESTDeleteQosConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteQosConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteQosConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetQosConfig a REST Get request.
func (service *ContrailService) RESTGetQosConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetQosConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListQosConfig handles a List REST service Request.
func (service *ContrailService) RESTListQosConfig(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListQosConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
