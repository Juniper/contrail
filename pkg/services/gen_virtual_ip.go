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

//RESTCreateVirtualIP handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateVirtualIP(c echo.Context) error {
	requestData := &models.CreateVirtualIPRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_ip",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualIP(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualIP handle a Create API
// nolint
func (service *ContrailService) CreateVirtualIP(
	ctx context.Context,
	request *models.CreateVirtualIPRequest) (*models.CreateVirtualIPResponse, error) {
	model := request.VirtualIP
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

	return service.Next().CreateVirtualIP(ctx, request)
}

//RESTUpdateVirtualIP handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateVirtualIP(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualIPRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_ip",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualIP handles a Update request.
// nolint
func (service *ContrailService) UpdateVirtualIP(
	ctx context.Context,
	request *models.UpdateVirtualIPRequest) (*models.UpdateVirtualIPResponse, error) {
	model := request.VirtualIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateVirtualIP(ctx, request)
}

//RESTDeleteVirtualIP delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteVirtualIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetVirtualIP a REST Get request.
// nolint
func (service *ContrailService) RESTGetVirtualIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListVirtualIP handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListVirtualIP(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListVirtualIPRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
