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

//RESTOpenstackComputeNodeRoleUpdateRequest for update request for REST.
type RESTOpenstackComputeNodeRoleUpdateRequest struct {
    Data map[string]interface{} `json:"openstack-compute-node-role"`
}

//RESTCreateOpenstackComputeNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateOpenstackComputeNodeRole(c echo.Context) error {
    requestData := &models.OpenstackComputeNodeRoleCreateRequest{
        OpenstackComputeNodeRole: models.MakeOpenstackComputeNodeRole(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "openstack_compute_node_role",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateOpenstackComputeNodeRole(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateOpenstackComputeNodeRole handle a Create API
func (service *ContrailService) CreateOpenstackComputeNodeRole(
    ctx context.Context, 
    request *models.OpenstackComputeNodeRoleCreateRequest) (*models.OpenstackComputeNodeRoleCreateResponse, error) {
    model := request.OpenstackComputeNodeRole
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
            return db.CreateOpenstackComputeNodeRole(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "openstack_compute_node_role",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.OpenstackComputeNodeRoleCreateResponse{
        OpenstackComputeNodeRole: request.OpenstackComputeNodeRole,
    }, nil
}

//RESTUpdateOpenstackComputeNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateOpenstackComputeNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.OpenstackComputeNodeRoleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "openstack_compute_node_role",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateOpenstackComputeNodeRole(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackComputeNodeRole handles a Update request.
func (service *ContrailService) UpdateOpenstackComputeNodeRole(ctx context.Context, request *models.OpenstackComputeNodeRoleUpdateRequest) (*models.OpenstackComputeNodeRoleUpdateResponse, error) {
    id = request.ID
    model = request.OpenstackComputeNodeRole
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
            return db.UpdateOpenstackComputeNodeRole(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "openstack_compute_node_role",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.OpenstackComputeNodeRole.UpdateResponse{
        OpenstackComputeNodeRole: model,
    }, nil
}

//RESTDeleteOpenstackComputeNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteOpenstackComputeNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.OpenstackComputeNodeRoleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteOpenstackComputeNodeRole(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteOpenstackComputeNodeRole delete a resource.
func (service *ContrailService) DeleteOpenstackComputeNodeRole(ctx context.Context, request *models.OpenstackComputeNodeRoleDeleteRequest) (*models.OpenstackComputeNodeRoleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteOpenstackComputeNodeRole(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.OpenstackComputeNodeRoleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowOpenstackComputeNodeRole a REST Show request.
func (service *ContrailService) RESTShowOpenstackComputeNodeRole(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.OpenstackComputeNodeRole
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListOpenstackComputeNodeRole(tx, &common.ListSpec{
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
        "openstack_compute_node_role": result,
    })
}

//RESTListOpenstackComputeNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListOpenstackComputeNodeRole(c echo.Context) (error) {
    var result []*models.OpenstackComputeNodeRole
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListOpenstackComputeNodeRole(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "openstack-compute-node-roles": result,
    })
}