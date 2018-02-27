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

//RESTCreateServiceGroup handle a Create REST service.
func (service *ContrailService) RESTCreateServiceGroup(c echo.Context) error {
	requestData := &models.CreateServiceGroupRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_group",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceGroup(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceGroup handle a Create API
func (service *ContrailService) CreateServiceGroup(
	ctx context.Context,
	request *models.CreateServiceGroupRequest) (*models.CreateServiceGroupResponse, error) {
	model := request.ServiceGroup
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

	return service.Next().CreateServiceGroup(ctx, request)
}

//RESTUpdateServiceGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceGroup(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceGroupRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_group",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceGroup handles a Update request.
func (service *ContrailService) UpdateServiceGroup(
	ctx context.Context,
	request *models.UpdateServiceGroupRequest) (*models.UpdateServiceGroupResponse, error) {
	model := request.ServiceGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateServiceGroup(ctx, request)
}

//RESTDeleteServiceGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetServiceGroup a REST Get request.
func (service *ContrailService) RESTGetServiceGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListServiceGroup handles a List REST service Request.
func (service *ContrailService) RESTListServiceGroup(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceGroupRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
