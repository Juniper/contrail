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

//RESTCreateLoadbalancerHealthmonitor handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancerHealthmonitor(c echo.Context) error {
    requestData := &models.LoadbalancerHealthmonitorCreateRequest{
        LoadbalancerHealthmonitor: models.MakeLoadbalancerHealthmonitor(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_healthmonitor",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateLoadbalancerHealthmonitor(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerHealthmonitor handle a Create API
func (service *ContrailService) CreateLoadbalancerHealthmonitor(
    ctx context.Context, 
    request *models.LoadbalancerHealthmonitorCreateRequest) (*models.LoadbalancerHealthmonitorCreateResponse, error) {
    model := request.LoadbalancerHealthmonitor
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
            return db.CreateLoadbalancerHealthmonitor(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_healthmonitor",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.LoadbalancerHealthmonitorCreateResponse{
        LoadbalancerHealthmonitor: request.LoadbalancerHealthmonitor,
    }, nil
}

//RESTUpdateLoadbalancerHealthmonitor handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancerHealthmonitor(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerHealthmonitorUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "loadbalancer_healthmonitor",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateLoadbalancerHealthmonitor(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerHealthmonitor handles a Update request.
func (service *ContrailService) UpdateLoadbalancerHealthmonitor(ctx context.Context, request *models.LoadbalancerHealthmonitorUpdateRequest) (*models.LoadbalancerHealthmonitorUpdateResponse, error) {
    id = request.ID
    model = request.LoadbalancerHealthmonitor
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
            return db.UpdateLoadbalancerHealthmonitor(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_healthmonitor",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.LoadbalancerHealthmonitor.UpdateResponse{
        LoadbalancerHealthmonitor: model,
    }, nil
}

//RESTDeleteLoadbalancerHealthmonitor delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancerHealthmonitor(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerHealthmonitorDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteLoadbalancerHealthmonitor(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteLoadbalancerHealthmonitor delete a resource.
func (service *ContrailService) DeleteLoadbalancerHealthmonitor(ctx context.Context, request *models.LoadbalancerHealthmonitorDeleteRequest) (*models.LoadbalancerHealthmonitorDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteLoadbalancerHealthmonitor(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.LoadbalancerHealthmonitorDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowLoadbalancerHealthmonitor a REST Show request.
func (service *ContrailService) RESTShowLoadbalancerHealthmonitor(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.LoadbalancerHealthmonitor
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancerHealthmonitor(tx, &common.ListSpec{
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
        "loadbalancer_healthmonitor": result,
    })
}

//RESTListLoadbalancerHealthmonitor handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancerHealthmonitor(c echo.Context) (error) {
    var result []*models.LoadbalancerHealthmonitor
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancerHealthmonitor(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "loadbalancer-healthmonitors": result,
    })
}