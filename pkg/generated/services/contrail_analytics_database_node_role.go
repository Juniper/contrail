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

//RESTContrailAnalyticsDatabaseNodeRoleUpdateRequest for update request for REST.
type RESTContrailAnalyticsDatabaseNodeRoleUpdateRequest struct {
	Data map[string]interface{} `json:"contrail-analytics-database-node-role"`
}

//RESTCreateContrailAnalyticsDatabaseNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateContrailAnalyticsDatabaseNodeRole(c echo.Context) error {
	requestData := &models.CreateContrailAnalyticsDatabaseNodeRoleRequest{
		ContrailAnalyticsDatabaseNodeRole: models.MakeContrailAnalyticsDatabaseNodeRole(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_database_node_role",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailAnalyticsDatabaseNodeRole(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailAnalyticsDatabaseNodeRole handle a Create API
func (service *ContrailService) CreateContrailAnalyticsDatabaseNodeRole(
	ctx context.Context,
	request *models.CreateContrailAnalyticsDatabaseNodeRoleRequest) (*models.CreateContrailAnalyticsDatabaseNodeRoleResponse, error) {
	model := request.ContrailAnalyticsDatabaseNodeRole
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
			return db.CreateContrailAnalyticsDatabaseNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_database_node_role",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailAnalyticsDatabaseNodeRoleResponse{
		ContrailAnalyticsDatabaseNodeRole: request.ContrailAnalyticsDatabaseNodeRole,
	}, nil
}

//RESTUpdateContrailAnalyticsDatabaseNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailAnalyticsDatabaseNodeRole(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailAnalyticsDatabaseNodeRoleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_database_node_role",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailAnalyticsDatabaseNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailAnalyticsDatabaseNodeRole handles a Update request.
func (service *ContrailService) UpdateContrailAnalyticsDatabaseNodeRole(
	ctx context.Context,
	request *models.UpdateContrailAnalyticsDatabaseNodeRoleRequest) (*models.UpdateContrailAnalyticsDatabaseNodeRoleResponse, error) {
	model := request.ContrailAnalyticsDatabaseNodeRole
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateContrailAnalyticsDatabaseNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_database_node_role",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailAnalyticsDatabaseNodeRoleResponse{
		ContrailAnalyticsDatabaseNodeRole: model,
	}, nil
}

//RESTDeleteContrailAnalyticsDatabaseNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailAnalyticsDatabaseNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailAnalyticsDatabaseNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailAnalyticsDatabaseNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailAnalyticsDatabaseNodeRole delete a resource.
func (service *ContrailService) DeleteContrailAnalyticsDatabaseNodeRole(ctx context.Context, request *models.DeleteContrailAnalyticsDatabaseNodeRoleRequest) (*models.DeleteContrailAnalyticsDatabaseNodeRoleResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteContrailAnalyticsDatabaseNodeRole(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailAnalyticsDatabaseNodeRoleResponse{
		ID: request.ID,
	}, nil
}

//RESTGetContrailAnalyticsDatabaseNodeRole a REST Get request.
func (service *ContrailService) RESTGetContrailAnalyticsDatabaseNodeRole(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailAnalyticsDatabaseNodeRoleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailAnalyticsDatabaseNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetContrailAnalyticsDatabaseNodeRole a Get request.
func (service *ContrailService) GetContrailAnalyticsDatabaseNodeRole(ctx context.Context, request *models.GetContrailAnalyticsDatabaseNodeRoleRequest) (response *models.GetContrailAnalyticsDatabaseNodeRoleResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListContrailAnalyticsDatabaseNodeRoleRequest{
		Spec: spec,
	}
	var result *models.ListContrailAnalyticsDatabaseNodeRoleResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailAnalyticsDatabaseNodeRole(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailAnalyticsDatabaseNodeRoles) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailAnalyticsDatabaseNodeRoleResponse{
		ContrailAnalyticsDatabaseNodeRole: result.ContrailAnalyticsDatabaseNodeRoles[0],
	}
	return response, nil
}

//RESTListContrailAnalyticsDatabaseNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListContrailAnalyticsDatabaseNodeRole(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailAnalyticsDatabaseNodeRoleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailAnalyticsDatabaseNodeRole(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListContrailAnalyticsDatabaseNodeRole handles a List service Request.
func (service *ContrailService) ListContrailAnalyticsDatabaseNodeRole(
	ctx context.Context,
	request *models.ListContrailAnalyticsDatabaseNodeRoleRequest) (response *models.ListContrailAnalyticsDatabaseNodeRoleResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListContrailAnalyticsDatabaseNodeRole(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
