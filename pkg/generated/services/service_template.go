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

//RESTServiceTemplateUpdateRequest for update request for REST.
type RESTServiceTemplateUpdateRequest struct {
	Data map[string]interface{} `json:"service-template"`
}

//RESTCreateServiceTemplate handle a Create REST service.
func (service *ContrailService) RESTCreateServiceTemplate(c echo.Context) error {
	requestData := &models.CreateServiceTemplateRequest{
		ServiceTemplate: models.MakeServiceTemplate(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_template",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceTemplate(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceTemplate handle a Create API
func (service *ContrailService) CreateServiceTemplate(
	ctx context.Context,
	request *models.CreateServiceTemplateRequest) (*models.CreateServiceTemplateResponse, error) {
	model := request.ServiceTemplate
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
			return db.CreateServiceTemplate(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_template",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceTemplateResponse{
		ServiceTemplate: request.ServiceTemplate,
	}, nil
}

//RESTUpdateServiceTemplate handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceTemplate(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceTemplateRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_template",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceTemplate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceTemplate handles a Update request.
func (service *ContrailService) UpdateServiceTemplate(
	ctx context.Context,
	request *models.UpdateServiceTemplateRequest) (*models.UpdateServiceTemplateResponse, error) {
	model := request.ServiceTemplate
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateServiceTemplate(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_template",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceTemplateResponse{
		ServiceTemplate: model,
	}, nil
}

//RESTDeleteServiceTemplate delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceTemplate(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceTemplateRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceTemplate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceTemplate delete a resource.
func (service *ContrailService) DeleteServiceTemplate(ctx context.Context, request *models.DeleteServiceTemplateRequest) (*models.DeleteServiceTemplateResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServiceTemplate(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceTemplateResponse{
		ID: request.ID,
	}, nil
}

//RESTGetServiceTemplate a REST Get request.
func (service *ContrailService) RESTGetServiceTemplate(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceTemplateRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceTemplate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetServiceTemplate a Get request.
func (service *ContrailService) GetServiceTemplate(ctx context.Context, request *models.GetServiceTemplateRequest) (response *models.GetServiceTemplateResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListServiceTemplateRequest{
		Spec: spec,
	}
	var result *models.ListServiceTemplateResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceTemplate(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceTemplates) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceTemplateResponse{
		ServiceTemplate: result.ServiceTemplates[0],
	}
	return response, nil
}

//RESTListServiceTemplate handles a List REST service Request.
func (service *ContrailService) RESTListServiceTemplate(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceTemplateRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceTemplate(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListServiceTemplate handles a List service Request.
func (service *ContrailService) ListServiceTemplate(
	ctx context.Context,
	request *models.ListServiceTemplateRequest) (response *models.ListServiceTemplateResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListServiceTemplate(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
