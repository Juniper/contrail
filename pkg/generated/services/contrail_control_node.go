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

//RESTCreateContrailControlNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailControlNode(c echo.Context) error {
	requestData := &models.CreateContrailControlNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_control_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailControlNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailControlNode handle a Create API
func (service *ContrailService) CreateContrailControlNode(
	ctx context.Context,
	request *models.CreateContrailControlNodeRequest) (*models.CreateContrailControlNodeResponse, error) {
	model := request.ContrailControlNode
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
			return db.CreateContrailControlNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_control_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailControlNodeResponse{
		ContrailControlNode: request.ContrailControlNode,
	}, nil
}

//RESTUpdateContrailControlNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailControlNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailControlNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_control_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailControlNode handles a Update request.
func (service *ContrailService) UpdateContrailControlNode(
	ctx context.Context,
	request *models.UpdateContrailControlNodeRequest) (*models.UpdateContrailControlNodeResponse, error) {
	model := request.ContrailControlNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateContrailControlNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_control_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailControlNodeResponse{
		ContrailControlNode: model,
	}, nil
}

//RESTDeleteContrailControlNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailControlNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailControlNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailControlNode delete a resource.
func (service *ContrailService) DeleteContrailControlNode(ctx context.Context, request *models.DeleteContrailControlNodeRequest) (*models.DeleteContrailControlNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteContrailControlNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailControlNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetContrailControlNode a REST Get request.
func (service *ContrailService) RESTGetContrailControlNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailControlNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetContrailControlNode a Get request.
func (service *ContrailService) GetContrailControlNode(ctx context.Context, request *models.GetContrailControlNodeRequest) (response *models.GetContrailControlNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListContrailControlNodeRequest{
		Spec: spec,
	}
	var result *models.ListContrailControlNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailControlNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailControlNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailControlNodeResponse{
		ContrailControlNode: result.ContrailControlNodes[0],
	}
	return response, nil
}

//RESTListContrailControlNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailControlNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailControlNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListContrailControlNode handles a List service Request.
func (service *ContrailService) ListContrailControlNode(
	ctx context.Context,
	request *models.ListContrailControlNodeRequest) (response *models.ListContrailControlNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListContrailControlNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
