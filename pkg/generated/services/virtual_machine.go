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

//RESTCreateVirtualMachine handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualMachine(c echo.Context) error {
	requestData := &models.CreateVirtualMachineRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualMachine(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualMachine handle a Create API
func (service *ContrailService) CreateVirtualMachine(
	ctx context.Context,
	request *models.CreateVirtualMachineRequest) (*models.CreateVirtualMachineResponse, error) {
	model := request.VirtualMachine
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

	return service.Next().CreateVirtualMachine(ctx, request)
}

//RESTUpdateVirtualMachine handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualMachine(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualMachineRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualMachine(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualMachine handles a Update request.
func (service *ContrailService) UpdateVirtualMachine(
	ctx context.Context,
	request *models.UpdateVirtualMachineRequest) (*models.UpdateVirtualMachineResponse, error) {
	model := request.VirtualMachine
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateVirtualMachine(ctx, request)
}

//RESTDeleteVirtualMachine delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualMachine(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualMachineRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualMachine(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetVirtualMachine a REST Get request.
func (service *ContrailService) RESTGetVirtualMachine(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualMachineRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualMachine(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListVirtualMachine handles a List REST service Request.
func (service *ContrailService) RESTListVirtualMachine(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVirtualMachineRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualMachine(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
