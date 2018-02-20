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

//RESTLogicalRouterUpdateRequest for update request for REST.
type RESTLogicalRouterUpdateRequest struct {
	Data map[string]interface{} `json:"logical-router"`
}

//RESTCreateLogicalRouter handle a Create REST service.
func (service *ContrailService) RESTCreateLogicalRouter(c echo.Context) error {
	requestData := &models.CreateLogicalRouterRequest{
		LogicalRouter: models.MakeLogicalRouter(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_router",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLogicalRouter(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLogicalRouter handle a Create API
func (service *ContrailService) CreateLogicalRouter(
	ctx context.Context,
	request *models.CreateLogicalRouterRequest) (*models.CreateLogicalRouterResponse, error) {
	model := request.LogicalRouter
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	if model.FQName == nil {
		return nil, common.ErrorBadRequest("Missing fq_name")
	}

	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateLogicalRouter(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_router",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLogicalRouterResponse{
		LogicalRouter: request.LogicalRouter,
	}, nil
}

//RESTUpdateLogicalRouter handles a REST Update request.
func (service *ContrailService) RESTUpdateLogicalRouter(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLogicalRouterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_router",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLogicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLogicalRouter handles a Update request.
func (service *ContrailService) UpdateLogicalRouter(
	ctx context.Context,
	request *models.UpdateLogicalRouterRequest) (*models.UpdateLogicalRouterResponse, error) {
	model := request.LogicalRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateLogicalRouter(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_router",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLogicalRouterResponse{
		LogicalRouter: model,
	}, nil
}

//RESTDeleteLogicalRouter delete a resource using REST service.
func (service *ContrailService) RESTDeleteLogicalRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLogicalRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLogicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteLogicalRouter delete a resource.
func (service *ContrailService) DeleteLogicalRouter(ctx context.Context, request *models.DeleteLogicalRouterRequest) (*models.DeleteLogicalRouterResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteLogicalRouter(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLogicalRouterResponse{
		ID: request.ID,
	}, nil
}

//RESTGetLogicalRouter a REST Get request.
func (service *ContrailService) RESTGetLogicalRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLogicalRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLogicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetLogicalRouter a Get request.
func (service *ContrailService) GetLogicalRouter(ctx context.Context, request *models.GetLogicalRouterRequest) (response *models.GetLogicalRouterResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListLogicalRouterRequest{
		Spec: spec,
	}
	var result *models.ListLogicalRouterResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListLogicalRouter(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LogicalRouters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLogicalRouterResponse{
		LogicalRouter: result.LogicalRouters[0],
	}
	return response, nil
}

//RESTListLogicalRouter handles a List REST service Request.
func (service *ContrailService) RESTListLogicalRouter(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLogicalRouterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLogicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListLogicalRouter handles a List service Request.
func (service *ContrailService) ListLogicalRouter(
	ctx context.Context,
	request *models.ListLogicalRouterRequest) (response *models.ListLogicalRouterResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListLogicalRouter(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
