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

//RESTCreateGlobalQosConfig handle a Create REST service.
func (service *ContrailService) RESTCreateGlobalQosConfig(c echo.Context) error {
    requestData := &models.GlobalQosConfigCreateRequest{
        GlobalQosConfig: models.MakeGlobalQosConfig(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "global_qos_config",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateGlobalQosConfig(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateGlobalQosConfig handle a Create API
func (service *ContrailService) CreateGlobalQosConfig(
    ctx context.Context, 
    request *models.GlobalQosConfigCreateRequest) (*models.GlobalQosConfigCreateResponse, error) {
    model := request.GlobalQosConfig
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
            return db.CreateGlobalQosConfig(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "global_qos_config",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.GlobalQosConfigCreateResponse{
        GlobalQosConfig: request.GlobalQosConfig,
    }, nil
}

//RESTUpdateGlobalQosConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateGlobalQosConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.GlobalQosConfigUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "global_qos_config",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateGlobalQosConfig(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateGlobalQosConfig handles a Update request.
func (service *ContrailService) UpdateGlobalQosConfig(ctx context.Context, request *models.GlobalQosConfigUpdateRequest) (*models.GlobalQosConfigUpdateResponse, error) {
    id = request.ID
    model = request.GlobalQosConfig
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
            return db.UpdateGlobalQosConfig(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "global_qos_config",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.GlobalQosConfig.UpdateResponse{
        GlobalQosConfig: model,
    }, nil
}

//RESTDeleteGlobalQosConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteGlobalQosConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.GlobalQosConfigDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteGlobalQosConfig(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteGlobalQosConfig delete a resource.
func (service *ContrailService) DeleteGlobalQosConfig(ctx context.Context, request *models.GlobalQosConfigDeleteRequest) (*models.GlobalQosConfigDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteGlobalQosConfig(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.GlobalQosConfigDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowGlobalQosConfig a REST Show request.
func (service *ContrailService) RESTShowGlobalQosConfig(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.GlobalQosConfig
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListGlobalQosConfig(tx, &common.ListSpec{
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
        "global_qos_config": result,
    })
}

//RESTListGlobalQosConfig handles a List REST service Request.
func (service *ContrailService) RESTListGlobalQosConfig(c echo.Context) (error) {
    var result []*models.GlobalQosConfig
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListGlobalQosConfig(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "global-qos-configs": result,
    })
}