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

//RESTAliasIPPoolUpdateRequest for update request for REST.
type RESTAliasIPPoolUpdateRequest struct {
    Data map[string]interface{} `json:"alias-ip-pool"`
}

//RESTCreateAliasIPPool handle a Create REST service.
func (service *ContrailService) RESTCreateAliasIPPool(c echo.Context) error {
    requestData := &models.AliasIPPoolCreateRequest{
        AliasIPPool: models.MakeAliasIPPool(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "alias_ip_pool",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateAliasIPPool(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateAliasIPPool handle a Create API
func (service *ContrailService) CreateAliasIPPool(
    ctx context.Context, 
    request *models.AliasIPPoolCreateRequest) (*models.AliasIPPoolCreateResponse, error) {
    model := request.AliasIPPool
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
            return db.CreateAliasIPPool(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "alias_ip_pool",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.AliasIPPoolCreateResponse{
        AliasIPPool: request.AliasIPPool,
    }, nil
}

//RESTUpdateAliasIPPool handles a REST Update request.
func (service *ContrailService) RESTUpdateAliasIPPool(c echo.Context) error {
    id := c.Param("id")
    request := &models.AliasIPPoolUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "alias_ip_pool",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateAliasIPPool(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateAliasIPPool handles a Update request.
func (service *ContrailService) UpdateAliasIPPool(ctx context.Context, request *models.AliasIPPoolUpdateRequest) (*models.AliasIPPoolUpdateResponse, error) {
    id = request.ID
    model = request.AliasIPPool
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
            return db.UpdateAliasIPPool(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "alias_ip_pool",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.AliasIPPool.UpdateResponse{
        AliasIPPool: model,
    }, nil
}

//RESTDeleteAliasIPPool delete a resource using REST service.
func (service *ContrailService) RESTDeleteAliasIPPool(c echo.Context) error {
    id := c.Param("id")
    request := &models.AliasIPPoolDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteAliasIPPool(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteAliasIPPool delete a resource.
func (service *ContrailService) DeleteAliasIPPool(ctx context.Context, request *models.AliasIPPoolDeleteRequest) (*models.AliasIPPoolDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteAliasIPPool(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.AliasIPPoolDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowAliasIPPool a REST Show request.
func (service *ContrailService) RESTShowAliasIPPool(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.AliasIPPool
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAliasIPPool(tx, &common.ListSpec{
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
        "alias_ip_pool": result,
    })
}

//RESTListAliasIPPool handles a List REST service Request.
func (service *ContrailService) RESTListAliasIPPool(c echo.Context) (error) {
    var result []*models.AliasIPPool
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAliasIPPool(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "alias-ip-pools": result,
    })
}