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

//RESTQosQueueUpdateRequest for update request for REST.
type RESTQosQueueUpdateRequest struct {
    Data map[string]interface{} `json:"qos-queue"`
}

//RESTCreateQosQueue handle a Create REST service.
func (service *ContrailService) RESTCreateQosQueue(c echo.Context) error {
    requestData := &models.QosQueueCreateRequest{
        QosQueue: models.MakeQosQueue(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "qos_queue",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateQosQueue(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateQosQueue handle a Create API
func (service *ContrailService) CreateQosQueue(
    ctx context.Context, 
    request *models.QosQueueCreateRequest) (*models.QosQueueCreateResponse, error) {
    model := request.QosQueue
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
            return db.CreateQosQueue(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "qos_queue",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.QosQueueCreateResponse{
        QosQueue: request.QosQueue,
    }, nil
}

//RESTUpdateQosQueue handles a REST Update request.
func (service *ContrailService) RESTUpdateQosQueue(c echo.Context) error {
    id := c.Param("id")
    request := &models.QosQueueUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "qos_queue",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateQosQueue(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateQosQueue handles a Update request.
func (service *ContrailService) UpdateQosQueue(ctx context.Context, request *models.QosQueueUpdateRequest) (*models.QosQueueUpdateResponse, error) {
    id = request.ID
    model = request.QosQueue
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
            return db.UpdateQosQueue(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "qos_queue",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.QosQueue.UpdateResponse{
        QosQueue: model,
    }, nil
}

//RESTDeleteQosQueue delete a resource using REST service.
func (service *ContrailService) RESTDeleteQosQueue(c echo.Context) error {
    id := c.Param("id")
    request := &models.QosQueueDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteQosQueue(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteQosQueue delete a resource.
func (service *ContrailService) DeleteQosQueue(ctx context.Context, request *models.QosQueueDeleteRequest) (*models.QosQueueDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteQosQueue(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.QosQueueDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowQosQueue a REST Show request.
func (service *ContrailService) RESTShowQosQueue(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.QosQueue
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListQosQueue(tx, &common.ListSpec{
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
        "qos_queue": result,
    })
}

//RESTListQosQueue handles a List REST service Request.
func (service *ContrailService) RESTListQosQueue(c echo.Context) (error) {
    var result []*models.QosQueue
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListQosQueue(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "qos-queues": result,
    })
}