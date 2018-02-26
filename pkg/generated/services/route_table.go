package services

import (
	"context"
	"database/sql"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateRouteTable handle a Create REST service.
func (service *ContrailService) RESTCreateRouteTable(c echo.Context) error {
	requestData := &models.CreateRouteTableRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_table",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateRouteTable(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateRouteTable handle a Create API
func (service *ContrailService) CreateRouteTable(
	ctx context.Context,
	request *models.CreateRouteTableRequest) (*models.CreateRouteTableResponse, error) {
	model := request.RouteTable
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
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateRouteTable(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_table",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateRouteTableResponse{
		RouteTable: request.RouteTable,
	}, nil
}

//RESTUpdateRouteTable handles a REST Update request.
func (service *ContrailService) RESTUpdateRouteTable(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateRouteTableRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_table",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateRouteTable handles a Update request.
func (service *ContrailService) UpdateRouteTable(
	ctx context.Context,
	request *models.UpdateRouteTableRequest) (*models.UpdateRouteTableResponse, error) {
	model := request.RouteTable
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateRouteTable(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_table",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateRouteTableResponse{
		RouteTable: model,
	}, nil
}

//RESTDeleteRouteTable delete a resource using REST service.
func (service *ContrailService) RESTDeleteRouteTable(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteRouteTableRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteRouteTable delete a resource.
func (service *ContrailService) DeleteRouteTable(ctx context.Context, request *models.DeleteRouteTableRequest) (*models.DeleteRouteTableResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteRouteTable(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteRouteTableResponse{
		ID: request.ID,
	}, nil
}

//RESTGetRouteTable a REST Get request.
func (service *ContrailService) RESTGetRouteTable(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetRouteTableRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetRouteTable a Get request.
func (service *ContrailService) GetRouteTable(ctx context.Context, request *models.GetRouteTableRequest) (response *models.GetRouteTableResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListRouteTableRequest{
		Spec: spec,
	}
	var result *models.ListRouteTableResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListRouteTable(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.RouteTables) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetRouteTableResponse{
		RouteTable: result.RouteTables[0],
	}
	return response, nil
}

//RESTListRouteTable handles a List REST service Request.
func (service *ContrailService) RESTListRouteTable(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListRouteTableRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListRouteTable handles a List service Request.
func (service *ContrailService) ListRouteTable(
	ctx context.Context,
	request *models.ListRouteTableRequest) (response *models.ListRouteTableResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListRouteTable(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
