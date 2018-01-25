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

//RESTServiceConnectionModuleUpdateRequest for update request for REST.
type RESTServiceConnectionModuleUpdateRequest struct {
	Data map[string]interface{} `json:"service-connection-module"`
}

//RESTCreateServiceConnectionModule handle a Create REST service.
func (service *ContrailService) RESTCreateServiceConnectionModule(c echo.Context) error {
	requestData := &models.CreateServiceConnectionModuleRequest{
		ServiceConnectionModule: models.MakeServiceConnectionModule(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceConnectionModule(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceConnectionModule handle a Create API
func (service *ContrailService) CreateServiceConnectionModule(
	ctx context.Context,
	request *models.CreateServiceConnectionModuleRequest) (*models.CreateServiceConnectionModuleResponse, error) {
	model := request.ServiceConnectionModule
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
			return db.CreateServiceConnectionModule(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceConnectionModuleResponse{
		ServiceConnectionModule: request.ServiceConnectionModule,
	}, nil
}

//RESTUpdateServiceConnectionModule handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceConnectionModule(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceConnectionModuleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceConnectionModule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceConnectionModule handles a Update request.
func (service *ContrailService) UpdateServiceConnectionModule(
	ctx context.Context,
	request *models.UpdateServiceConnectionModuleRequest) (*models.UpdateServiceConnectionModuleResponse, error) {
	model := request.ServiceConnectionModule
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateServiceConnectionModule(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceConnectionModuleResponse{
		ServiceConnectionModule: model,
	}, nil
}

//RESTDeleteServiceConnectionModule delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceConnectionModule(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceConnectionModuleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceConnectionModule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceConnectionModule delete a resource.
func (service *ContrailService) DeleteServiceConnectionModule(ctx context.Context, request *models.DeleteServiceConnectionModuleRequest) (*models.DeleteServiceConnectionModuleResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServiceConnectionModule(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceConnectionModuleResponse{
		ID: request.ID,
	}, nil
}

//RESTGetServiceConnectionModule a REST Get request.
func (service *ContrailService) RESTGetServiceConnectionModule(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceConnectionModuleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceConnectionModule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetServiceConnectionModule a Get request.
func (service *ContrailService) GetServiceConnectionModule(ctx context.Context, request *models.GetServiceConnectionModuleRequest) (response *models.GetServiceConnectionModuleResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListServiceConnectionModuleRequest{
		Spec: spec,
	}
	var result *models.ListServiceConnectionModuleResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceConnectionModule(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceConnectionModules) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceConnectionModuleResponse{
		ServiceConnectionModule: result.ServiceConnectionModules[0],
	}
	return response, nil
}

//RESTListServiceConnectionModule handles a List REST service Request.
func (service *ContrailService) RESTListServiceConnectionModule(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceConnectionModuleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceConnectionModule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListServiceConnectionModule handles a List service Request.
func (service *ContrailService) ListServiceConnectionModule(
	ctx context.Context,
	request *models.ListServiceConnectionModuleRequest) (response *models.ListServiceConnectionModuleResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListServiceConnectionModule(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
