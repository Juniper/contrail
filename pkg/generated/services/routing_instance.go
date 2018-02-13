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

//RESTRoutingInstanceUpdateRequest for update request for REST.
type RESTRoutingInstanceUpdateRequest struct {
    Data map[string]interface{} `json:"routing-instance"`
}

//RESTCreateRoutingInstance handle a Create REST service.
func (service *ContrailService) RESTCreateRoutingInstance(c echo.Context) error {
    requestData := &models.RoutingInstanceCreateRequest{
        RoutingInstance: models.MakeRoutingInstance(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "routing_instance",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateRoutingInstance(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateRoutingInstance handle a Create API
func (service *ContrailService) CreateRoutingInstance(
    ctx context.Context, 
    request *models.RoutingInstanceCreateRequest) (*models.RoutingInstanceCreateResponse, error) {
    model := request.RoutingInstance
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
            return db.CreateRoutingInstance(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "routing_instance",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.RoutingInstanceCreateResponse{
        RoutingInstance: request.RoutingInstance,
    }, nil
}

//RESTUpdateRoutingInstance handles a REST Update request.
func (service *ContrailService) RESTUpdateRoutingInstance(c echo.Context) error {
    id := c.Param("id")
    request := &models.RoutingInstanceUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "routing_instance",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateRoutingInstance(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateRoutingInstance handles a Update request.
func (service *ContrailService) UpdateRoutingInstance(ctx context.Context, request *models.RoutingInstanceUpdateRequest) (*models.RoutingInstanceUpdateResponse, error) {
    id = request.ID
    model = request.RoutingInstance
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
            return db.UpdateRoutingInstance(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "routing_instance",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.RoutingInstance.UpdateResponse{
        RoutingInstance: model,
    }, nil
}

//RESTDeleteRoutingInstance delete a resource using REST service.
func (service *ContrailService) RESTDeleteRoutingInstance(c echo.Context) error {
    id := c.Param("id")
    request := &models.RoutingInstanceDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteRoutingInstance(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteRoutingInstance delete a resource.
func (service *ContrailService) DeleteRoutingInstance(ctx context.Context, request *models.RoutingInstanceDeleteRequest) (*models.RoutingInstanceDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteRoutingInstance(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.RoutingInstanceDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowRoutingInstance a REST Show request.
func (service *ContrailService) RESTShowRoutingInstance(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.RoutingInstance
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListRoutingInstance(tx, &common.ListSpec{
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
        "routing_instance": result,
    })
}

//RESTListRoutingInstance handles a List REST service Request.
func (service *ContrailService) RESTListRoutingInstance(c echo.Context) (error) {
    var result []*models.RoutingInstance
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListRoutingInstance(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "routing-instances": result,
    })
}