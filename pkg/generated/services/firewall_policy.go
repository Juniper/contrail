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

//RESTFirewallPolicyUpdateRequest for update request for REST.
type RESTFirewallPolicyUpdateRequest struct {
    Data map[string]interface{} `json:"firewall-policy"`
}

//RESTCreateFirewallPolicy handle a Create REST service.
func (service *ContrailService) RESTCreateFirewallPolicy(c echo.Context) error {
    requestData := &models.FirewallPolicyCreateRequest{
        FirewallPolicy: models.MakeFirewallPolicy(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "firewall_policy",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateFirewallPolicy(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateFirewallPolicy handle a Create API
func (service *ContrailService) CreateFirewallPolicy(
    ctx context.Context, 
    request *models.FirewallPolicyCreateRequest) (*models.FirewallPolicyCreateResponse, error) {
    model := request.FirewallPolicy
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
            return db.CreateFirewallPolicy(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "firewall_policy",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.FirewallPolicyCreateResponse{
        FirewallPolicy: request.FirewallPolicy,
    }, nil
}

//RESTUpdateFirewallPolicy handles a REST Update request.
func (service *ContrailService) RESTUpdateFirewallPolicy(c echo.Context) error {
    id := c.Param("id")
    request := &models.FirewallPolicyUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "firewall_policy",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateFirewallPolicy(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateFirewallPolicy handles a Update request.
func (service *ContrailService) UpdateFirewallPolicy(ctx context.Context, request *models.FirewallPolicyUpdateRequest) (*models.FirewallPolicyUpdateResponse, error) {
    id = request.ID
    model = request.FirewallPolicy
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
            return db.UpdateFirewallPolicy(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "firewall_policy",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.FirewallPolicy.UpdateResponse{
        FirewallPolicy: model,
    }, nil
}

//RESTDeleteFirewallPolicy delete a resource using REST service.
func (service *ContrailService) RESTDeleteFirewallPolicy(c echo.Context) error {
    id := c.Param("id")
    request := &models.FirewallPolicyDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteFirewallPolicy(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteFirewallPolicy delete a resource.
func (service *ContrailService) DeleteFirewallPolicy(ctx context.Context, request *models.FirewallPolicyDeleteRequest) (*models.FirewallPolicyDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteFirewallPolicy(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.FirewallPolicyDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowFirewallPolicy a REST Show request.
func (service *ContrailService) RESTShowFirewallPolicy(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.FirewallPolicy
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListFirewallPolicy(tx, &common.ListSpec{
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
        "firewall_policy": result,
    })
}

//RESTListFirewallPolicy handles a List REST service Request.
func (service *ContrailService) RESTListFirewallPolicy(c echo.Context) (error) {
    var result []*models.FirewallPolicy
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListFirewallPolicy(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "firewall-policys": result,
    })
}