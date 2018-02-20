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

//RESTServiceEndpointUpdateRequest for update request for REST.
type RESTServiceEndpointUpdateRequest struct {
	Data map[string]interface{} `json:"service-endpoint"`
}

//RESTCreateServiceEndpoint handle a Create REST service.
func (service *ContrailService) RESTCreateServiceEndpoint(c echo.Context) error {
	requestData := &models.CreateServiceEndpointRequest{
		ServiceEndpoint: models.MakeServiceEndpoint(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_endpoint",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceEndpoint(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceEndpoint handle a Create API
func (service *ContrailService) CreateServiceEndpoint(
	ctx context.Context,
	request *models.CreateServiceEndpointRequest) (*models.CreateServiceEndpointResponse, error) {
	model := request.ServiceEndpoint
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
			return db.CreateServiceEndpoint(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_endpoint",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceEndpointResponse{
		ServiceEndpoint: request.ServiceEndpoint,
	}, nil
}

//RESTUpdateServiceEndpoint handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceEndpoint(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceEndpointRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_endpoint",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceEndpoint(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceEndpoint handles a Update request.
func (service *ContrailService) UpdateServiceEndpoint(
	ctx context.Context,
	request *models.UpdateServiceEndpointRequest) (*models.UpdateServiceEndpointResponse, error) {
	model := request.ServiceEndpoint
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateServiceEndpoint(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_endpoint",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceEndpointResponse{
		ServiceEndpoint: model,
	}, nil
}

//RESTDeleteServiceEndpoint delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceEndpoint(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceEndpointRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceEndpoint(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceEndpoint delete a resource.
func (service *ContrailService) DeleteServiceEndpoint(ctx context.Context, request *models.DeleteServiceEndpointRequest) (*models.DeleteServiceEndpointResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServiceEndpoint(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceEndpointResponse{
		ID: request.ID,
	}, nil
}

//RESTGetServiceEndpoint a REST Get request.
func (service *ContrailService) RESTGetServiceEndpoint(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceEndpointRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceEndpoint(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetServiceEndpoint a Get request.
func (service *ContrailService) GetServiceEndpoint(ctx context.Context, request *models.GetServiceEndpointRequest) (response *models.GetServiceEndpointResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListServiceEndpointRequest{
		Spec: spec,
	}
	var result *models.ListServiceEndpointResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceEndpoint(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceEndpoints) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceEndpointResponse{
		ServiceEndpoint: result.ServiceEndpoints[0],
	}
	return response, nil
}

//RESTListServiceEndpoint handles a List REST service Request.
func (service *ContrailService) RESTListServiceEndpoint(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceEndpointRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceEndpoint(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListServiceEndpoint handles a List service Request.
func (service *ContrailService) ListServiceEndpoint(
	ctx context.Context,
	request *models.ListServiceEndpointRequest) (response *models.ListServiceEndpointResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListServiceEndpoint(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
