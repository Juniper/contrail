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

//RESTCreateVirtualNetwork handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualNetwork(c echo.Context) error {
	requestData := &models.CreateVirtualNetworkRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_network",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualNetwork(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualNetwork handle a Create API
func (service *ContrailService) CreateVirtualNetwork(
	ctx context.Context,
	request *models.CreateVirtualNetworkRequest) (*models.CreateVirtualNetworkResponse, error) {
	model := request.VirtualNetwork
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

	return service.Next().CreateVirtualNetwork(ctx, request)
}

//RESTUpdateVirtualNetwork handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualNetwork(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualNetworkRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_network",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualNetwork(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualNetwork handles a Update request.
func (service *ContrailService) UpdateVirtualNetwork(
	ctx context.Context,
	request *models.UpdateVirtualNetworkRequest) (*models.UpdateVirtualNetworkResponse, error) {
	model := request.VirtualNetwork
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateVirtualNetwork(ctx, request)
}

//RESTDeleteVirtualNetwork delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualNetwork(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualNetworkRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualNetwork(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetVirtualNetwork a REST Get request.
func (service *ContrailService) RESTGetVirtualNetwork(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualNetworkRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualNetwork(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListVirtualNetwork handles a List REST service Request.
func (service *ContrailService) RESTListVirtualNetwork(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVirtualNetworkRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualNetwork(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
