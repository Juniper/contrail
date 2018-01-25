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

//RESTCreateLoadbalancerListener handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancerListener(c echo.Context) error {
    requestData := &models.LoadbalancerListenerCreateRequest{
        LoadbalancerListener: models.MakeLoadbalancerListener(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_listener",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateLoadbalancerListener(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerListener handle a Create API
func (service *ContrailService) CreateLoadbalancerListener(
    ctx context.Context, 
    request *models.LoadbalancerListenerCreateRequest) (*models.LoadbalancerListenerCreateResponse, error) {
    model := request.LoadbalancerListener
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
            return db.CreateLoadbalancerListener(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_listener",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.LoadbalancerListenerCreateResponse{
        LoadbalancerListener: request.LoadbalancerListener,
    }, nil
}

//RESTUpdateLoadbalancerListener handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancerListener(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerListenerUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "loadbalancer_listener",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateLoadbalancerListener(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerListener handles a Update request.
func (service *ContrailService) UpdateLoadbalancerListener(ctx context.Context, request *models.LoadbalancerListenerUpdateRequest) (*models.LoadbalancerListenerUpdateResponse, error) {
    id = request.ID
    model = request.LoadbalancerListener
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
            return db.UpdateLoadbalancerListener(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_listener",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.LoadbalancerListener.UpdateResponse{
        LoadbalancerListener: model,
    }, nil
}

//RESTDeleteLoadbalancerListener delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancerListener(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerListenerDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteLoadbalancerListener(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteLoadbalancerListener delete a resource.
func (service *ContrailService) DeleteLoadbalancerListener(ctx context.Context, request *models.LoadbalancerListenerDeleteRequest) (*models.LoadbalancerListenerDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteLoadbalancerListener(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.LoadbalancerListenerDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowLoadbalancerListener a REST Show request.
func (service *ContrailService) RESTShowLoadbalancerListener(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.LoadbalancerListener
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancerListener(tx, &common.ListSpec{
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
        "loadbalancer_listener": result,
    })
}

//RESTListLoadbalancerListener handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancerListener(c echo.Context) (error) {
    var result []*models.LoadbalancerListener
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancerListener(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "loadbalancer-listeners": result,
    })
}