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

//RESTPeeringPolicyUpdateRequest for update request for REST.
type RESTPeeringPolicyUpdateRequest struct {
    Data map[string]interface{} `json:"peering-policy"`
}

//RESTCreatePeeringPolicy handle a Create REST service.
func (service *ContrailService) RESTCreatePeeringPolicy(c echo.Context) error {
    requestData := &models.PeeringPolicyCreateRequest{
        PeeringPolicy: models.MakePeeringPolicy(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "peering_policy",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreatePeeringPolicy(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreatePeeringPolicy handle a Create API
func (service *ContrailService) CreatePeeringPolicy(
    ctx context.Context, 
    request *models.PeeringPolicyCreateRequest) (*models.PeeringPolicyCreateResponse, error) {
    model := request.PeeringPolicy
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
            return db.CreatePeeringPolicy(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "peering_policy",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.PeeringPolicyCreateResponse{
        PeeringPolicy: request.PeeringPolicy,
    }, nil
}

//RESTUpdatePeeringPolicy handles a REST Update request.
func (service *ContrailService) RESTUpdatePeeringPolicy(c echo.Context) error {
    id := c.Param("id")
    request := &models.PeeringPolicyUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "peering_policy",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdatePeeringPolicy(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdatePeeringPolicy handles a Update request.
func (service *ContrailService) UpdatePeeringPolicy(ctx context.Context, request *models.PeeringPolicyUpdateRequest) (*models.PeeringPolicyUpdateResponse, error) {
    id = request.ID
    model = request.PeeringPolicy
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
            return db.UpdatePeeringPolicy(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "peering_policy",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.PeeringPolicy.UpdateResponse{
        PeeringPolicy: model,
    }, nil
}

//RESTDeletePeeringPolicy delete a resource using REST service.
func (service *ContrailService) RESTDeletePeeringPolicy(c echo.Context) error {
    id := c.Param("id")
    request := &models.PeeringPolicyDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeletePeeringPolicy(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeletePeeringPolicy delete a resource.
func (service *ContrailService) DeletePeeringPolicy(ctx context.Context, request *models.PeeringPolicyDeleteRequest) (*models.PeeringPolicyDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeletePeeringPolicy(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.PeeringPolicyDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowPeeringPolicy a REST Show request.
func (service *ContrailService) RESTShowPeeringPolicy(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.PeeringPolicy
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPeeringPolicy(tx, &common.ListSpec{
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
        "peering_policy": result,
    })
}

//RESTListPeeringPolicy handles a List REST service Request.
func (service *ContrailService) RESTListPeeringPolicy(c echo.Context) (error) {
    var result []*models.PeeringPolicy
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPeeringPolicy(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "peering-policys": result,
    })
}