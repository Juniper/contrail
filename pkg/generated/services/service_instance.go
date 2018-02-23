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

//RESTCreateServiceInstance handle a Create REST service.
func (service *ContrailService) RESTCreateServiceInstance(c echo.Context) error {
	requestData := &models.CreateServiceInstanceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_instance",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceInstance(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceInstance handle a Create API
func (service *ContrailService) CreateServiceInstance(
	ctx context.Context,
	request *models.CreateServiceInstanceRequest) (*models.CreateServiceInstanceResponse, error) {
	model := request.ServiceInstance
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
			return db.CreateServiceInstance(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_instance",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceInstanceResponse{
		ServiceInstance: request.ServiceInstance,
	}, nil
}

//RESTUpdateServiceInstance handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceInstance(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceInstanceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_instance",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceInstance handles a Update request.
func (service *ContrailService) UpdateServiceInstance(
	ctx context.Context,
	request *models.UpdateServiceInstanceRequest) (*models.UpdateServiceInstanceResponse, error) {
	model := request.ServiceInstance
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateServiceInstance(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_instance",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceInstanceResponse{
		ServiceInstance: model,
	}, nil
}

//RESTDeleteServiceInstance delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceInstance(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceInstanceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceInstance delete a resource.
func (service *ContrailService) DeleteServiceInstance(ctx context.Context, request *models.DeleteServiceInstanceRequest) (*models.DeleteServiceInstanceResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServiceInstance(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceInstanceResponse{
		ID: request.ID,
	}, nil
}

//RESTGetServiceInstance a REST Get request.
func (service *ContrailService) RESTGetServiceInstance(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceInstanceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetServiceInstance a Get request.
func (service *ContrailService) GetServiceInstance(ctx context.Context, request *models.GetServiceInstanceRequest) (response *models.GetServiceInstanceResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListServiceInstanceRequest{
		Spec: spec,
	}
	var result *models.ListServiceInstanceResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceInstance(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceInstances) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceInstanceResponse{
		ServiceInstance: result.ServiceInstances[0],
	}
	return response, nil
}

//RESTListServiceInstance handles a List REST service Request.
func (service *ContrailService) RESTListServiceInstance(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceInstanceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListServiceInstance handles a List service Request.
func (service *ContrailService) ListServiceInstance(
	ctx context.Context,
	request *models.ListServiceInstanceRequest) (response *models.ListServiceInstanceResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListServiceInstance(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
