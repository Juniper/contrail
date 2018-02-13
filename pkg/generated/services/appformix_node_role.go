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

//RESTAppformixNodeRoleUpdateRequest for update request for REST.
type RESTAppformixNodeRoleUpdateRequest struct {
    Data map[string]interface{} `json:"appformix-node-role"`
}

//RESTCreateAppformixNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateAppformixNodeRole(c echo.Context) error {
    requestData := &models.AppformixNodeRoleCreateRequest{
        AppformixNodeRole: models.MakeAppformixNodeRole(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "appformix_node_role",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateAppformixNodeRole(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateAppformixNodeRole handle a Create API
func (service *ContrailService) CreateAppformixNodeRole(
    ctx context.Context, 
    request *models.AppformixNodeRoleCreateRequest) (*models.AppformixNodeRoleCreateResponse, error) {
    model := request.AppformixNodeRole
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
            return db.CreateAppformixNodeRole(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "appformix_node_role",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.AppformixNodeRoleCreateResponse{
        AppformixNodeRole: request.AppformixNodeRole,
    }, nil
}

//RESTUpdateAppformixNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateAppformixNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.AppformixNodeRoleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "appformix_node_role",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateAppformixNodeRole(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateAppformixNodeRole handles a Update request.
func (service *ContrailService) UpdateAppformixNodeRole(ctx context.Context, request *models.AppformixNodeRoleUpdateRequest) (*models.AppformixNodeRoleUpdateResponse, error) {
    id = request.ID
    model = request.AppformixNodeRole
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
            return db.UpdateAppformixNodeRole(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "appformix_node_role",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.AppformixNodeRole.UpdateResponse{
        AppformixNodeRole: model,
    }, nil
}

//RESTDeleteAppformixNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteAppformixNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.AppformixNodeRoleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteAppformixNodeRole(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteAppformixNodeRole delete a resource.
func (service *ContrailService) DeleteAppformixNodeRole(ctx context.Context, request *models.AppformixNodeRoleDeleteRequest) (*models.AppformixNodeRoleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteAppformixNodeRole(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.AppformixNodeRoleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowAppformixNodeRole a REST Show request.
func (service *ContrailService) RESTShowAppformixNodeRole(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.AppformixNodeRole
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAppformixNodeRole(tx, &common.ListSpec{
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
        "appformix_node_role": result,
    })
}

//RESTListAppformixNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListAppformixNodeRole(c echo.Context) (error) {
    var result []*models.AppformixNodeRole
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAppformixNodeRole(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "appformix-node-roles": result,
    })
}