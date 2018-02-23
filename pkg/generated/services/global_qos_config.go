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

//RESTCreateGlobalQosConfig handle a Create REST service.
func (service *ContrailService) RESTCreateGlobalQosConfig(c echo.Context) error {
	requestData := &models.CreateGlobalQosConfigRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_qos_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateGlobalQosConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateGlobalQosConfig handle a Create API
func (service *ContrailService) CreateGlobalQosConfig(
	ctx context.Context,
	request *models.CreateGlobalQosConfigRequest) (*models.CreateGlobalQosConfigResponse, error) {
	model := request.GlobalQosConfig
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
			return db.CreateGlobalQosConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_qos_config",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateGlobalQosConfigResponse{
		GlobalQosConfig: request.GlobalQosConfig,
	}, nil
}

//RESTUpdateGlobalQosConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateGlobalQosConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateGlobalQosConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_qos_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateGlobalQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateGlobalQosConfig handles a Update request.
func (service *ContrailService) UpdateGlobalQosConfig(
	ctx context.Context,
	request *models.UpdateGlobalQosConfigRequest) (*models.UpdateGlobalQosConfigResponse, error) {
	model := request.GlobalQosConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateGlobalQosConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_qos_config",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateGlobalQosConfigResponse{
		GlobalQosConfig: model,
	}, nil
}

//RESTDeleteGlobalQosConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteGlobalQosConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteGlobalQosConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteGlobalQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteGlobalQosConfig delete a resource.
func (service *ContrailService) DeleteGlobalQosConfig(ctx context.Context, request *models.DeleteGlobalQosConfigRequest) (*models.DeleteGlobalQosConfigResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteGlobalQosConfig(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteGlobalQosConfigResponse{
		ID: request.ID,
	}, nil
}

//RESTGetGlobalQosConfig a REST Get request.
func (service *ContrailService) RESTGetGlobalQosConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetGlobalQosConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetGlobalQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetGlobalQosConfig a Get request.
func (service *ContrailService) GetGlobalQosConfig(ctx context.Context, request *models.GetGlobalQosConfigRequest) (response *models.GetGlobalQosConfigResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListGlobalQosConfigRequest{
		Spec: spec,
	}
	var result *models.ListGlobalQosConfigResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListGlobalQosConfig(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.GlobalQosConfigs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetGlobalQosConfigResponse{
		GlobalQosConfig: result.GlobalQosConfigs[0],
	}
	return response, nil
}

//RESTListGlobalQosConfig handles a List REST service Request.
func (service *ContrailService) RESTListGlobalQosConfig(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListGlobalQosConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListGlobalQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListGlobalQosConfig handles a List service Request.
func (service *ContrailService) ListGlobalQosConfig(
	ctx context.Context,
	request *models.ListGlobalQosConfigRequest) (response *models.ListGlobalQosConfigResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListGlobalQosConfig(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
