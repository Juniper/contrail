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

//RESTDatabaseNodeUpdateRequest for update request for REST.
type RESTDatabaseNodeUpdateRequest struct {
	Data map[string]interface{} `json:"database-node"`
}

//RESTCreateDatabaseNode handle a Create REST service.
func (service *ContrailService) RESTCreateDatabaseNode(c echo.Context) error {
	requestData := &models.CreateDatabaseNodeRequest{
		DatabaseNode: models.MakeDatabaseNode(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "database_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateDatabaseNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateDatabaseNode handle a Create API
func (service *ContrailService) CreateDatabaseNode(
	ctx context.Context,
	request *models.CreateDatabaseNodeRequest) (*models.CreateDatabaseNodeResponse, error) {
	model := request.DatabaseNode
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
			return db.CreateDatabaseNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "database_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateDatabaseNodeResponse{
		DatabaseNode: request.DatabaseNode,
	}, nil
}

//RESTUpdateDatabaseNode handles a REST Update request.
func (service *ContrailService) RESTUpdateDatabaseNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateDatabaseNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "database_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateDatabaseNode handles a Update request.
func (service *ContrailService) UpdateDatabaseNode(
	ctx context.Context,
	request *models.UpdateDatabaseNodeRequest) (*models.UpdateDatabaseNodeResponse, error) {
	model := request.DatabaseNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateDatabaseNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "database_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateDatabaseNodeResponse{
		DatabaseNode: model,
	}, nil
}

//RESTDeleteDatabaseNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteDatabaseNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteDatabaseNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteDatabaseNode delete a resource.
func (service *ContrailService) DeleteDatabaseNode(ctx context.Context, request *models.DeleteDatabaseNodeRequest) (*models.DeleteDatabaseNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteDatabaseNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteDatabaseNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetDatabaseNode a REST Get request.
func (service *ContrailService) RESTGetDatabaseNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetDatabaseNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetDatabaseNode a Get request.
func (service *ContrailService) GetDatabaseNode(ctx context.Context, request *models.GetDatabaseNodeRequest) (response *models.GetDatabaseNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListDatabaseNodeRequest{
		Spec: spec,
	}
	var result *models.ListDatabaseNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListDatabaseNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.DatabaseNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetDatabaseNodeResponse{
		DatabaseNode: result.DatabaseNodes[0],
	}
	return response, nil
}

//RESTListDatabaseNode handles a List REST service Request.
func (service *ContrailService) RESTListDatabaseNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListDatabaseNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListDatabaseNode handles a List service Request.
func (service *ContrailService) ListDatabaseNode(
	ctx context.Context,
	request *models.ListDatabaseNodeRequest) (response *models.ListDatabaseNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListDatabaseNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
