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

//RESTLoadbalancerUpdateRequest for update request for REST.
type RESTLoadbalancerUpdateRequest struct {
    Data map[string]interface{} `json:"loadbalancer"`
}

//RESTCreateLoadbalancer handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancer(c echo.Context) error {
    requestData := &models.LoadbalancerCreateRequest{
        Loadbalancer: models.MakeLoadbalancer(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateLoadbalancer(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancer handle a Create API
func (service *ContrailService) CreateLoadbalancer(
    ctx context.Context, 
    request *models.LoadbalancerCreateRequest) (*models.LoadbalancerCreateResponse, error) {
    model := request.Loadbalancer
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
            return db.CreateLoadbalancer(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.LoadbalancerCreateResponse{
        Loadbalancer: request.Loadbalancer,
    }, nil
}

//RESTUpdateLoadbalancer handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancer(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "loadbalancer",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateLoadbalancer(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancer handles a Update request.
func (service *ContrailService) UpdateLoadbalancer(ctx context.Context, request *models.LoadbalancerUpdateRequest) (*models.LoadbalancerUpdateResponse, error) {
    id = request.ID
    model = request.Loadbalancer
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
            return db.UpdateLoadbalancer(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.Loadbalancer.UpdateResponse{
        Loadbalancer: model,
    }, nil
}

//RESTDeleteLoadbalancer delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancer(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteLoadbalancer(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteLoadbalancer delete a resource.
func (service *ContrailService) DeleteLoadbalancer(ctx context.Context, request *models.LoadbalancerDeleteRequest) (*models.LoadbalancerDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteLoadbalancer(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.LoadbalancerDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowLoadbalancer a REST Show request.
func (service *ContrailService) RESTShowLoadbalancer(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.Loadbalancer
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancer(tx, &common.ListSpec{
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
        "loadbalancer": result,
    })
}

//RESTListLoadbalancer handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancer(c echo.Context) (error) {
    var result []*models.Loadbalancer
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancer(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "loadbalancers": result,
    })
}