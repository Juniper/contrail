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

//RESTFirewallRuleUpdateRequest for update request for REST.
type RESTFirewallRuleUpdateRequest struct {
    Data map[string]interface{} `json:"firewall-rule"`
}

//RESTCreateFirewallRule handle a Create REST service.
func (service *ContrailService) RESTCreateFirewallRule(c echo.Context) error {
    requestData := &models.FirewallRuleCreateRequest{
        FirewallRule: models.MakeFirewallRule(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "firewall_rule",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateFirewallRule(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateFirewallRule handle a Create API
func (service *ContrailService) CreateFirewallRule(
    ctx context.Context, 
    request *models.FirewallRuleCreateRequest) (*models.FirewallRuleCreateResponse, error) {
    model := request.FirewallRule
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
            return db.CreateFirewallRule(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "firewall_rule",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.FirewallRuleCreateResponse{
        FirewallRule: request.FirewallRule,
    }, nil
}

//RESTUpdateFirewallRule handles a REST Update request.
func (service *ContrailService) RESTUpdateFirewallRule(c echo.Context) error {
    id := c.Param("id")
    request := &models.FirewallRuleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "firewall_rule",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateFirewallRule(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateFirewallRule handles a Update request.
func (service *ContrailService) UpdateFirewallRule(ctx context.Context, request *models.FirewallRuleUpdateRequest) (*models.FirewallRuleUpdateResponse, error) {
    id = request.ID
    model = request.FirewallRule
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
            return db.UpdateFirewallRule(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "firewall_rule",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.FirewallRule.UpdateResponse{
        FirewallRule: model,
    }, nil
}

//RESTDeleteFirewallRule delete a resource using REST service.
func (service *ContrailService) RESTDeleteFirewallRule(c echo.Context) error {
    id := c.Param("id")
    request := &models.FirewallRuleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteFirewallRule(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteFirewallRule delete a resource.
func (service *ContrailService) DeleteFirewallRule(ctx context.Context, request *models.FirewallRuleDeleteRequest) (*models.FirewallRuleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteFirewallRule(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.FirewallRuleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowFirewallRule a REST Show request.
func (service *ContrailService) RESTShowFirewallRule(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.FirewallRule
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListFirewallRule(tx, &common.ListSpec{
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
        "firewall_rule": result,
    })
}

//RESTListFirewallRule handles a List REST service Request.
func (service *ContrailService) RESTListFirewallRule(c echo.Context) (error) {
    var result []*models.FirewallRule
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListFirewallRule(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "firewall-rules": result,
    })
}