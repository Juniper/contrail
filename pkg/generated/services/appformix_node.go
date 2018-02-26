package services

import (
	"context"
	"database/sql"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateAppformixNode handle a Create REST service.
func (service *ContrailService) RESTCreateAppformixNode(c echo.Context) error {
	requestData := &models.CreateAppformixNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "appformix_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAppformixNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAppformixNode handle a Create API
func (service *ContrailService) CreateAppformixNode(
	ctx context.Context,
	request *models.CreateAppformixNodeRequest) (*models.CreateAppformixNodeResponse, error) {
	model := request.AppformixNode
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}

	if model.FQName == nil {
		if model.DisplayName != "" {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
		} else {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.UUID}
		}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateAppformixNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "appformix_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAppformixNodeResponse{
		AppformixNode: request.AppformixNode,
	}, nil
}

//RESTUpdateAppformixNode handles a REST Update request.
func (service *ContrailService) RESTUpdateAppformixNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAppformixNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "appformix_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAppformixNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAppformixNode handles a Update request.
func (service *ContrailService) UpdateAppformixNode(
	ctx context.Context,
	request *models.UpdateAppformixNodeRequest) (*models.UpdateAppformixNodeResponse, error) {
	model := request.AppformixNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateAppformixNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "appformix_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAppformixNodeResponse{
		AppformixNode: model,
	}, nil
}

//RESTDeleteAppformixNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteAppformixNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAppformixNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAppformixNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteAppformixNode delete a resource.
func (service *ContrailService) DeleteAppformixNode(ctx context.Context, request *models.DeleteAppformixNodeRequest) (*models.DeleteAppformixNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteAppformixNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAppformixNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetAppformixNode a REST Get request.
func (service *ContrailService) RESTGetAppformixNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAppformixNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAppformixNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetAppformixNode a Get request.
func (service *ContrailService) GetAppformixNode(ctx context.Context, request *models.GetAppformixNodeRequest) (response *models.GetAppformixNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListAppformixNodeRequest{
		Spec: spec,
	}
	var result *models.ListAppformixNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListAppformixNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.AppformixNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAppformixNodeResponse{
		AppformixNode: result.AppformixNodes[0],
	}
	return response, nil
}

//RESTListAppformixNode handles a List REST service Request.
func (service *ContrailService) RESTListAppformixNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAppformixNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAppformixNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListAppformixNode handles a List service Request.
func (service *ContrailService) ListAppformixNode(
	ctx context.Context,
	request *models.ListAppformixNodeRequest) (response *models.ListAppformixNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListAppformixNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
