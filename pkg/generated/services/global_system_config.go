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

//RESTCreateGlobalSystemConfig handle a Create REST service.
func (service *ContrailService) RESTCreateGlobalSystemConfig(c echo.Context) error {
	requestData := &models.CreateGlobalSystemConfigRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_system_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateGlobalSystemConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateGlobalSystemConfig handle a Create API
func (service *ContrailService) CreateGlobalSystemConfig(
	ctx context.Context,
	request *models.CreateGlobalSystemConfigRequest) (*models.CreateGlobalSystemConfigResponse, error) {
	model := request.GlobalSystemConfig
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
			return db.CreateGlobalSystemConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_system_config",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateGlobalSystemConfigResponse{
		GlobalSystemConfig: request.GlobalSystemConfig,
	}, nil
}

//RESTUpdateGlobalSystemConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateGlobalSystemConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateGlobalSystemConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_system_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateGlobalSystemConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateGlobalSystemConfig handles a Update request.
func (service *ContrailService) UpdateGlobalSystemConfig(
	ctx context.Context,
	request *models.UpdateGlobalSystemConfigRequest) (*models.UpdateGlobalSystemConfigResponse, error) {
	model := request.GlobalSystemConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateGlobalSystemConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_system_config",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateGlobalSystemConfigResponse{
		GlobalSystemConfig: model,
	}, nil
}

//RESTDeleteGlobalSystemConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteGlobalSystemConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteGlobalSystemConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteGlobalSystemConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteGlobalSystemConfig delete a resource.
func (service *ContrailService) DeleteGlobalSystemConfig(ctx context.Context, request *models.DeleteGlobalSystemConfigRequest) (*models.DeleteGlobalSystemConfigResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteGlobalSystemConfig(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteGlobalSystemConfigResponse{
		ID: request.ID,
	}, nil
}

//RESTGetGlobalSystemConfig a REST Get request.
func (service *ContrailService) RESTGetGlobalSystemConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetGlobalSystemConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetGlobalSystemConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetGlobalSystemConfig a Get request.
func (service *ContrailService) GetGlobalSystemConfig(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (response *models.GetGlobalSystemConfigResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListGlobalSystemConfigRequest{
		Spec: spec,
	}
	var result *models.ListGlobalSystemConfigResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListGlobalSystemConfig(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.GlobalSystemConfigs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetGlobalSystemConfigResponse{
		GlobalSystemConfig: result.GlobalSystemConfigs[0],
	}
	return response, nil
}

//RESTListGlobalSystemConfig handles a List REST service Request.
func (service *ContrailService) RESTListGlobalSystemConfig(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListGlobalSystemConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListGlobalSystemConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListGlobalSystemConfig handles a List service Request.
func (service *ContrailService) ListGlobalSystemConfig(
	ctx context.Context,
	request *models.ListGlobalSystemConfigRequest) (response *models.ListGlobalSystemConfigResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListGlobalSystemConfig(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
