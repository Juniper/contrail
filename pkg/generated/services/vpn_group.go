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

//RESTCreateVPNGroup handle a Create REST service.
func (service *ContrailService) RESTCreateVPNGroup(c echo.Context) error {
	requestData := &models.CreateVPNGroupRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "vpn_group",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVPNGroup(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVPNGroup handle a Create API
func (service *ContrailService) CreateVPNGroup(
	ctx context.Context,
	request *models.CreateVPNGroupRequest) (*models.CreateVPNGroupResponse, error) {
	model := request.VPNGroup
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

	return service.Next().CreateVPNGroup(ctx, request)
}

//RESTUpdateVPNGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateVPNGroup(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVPNGroupRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "vpn_group",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVPNGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVPNGroup handles a Update request.
func (service *ContrailService) UpdateVPNGroup(
	ctx context.Context,
	request *models.UpdateVPNGroupRequest) (*models.UpdateVPNGroupResponse, error) {
	model := request.VPNGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateVPNGroup(ctx, request)
}

//RESTDeleteVPNGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteVPNGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVPNGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVPNGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetVPNGroup a REST Get request.
func (service *ContrailService) RESTGetVPNGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVPNGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVPNGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListVPNGroup handles a List REST service Request.
func (service *ContrailService) RESTListVPNGroup(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVPNGroupRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVPNGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
