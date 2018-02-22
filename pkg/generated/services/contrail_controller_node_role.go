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

//RESTContrailControllerNodeRoleUpdateRequest for update request for REST.
type RESTContrailControllerNodeRoleUpdateRequest struct {
	Data map[string]interface{} `json:"contrail-controller-node-role"`
}

//RESTCreateContrailControllerNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateContrailControllerNodeRole(c echo.Context) error {
	requestData := &models.CreateContrailControllerNodeRoleRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
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
	request *models.CreateContrailControllerNodeRoleRequest) (*models.CreateContrailControllerNodeRoleResponse, error) {
	model := request.ContrailControllerNodeRole
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
			return db.CreateContrailControllerNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node_role",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailControllerNodeRoleResponse{
		ContrailControllerNodeRole: request.ContrailControllerNodeRole,
	}, nil
}

//RESTUpdateContrailControllerNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailControllerNodeRole(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailControllerNodeRoleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node_role",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailControllerNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailControllerNodeRole handles a Update request.
func (service *ContrailService) UpdateContrailControllerNodeRole(
	ctx context.Context,
	request *models.UpdateContrailControllerNodeRoleRequest) (*models.UpdateContrailControllerNodeRoleResponse, error) {
	model := request.ContrailControllerNodeRole
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateContrailControllerNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node_role",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailControllerNodeRoleResponse{
		ContrailControllerNodeRole: model,
	}, nil
}

//RESTDeleteContrailControllerNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailControllerNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailControllerNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailControllerNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailControllerNodeRole delete a resource.
func (service *ContrailService) DeleteContrailControllerNodeRole(ctx context.Context, request *models.DeleteContrailControllerNodeRoleRequest) (*models.DeleteContrailControllerNodeRoleResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteContrailControllerNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailControllerNodeRoleResponse{
		ID: request.ID,
	}, nil
}

//RESTGetContrailControllerNodeRole a REST Get request.
func (service *ContrailService) RESTGetContrailControllerNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailControllerNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailControllerNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetContrailControllerNodeRole a Get request.
func (service *ContrailService) GetContrailControllerNodeRole(ctx context.Context, request *models.GetContrailControllerNodeRoleRequest) (response *models.GetContrailControllerNodeRoleResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListContrailControllerNodeRoleRequest{
		Spec: spec,
	}
	var result *models.ListContrailControllerNodeRoleResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailControllerNodeRole(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailControllerNodeRoles) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailControllerNodeRoleResponse{
		ContrailControllerNodeRole: result.ContrailControllerNodeRoles[0],
	}
	return response, nil
}

//RESTListContrailControllerNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListContrailControllerNodeRole(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailControllerNodeRoleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailControllerNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListContrailControllerNodeRole handles a List service Request.
func (service *ContrailService) ListContrailControllerNodeRole(
	ctx context.Context,
	request *models.ListContrailControllerNodeRoleRequest) (response *models.ListContrailControllerNodeRoleResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListContrailControllerNodeRole(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
