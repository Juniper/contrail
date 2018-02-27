package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateCustomerAttachment handle a Create REST service.
func (service *ContrailService) RESTCreateCustomerAttachment(c echo.Context) error {
	requestData := &models.CreateCustomerAttachmentRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "customer_attachment",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateCustomerAttachment(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateCustomerAttachment handle a Create API
func (service *ContrailService) CreateCustomerAttachment(
	ctx context.Context,
	request *models.CreateCustomerAttachmentRequest) (*models.CreateCustomerAttachmentResponse, error) {
	model := request.CustomerAttachment
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

	return service.Next().CreateCustomerAttachment(ctx, request)
}

//RESTUpdateCustomerAttachment handles a REST Update request.
func (service *ContrailService) RESTUpdateCustomerAttachment(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateCustomerAttachmentRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "customer_attachment",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateCustomerAttachment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateCustomerAttachment handles a Update request.
func (service *ContrailService) UpdateCustomerAttachment(
	ctx context.Context,
	request *models.UpdateCustomerAttachmentRequest) (*models.UpdateCustomerAttachmentResponse, error) {
	model := request.CustomerAttachment
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateCustomerAttachment(ctx, request)
}

//RESTDeleteCustomerAttachment delete a resource using REST service.
func (service *ContrailService) RESTDeleteCustomerAttachment(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteCustomerAttachmentRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteCustomerAttachment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetCustomerAttachment a REST Get request.
func (service *ContrailService) RESTGetCustomerAttachment(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetCustomerAttachmentRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetCustomerAttachment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListCustomerAttachment handles a List REST service Request.
func (service *ContrailService) RESTListCustomerAttachment(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListCustomerAttachmentRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListCustomerAttachment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
