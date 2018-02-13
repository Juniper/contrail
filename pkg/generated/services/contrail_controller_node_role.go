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

//RESTContrailControllerNodeRoleUpdateRequest for update request for REST.
type RESTContrailControllerNodeRoleUpdateRequest struct {
    Data map[string]interface{} `json:"contrail-controller-node-role"`
}

//RESTCreateContrailControllerNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateContrailControllerNodeRole(c echo.Context) error {
    requestData := &models.ContrailControllerNodeRoleCreateRequest{
        ContrailControllerNodeRole: models.MakeContrailControllerNodeRole(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_controller_node_role",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateContrailControllerNodeRole(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateContrailControllerNodeRole handle a Create API
func (service *ContrailService) CreateContrailControllerNodeRole(
    ctx context.Context, 
    request *models.ContrailControllerNodeRoleCreateRequest) (*models.ContrailControllerNodeRoleCreateResponse, error) {
    model := request.ContrailControllerNodeRole
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
            return db.CreateContrailControllerNodeRole(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_controller_node_role",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ContrailControllerNodeRoleCreateResponse{
        ContrailControllerNodeRole: request.ContrailControllerNodeRole,
    }, nil
}

//RESTUpdateContrailControllerNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailControllerNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.ContrailControllerNodeRoleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "contrail_controller_node_role",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateContrailControllerNodeRole(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateContrailControllerNodeRole handles a Update request.
func (service *ContrailService) UpdateContrailControllerNodeRole(ctx context.Context, request *models.ContrailControllerNodeRoleUpdateRequest) (*models.ContrailControllerNodeRoleUpdateResponse, error) {
    id = request.ID
    model = request.ContrailControllerNodeRole
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
            return db.UpdateContrailControllerNodeRole(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_controller_node_role",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ContrailControllerNodeRole.UpdateResponse{
        ContrailControllerNodeRole: model,
    }, nil
}

//RESTDeleteContrailControllerNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailControllerNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.ContrailControllerNodeRoleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteContrailControllerNodeRole(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailControllerNodeRole delete a resource.
func (service *ContrailService) DeleteContrailControllerNodeRole(ctx context.Context, request *models.ContrailControllerNodeRoleDeleteRequest) (*models.ContrailControllerNodeRoleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteContrailControllerNodeRole(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ContrailControllerNodeRoleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowContrailControllerNodeRole a REST Show request.
func (service *ContrailService) RESTShowContrailControllerNodeRole(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ContrailControllerNodeRole
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailControllerNodeRole(tx, &common.ListSpec{
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
        "contrail_controller_node_role": result,
    })
}

//RESTListContrailControllerNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListContrailControllerNodeRole(c echo.Context) (error) {
    var result []*models.ContrailControllerNodeRole
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailControllerNodeRole(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "contrail-controller-node-roles": result,
    })
}