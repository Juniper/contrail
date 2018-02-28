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

//RESTCreatePhysicalInterface handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreatePhysicalInterface(c echo.Context) error {
	requestData := &models.CreatePhysicalInterfaceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_interface",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreatePhysicalInterface(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreatePhysicalInterface handle a Create API
// nolint
func (service *ContrailService) CreatePhysicalInterface(
	ctx context.Context,
	request *models.CreatePhysicalInterfaceRequest) (*models.CreatePhysicalInterfaceResponse, error) {
	model := request.PhysicalInterface
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

	return service.Next().CreatePhysicalInterface(ctx, request)
}

//RESTUpdatePhysicalInterface handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdatePhysicalInterface(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdatePhysicalInterfaceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_interface",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdatePhysicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdatePhysicalInterface handles a Update request.
// nolint
func (service *ContrailService) UpdatePhysicalInterface(
	ctx context.Context,
	request *models.UpdatePhysicalInterfaceRequest) (*models.UpdatePhysicalInterfaceResponse, error) {
	model := request.PhysicalInterface
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdatePhysicalInterface(ctx, request)
}

//RESTDeletePhysicalInterface delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeletePhysicalInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeletePhysicalInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeletePhysicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetPhysicalInterface a REST Get request.
// nolint
func (service *ContrailService) RESTGetPhysicalInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetPhysicalInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetPhysicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListPhysicalInterface handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListPhysicalInterface(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListPhysicalInterfaceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListPhysicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
