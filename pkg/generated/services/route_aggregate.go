package services 

import (
    "context"
    "net/http"
    "database/sql"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/Juniper/contrail/pkg/generated/db"
    "github.com/satori/go.uuid"
    "github.com/labstack/echo"
    "github.com/Juniper/contrail/pkg/common"

	log "github.com/sirupsen/logrus"
)

//RESTCreateRouteAggregate handle a Create REST service.
func (service *ContrailService) RESTCreateRouteAggregate(c echo.Context) error {
    requestData := &models.RouteAggregateCreateRequest{
        RouteAggregate: models.MakeRouteAggregate(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
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
    request *models.RouteAggregateCreateRequest) (*models.RouteAggregateCreateResponse, error) {
    model := request.RouteAggregate
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
    model.Perms2.Owner = auth.ProjectID()
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.CreateRouteAggregate(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "route_aggregate",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.RouteAggregateCreateResponse{
        RouteAggregate: request.RouteAggregate,
    }, nil
}

//RESTUpdateRouteAggregate handles a REST Update request.
func (service *ContrailService) RESTUpdateRouteAggregate(c echo.Context) error {
    id := c.Param("id")
    request := &models.RouteAggregateUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "route_aggregate",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateRouteAggregate(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateRouteAggregate handles a Update request.
func (service *ContrailService) UpdateRouteAggregate(ctx context.Context, request *models.RouteAggregateUpdateRequest) (*models.RouteAggregateUpdateResponse, error) {
    id = request.ID
    model = request.RouteAggregate
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    auth := common.GetAuthCTX(ctx)
    ok := common.SetValueByPath(model, "Perms2.Owner", ".", auth.ProjectID())
    if !ok {
        return nil, common.ErrorBadRequest("Invalid JSON format")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateRouteAggregate(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "route_aggregate",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.RouteAggregate.UpdateResponse{
        RouteAggregate: model,
    }, nil
}

//RESTDeleteRouteAggregate delete a resource using REST service.
func (service *ContrailService) RESTDeleteRouteAggregate(c echo.Context) error {
    id := c.Param("id")
    request := &models.RouteAggregateDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteRouteAggregate(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteRouteAggregate delete a resource.
func (service *ContrailService) DeleteRouteAggregate(ctx context.Context, request *models.RouteAggregateDeleteRequest) (*models.RouteAggregateDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteRouteAggregate(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.RouteAggregateDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowRouteAggregate a REST Show request.
func (service *ContrailService) RESTShowRouteAggregate(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.RouteAggregate
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListRouteAggregate(tx, &common.ListSpec{
                Limit: 1,
                Auth: auth,
                Filter: common.Filter{
                    "uuid": []string{id},
                },
            })
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "route_aggregate": result,
    })
}

//RESTListRouteAggregate handles a List REST service Request.
func (service *ContrailService) RESTListRouteAggregate(c echo.Context) (error) {
    var result []*models.RouteAggregate
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListRouteAggregate(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "route-aggregates": result,
    })
}