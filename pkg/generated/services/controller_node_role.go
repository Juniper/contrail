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

//RESTCreateControllerNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateControllerNodeRole(c echo.Context) error {
    requestData := &models.ControllerNodeRoleCreateRequest{
        ControllerNodeRole: models.MakeControllerNodeRole(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "controller_node_role",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateControllerNodeRole(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateControllerNodeRole handle a Create API
func (service *ContrailService) CreateControllerNodeRole(
    ctx context.Context, 
    request *models.ControllerNodeRoleCreateRequest) (*models.ControllerNodeRoleCreateResponse, error) {
    model := request.ControllerNodeRole
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
            return db.CreateControllerNodeRole(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "controller_node_role",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ControllerNodeRoleCreateResponse{
        ControllerNodeRole: request.ControllerNodeRole,
    }, nil
}

//RESTUpdateControllerNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateControllerNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.ControllerNodeRoleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "controller_node_role",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateControllerNodeRole(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateControllerNodeRole handles a Update request.
func (service *ContrailService) UpdateControllerNodeRole(ctx context.Context, request *models.ControllerNodeRoleUpdateRequest) (*models.ControllerNodeRoleUpdateResponse, error) {
    id = request.ID
    model = request.ControllerNodeRole
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
            return db.UpdateControllerNodeRole(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "controller_node_role",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ControllerNodeRole.UpdateResponse{
        ControllerNodeRole: model,
    }, nil
}

//RESTDeleteControllerNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteControllerNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.ControllerNodeRoleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteControllerNodeRole(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteControllerNodeRole delete a resource.
func (service *ContrailService) DeleteControllerNodeRole(ctx context.Context, request *models.ControllerNodeRoleDeleteRequest) (*models.ControllerNodeRoleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteControllerNodeRole(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ControllerNodeRoleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowControllerNodeRole a REST Show request.
func (service *ContrailService) RESTShowControllerNodeRole(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ControllerNodeRole
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListControllerNodeRole(tx, &common.ListSpec{
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
        "controller_node_role": result,
    })
}

//RESTListControllerNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListControllerNodeRole(c echo.Context) (error) {
    var result []*models.ControllerNodeRole
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListControllerNodeRole(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "controller-node-roles": result,
    })
}