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

//RESTCreateBaremetalNode handle a Create REST service.
func (service *ContrailService) RESTCreateBaremetalNode(c echo.Context) error {
	requestData := &models.CreateBaremetalNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBaremetalNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBaremetalNode handle a Create API
func (service *ContrailService) CreateBaremetalNode(
	ctx context.Context,
	request *models.CreateBaremetalNodeRequest) (*models.CreateBaremetalNodeResponse, error) {
	model := request.BaremetalNode
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
			return db.CreateBaremetalNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBaremetalNodeResponse{
		BaremetalNode: request.BaremetalNode,
	}, nil
}

//RESTUpdateBaremetalNode handles a REST Update request.
func (service *ContrailService) RESTUpdateBaremetalNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBaremetalNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBaremetalNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBaremetalNode handles a Update request.
func (service *ContrailService) UpdateBaremetalNode(
	ctx context.Context,
	request *models.UpdateBaremetalNodeRequest) (*models.UpdateBaremetalNodeResponse, error) {
	model := request.BaremetalNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateBaremetalNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBaremetalNodeResponse{
		BaremetalNode: model,
	}, nil
}

//RESTDeleteBaremetalNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteBaremetalNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBaremetalNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBaremetalNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteBaremetalNode delete a resource.
func (service *ContrailService) DeleteBaremetalNode(ctx context.Context, request *models.DeleteBaremetalNodeRequest) (*models.DeleteBaremetalNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteBaremetalNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBaremetalNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetBaremetalNode a REST Get request.
func (service *ContrailService) RESTGetBaremetalNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBaremetalNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBaremetalNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetBaremetalNode a Get request.
func (service *ContrailService) GetBaremetalNode(ctx context.Context, request *models.GetBaremetalNodeRequest) (response *models.GetBaremetalNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListBaremetalNodeRequest{
		Spec: spec,
	}
	var result *models.ListBaremetalNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListBaremetalNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BaremetalNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBaremetalNodeResponse{
		BaremetalNode: result.BaremetalNodes[0],
	}
	return response, nil
}

//RESTListBaremetalNode handles a List REST service Request.
func (service *ContrailService) RESTListBaremetalNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListBaremetalNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBaremetalNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListBaremetalNode handles a List service Request.
func (service *ContrailService) ListBaremetalNode(
	ctx context.Context,
	request *models.ListBaremetalNodeRequest) (response *models.ListBaremetalNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListBaremetalNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
