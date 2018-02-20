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

//RESTCreateDsaRule handle a Create REST service.
func (service *ContrailService) RESTCreateDsaRule(c echo.Context) error {
    requestData := &models.CreateDsaRuleRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "dsa_rule",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateDsaRule(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateDsaRule handle a Create API
func (service *ContrailService) CreateDsaRule(
    ctx context.Context, 
    request *models.CreateDsaRuleRequest) (*models.CreateDsaRuleResponse, error) {
    model := request.DsaRule
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
            return db.CreateDsaRule(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "dsa_rule",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateDsaRuleResponse{
        DsaRule: request.DsaRule,
    }, nil
}

//RESTUpdateDsaRule handles a REST Update request.
func (service *ContrailService) RESTUpdateDsaRule(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateDsaRuleRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "dsa_rule",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateDsaRule(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateDsaRule handles a Update request.
func (service *ContrailService) UpdateDsaRule(
    ctx context.Context, 
    request *models.UpdateDsaRuleRequest) (*models.UpdateDsaRuleResponse, error) {
    model := request.DsaRule
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateDsaRule(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "dsa_rule",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateDsaRuleResponse{
        DsaRule: model,
    }, nil
}

//RESTDeleteDsaRule delete a resource using REST service.
func (service *ContrailService) RESTDeleteDsaRule(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteDsaRuleRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteDsaRule(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteDsaRule delete a resource.
func (service *ContrailService) DeleteDsaRule(ctx context.Context, request *models.DeleteDsaRuleRequest) (*models.DeleteDsaRuleResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteDsaRule(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteDsaRuleResponse{
        ID: request.ID,
    }, nil
}

//RESTGetDsaRule a REST Get request.
func (service *ContrailService) RESTGetDsaRule(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetDsaRuleRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetDsaRule(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetDsaRule a Get request.
func (service *ContrailService) GetDsaRule(ctx context.Context, request *models.GetDsaRuleRequest) (response *models.GetDsaRuleResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListDsaRuleRequest{
        Spec: spec,
    }
    var result *models.ListDsaRuleResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDsaRule(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.DsaRules) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetDsaRuleResponse{
       DsaRule: result.DsaRules[0],
    }
    return response, nil
}

//RESTListDsaRule handles a List REST service Request.
func (service *ContrailService) RESTListDsaRule(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListDsaRuleRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListDsaRule(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListDsaRule handles a List service Request.
func (service *ContrailService) ListDsaRule(
    ctx context.Context, 
    request *models.ListDsaRuleRequest) (response *models.ListDsaRuleResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListDsaRule(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}