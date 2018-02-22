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

//RESTVirtualDNSRecordUpdateRequest for update request for REST.
type RESTVirtualDNSRecordUpdateRequest struct {
	Data map[string]interface{} `json:"virtual-DNS-record"`
}

//RESTCreateVirtualDNSRecord handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualDNSRecord(c echo.Context) error {
	requestData := &models.CreateVirtualDNSRecordRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS_record",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualDNSRecord(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualDNSRecord handle a Create API
func (service *ContrailService) CreateVirtualDNSRecord(
	ctx context.Context,
	request *models.CreateVirtualDNSRecordRequest) (*models.CreateVirtualDNSRecordResponse, error) {
	model := request.VirtualDNSRecord
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
			return db.CreateVirtualDNSRecord(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS_record",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualDNSRecordResponse{
		VirtualDNSRecord: request.VirtualDNSRecord,
	}, nil
}

//RESTUpdateVirtualDNSRecord handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualDNSRecord(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualDNSRecordRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS_record",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualDNSRecord(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualDNSRecord handles a Update request.
func (service *ContrailService) UpdateVirtualDNSRecord(
	ctx context.Context,
	request *models.UpdateVirtualDNSRecordRequest) (*models.UpdateVirtualDNSRecordResponse, error) {
	model := request.VirtualDNSRecord
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateVirtualDNSRecord(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS_record",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualDNSRecordResponse{
		VirtualDNSRecord: model,
	}, nil
}

//RESTDeleteVirtualDNSRecord delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualDNSRecord(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualDNSRecordRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualDNSRecord(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualDNSRecord delete a resource.
func (service *ContrailService) DeleteVirtualDNSRecord(ctx context.Context, request *models.DeleteVirtualDNSRecordRequest) (*models.DeleteVirtualDNSRecordResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteVirtualDNSRecord(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualDNSRecordResponse{
		ID: request.ID,
	}, nil
}

//RESTGetVirtualDNSRecord a REST Get request.
func (service *ContrailService) RESTGetVirtualDNSRecord(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualDNSRecordRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualDNSRecord(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetVirtualDNSRecord a Get request.
func (service *ContrailService) GetVirtualDNSRecord(ctx context.Context, request *models.GetVirtualDNSRecordRequest) (response *models.GetVirtualDNSRecordResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListVirtualDNSRecordRequest{
		Spec: spec,
	}
	var result *models.ListVirtualDNSRecordResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListVirtualDNSRecord(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualDNSRecords) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualDNSRecordResponse{
		VirtualDNSRecord: result.VirtualDNSRecords[0],
	}
	return response, nil
}

//RESTListVirtualDNSRecord handles a List REST service Request.
func (service *ContrailService) RESTListVirtualDNSRecord(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVirtualDNSRecordRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualDNSRecord(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListVirtualDNSRecord handles a List service Request.
func (service *ContrailService) ListVirtualDNSRecord(
	ctx context.Context,
	request *models.ListVirtualDNSRecordRequest) (response *models.ListVirtualDNSRecordResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListVirtualDNSRecord(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
