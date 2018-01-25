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

//RESTCreatePhysicalRouter handle a Create REST service.
func (service *ContrailService) RESTCreatePhysicalRouter(c echo.Context) error {
    requestData := &models.PhysicalRouterCreateRequest{
        PhysicalRouter: models.MakePhysicalRouter(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "physical_router",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreatePhysicalRouter(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreatePhysicalRouter handle a Create API
func (service *ContrailService) CreatePhysicalRouter(
    ctx context.Context, 
    request *models.PhysicalRouterCreateRequest) (*models.PhysicalRouterCreateResponse, error) {
    model := request.PhysicalRouter
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
            return db.CreatePhysicalRouter(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "physical_router",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.PhysicalRouterCreateResponse{
        PhysicalRouter: request.PhysicalRouter,
    }, nil
}

//RESTUpdatePhysicalRouter handles a REST Update request.
func (service *ContrailService) RESTUpdatePhysicalRouter(c echo.Context) error {
    id := c.Param("id")
    request := &models.PhysicalRouterUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "physical_router",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdatePhysicalRouter(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdatePhysicalRouter handles a Update request.
func (service *ContrailService) UpdatePhysicalRouter(ctx context.Context, request *models.PhysicalRouterUpdateRequest) (*models.PhysicalRouterUpdateResponse, error) {
    id = request.ID
    model = request.PhysicalRouter
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
            return db.UpdatePhysicalRouter(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "physical_router",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.PhysicalRouter.UpdateResponse{
        PhysicalRouter: model,
    }, nil
}

//RESTDeletePhysicalRouter delete a resource using REST service.
func (service *ContrailService) RESTDeletePhysicalRouter(c echo.Context) error {
    id := c.Param("id")
    request := &models.PhysicalRouterDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeletePhysicalRouter(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeletePhysicalRouter delete a resource.
func (service *ContrailService) DeletePhysicalRouter(ctx context.Context, request *models.PhysicalRouterDeleteRequest) (*models.PhysicalRouterDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeletePhysicalRouter(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.PhysicalRouterDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowPhysicalRouter a REST Show request.
func (service *ContrailService) RESTShowPhysicalRouter(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.PhysicalRouter
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPhysicalRouter(tx, &common.ListSpec{
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
        "physical_router": result,
    })
}

//RESTListPhysicalRouter handles a List REST service Request.
func (service *ContrailService) RESTListPhysicalRouter(c echo.Context) (error) {
    var result []*models.PhysicalRouter
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPhysicalRouter(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "physical-routers": result,
    })
}