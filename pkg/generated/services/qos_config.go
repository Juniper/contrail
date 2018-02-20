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

//RESTQosConfigUpdateRequest for update request for REST.
type RESTQosConfigUpdateRequest struct {
	Data map[string]interface{} `json:"qos-config"`
}

//RESTCreateQosConfig handle a Create REST service.
func (service *ContrailService) RESTCreateQosConfig(c echo.Context) error {
	requestData := &models.CreateQosConfigRequest{
		QosConfig: models.MakeQosConfig(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_config",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateQosConfig(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateQosConfig handle a Create API
func (service *ContrailService) CreateQosConfig(
	ctx context.Context,
	request *models.CreateQosConfigRequest) (*models.CreateQosConfigResponse, error) {
	model := request.QosConfig
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
			return db.CreateQosConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_config",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateQosConfigResponse{
		QosConfig: request.QosConfig,
	}, nil
}

//RESTUpdateQosConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateQosConfig(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateQosConfigRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_config",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateQosConfig handles a Update request.
func (service *ContrailService) UpdateQosConfig(
	ctx context.Context,
	request *models.UpdateQosConfigRequest) (*models.UpdateQosConfigResponse, error) {
	model := request.QosConfig
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateQosConfig(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_config",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateQosConfigResponse{
		QosConfig: model,
	}, nil
}

//RESTDeleteQosConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteQosConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteQosConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteQosConfig delete a resource.
func (service *ContrailService) DeleteQosConfig(ctx context.Context, request *models.DeleteQosConfigRequest) (*models.DeleteQosConfigResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteQosConfig(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteQosConfigResponse{
		ID: request.ID,
	}, nil
}

//RESTGetQosConfig a REST Get request.
func (service *ContrailService) RESTGetQosConfig(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetQosConfigRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetQosConfig a Get request.
func (service *ContrailService) GetQosConfig(ctx context.Context, request *models.GetQosConfigRequest) (response *models.GetQosConfigResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListQosConfigRequest{
		Spec: spec,
	}
	var result *models.ListQosConfigResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListQosConfig(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.QosConfigs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetQosConfigResponse{
		QosConfig: result.QosConfigs[0],
	}
	return response, nil
}

//RESTListQosConfig handles a List REST service Request.
func (service *ContrailService) RESTListQosConfig(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListQosConfigRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListQosConfig(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListQosConfig handles a List service Request.
func (service *ContrailService) ListQosConfig(
	ctx context.Context,
	request *models.ListQosConfigRequest) (response *models.ListQosConfigResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListQosConfig(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
