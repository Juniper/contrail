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

//RESTCreateServiceApplianceSet handle a Create REST service.
func (service *ContrailService) RESTCreateServiceApplianceSet(c echo.Context) error {
	requestData := &models.CreateServiceApplianceSetRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance_set",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceApplianceSet(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceApplianceSet handle a Create API
func (service *ContrailService) CreateServiceApplianceSet(
	ctx context.Context,
	request *models.CreateServiceApplianceSetRequest) (*models.CreateServiceApplianceSetResponse, error) {
	model := request.ServiceApplianceSet
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

	return service.Next().CreateServiceApplianceSet(ctx, request)
}

//RESTUpdateServiceApplianceSet handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceApplianceSet(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceApplianceSetRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance_set",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceApplianceSet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceApplianceSet handles a Update request.
func (service *ContrailService) UpdateServiceApplianceSet(
	ctx context.Context,
	request *models.UpdateServiceApplianceSetRequest) (*models.UpdateServiceApplianceSetResponse, error) {
	model := request.ServiceApplianceSet
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateServiceApplianceSet(ctx, request)
}

//RESTDeleteServiceApplianceSet delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceApplianceSet(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceApplianceSetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceApplianceSet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetServiceApplianceSet a REST Get request.
func (service *ContrailService) RESTGetServiceApplianceSet(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceApplianceSetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceApplianceSet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListServiceApplianceSet handles a List REST service Request.
func (service *ContrailService) RESTListServiceApplianceSet(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceApplianceSetRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceApplianceSet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
