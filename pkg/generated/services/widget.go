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

//RESTWidgetUpdateRequest for update request for REST.
type RESTWidgetUpdateRequest struct {
	Data map[string]interface{} `json:"widget"`
}

//RESTCreateWidget handle a Create REST service.
func (service *ContrailService) RESTCreateWidget(c echo.Context) error {
	requestData := &models.CreateWidgetRequest{
		Widget: models.MakeWidget(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "widget",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateWidget(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateWidget handle a Create API
func (service *ContrailService) CreateWidget(
	ctx context.Context,
	request *models.CreateWidgetRequest) (*models.CreateWidgetResponse, error) {
	model := request.Widget
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
			return db.CreateWidget(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "widget",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateWidgetResponse{
		Widget: request.Widget,
	}, nil
}

//RESTUpdateWidget handles a REST Update request.
func (service *ContrailService) RESTUpdateWidget(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateWidgetRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "widget",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateWidget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateWidget handles a Update request.
func (service *ContrailService) UpdateWidget(
	ctx context.Context,
	request *models.UpdateWidgetRequest) (*models.UpdateWidgetResponse, error) {
	model := request.Widget
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateWidget(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "widget",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateWidgetResponse{
		Widget: model,
	}, nil
}

//RESTDeleteWidget delete a resource using REST service.
func (service *ContrailService) RESTDeleteWidget(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteWidgetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteWidget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteWidget delete a resource.
func (service *ContrailService) DeleteWidget(ctx context.Context, request *models.DeleteWidgetRequest) (*models.DeleteWidgetResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteWidget(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteWidgetResponse{
		ID: request.ID,
	}, nil
}

//RESTGetWidget a REST Get request.
func (service *ContrailService) RESTGetWidget(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetWidgetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetWidget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetWidget a Get request.
func (service *ContrailService) GetWidget(ctx context.Context, request *models.GetWidgetRequest) (response *models.GetWidgetResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListWidgetRequest{
		Spec: spec,
	}
	var result *models.ListWidgetResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListWidget(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Widgets) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetWidgetResponse{
		Widget: result.Widgets[0],
	}
	return response, nil
}

//RESTListWidget handles a List REST service Request.
func (service *ContrailService) RESTListWidget(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListWidgetRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListWidget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListWidget handles a List service Request.
func (service *ContrailService) ListWidget(
	ctx context.Context,
	request *models.ListWidgetRequest) (response *models.ListWidgetResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListWidget(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
