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

//RESTRouteTargetUpdateRequest for update request for REST.
type RESTRouteTargetUpdateRequest struct {
    Data map[string]interface{} `json:"route-target"`
}

//RESTCreateRouteTarget handle a Create REST service.
func (service *ContrailService) RESTCreateRouteTarget(c echo.Context) error {
    requestData := &models.RouteTargetCreateRequest{
        RouteTarget: models.MakeRouteTarget(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "route_target",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateRouteTarget(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateRouteTarget handle a Create API
func (service *ContrailService) CreateRouteTarget(
    ctx context.Context, 
    request *models.RouteTargetCreateRequest) (*models.RouteTargetCreateResponse, error) {
    model := request.RouteTarget
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
            return db.CreateRouteTarget(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "route_target",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.RouteTargetCreateResponse{
        RouteTarget: request.RouteTarget,
    }, nil
}

//RESTUpdateRouteTarget handles a REST Update request.
func (service *ContrailService) RESTUpdateRouteTarget(c echo.Context) error {
    id := c.Param("id")
    request := &models.RouteTargetUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "route_target",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateRouteTarget(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateRouteTarget handles a Update request.
func (service *ContrailService) UpdateRouteTarget(ctx context.Context, request *models.RouteTargetUpdateRequest) (*models.RouteTargetUpdateResponse, error) {
    id = request.ID
    model = request.RouteTarget
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
            return db.UpdateRouteTarget(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "route_target",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.RouteTarget.UpdateResponse{
        RouteTarget: model,
    }, nil
}

//RESTDeleteRouteTarget delete a resource using REST service.
func (service *ContrailService) RESTDeleteRouteTarget(c echo.Context) error {
    id := c.Param("id")
    request := &models.RouteTargetDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteRouteTarget(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteRouteTarget delete a resource.
func (service *ContrailService) DeleteRouteTarget(ctx context.Context, request *models.RouteTargetDeleteRequest) (*models.RouteTargetDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteRouteTarget(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.RouteTargetDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowRouteTarget a REST Show request.
func (service *ContrailService) RESTShowRouteTarget(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.RouteTarget
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListRouteTarget(tx, &common.ListSpec{
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
        "route_target": result,
    })
}

//RESTListRouteTarget handles a List REST service Request.
func (service *ContrailService) RESTListRouteTarget(c echo.Context) (error) {
    var result []*models.RouteTarget
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListRouteTarget(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "route-targets": result,
    })
}