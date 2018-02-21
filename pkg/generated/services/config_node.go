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

//RESTConfigNodeUpdateRequest for update request for REST.
type RESTConfigNodeUpdateRequest struct {
	Data map[string]interface{} `json:"config-node"`
}

//RESTCreateConfigNode handle a Create REST service.
func (service *ContrailService) RESTCreateConfigNode(c echo.Context) error {
	requestData := &models.CreateConfigNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateConfigNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateConfigNode handle a Create API
func (service *ContrailService) CreateConfigNode(
	ctx context.Context,
	request *models.CreateConfigNodeRequest) (*models.CreateConfigNodeResponse, error) {
	model := request.ConfigNode
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
			return db.CreateConfigNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateConfigNodeResponse{
		ConfigNode: request.ConfigNode,
	}, nil
}

//RESTUpdateConfigNode handles a REST Update request.
func (service *ContrailService) RESTUpdateConfigNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateConfigNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateConfigNode handles a Update request.
func (service *ContrailService) UpdateConfigNode(
	ctx context.Context,
	request *models.UpdateConfigNodeRequest) (*models.UpdateConfigNodeResponse, error) {
	model := request.ConfigNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateConfigNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateConfigNodeResponse{
		ConfigNode: model,
	}, nil
}

//RESTDeleteConfigNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteConfigNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteConfigNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteConfigNode delete a resource.
func (service *ContrailService) DeleteConfigNode(ctx context.Context, request *models.DeleteConfigNodeRequest) (*models.DeleteConfigNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteConfigNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteConfigNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetConfigNode a REST Get request.
func (service *ContrailService) RESTGetConfigNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetConfigNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetConfigNode a Get request.
func (service *ContrailService) GetConfigNode(ctx context.Context, request *models.GetConfigNodeRequest) (response *models.GetConfigNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListConfigNodeRequest{
		Spec: spec,
	}
	var result *models.ListConfigNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListConfigNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ConfigNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetConfigNodeResponse{
		ConfigNode: result.ConfigNodes[0],
	}
	return response, nil
}

//RESTListConfigNode handles a List REST service Request.
func (service *ContrailService) RESTListConfigNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListConfigNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListConfigNode handles a List service Request.
func (service *ContrailService) ListConfigNode(
	ctx context.Context,
	request *models.ListConfigNodeRequest) (response *models.ListConfigNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListConfigNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
