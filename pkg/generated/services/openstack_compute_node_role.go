package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//RESTOpenstackComputeNodeRoleUpdateRequest for update request for REST.
type RESTOpenstackComputeNodeRoleUpdateRequest struct {
	Data map[string]interface{} `json:"openstack-compute-node-role"`
}

//RESTCreateOpenstackComputeNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateOpenstackComputeNodeRole(c echo.Context) error {
	requestData := &models.CreateOpenstackComputeNodeRoleRequest{
		OpenstackComputeNodeRole: models.MakeOpenstackComputeNodeRole(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
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
	request *models.CreateOpenstackComputeNodeRoleRequest) (*models.CreateOpenstackComputeNodeRoleResponse, error) {
	model := request.OpenstackComputeNodeRole
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
		func(tx *sql.Tx) error {
			return db.CreateOpenstackComputeNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_compute_node_role",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateOpenstackComputeNodeRoleResponse{
		OpenstackComputeNodeRole: request.OpenstackComputeNodeRole,
	}, nil
}

//RESTUpdateOpenstackComputeNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateOpenstackComputeNodeRole(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateOpenstackComputeNodeRoleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_compute_node_role",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateOpenstackComputeNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackComputeNodeRole handles a Update request.
func (service *ContrailService) UpdateOpenstackComputeNodeRole(
	ctx context.Context,
	request *models.UpdateOpenstackComputeNodeRoleRequest) (*models.UpdateOpenstackComputeNodeRoleResponse, error) {
	model := request.OpenstackComputeNodeRole
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateOpenstackComputeNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_compute_node_role",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateOpenstackComputeNodeRoleResponse{
		OpenstackComputeNodeRole: model,
	}, nil
}

//RESTDeleteOpenstackComputeNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteOpenstackComputeNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteOpenstackComputeNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteOpenstackComputeNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteOpenstackComputeNodeRole delete a resource.
func (service *ContrailService) DeleteOpenstackComputeNodeRole(ctx context.Context, request *models.DeleteOpenstackComputeNodeRoleRequest) (*models.DeleteOpenstackComputeNodeRoleResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteOpenstackComputeNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteOpenstackComputeNodeRoleResponse{
		ID: request.ID,
	}, nil
}

//RESTGetOpenstackComputeNodeRole a REST Get request.
func (service *ContrailService) RESTGetOpenstackComputeNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetOpenstackComputeNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetOpenstackComputeNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetOpenstackComputeNodeRole a Get request.
func (service *ContrailService) GetOpenstackComputeNodeRole(ctx context.Context, request *models.GetOpenstackComputeNodeRoleRequest) (response *models.GetOpenstackComputeNodeRoleResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListOpenstackComputeNodeRoleRequest{
		Spec: spec,
	}
	var result *models.ListOpenstackComputeNodeRoleResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListOpenstackComputeNodeRole(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.OpenstackComputeNodeRoles) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetOpenstackComputeNodeRoleResponse{
		OpenstackComputeNodeRole: result.OpenstackComputeNodeRoles[0],
	}
	return response, nil
}

//RESTListOpenstackComputeNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListOpenstackComputeNodeRole(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListOpenstackComputeNodeRoleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListOpenstackComputeNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListOpenstackComputeNodeRole handles a List service Request.
func (service *ContrailService) ListOpenstackComputeNodeRole(
	ctx context.Context,
	request *models.ListOpenstackComputeNodeRoleRequest) (response *models.ListOpenstackComputeNodeRoleResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListOpenstackComputeNodeRole(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
