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

//RESTServiceObjectUpdateRequest for update request for REST.
type RESTServiceObjectUpdateRequest struct {
	Data map[string]interface{} `json:"service-object"`
}

//RESTCreateServiceObject handle a Create REST service.
func (service *ContrailService) RESTCreateServiceObject(c echo.Context) error {
	requestData := &models.CreateServiceObjectRequest{
		ServiceObject: models.MakeServiceObject(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceObject(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceObject handle a Create API
func (service *ContrailService) CreateServiceObject(
	ctx context.Context,
	request *models.CreateServiceObjectRequest) (*models.CreateServiceObjectResponse, error) {
	model := request.ServiceObject
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
			return db.CreateServiceObject(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceObjectResponse{
		ServiceObject: request.ServiceObject,
	}, nil
}

//RESTUpdateServiceObject handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceObject(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceObjectRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceObject handles a Update request.
func (service *ContrailService) UpdateServiceObject(
	ctx context.Context,
	request *models.UpdateServiceObjectRequest) (*models.UpdateServiceObjectResponse, error) {
	model := request.ServiceObject
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateServiceObject(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_object",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceObjectResponse{
		ServiceObject: model,
	}, nil
}

//RESTDeleteServiceObject delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceObject(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceObjectRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceObject delete a resource.
func (service *ContrailService) DeleteServiceObject(ctx context.Context, request *models.DeleteServiceObjectRequest) (*models.DeleteServiceObjectResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServiceObject(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceObjectResponse{
		ID: request.ID,
	}, nil
}

//RESTGetServiceObject a REST Get request.
func (service *ContrailService) RESTGetServiceObject(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceObjectRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetServiceObject a Get request.
func (service *ContrailService) GetServiceObject(ctx context.Context, request *models.GetServiceObjectRequest) (response *models.GetServiceObjectResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListServiceObjectRequest{
		Spec: spec,
	}
	var result *models.ListServiceObjectResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceObject(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceObjects) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceObjectResponse{
		ServiceObject: result.ServiceObjects[0],
	}
	return response, nil
}

//RESTListServiceObject handles a List REST service Request.
func (service *ContrailService) RESTListServiceObject(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceObjectRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListServiceObject handles a List service Request.
func (service *ContrailService) ListServiceObject(
	ctx context.Context,
	request *models.ListServiceObjectRequest) (response *models.ListServiceObjectResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListServiceObject(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
