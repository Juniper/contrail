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

//RESTCreateContrailControllerNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailControllerNode(c echo.Context) error {
	requestData := &models.CreateContrailControllerNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailControllerNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailControllerNode handle a Create API
func (service *ContrailService) CreateContrailControllerNode(
	ctx context.Context,
	request *models.CreateContrailControllerNodeRequest) (*models.CreateContrailControllerNodeResponse, error) {
	model := request.ContrailControllerNode
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
			return db.CreateContrailControllerNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailControllerNodeResponse{
		ContrailControllerNode: request.ContrailControllerNode,
	}, nil
}

//RESTUpdateContrailControllerNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailControllerNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailControllerNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailControllerNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailControllerNode handles a Update request.
func (service *ContrailService) UpdateContrailControllerNode(
	ctx context.Context,
	request *models.UpdateContrailControllerNodeRequest) (*models.UpdateContrailControllerNodeResponse, error) {
	model := request.ContrailControllerNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateContrailControllerNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailControllerNodeResponse{
		ContrailControllerNode: model,
	}, nil
}

//RESTDeleteContrailControllerNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailControllerNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailControllerNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailControllerNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailControllerNode delete a resource.
func (service *ContrailService) DeleteContrailControllerNode(ctx context.Context, request *models.DeleteContrailControllerNodeRequest) (*models.DeleteContrailControllerNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteContrailControllerNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailControllerNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetContrailControllerNode a REST Get request.
func (service *ContrailService) RESTGetContrailControllerNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailControllerNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailControllerNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetContrailControllerNode a Get request.
func (service *ContrailService) GetContrailControllerNode(ctx context.Context, request *models.GetContrailControllerNodeRequest) (response *models.GetContrailControllerNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListContrailControllerNodeRequest{
		Spec: spec,
	}
	var result *models.ListContrailControllerNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailControllerNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailControllerNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailControllerNodeResponse{
		ContrailControllerNode: result.ContrailControllerNodes[0],
	}
	return response, nil
}

//RESTListContrailControllerNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailControllerNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailControllerNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailControllerNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListContrailControllerNode handles a List service Request.
func (service *ContrailService) ListContrailControllerNode(
	ctx context.Context,
	request *models.ListContrailControllerNodeRequest) (response *models.ListContrailControllerNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListContrailControllerNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
