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

//RESTGlobalSystemConfigUpdateRequest for update request for REST.
type RESTGlobalSystemConfigUpdateRequest struct {
    Data map[string]interface{} `json:"global-system-config"`
}

//RESTCreateGlobalSystemConfig handle a Create REST service.
func (service *ContrailService) RESTCreateGlobalSystemConfig(c echo.Context) error {
    requestData := &models.GlobalSystemConfigCreateRequest{
        GlobalSystemConfig: models.MakeGlobalSystemConfig(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "global_system_config",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateGlobalSystemConfig(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateGlobalSystemConfig handle a Create API
func (service *ContrailService) CreateGlobalSystemConfig(
    ctx context.Context, 
    request *models.GlobalSystemConfigCreateRequest) (*models.GlobalSystemConfigCreateResponse, error) {
    model := request.GlobalSystemConfig
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
            return db.CreateGlobalSystemConfig(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "global_system_config",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.GlobalSystemConfigCreateResponse{
        GlobalSystemConfig: request.GlobalSystemConfig,
    }, nil
}

//RESTUpdateGlobalSystemConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateGlobalSystemConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.GlobalSystemConfigUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "global_system_config",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateGlobalSystemConfig(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateGlobalSystemConfig handles a Update request.
func (service *ContrailService) UpdateGlobalSystemConfig(ctx context.Context, request *models.GlobalSystemConfigUpdateRequest) (*models.GlobalSystemConfigUpdateResponse, error) {
    id = request.ID
    model = request.GlobalSystemConfig
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
            return db.UpdateGlobalSystemConfig(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "global_system_config",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.GlobalSystemConfig.UpdateResponse{
        GlobalSystemConfig: model,
    }, nil
}

//RESTDeleteGlobalSystemConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteGlobalSystemConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.GlobalSystemConfigDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteGlobalSystemConfig(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteGlobalSystemConfig delete a resource.
func (service *ContrailService) DeleteGlobalSystemConfig(ctx context.Context, request *models.GlobalSystemConfigDeleteRequest) (*models.GlobalSystemConfigDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteGlobalSystemConfig(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.GlobalSystemConfigDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowGlobalSystemConfig a REST Show request.
func (service *ContrailService) RESTShowGlobalSystemConfig(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.GlobalSystemConfig
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListGlobalSystemConfig(tx, &common.ListSpec{
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
        "global_system_config": result,
    })
}

//RESTListGlobalSystemConfig handles a List REST service Request.
func (service *ContrailService) RESTListGlobalSystemConfig(c echo.Context) (error) {
    var result []*models.GlobalSystemConfig
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListGlobalSystemConfig(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "global-system-configs": result,
    })
}