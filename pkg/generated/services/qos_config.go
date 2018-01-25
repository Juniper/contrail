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

//RESTCreateQosConfig handle a Create REST service.
func (service *ContrailService) RESTCreateQosConfig(c echo.Context) error {
    requestData := &models.QosConfigCreateRequest{
        QosConfig: models.MakeQosConfig(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "qos_config",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateQosConfig(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateQosConfig handle a Create API
func (service *ContrailService) CreateQosConfig(
    ctx context.Context, 
    request *models.QosConfigCreateRequest) (*models.QosConfigCreateResponse, error) {
    model := request.QosConfig
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
            return db.CreateQosConfig(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "qos_config",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.QosConfigCreateResponse{
        QosConfig: request.QosConfig,
    }, nil
}

//RESTUpdateQosConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateQosConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.QosConfigUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "qos_config",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateQosConfig(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateQosConfig handles a Update request.
func (service *ContrailService) UpdateQosConfig(ctx context.Context, request *models.QosConfigUpdateRequest) (*models.QosConfigUpdateResponse, error) {
    id = request.ID
    model = request.QosConfig
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
            return db.UpdateQosConfig(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "qos_config",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.QosConfig.UpdateResponse{
        QosConfig: model,
    }, nil
}

//RESTDeleteQosConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteQosConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.QosConfigDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteQosConfig(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteQosConfig delete a resource.
func (service *ContrailService) DeleteQosConfig(ctx context.Context, request *models.QosConfigDeleteRequest) (*models.QosConfigDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteQosConfig(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.QosConfigDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowQosConfig a REST Show request.
func (service *ContrailService) RESTShowQosConfig(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.QosConfig
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListQosConfig(tx, &common.ListSpec{
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
        "qos_config": result,
    })
}

//RESTListQosConfig handles a List REST service Request.
func (service *ContrailService) RESTListQosConfig(c echo.Context) (error) {
    var result []*models.QosConfig
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListQosConfig(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "qos-configs": result,
    })
}