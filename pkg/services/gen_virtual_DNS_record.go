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

//RESTCreateVirtualDNSRecord handle a Create REST service.
// nolint
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
// nolint
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
		if model.DisplayName != "" {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
		} else {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.UUID}
		}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()

	return service.Next().CreateVirtualDNSRecord(ctx, request)
}

//RESTUpdateVirtualDNSRecord handles a REST Update request.
// nolint
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
// nolint
func (service *ContrailService) UpdateVirtualDNSRecord(
	ctx context.Context,
	request *models.UpdateVirtualDNSRecordRequest) (*models.UpdateVirtualDNSRecordResponse, error) {
	model := request.VirtualDNSRecord
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateVirtualDNSRecord(ctx, request)
}

//RESTDeleteVirtualDNSRecord delete a resource using REST service.
// nolint
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

//RESTGetVirtualDNSRecord a REST Get request.
// nolint
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

//RESTListVirtualDNSRecord handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListVirtualDNSRecord(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
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
