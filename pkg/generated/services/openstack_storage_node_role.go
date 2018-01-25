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

//RESTCreateOpenstackStorageNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateOpenstackStorageNodeRole(c echo.Context) error {
    requestData := &models.OpenstackStorageNodeRoleCreateRequest{
        OpenstackStorageNodeRole: models.MakeOpenstackStorageNodeRole(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "openstack_storage_node_role",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateOpenstackStorageNodeRole(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateOpenstackStorageNodeRole handle a Create API
func (service *ContrailService) CreateOpenstackStorageNodeRole(
    ctx context.Context, 
    request *models.OpenstackStorageNodeRoleCreateRequest) (*models.OpenstackStorageNodeRoleCreateResponse, error) {
    model := request.OpenstackStorageNodeRole
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
            return db.CreateOpenstackStorageNodeRole(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "openstack_storage_node_role",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.OpenstackStorageNodeRoleCreateResponse{
        OpenstackStorageNodeRole: request.OpenstackStorageNodeRole,
    }, nil
}

//RESTUpdateOpenstackStorageNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateOpenstackStorageNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.OpenstackStorageNodeRoleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "openstack_storage_node_role",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateOpenstackStorageNodeRole(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackStorageNodeRole handles a Update request.
func (service *ContrailService) UpdateOpenstackStorageNodeRole(ctx context.Context, request *models.OpenstackStorageNodeRoleUpdateRequest) (*models.OpenstackStorageNodeRoleUpdateResponse, error) {
    id = request.ID
    model = request.OpenstackStorageNodeRole
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
            return db.UpdateOpenstackStorageNodeRole(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "openstack_storage_node_role",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.OpenstackStorageNodeRole.UpdateResponse{
        OpenstackStorageNodeRole: model,
    }, nil
}

//RESTDeleteOpenstackStorageNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteOpenstackStorageNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.OpenstackStorageNodeRoleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteOpenstackStorageNodeRole(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteOpenstackStorageNodeRole delete a resource.
func (service *ContrailService) DeleteOpenstackStorageNodeRole(ctx context.Context, request *models.OpenstackStorageNodeRoleDeleteRequest) (*models.OpenstackStorageNodeRoleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteOpenstackStorageNodeRole(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.OpenstackStorageNodeRoleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowOpenstackStorageNodeRole a REST Show request.
func (service *ContrailService) RESTShowOpenstackStorageNodeRole(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.OpenstackStorageNodeRole
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListOpenstackStorageNodeRole(tx, &common.ListSpec{
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
        "openstack_storage_node_role": result,
    })
}

//RESTListOpenstackStorageNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListOpenstackStorageNodeRole(c echo.Context) (error) {
    var result []*models.OpenstackStorageNodeRole
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListOpenstackStorageNodeRole(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "openstack-storage-node-roles": result,
    })
}