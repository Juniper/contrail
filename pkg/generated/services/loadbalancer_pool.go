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

//RESTCreateLoadbalancerPool handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancerPool(c echo.Context) error {
    requestData := &models.LoadbalancerPoolCreateRequest{
        LoadbalancerPool: models.MakeLoadbalancerPool(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_pool",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateLoadbalancerPool(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerPool handle a Create API
func (service *ContrailService) CreateLoadbalancerPool(
    ctx context.Context, 
    request *models.LoadbalancerPoolCreateRequest) (*models.LoadbalancerPoolCreateResponse, error) {
    model := request.LoadbalancerPool
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
            return db.CreateLoadbalancerPool(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_pool",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.LoadbalancerPoolCreateResponse{
        LoadbalancerPool: request.LoadbalancerPool,
    }, nil
}

//RESTUpdateLoadbalancerPool handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancerPool(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerPoolUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "loadbalancer_pool",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateLoadbalancerPool(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerPool handles a Update request.
func (service *ContrailService) UpdateLoadbalancerPool(ctx context.Context, request *models.LoadbalancerPoolUpdateRequest) (*models.LoadbalancerPoolUpdateResponse, error) {
    id = request.ID
    model = request.LoadbalancerPool
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
            return db.UpdateLoadbalancerPool(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_pool",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.LoadbalancerPool.UpdateResponse{
        LoadbalancerPool: model,
    }, nil
}

//RESTDeleteLoadbalancerPool delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancerPool(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerPoolDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteLoadbalancerPool(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteLoadbalancerPool delete a resource.
func (service *ContrailService) DeleteLoadbalancerPool(ctx context.Context, request *models.LoadbalancerPoolDeleteRequest) (*models.LoadbalancerPoolDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteLoadbalancerPool(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.LoadbalancerPoolDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowLoadbalancerPool a REST Show request.
func (service *ContrailService) RESTShowLoadbalancerPool(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.LoadbalancerPool
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancerPool(tx, &common.ListSpec{
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
        "loadbalancer_pool": result,
    })
}

//RESTListLoadbalancerPool handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancerPool(c echo.Context) (error) {
    var result []*models.LoadbalancerPool
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancerPool(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "loadbalancer-pools": result,
    })
}