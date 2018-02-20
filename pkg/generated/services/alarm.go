package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//RESTAlarmUpdateRequest for update request for REST.
type RESTAlarmUpdateRequest struct {
	Data map[string]interface{} `json:"alarm"`
}

//RESTCreateAlarm handle a Create REST service.
func (service *ContrailService) RESTCreateAlarm(c echo.Context) error {
	requestData := &models.CreateAlarmRequest{
		Alarm: models.MakeAlarm(),
	}
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
		if model.DisplayName == "" {
			return nil, common.ErrorBadRequest("Both of FQName and Display Name is empty")
		}
		model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateAlarm(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alarm",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAlarmResponse{
		Alarm: request.Alarm,
	}, nil
}

//RESTUpdateAlarm handles a REST Update request.
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
func (service *ContrailService) UpdateAlarm(
	ctx context.Context,
	request *models.UpdateAlarmRequest) (*models.UpdateAlarmResponse, error) {
	model := request.Alarm
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateAlarm(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alarm",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAlarmResponse{
		Alarm: model,
	}, nil
}

//RESTDeleteAlarm delete a resource using REST service.
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

//DeleteAlarm delete a resource.
func (service *ContrailService) DeleteAlarm(ctx context.Context, request *models.DeleteAlarmRequest) (*models.DeleteAlarmResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteAlarm(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAlarmResponse{
		ID: request.ID,
	}, nil
}

//RESTGetAlarm a REST Get request.
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

//GetAlarm a Get request.
func (service *ContrailService) GetAlarm(ctx context.Context, request *models.GetAlarmRequest) (response *models.GetAlarmResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListAlarmRequest{
		Spec: spec,
	}
	var result *models.ListAlarmResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListAlarm(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Alarms) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAlarmResponse{
		Alarm: result.Alarms[0],
	}
	return response, nil
}

//RESTListAlarm handles a List REST service Request.
func (service *ContrailService) RESTListAlarm(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
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

//ListAlarm handles a List service Request.
func (service *ContrailService) ListAlarm(
	ctx context.Context,
	request *models.ListAlarmRequest) (response *models.ListAlarmResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListAlarm(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
