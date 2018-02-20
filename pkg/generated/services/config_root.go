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

//RESTConfigRootUpdateRequest for update request for REST.
type RESTConfigRootUpdateRequest struct {
	Data map[string]interface{} `json:"config-root"`
}

//RESTCreateConfigRoot handle a Create REST service.
func (service *ContrailService) RESTCreateConfigRoot(c echo.Context) error {
	requestData := &models.CreateConfigRootRequest{
		ConfigRoot: models.MakeConfigRoot(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_root",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateConfigRoot(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateConfigRoot handle a Create API
func (service *ContrailService) CreateConfigRoot(
	ctx context.Context,
	request *models.CreateConfigRootRequest) (*models.CreateConfigRootResponse, error) {
	model := request.ConfigRoot
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
			return db.CreateConfigRoot(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_root",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateConfigRootResponse{
		ConfigRoot: request.ConfigRoot,
	}, nil
}

//RESTUpdateConfigRoot handles a REST Update request.
func (service *ContrailService) RESTUpdateConfigRoot(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateConfigRootRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_root",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateConfigRoot(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateConfigRoot handles a Update request.
func (service *ContrailService) UpdateConfigRoot(
	ctx context.Context,
	request *models.UpdateConfigRootRequest) (*models.UpdateConfigRootResponse, error) {
	model := request.ConfigRoot
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateConfigRoot(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_root",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateConfigRootResponse{
		ConfigRoot: model,
	}, nil
}

//RESTDeleteConfigRoot delete a resource using REST service.
func (service *ContrailService) RESTDeleteConfigRoot(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteConfigRootRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteConfigRoot(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteConfigRoot delete a resource.
func (service *ContrailService) DeleteConfigRoot(ctx context.Context, request *models.DeleteConfigRootRequest) (*models.DeleteConfigRootResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteConfigRoot(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteConfigRootResponse{
		ID: request.ID,
	}, nil
}

//RESTGetConfigRoot a REST Get request.
func (service *ContrailService) RESTGetConfigRoot(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetConfigRootRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetConfigRoot(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetConfigRoot a Get request.
func (service *ContrailService) GetConfigRoot(ctx context.Context, request *models.GetConfigRootRequest) (response *models.GetConfigRootResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListConfigRootRequest{
		Spec: spec,
	}
	var result *models.ListConfigRootResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListConfigRoot(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ConfigRoots) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetConfigRootResponse{
		ConfigRoot: result.ConfigRoots[0],
	}
	return response, nil
}

//RESTListConfigRoot handles a List REST service Request.
func (service *ContrailService) RESTListConfigRoot(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListConfigRootRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListConfigRoot(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListConfigRoot handles a List service Request.
func (service *ContrailService) ListConfigRoot(
	ctx context.Context,
	request *models.ListConfigRootRequest) (response *models.ListConfigRootResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListConfigRoot(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
