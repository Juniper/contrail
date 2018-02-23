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

//RESTCreateFirewallPolicy handle a Create REST service.
func (service *ContrailService) RESTCreateFirewallPolicy(c echo.Context) error {
    requestData := &models.CreateFirewallPolicyRequest{
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
    request *models.CreateFirewallPolicyRequest) (*models.CreateFirewallPolicyResponse, error) {
    model := request.FirewallPolicy
    if model.UUID == "" {
        model.UUID = uuid.NewV4().String()
    }
    auth := common.GetAuthCTX(ctx)
    if auth == nil {
        return nil, common.ErrorUnauthenticated
    }

    if model.FQName == nil {
        if model.DisplayName == "" {
        return nil, common.ErrorBadRequest("Both of FQName and Display Name is empty") 
        }
        model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
    }
    model.Perms2 = &models.PermType2{}
    model.Perms2.Owner = auth.ProjectID()
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.CreateFirewallPolicy(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "firewall_policy",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateFirewallPolicyResponse{
        FirewallPolicy: request.FirewallPolicy,
    }, nil
}

//RESTUpdateFirewallPolicy handles a REST Update request.
func (service *ContrailService) RESTUpdateFirewallPolicy(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateFirewallPolicyRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "firewall_policy",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateFirewallPolicy(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateFirewallPolicy handles a Update request.
func (service *ContrailService) UpdateFirewallPolicy(
    ctx context.Context, 
    request *models.UpdateFirewallPolicyRequest) (*models.UpdateFirewallPolicyResponse, error) {
    model := request.FirewallPolicy
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateFirewallPolicy(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "firewall_policy",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateFirewallPolicyResponse{
        FirewallPolicy: model,
    }, nil
}

//RESTDeleteFirewallPolicy delete a resource using REST service.
func (service *ContrailService) RESTDeleteFirewallPolicy(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteFirewallPolicyRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteFirewallPolicy(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteFirewallPolicy delete a resource.
func (service *ContrailService) DeleteFirewallPolicy(ctx context.Context, request *models.DeleteFirewallPolicyRequest) (*models.DeleteFirewallPolicyResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteFirewallPolicy(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteFirewallPolicyResponse{
        ID: request.ID,
    }, nil
}

//RESTGetFirewallPolicy a REST Get request.
func (service *ContrailService) RESTGetFirewallPolicy(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetFirewallPolicyRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetFirewallPolicy(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetFirewallPolicy a Get request.
func (service *ContrailService) GetFirewallPolicy(ctx context.Context, request *models.GetFirewallPolicyRequest) (response *models.GetFirewallPolicyResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListFirewallPolicyRequest{
        Spec: spec,
    }
    var result *models.ListFirewallPolicyResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListFirewallPolicy(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.FirewallPolicys) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetFirewallPolicyResponse{
       FirewallPolicy: result.FirewallPolicys[0],
    }
    return response, nil
}

//RESTListFirewallPolicy handles a List REST service Request.
func (service *ContrailService) RESTListFirewallPolicy(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListFirewallPolicyRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListFirewallPolicy(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListFirewallPolicy handles a List service Request.
func (service *ContrailService) ListFirewallPolicy(
    ctx context.Context, 
    request *models.ListFirewallPolicyRequest) (response *models.ListFirewallPolicyResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListFirewallPolicy(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}