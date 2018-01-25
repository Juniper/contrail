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

//RESTCreateRouteTable handle a Create REST service.
func (service *ContrailService) RESTCreateRouteTable(c echo.Context) error {
    requestData := &models.RouteTableCreateRequest{
        RouteTable: models.MakeRouteTable(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
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
    request *models.RouteTableCreateRequest) (*models.RouteTableCreateResponse, error) {
    model := request.RouteTable
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
            return db.CreateRouteTable(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "route_table",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.RouteTableCreateResponse{
        RouteTable: request.RouteTable,
    }, nil
}

//RESTUpdateRouteTable handles a REST Update request.
func (service *ContrailService) RESTUpdateRouteTable(c echo.Context) error {
    id := c.Param("id")
    request := &models.RouteTableUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "route_table",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateRouteTable(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateRouteTable handles a Update request.
func (service *ContrailService) UpdateRouteTable(ctx context.Context, request *models.RouteTableUpdateRequest) (*models.RouteTableUpdateResponse, error) {
    id = request.ID
    model = request.RouteTable
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
            return db.UpdateRouteTable(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "route_table",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.RouteTable.UpdateResponse{
        RouteTable: model,
    }, nil
}

//RESTDeleteRouteTable delete a resource using REST service.
func (service *ContrailService) RESTDeleteRouteTable(c echo.Context) error {
    id := c.Param("id")
    request := &models.RouteTableDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteRouteTable(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteRouteTable delete a resource.
func (service *ContrailService) DeleteRouteTable(ctx context.Context, request *models.RouteTableDeleteRequest) (*models.RouteTableDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteRouteTable(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.RouteTableDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowRouteTable a REST Show request.
func (service *ContrailService) RESTShowRouteTable(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.RouteTable
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListRouteTable(tx, &common.ListSpec{
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
        "route_table": result,
    })
}

//RESTListRouteTable handles a List REST service Request.
func (service *ContrailService) RESTListRouteTable(c echo.Context) (error) {
    var result []*models.RouteTable
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListRouteTable(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "route-tables": result,
    })
}