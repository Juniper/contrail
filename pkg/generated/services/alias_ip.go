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

//RESTAliasIPUpdateRequest for update request for REST.
type RESTAliasIPUpdateRequest struct {
    Data map[string]interface{} `json:"alias-ip"`
}

//RESTCreateAliasIP handle a Create REST service.
func (service *ContrailService) RESTCreateAliasIP(c echo.Context) error {
    requestData := &models.AliasIPCreateRequest{
        AliasIP: models.MakeAliasIP(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "alias_ip",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateAliasIP(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateAliasIP handle a Create API
func (service *ContrailService) CreateAliasIP(
    ctx context.Context, 
    request *models.AliasIPCreateRequest) (*models.AliasIPCreateResponse, error) {
    model := request.AliasIP
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
            return db.CreateAliasIP(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "alias_ip",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.AliasIPCreateResponse{
        AliasIP: request.AliasIP,
    }, nil
}

//RESTUpdateAliasIP handles a REST Update request.
func (service *ContrailService) RESTUpdateAliasIP(c echo.Context) error {
    id := c.Param("id")
    request := &models.AliasIPUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "alias_ip",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateAliasIP(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateAliasIP handles a Update request.
func (service *ContrailService) UpdateAliasIP(ctx context.Context, request *models.AliasIPUpdateRequest) (*models.AliasIPUpdateResponse, error) {
    id = request.ID
    model = request.AliasIP
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
            return db.UpdateAliasIP(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "alias_ip",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.AliasIP.UpdateResponse{
        AliasIP: model,
    }, nil
}

//RESTDeleteAliasIP delete a resource using REST service.
func (service *ContrailService) RESTDeleteAliasIP(c echo.Context) error {
    id := c.Param("id")
    request := &models.AliasIPDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteAliasIP(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteAliasIP delete a resource.
func (service *ContrailService) DeleteAliasIP(ctx context.Context, request *models.AliasIPDeleteRequest) (*models.AliasIPDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteAliasIP(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.AliasIPDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowAliasIP a REST Show request.
func (service *ContrailService) RESTShowAliasIP(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.AliasIP
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAliasIP(tx, &common.ListSpec{
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
        "alias_ip": result,
    })
}

//RESTListAliasIP handles a List REST service Request.
func (service *ContrailService) RESTListAliasIP(c echo.Context) (error) {
    var result []*models.AliasIP
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAliasIP(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "alias-ips": result,
    })
}