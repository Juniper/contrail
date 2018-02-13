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

//RESTNetworkPolicyUpdateRequest for update request for REST.
type RESTNetworkPolicyUpdateRequest struct {
    Data map[string]interface{} `json:"network-policy"`
}

//RESTCreateNetworkPolicy handle a Create REST service.
func (service *ContrailService) RESTCreateNetworkPolicy(c echo.Context) error {
    requestData := &models.NetworkPolicyCreateRequest{
        NetworkPolicy: models.MakeNetworkPolicy(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "network_policy",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateNetworkPolicy(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateNetworkPolicy handle a Create API
func (service *ContrailService) CreateNetworkPolicy(
    ctx context.Context, 
    request *models.NetworkPolicyCreateRequest) (*models.NetworkPolicyCreateResponse, error) {
    model := request.NetworkPolicy
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
            return db.CreateNetworkPolicy(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "network_policy",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.NetworkPolicyCreateResponse{
        NetworkPolicy: request.NetworkPolicy,
    }, nil
}

//RESTUpdateNetworkPolicy handles a REST Update request.
func (service *ContrailService) RESTUpdateNetworkPolicy(c echo.Context) error {
    id := c.Param("id")
    request := &models.NetworkPolicyUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "network_policy",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateNetworkPolicy(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateNetworkPolicy handles a Update request.
func (service *ContrailService) UpdateNetworkPolicy(ctx context.Context, request *models.NetworkPolicyUpdateRequest) (*models.NetworkPolicyUpdateResponse, error) {
    id = request.ID
    model = request.NetworkPolicy
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
            return db.UpdateNetworkPolicy(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "network_policy",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.NetworkPolicy.UpdateResponse{
        NetworkPolicy: model,
    }, nil
}

//RESTDeleteNetworkPolicy delete a resource using REST service.
func (service *ContrailService) RESTDeleteNetworkPolicy(c echo.Context) error {
    id := c.Param("id")
    request := &models.NetworkPolicyDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteNetworkPolicy(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteNetworkPolicy delete a resource.
func (service *ContrailService) DeleteNetworkPolicy(ctx context.Context, request *models.NetworkPolicyDeleteRequest) (*models.NetworkPolicyDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteNetworkPolicy(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.NetworkPolicyDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowNetworkPolicy a REST Show request.
func (service *ContrailService) RESTShowNetworkPolicy(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.NetworkPolicy
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListNetworkPolicy(tx, &common.ListSpec{
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
        "network_policy": result,
    })
}

//RESTListNetworkPolicy handles a List REST service Request.
func (service *ContrailService) RESTListNetworkPolicy(c echo.Context) (error) {
    var result []*models.NetworkPolicy
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListNetworkPolicy(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "network-policys": result,
    })
}