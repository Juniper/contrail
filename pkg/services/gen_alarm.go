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

//RESTCreateAlarm handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateAlarm(c echo.Context) error {
	requestData := &models.CreateAlarmRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alarm",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAlarm(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAlarm handle a Create API
// nolint
func (service *ContrailService) CreateAlarm(
	ctx context.Context,
	request *models.CreateAlarmRequest) (*models.CreateAlarmResponse, error) {
	model := request.Alarm
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

	return service.Next().CreateAlarm(ctx, request)
}

//RESTUpdateAlarm handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateAlarm(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAlarmRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alarm",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAlarm(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAlarm handles a Update request.
// nolint
func (service *ContrailService) UpdateAlarm(
	ctx context.Context,
	request *models.UpdateAlarmRequest) (*models.UpdateAlarmResponse, error) {
	model := request.Alarm
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateAlarm(ctx, request)
}

//RESTDeleteAlarm delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteAlarm(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAlarmRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAlarm(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetAlarm a REST Get request.
// nolint
func (service *ContrailService) RESTGetAlarm(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAlarmRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAlarm(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListAlarm handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListAlarm(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListAlarmRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAlarm(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
