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

//RESTCreateNetworkDeviceConfig handle a Create REST service.
func (service *ContrailService) RESTCreateNetworkDeviceConfig(c echo.Context) error {
	requestData := &models.CreateNetworkDeviceConfigRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_device_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateNetworkDeviceConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateNetworkDeviceConfig handle a Create API
func (service *ContrailService) CreateNetworkDeviceConfig(
	ctx context.Context,
	request *models.CreateNetworkDeviceConfigRequest) (*models.CreateNetworkDeviceConfigResponse, error) {
	model := request.NetworkDeviceConfig
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
			return db.CreateNetworkDeviceConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_device_config",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateNetworkDeviceConfigResponse{
		NetworkDeviceConfig: request.NetworkDeviceConfig,
	}, nil
}

//RESTUpdateNetworkDeviceConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateNetworkDeviceConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateNetworkDeviceConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_device_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateNetworkDeviceConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateNetworkDeviceConfig handles a Update request.
func (service *ContrailService) UpdateNetworkDeviceConfig(
	ctx context.Context,
	request *models.UpdateNetworkDeviceConfigRequest) (*models.UpdateNetworkDeviceConfigResponse, error) {
	model := request.NetworkDeviceConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateNetworkDeviceConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_device_config",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateNetworkDeviceConfigResponse{
		NetworkDeviceConfig: model,
	}, nil
}

//RESTDeleteNetworkDeviceConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteNetworkDeviceConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteNetworkDeviceConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteNetworkDeviceConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteNetworkDeviceConfig delete a resource.
func (service *ContrailService) DeleteNetworkDeviceConfig(ctx context.Context, request *models.DeleteNetworkDeviceConfigRequest) (*models.DeleteNetworkDeviceConfigResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteNetworkDeviceConfig(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteNetworkDeviceConfigResponse{
		ID: request.ID,
	}, nil
}

//RESTGetNetworkDeviceConfig a REST Get request.
func (service *ContrailService) RESTGetNetworkDeviceConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetNetworkDeviceConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetNetworkDeviceConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetNetworkDeviceConfig a Get request.
func (service *ContrailService) GetNetworkDeviceConfig(ctx context.Context, request *models.GetNetworkDeviceConfigRequest) (response *models.GetNetworkDeviceConfigResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListNetworkDeviceConfigRequest{
		Spec: spec,
	}
	var result *models.ListNetworkDeviceConfigResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListNetworkDeviceConfig(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.NetworkDeviceConfigs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetNetworkDeviceConfigResponse{
		NetworkDeviceConfig: result.NetworkDeviceConfigs[0],
	}
	return response, nil
}

//RESTListNetworkDeviceConfig handles a List REST service Request.
func (service *ContrailService) RESTListNetworkDeviceConfig(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListNetworkDeviceConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListNetworkDeviceConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListNetworkDeviceConfig handles a List service Request.
func (service *ContrailService) ListNetworkDeviceConfig(
	ctx context.Context,
	request *models.ListNetworkDeviceConfigRequest) (response *models.ListNetworkDeviceConfigResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListNetworkDeviceConfig(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
