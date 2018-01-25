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

//RESTCreateBGPRouter handle a Create REST service.
func (service *ContrailService) RESTCreateBGPRouter(c echo.Context) error {
    requestData := &models.BGPRouterCreateRequest{
        BGPRouter: models.MakeBGPRouter(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bgp_router",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateBGPRouter(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateBGPRouter handle a Create API
func (service *ContrailService) CreateBGPRouter(
    ctx context.Context, 
    request *models.BGPRouterCreateRequest) (*models.BGPRouterCreateResponse, error) {
    model := request.BGPRouter
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
            return db.CreateBGPRouter(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bgp_router",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.BGPRouterCreateResponse{
        BGPRouter: request.BGPRouter,
    }, nil
}

//RESTUpdateBGPRouter handles a REST Update request.
func (service *ContrailService) RESTUpdateBGPRouter(c echo.Context) error {
    id := c.Param("id")
    request := &models.BGPRouterUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "bgp_router",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateBGPRouter(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateBGPRouter handles a Update request.
func (service *ContrailService) UpdateBGPRouter(ctx context.Context, request *models.BGPRouterUpdateRequest) (*models.BGPRouterUpdateResponse, error) {
    id = request.ID
    model = request.BGPRouter
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
            return db.UpdateBGPRouter(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bgp_router",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.BGPRouter.UpdateResponse{
        BGPRouter: model,
    }, nil
}

//RESTDeleteBGPRouter delete a resource using REST service.
func (service *ContrailService) RESTDeleteBGPRouter(c echo.Context) error {
    id := c.Param("id")
    request := &models.BGPRouterDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteBGPRouter(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteBGPRouter delete a resource.
func (service *ContrailService) DeleteBGPRouter(ctx context.Context, request *models.BGPRouterDeleteRequest) (*models.BGPRouterDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteBGPRouter(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.BGPRouterDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowBGPRouter a REST Show request.
func (service *ContrailService) RESTShowBGPRouter(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.BGPRouter
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBGPRouter(tx, &common.ListSpec{
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
        "bgp_router": result,
    })
}

//RESTListBGPRouter handles a List REST service Request.
func (service *ContrailService) RESTListBGPRouter(c echo.Context) (error) {
    var result []*models.BGPRouter
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBGPRouter(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "bgp-routers": result,
    })
}