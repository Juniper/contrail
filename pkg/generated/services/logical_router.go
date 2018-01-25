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

//RESTCreateLogicalRouter handle a Create REST service.
func (service *ContrailService) RESTCreateLogicalRouter(c echo.Context) error {
    requestData := &models.LogicalRouterCreateRequest{
        LogicalRouter: models.MakeLogicalRouter(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
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
    request *models.LogicalRouterCreateRequest) (*models.LogicalRouterCreateResponse, error) {
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
    model.Perms2.Owner = auth.ProjectID()
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.CreateLogicalRouter(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "logical_router",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.LogicalRouterCreateResponse{
        LogicalRouter: request.LogicalRouter,
    }, nil
}

//RESTUpdateLogicalRouter handles a REST Update request.
func (service *ContrailService) RESTUpdateLogicalRouter(c echo.Context) error {
    id := c.Param("id")
    request := &models.LogicalRouterUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "logical_router",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateLogicalRouter(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateLogicalRouter handles a Update request.
func (service *ContrailService) UpdateLogicalRouter(ctx context.Context, request *models.LogicalRouterUpdateRequest) (*models.LogicalRouterUpdateResponse, error) {
    id = request.ID
    model = request.LogicalRouter
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
            return db.UpdateLogicalRouter(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "logical_router",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.LogicalRouter.UpdateResponse{
        LogicalRouter: model,
    }, nil
}

//RESTDeleteLogicalRouter delete a resource using REST service.
func (service *ContrailService) RESTDeleteLogicalRouter(c echo.Context) error {
    id := c.Param("id")
    request := &models.LogicalRouterDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteLogicalRouter(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteLogicalRouter delete a resource.
func (service *ContrailService) DeleteLogicalRouter(ctx context.Context, request *models.LogicalRouterDeleteRequest) (*models.LogicalRouterDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteLogicalRouter(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.LogicalRouterDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowLogicalRouter a REST Show request.
func (service *ContrailService) RESTShowLogicalRouter(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.LogicalRouter
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLogicalRouter(tx, &common.ListSpec{
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
        "logical_router": result,
    })
}

//RESTListLogicalRouter handles a List REST service Request.
func (service *ContrailService) RESTListLogicalRouter(c echo.Context) (error) {
    var result []*models.LogicalRouter
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLogicalRouter(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "logical-routers": result,
    })
}