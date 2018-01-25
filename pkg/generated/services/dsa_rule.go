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
    requestData := &models.DsaRuleCreateRequest{
        DsaRule: models.MakeDsaRule(),
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
    request *models.DsaRuleCreateRequest) (*models.DsaRuleCreateResponse, error) {
    model := request.DsaRule
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
            return db.CreateDsaRule(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "dsa_rule",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.DsaRuleCreateResponse{
        DsaRule: request.DsaRule,
    }, nil
}

//RESTUpdateDsaRule handles a REST Update request.
func (service *ContrailService) RESTUpdateDsaRule(c echo.Context) error {
    id := c.Param("id")
    request := &models.DsaRuleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "dsa_rule",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateDsaRule(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateDsaRule handles a Update request.
func (service *ContrailService) UpdateDsaRule(ctx context.Context, request *models.DsaRuleUpdateRequest) (*models.DsaRuleUpdateResponse, error) {
    id = request.ID
    model = request.DsaRule
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
            return db.UpdateDsaRule(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "dsa_rule",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.DsaRule.UpdateResponse{
        DsaRule: model,
    }, nil
}

//RESTDeleteDsaRule delete a resource using REST service.
func (service *ContrailService) RESTDeleteDsaRule(c echo.Context) error {
    id := c.Param("id")
    request := &models.DsaRuleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteDsaRule(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteDsaRule delete a resource.
func (service *ContrailService) DeleteDsaRule(ctx context.Context, request *models.DsaRuleDeleteRequest) (*models.DsaRuleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteDsaRule(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DsaRuleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowDsaRule a REST Show request.
func (service *ContrailService) RESTShowDsaRule(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.DsaRule
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDsaRule(tx, &common.ListSpec{
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
        "dsa_rule": result,
    })
}

//RESTListDsaRule handles a List REST service Request.
func (service *ContrailService) RESTListDsaRule(c echo.Context) (error) {
    var result []*models.DsaRule
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDsaRule(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "dsa-rules": result,
    })
}