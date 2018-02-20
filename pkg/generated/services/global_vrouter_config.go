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

//RESTGlobalVrouterConfigUpdateRequest for update request for REST.
type RESTGlobalVrouterConfigUpdateRequest struct {
	Data map[string]interface{} `json:"global-vrouter-config"`
}

//RESTCreateGlobalVrouterConfig handle a Create REST service.
func (service *ContrailService) RESTCreateGlobalVrouterConfig(c echo.Context) error {
	requestData := &models.CreateGlobalVrouterConfigRequest{
		GlobalVrouterConfig: models.MakeGlobalVrouterConfig(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_vrouter_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateGlobalVrouterConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateGlobalVrouterConfig handle a Create API
func (service *ContrailService) CreateGlobalVrouterConfig(
	ctx context.Context,
	request *models.CreateGlobalVrouterConfigRequest) (*models.CreateGlobalVrouterConfigResponse, error) {
	model := request.GlobalVrouterConfig
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
			return db.CreateGlobalVrouterConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_vrouter_config",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateGlobalVrouterConfigResponse{
		GlobalVrouterConfig: request.GlobalVrouterConfig,
	}, nil
}

//RESTUpdateGlobalVrouterConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateGlobalVrouterConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateGlobalVrouterConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_vrouter_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateGlobalVrouterConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateGlobalVrouterConfig handles a Update request.
func (service *ContrailService) UpdateGlobalVrouterConfig(
	ctx context.Context,
	request *models.UpdateGlobalVrouterConfigRequest) (*models.UpdateGlobalVrouterConfigResponse, error) {
	model := request.GlobalVrouterConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateGlobalVrouterConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_vrouter_config",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateGlobalVrouterConfigResponse{
		GlobalVrouterConfig: model,
	}, nil
}

//RESTDeleteGlobalVrouterConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteGlobalVrouterConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteGlobalVrouterConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteGlobalVrouterConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteGlobalVrouterConfig delete a resource.
func (service *ContrailService) DeleteGlobalVrouterConfig(ctx context.Context, request *models.DeleteGlobalVrouterConfigRequest) (*models.DeleteGlobalVrouterConfigResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteGlobalVrouterConfig(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteGlobalVrouterConfigResponse{
		ID: request.ID,
	}, nil
}

//RESTGetGlobalVrouterConfig a REST Get request.
func (service *ContrailService) RESTGetGlobalVrouterConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetGlobalVrouterConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetGlobalVrouterConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetGlobalVrouterConfig a Get request.
func (service *ContrailService) GetGlobalVrouterConfig(ctx context.Context, request *models.GetGlobalVrouterConfigRequest) (response *models.GetGlobalVrouterConfigResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListGlobalVrouterConfigRequest{
		Spec: spec,
	}
	var result *models.ListGlobalVrouterConfigResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListGlobalVrouterConfig(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.GlobalVrouterConfigs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetGlobalVrouterConfigResponse{
		GlobalVrouterConfig: result.GlobalVrouterConfigs[0],
	}
	return response, nil
}

//RESTListGlobalVrouterConfig handles a List REST service Request.
func (service *ContrailService) RESTListGlobalVrouterConfig(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListGlobalVrouterConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListGlobalVrouterConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListGlobalVrouterConfig handles a List service Request.
func (service *ContrailService) ListGlobalVrouterConfig(
	ctx context.Context,
	request *models.ListGlobalVrouterConfigRequest) (response *models.ListGlobalVrouterConfigResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListGlobalVrouterConfig(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
