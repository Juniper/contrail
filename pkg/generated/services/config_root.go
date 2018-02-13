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

//RESTConfigRootUpdateRequest for update request for REST.
type RESTConfigRootUpdateRequest struct {
    Data map[string]interface{} `json:"config-root"`
}

//RESTCreateConfigRoot handle a Create REST service.
func (service *ContrailService) RESTCreateConfigRoot(c echo.Context) error {
    requestData := &models.ConfigRootCreateRequest{
        ConfigRoot: models.MakeConfigRoot(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "config_root",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateConfigRoot(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateConfigRoot handle a Create API
func (service *ContrailService) CreateConfigRoot(
    ctx context.Context, 
    request *models.ConfigRootCreateRequest) (*models.ConfigRootCreateResponse, error) {
    model := request.ConfigRoot
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
            return db.CreateConfigRoot(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "config_root",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ConfigRootCreateResponse{
        ConfigRoot: request.ConfigRoot,
    }, nil
}

//RESTUpdateConfigRoot handles a REST Update request.
func (service *ContrailService) RESTUpdateConfigRoot(c echo.Context) error {
    id := c.Param("id")
    request := &models.ConfigRootUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "config_root",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateConfigRoot(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateConfigRoot handles a Update request.
func (service *ContrailService) UpdateConfigRoot(ctx context.Context, request *models.ConfigRootUpdateRequest) (*models.ConfigRootUpdateResponse, error) {
    id = request.ID
    model = request.ConfigRoot
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
            return db.UpdateConfigRoot(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "config_root",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ConfigRoot.UpdateResponse{
        ConfigRoot: model,
    }, nil
}

//RESTDeleteConfigRoot delete a resource using REST service.
func (service *ContrailService) RESTDeleteConfigRoot(c echo.Context) error {
    id := c.Param("id")
    request := &models.ConfigRootDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteConfigRoot(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteConfigRoot delete a resource.
func (service *ContrailService) DeleteConfigRoot(ctx context.Context, request *models.ConfigRootDeleteRequest) (*models.ConfigRootDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteConfigRoot(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ConfigRootDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowConfigRoot a REST Show request.
func (service *ContrailService) RESTShowConfigRoot(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ConfigRoot
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListConfigRoot(tx, &common.ListSpec{
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
        "config_root": result,
    })
}

//RESTListConfigRoot handles a List REST service Request.
func (service *ContrailService) RESTListConfigRoot(c echo.Context) (error) {
    var result []*models.ConfigRoot
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListConfigRoot(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "config-roots": result,
    })
}