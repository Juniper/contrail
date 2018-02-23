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

//RESTCreateServiceHealthCheck handle a Create REST service.
func (service *ContrailService) RESTCreateServiceHealthCheck(c echo.Context) error {
	requestData := &models.CreateServiceHealthCheckRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_health_check",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceHealthCheck(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceHealthCheck handle a Create API
func (service *ContrailService) CreateServiceHealthCheck(
	ctx context.Context,
	request *models.CreateServiceHealthCheckRequest) (*models.CreateServiceHealthCheckResponse, error) {
	model := request.ServiceHealthCheck
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
			return db.CreateServiceHealthCheck(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_health_check",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceHealthCheckResponse{
		ServiceHealthCheck: request.ServiceHealthCheck,
	}, nil
}

//RESTUpdateServiceHealthCheck handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceHealthCheck(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceHealthCheckRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_health_check",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceHealthCheck(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceHealthCheck handles a Update request.
func (service *ContrailService) UpdateServiceHealthCheck(
	ctx context.Context,
	request *models.UpdateServiceHealthCheckRequest) (*models.UpdateServiceHealthCheckResponse, error) {
	model := request.ServiceHealthCheck
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateServiceHealthCheck(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_health_check",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceHealthCheckResponse{
		ServiceHealthCheck: model,
	}, nil
}

//RESTDeleteServiceHealthCheck delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceHealthCheck(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceHealthCheckRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceHealthCheck(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceHealthCheck delete a resource.
func (service *ContrailService) DeleteServiceHealthCheck(ctx context.Context, request *models.DeleteServiceHealthCheckRequest) (*models.DeleteServiceHealthCheckResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServiceHealthCheck(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceHealthCheckResponse{
		ID: request.ID,
	}, nil
}

//RESTGetServiceHealthCheck a REST Get request.
func (service *ContrailService) RESTGetServiceHealthCheck(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceHealthCheckRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceHealthCheck(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetServiceHealthCheck a Get request.
func (service *ContrailService) GetServiceHealthCheck(ctx context.Context, request *models.GetServiceHealthCheckRequest) (response *models.GetServiceHealthCheckResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListServiceHealthCheckRequest{
		Spec: spec,
	}
	var result *models.ListServiceHealthCheckResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceHealthCheck(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceHealthChecks) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceHealthCheckResponse{
		ServiceHealthCheck: result.ServiceHealthChecks[0],
	}
	return response, nil
}

//RESTListServiceHealthCheck handles a List REST service Request.
func (service *ContrailService) RESTListServiceHealthCheck(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceHealthCheckRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceHealthCheck(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListServiceHealthCheck handles a List service Request.
func (service *ContrailService) ListServiceHealthCheck(
	ctx context.Context,
	request *models.ListServiceHealthCheckRequest) (response *models.ListServiceHealthCheckResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListServiceHealthCheck(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
