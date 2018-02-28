package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateProviderAttachment handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateProviderAttachment(c echo.Context) error {
	requestData := &models.CreateProviderAttachmentRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "provider_attachment",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateProviderAttachment(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateProviderAttachment handle a Create API
// nolint
func (service *ContrailService) CreateProviderAttachment(
	ctx context.Context,
	request *models.CreateProviderAttachmentRequest) (*models.CreateProviderAttachmentResponse, error) {
	model := request.ProviderAttachment
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

	return service.Next().CreateProviderAttachment(ctx, request)
}

//RESTUpdateProviderAttachment handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateProviderAttachment(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateProviderAttachmentRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "provider_attachment",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateProviderAttachment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateProviderAttachment handles a Update request.
// nolint
func (service *ContrailService) UpdateProviderAttachment(
	ctx context.Context,
	request *models.UpdateProviderAttachmentRequest) (*models.UpdateProviderAttachmentResponse, error) {
	model := request.ProviderAttachment
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateProviderAttachment(ctx, request)
}

//RESTDeleteProviderAttachment delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteProviderAttachment(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteProviderAttachmentRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteProviderAttachment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetProviderAttachment a REST Get request.
// nolint
func (service *ContrailService) RESTGetProviderAttachment(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetProviderAttachmentRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetProviderAttachment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListProviderAttachment handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListProviderAttachment(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListProviderAttachmentRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListProviderAttachment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
