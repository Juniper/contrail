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

//RESTAppformixNodeRoleUpdateRequest for update request for REST.
type RESTAppformixNodeRoleUpdateRequest struct {
	Data map[string]interface{} `json:"appformix-node-role"`
}

//RESTCreateAppformixNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateAppformixNodeRole(c echo.Context) error {
	requestData := &models.CreateAppformixNodeRoleRequest{
		AppformixNodeRole: models.MakeAppformixNodeRole(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
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
	request *models.CreateAppformixNodeRoleRequest) (*models.CreateAppformixNodeRoleResponse, error) {
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
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateAppformixNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "appformix_node_role",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAppformixNodeRoleResponse{
		AppformixNodeRole: request.AppformixNodeRole,
	}, nil
}

//RESTUpdateAppformixNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateAppformixNodeRole(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAppformixNodeRoleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "appformix_node_role",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAppformixNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAppformixNodeRole handles a Update request.
func (service *ContrailService) UpdateAppformixNodeRole(
	ctx context.Context,
	request *models.UpdateAppformixNodeRoleRequest) (*models.UpdateAppformixNodeRoleResponse, error) {
	model := request.AppformixNodeRole
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateAppformixNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "appformix_node_role",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAppformixNodeRoleResponse{
		AppformixNodeRole: model,
	}, nil
}

//RESTDeleteAppformixNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteAppformixNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAppformixNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAppformixNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteAppformixNodeRole delete a resource.
func (service *ContrailService) DeleteAppformixNodeRole(ctx context.Context, request *models.DeleteAppformixNodeRoleRequest) (*models.DeleteAppformixNodeRoleResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteAppformixNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAppformixNodeRoleResponse{
		ID: request.ID,
	}, nil
}

//RESTGetAppformixNodeRole a REST Get request.
func (service *ContrailService) RESTGetAppformixNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAppformixNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAppformixNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetAppformixNodeRole a Get request.
func (service *ContrailService) GetAppformixNodeRole(ctx context.Context, request *models.GetAppformixNodeRoleRequest) (response *models.GetAppformixNodeRoleResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListAppformixNodeRoleRequest{
		Spec: spec,
	}
	var result *models.ListAppformixNodeRoleResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListAppformixNodeRole(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.AppformixNodeRoles) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAppformixNodeRoleResponse{
		AppformixNodeRole: result.AppformixNodeRoles[0],
	}
	return response, nil
}

//RESTListAppformixNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListAppformixNodeRole(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAppformixNodeRoleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAppformixNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListAppformixNodeRole handles a List service Request.
func (service *ContrailService) ListAppformixNodeRole(
	ctx context.Context,
	request *models.ListAppformixNodeRoleRequest) (response *models.ListAppformixNodeRoleResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListAppformixNodeRole(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
