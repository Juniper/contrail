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

//RESTPolicyManagementUpdateRequest for update request for REST.
type RESTPolicyManagementUpdateRequest struct {
    Data map[string]interface{} `json:"policy-management"`
}

//RESTCreatePolicyManagement handle a Create REST service.
func (service *ContrailService) RESTCreatePolicyManagement(c echo.Context) error {
    requestData := &models.PolicyManagementCreateRequest{
        PolicyManagement: models.MakePolicyManagement(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "policy_management",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreatePolicyManagement(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreatePolicyManagement handle a Create API
func (service *ContrailService) CreatePolicyManagement(
    ctx context.Context, 
    request *models.PolicyManagementCreateRequest) (*models.PolicyManagementCreateResponse, error) {
    model := request.PolicyManagement
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
            return db.CreatePolicyManagement(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "policy_management",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.PolicyManagementCreateResponse{
        PolicyManagement: request.PolicyManagement,
    }, nil
}

//RESTUpdatePolicyManagement handles a REST Update request.
func (service *ContrailService) RESTUpdatePolicyManagement(c echo.Context) error {
    id := c.Param("id")
    request := &models.PolicyManagementUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "policy_management",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdatePolicyManagement(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdatePolicyManagement handles a Update request.
func (service *ContrailService) UpdatePolicyManagement(ctx context.Context, request *models.PolicyManagementUpdateRequest) (*models.PolicyManagementUpdateResponse, error) {
    id = request.ID
    model = request.PolicyManagement
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
            return db.UpdatePolicyManagement(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "policy_management",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.PolicyManagement.UpdateResponse{
        PolicyManagement: model,
    }, nil
}

//RESTDeletePolicyManagement delete a resource using REST service.
func (service *ContrailService) RESTDeletePolicyManagement(c echo.Context) error {
    id := c.Param("id")
    request := &models.PolicyManagementDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeletePolicyManagement(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeletePolicyManagement delete a resource.
func (service *ContrailService) DeletePolicyManagement(ctx context.Context, request *models.PolicyManagementDeleteRequest) (*models.PolicyManagementDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeletePolicyManagement(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.PolicyManagementDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowPolicyManagement a REST Show request.
func (service *ContrailService) RESTShowPolicyManagement(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.PolicyManagement
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPolicyManagement(tx, &common.ListSpec{
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
        "policy_management": result,
    })
}

//RESTListPolicyManagement handles a List REST service Request.
func (service *ContrailService) RESTListPolicyManagement(c echo.Context) (error) {
    var result []*models.PolicyManagement
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPolicyManagement(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "policy-managements": result,
    })
}