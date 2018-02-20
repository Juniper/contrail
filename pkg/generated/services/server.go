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

//RESTServerUpdateRequest for update request for REST.
type RESTServerUpdateRequest struct {
	Data map[string]interface{} `json:"server"`
}

//RESTCreateServer handle a Create REST service.
func (service *ContrailService) RESTCreateServer(c echo.Context) error {
	requestData := &models.CreateServerRequest{
		Server: models.MakeServer(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "server",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServer(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServer handle a Create API
func (service *ContrailService) CreateServer(
	ctx context.Context,
	request *models.CreateServerRequest) (*models.CreateServerResponse, error) {
	model := request.Server
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
			return db.CreateServer(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "server",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServerResponse{
		Server: request.Server,
	}, nil
}

//RESTUpdateServer handles a REST Update request.
func (service *ContrailService) RESTUpdateServer(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServerRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "server",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServer handles a Update request.
func (service *ContrailService) UpdateServer(
	ctx context.Context,
	request *models.UpdateServerRequest) (*models.UpdateServerResponse, error) {
	model := request.Server
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateServer(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "server",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServerResponse{
		Server: model,
	}, nil
}

//RESTDeleteServer delete a resource using REST service.
func (service *ContrailService) RESTDeleteServer(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteServer delete a resource.
func (service *ContrailService) DeleteServer(ctx context.Context, request *models.DeleteServerRequest) (*models.DeleteServerResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServer(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServerResponse{
		ID: request.ID,
	}, nil
}

//RESTGetServer a REST Get request.
func (service *ContrailService) RESTGetServer(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetServer a Get request.
func (service *ContrailService) GetServer(ctx context.Context, request *models.GetServerRequest) (response *models.GetServerResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListServerRequest{
		Spec: spec,
	}
	var result *models.ListServerResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServer(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Servers) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServerResponse{
		Server: result.Servers[0],
	}
	return response, nil
}

//RESTListServer handles a List REST service Request.
func (service *ContrailService) RESTListServer(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServerRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListServer handles a List service Request.
func (service *ContrailService) ListServer(
	ctx context.Context,
	request *models.ListServerRequest) (response *models.ListServerResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListServer(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
