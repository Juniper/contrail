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

//RESTProviderAttachmentUpdateRequest for update request for REST.
type RESTProviderAttachmentUpdateRequest struct {
	Data map[string]interface{} `json:"provider-attachment"`
}

//RESTCreateProviderAttachment handle a Create REST service.
func (service *ContrailService) RESTCreateProviderAttachment(c echo.Context) error {
	requestData := &models.CreateProviderAttachmentRequest{
		ProviderAttachment: models.MakeProviderAttachment(),
	}
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
			return db.CreateProviderAttachment(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "provider_attachment",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateProviderAttachmentResponse{
		ProviderAttachment: request.ProviderAttachment,
	}, nil
}

//RESTUpdateProviderAttachment handles a REST Update request.
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
func (service *ContrailService) UpdateProviderAttachment(
	ctx context.Context,
	request *models.UpdateProviderAttachmentRequest) (*models.UpdateProviderAttachmentResponse, error) {
	model := request.ProviderAttachment
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateProviderAttachment(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "provider_attachment",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateProviderAttachmentResponse{
		ProviderAttachment: model,
	}, nil
}

//RESTDeleteProviderAttachment delete a resource using REST service.
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

//DeleteProviderAttachment delete a resource.
func (service *ContrailService) DeleteProviderAttachment(ctx context.Context, request *models.DeleteProviderAttachmentRequest) (*models.DeleteProviderAttachmentResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteProviderAttachment(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteProviderAttachmentResponse{
		ID: request.ID,
	}, nil
}

//RESTGetProviderAttachment a REST Get request.
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

//GetProviderAttachment a Get request.
func (service *ContrailService) GetProviderAttachment(ctx context.Context, request *models.GetProviderAttachmentRequest) (response *models.GetProviderAttachmentResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListProviderAttachmentRequest{
		Spec: spec,
	}
	var result *models.ListProviderAttachmentResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListProviderAttachment(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ProviderAttachments) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetProviderAttachmentResponse{
		ProviderAttachment: result.ProviderAttachments[0],
	}
	return response, nil
}

//RESTListProviderAttachment handles a List REST service Request.
func (service *ContrailService) RESTListProviderAttachment(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
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

//ListProviderAttachment handles a List service Request.
func (service *ContrailService) ListProviderAttachment(
	ctx context.Context,
	request *models.ListProviderAttachmentRequest) (response *models.ListProviderAttachmentResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListProviderAttachment(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
