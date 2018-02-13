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

//RESTBridgeDomainUpdateRequest for update request for REST.
type RESTBridgeDomainUpdateRequest struct {
    Data map[string]interface{} `json:"bridge-domain"`
}

//RESTCreateBridgeDomain handle a Create REST service.
func (service *ContrailService) RESTCreateBridgeDomain(c echo.Context) error {
    requestData := &models.BridgeDomainCreateRequest{
        BridgeDomain: models.MakeBridgeDomain(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bridge_domain",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateBridgeDomain(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateBridgeDomain handle a Create API
func (service *ContrailService) CreateBridgeDomain(
    ctx context.Context, 
    request *models.BridgeDomainCreateRequest) (*models.BridgeDomainCreateResponse, error) {
    model := request.BridgeDomain
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
            return db.CreateBridgeDomain(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bridge_domain",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.BridgeDomainCreateResponse{
        BridgeDomain: request.BridgeDomain,
    }, nil
}

//RESTUpdateBridgeDomain handles a REST Update request.
func (service *ContrailService) RESTUpdateBridgeDomain(c echo.Context) error {
    id := c.Param("id")
    request := &models.BridgeDomainUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "bridge_domain",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateBridgeDomain(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateBridgeDomain handles a Update request.
func (service *ContrailService) UpdateBridgeDomain(ctx context.Context, request *models.BridgeDomainUpdateRequest) (*models.BridgeDomainUpdateResponse, error) {
    id = request.ID
    model = request.BridgeDomain
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
            return db.UpdateBridgeDomain(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bridge_domain",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.BridgeDomain.UpdateResponse{
        BridgeDomain: model,
    }, nil
}

//RESTDeleteBridgeDomain delete a resource using REST service.
func (service *ContrailService) RESTDeleteBridgeDomain(c echo.Context) error {
    id := c.Param("id")
    request := &models.BridgeDomainDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteBridgeDomain(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteBridgeDomain delete a resource.
func (service *ContrailService) DeleteBridgeDomain(ctx context.Context, request *models.BridgeDomainDeleteRequest) (*models.BridgeDomainDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteBridgeDomain(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.BridgeDomainDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowBridgeDomain a REST Show request.
func (service *ContrailService) RESTShowBridgeDomain(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.BridgeDomain
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBridgeDomain(tx, &common.ListSpec{
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
        "bridge_domain": result,
    })
}

//RESTListBridgeDomain handles a List REST service Request.
func (service *ContrailService) RESTListBridgeDomain(c echo.Context) (error) {
    var result []*models.BridgeDomain
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBridgeDomain(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "bridge-domains": result,
    })
}