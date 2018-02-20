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

//RESTControllerNodeRoleUpdateRequest for update request for REST.
type RESTControllerNodeRoleUpdateRequest struct {
	Data map[string]interface{} `json:"controller-node-role"`
}

//RESTCreateControllerNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateControllerNodeRole(c echo.Context) error {
	requestData := &models.CreateControllerNodeRoleRequest{
		ControllerNodeRole: models.MakeControllerNodeRole(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
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
	request *models.CreateControllerNodeRoleRequest) (*models.CreateControllerNodeRoleResponse, error) {
	model := request.ControllerNodeRole
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
			return db.CreateControllerNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "controller_node_role",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateControllerNodeRoleResponse{
		ControllerNodeRole: request.ControllerNodeRole,
	}, nil
}

//RESTUpdateControllerNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateControllerNodeRole(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateControllerNodeRoleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "controller_node_role",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateControllerNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateControllerNodeRole handles a Update request.
func (service *ContrailService) UpdateControllerNodeRole(
	ctx context.Context,
	request *models.UpdateControllerNodeRoleRequest) (*models.UpdateControllerNodeRoleResponse, error) {
	model := request.ControllerNodeRole
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateControllerNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "controller_node_role",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateControllerNodeRoleResponse{
		ControllerNodeRole: model,
	}, nil
}

//RESTDeleteControllerNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteControllerNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteControllerNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteControllerNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteControllerNodeRole delete a resource.
func (service *ContrailService) DeleteControllerNodeRole(ctx context.Context, request *models.DeleteControllerNodeRoleRequest) (*models.DeleteControllerNodeRoleResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteControllerNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteControllerNodeRoleResponse{
		ID: request.ID,
	}, nil
}

//RESTGetControllerNodeRole a REST Get request.
func (service *ContrailService) RESTGetControllerNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetControllerNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetControllerNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetControllerNodeRole a Get request.
func (service *ContrailService) GetControllerNodeRole(ctx context.Context, request *models.GetControllerNodeRoleRequest) (response *models.GetControllerNodeRoleResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListControllerNodeRoleRequest{
		Spec: spec,
	}
	var result *models.ListControllerNodeRoleResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListControllerNodeRole(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ControllerNodeRoles) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetControllerNodeRoleResponse{
		ControllerNodeRole: result.ControllerNodeRoles[0],
	}
	return response, nil
}

//RESTListControllerNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListControllerNodeRole(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListControllerNodeRoleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListControllerNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListControllerNodeRole handles a List service Request.
func (service *ContrailService) ListControllerNodeRole(
	ctx context.Context,
	request *models.ListControllerNodeRoleRequest) (response *models.ListControllerNodeRoleResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListControllerNodeRole(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
