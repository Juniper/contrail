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

//RESTCreateServiceAppliance handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateServiceAppliance(c echo.Context) error {
	requestData := &models.CreateServiceApplianceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceAppliance(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceAppliance handle a Create API
// nolint
func (service *ContrailService) CreateServiceAppliance(
	ctx context.Context,
	request *models.CreateServiceApplianceRequest) (*models.CreateServiceApplianceResponse, error) {
	model := request.ServiceAppliance
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

	return service.Next().CreateServiceAppliance(ctx, request)
}

//RESTUpdateServiceAppliance handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateServiceAppliance(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceApplianceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceAppliance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceAppliance handles a Update request.
// nolint
func (service *ContrailService) UpdateServiceAppliance(
	ctx context.Context,
	request *models.UpdateServiceApplianceRequest) (*models.UpdateServiceApplianceResponse, error) {
	model := request.ServiceAppliance
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateServiceAppliance(ctx, request)
}

//RESTDeleteServiceAppliance delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteServiceAppliance(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceApplianceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceAppliance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetServiceAppliance a REST Get request.
// nolint
func (service *ContrailService) RESTGetServiceAppliance(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceApplianceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceAppliance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListServiceAppliance handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListServiceAppliance(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListServiceApplianceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceAppliance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
