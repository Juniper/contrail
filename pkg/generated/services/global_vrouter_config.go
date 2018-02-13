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

//RESTGlobalVrouterConfigUpdateRequest for update request for REST.
type RESTGlobalVrouterConfigUpdateRequest struct {
    Data map[string]interface{} `json:"global-vrouter-config"`
}

//RESTCreateGlobalVrouterConfig handle a Create REST service.
func (service *ContrailService) RESTCreateGlobalVrouterConfig(c echo.Context) error {
    requestData := &models.GlobalVrouterConfigCreateRequest{
        GlobalVrouterConfig: models.MakeGlobalVrouterConfig(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "global_vrouter_config",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateGlobalVrouterConfig(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateGlobalVrouterConfig handle a Create API
func (service *ContrailService) CreateGlobalVrouterConfig(
    ctx context.Context, 
    request *models.GlobalVrouterConfigCreateRequest) (*models.GlobalVrouterConfigCreateResponse, error) {
    model := request.GlobalVrouterConfig
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
            return db.CreateGlobalVrouterConfig(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "global_vrouter_config",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.GlobalVrouterConfigCreateResponse{
        GlobalVrouterConfig: request.GlobalVrouterConfig,
    }, nil
}

//RESTUpdateGlobalVrouterConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateGlobalVrouterConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.GlobalVrouterConfigUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "global_vrouter_config",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateGlobalVrouterConfig(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateGlobalVrouterConfig handles a Update request.
func (service *ContrailService) UpdateGlobalVrouterConfig(ctx context.Context, request *models.GlobalVrouterConfigUpdateRequest) (*models.GlobalVrouterConfigUpdateResponse, error) {
    id = request.ID
    model = request.GlobalVrouterConfig
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
            return db.UpdateGlobalVrouterConfig(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "global_vrouter_config",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.GlobalVrouterConfig.UpdateResponse{
        GlobalVrouterConfig: model,
    }, nil
}

//RESTDeleteGlobalVrouterConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteGlobalVrouterConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.GlobalVrouterConfigDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteGlobalVrouterConfig(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteGlobalVrouterConfig delete a resource.
func (service *ContrailService) DeleteGlobalVrouterConfig(ctx context.Context, request *models.GlobalVrouterConfigDeleteRequest) (*models.GlobalVrouterConfigDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteGlobalVrouterConfig(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.GlobalVrouterConfigDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowGlobalVrouterConfig a REST Show request.
func (service *ContrailService) RESTShowGlobalVrouterConfig(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.GlobalVrouterConfig
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListGlobalVrouterConfig(tx, &common.ListSpec{
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
        "global_vrouter_config": result,
    })
}

//RESTListGlobalVrouterConfig handles a List REST service Request.
func (service *ContrailService) RESTListGlobalVrouterConfig(c echo.Context) (error) {
    var result []*models.GlobalVrouterConfig
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListGlobalVrouterConfig(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "global-vrouter-configs": result,
    })
}