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

//RESTRouteAggregateUpdateRequest for update request for REST.
type RESTRouteAggregateUpdateRequest struct {
	Data map[string]interface{} `json:"route-aggregate"`
}

//RESTCreateRouteAggregate handle a Create REST service.
func (service *ContrailService) RESTCreateRouteAggregate(c echo.Context) error {
	requestData := &models.CreateRouteAggregateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_aggregate",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateRouteAggregate(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateRouteAggregate handle a Create API
func (service *ContrailService) CreateRouteAggregate(
	ctx context.Context,
	request *models.CreateRouteAggregateRequest) (*models.CreateRouteAggregateResponse, error) {
	model := request.RouteAggregate
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
			return db.CreateRouteAggregate(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_aggregate",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateRouteAggregateResponse{
		RouteAggregate: request.RouteAggregate,
	}, nil
}

//RESTUpdateRouteAggregate handles a REST Update request.
func (service *ContrailService) RESTUpdateRouteAggregate(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateRouteAggregateRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_aggregate",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateRouteAggregate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateRouteAggregate handles a Update request.
func (service *ContrailService) UpdateRouteAggregate(
	ctx context.Context,
	request *models.UpdateRouteAggregateRequest) (*models.UpdateRouteAggregateResponse, error) {
	model := request.RouteAggregate
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateRouteAggregate(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_aggregate",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateRouteAggregateResponse{
		RouteAggregate: model,
	}, nil
}

//RESTDeleteRouteAggregate delete a resource using REST service.
func (service *ContrailService) RESTDeleteRouteAggregate(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteRouteAggregateRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteRouteAggregate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteRouteAggregate delete a resource.
func (service *ContrailService) DeleteRouteAggregate(ctx context.Context, request *models.DeleteRouteAggregateRequest) (*models.DeleteRouteAggregateResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteRouteAggregate(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteRouteAggregateResponse{
		ID: request.ID,
	}, nil
}

//RESTGetRouteAggregate a REST Get request.
func (service *ContrailService) RESTGetRouteAggregate(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetRouteAggregateRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetRouteAggregate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetRouteAggregate a Get request.
func (service *ContrailService) GetRouteAggregate(ctx context.Context, request *models.GetRouteAggregateRequest) (response *models.GetRouteAggregateResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListRouteAggregateRequest{
		Spec: spec,
	}
	var result *models.ListRouteAggregateResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListRouteAggregate(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.RouteAggregates) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetRouteAggregateResponse{
		RouteAggregate: result.RouteAggregates[0],
	}
	return response, nil
}

//RESTListRouteAggregate handles a List REST service Request.
func (service *ContrailService) RESTListRouteAggregate(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListRouteAggregateRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListRouteAggregate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListRouteAggregate handles a List service Request.
func (service *ContrailService) ListRouteAggregate(
	ctx context.Context,
	request *models.ListRouteAggregateRequest) (response *models.ListRouteAggregateResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListRouteAggregate(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
