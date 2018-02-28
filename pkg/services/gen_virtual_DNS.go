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

//RESTCreateVirtualDNS handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateVirtualDNS(c echo.Context) error {
	requestData := &models.CreateVirtualDNSRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualDNS(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualDNS handle a Create API
// nolint
func (service *ContrailService) CreateVirtualDNS(
	ctx context.Context,
	request *models.CreateVirtualDNSRequest) (*models.CreateVirtualDNSResponse, error) {
	model := request.VirtualDNS
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

	return service.Next().CreateVirtualDNS(ctx, request)
}

//RESTUpdateVirtualDNS handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateVirtualDNS(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualDNSRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualDNS(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualDNS handles a Update request.
// nolint
func (service *ContrailService) UpdateVirtualDNS(
	ctx context.Context,
	request *models.UpdateVirtualDNSRequest) (*models.UpdateVirtualDNSResponse, error) {
	model := request.VirtualDNS
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateVirtualDNS(ctx, request)
}

//RESTDeleteVirtualDNS delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteVirtualDNS(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualDNSRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualDNS(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetVirtualDNS a REST Get request.
// nolint
func (service *ContrailService) RESTGetVirtualDNS(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualDNSRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualDNS(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListVirtualDNS handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListVirtualDNS(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListVirtualDNSRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualDNS(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
