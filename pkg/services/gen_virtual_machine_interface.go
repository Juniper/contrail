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

//RESTCreateVirtualMachineInterface handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateVirtualMachineInterface(c echo.Context) error {
	requestData := &models.CreateVirtualMachineInterfaceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualMachineInterface(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualMachineInterface handle a Create API
// nolint
func (service *ContrailService) CreateVirtualMachineInterface(
	ctx context.Context,
	request *models.CreateVirtualMachineInterfaceRequest) (*models.CreateVirtualMachineInterfaceResponse, error) {
	model := request.VirtualMachineInterface
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

	return service.Next().CreateVirtualMachineInterface(ctx, request)
}

//RESTUpdateVirtualMachineInterface handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateVirtualMachineInterface(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualMachineInterfaceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualMachineInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualMachineInterface handles a Update request.
// nolint
func (service *ContrailService) UpdateVirtualMachineInterface(
	ctx context.Context,
	request *models.UpdateVirtualMachineInterfaceRequest) (*models.UpdateVirtualMachineInterfaceResponse, error) {
	model := request.VirtualMachineInterface
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateVirtualMachineInterface(ctx, request)
}

//RESTDeleteVirtualMachineInterface delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteVirtualMachineInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualMachineInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualMachineInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetVirtualMachineInterface a REST Get request.
// nolint
func (service *ContrailService) RESTGetVirtualMachineInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualMachineInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualMachineInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListVirtualMachineInterface handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListVirtualMachineInterface(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListVirtualMachineInterfaceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualMachineInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
