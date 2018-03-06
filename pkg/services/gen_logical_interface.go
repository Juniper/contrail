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

//RESTCreateLogicalInterface handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateLogicalInterface(c echo.Context) error {
	requestData := &models.CreateLogicalInterfaceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_interface",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLogicalInterface(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLogicalInterface handle a Create API
// nolint
func (service *ContrailService) CreateLogicalInterface(
	ctx context.Context,
	request *models.CreateLogicalInterfaceRequest) (*models.CreateLogicalInterfaceResponse, error) {
	model := request.LogicalInterface
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

	return service.Next().CreateLogicalInterface(ctx, request)
}

//RESTUpdateLogicalInterface handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateLogicalInterface(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLogicalInterfaceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_interface",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLogicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLogicalInterface handles a Update request.
// nolint
func (service *ContrailService) UpdateLogicalInterface(
	ctx context.Context,
	request *models.UpdateLogicalInterfaceRequest) (*models.UpdateLogicalInterfaceResponse, error) {
	model := request.LogicalInterface
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateLogicalInterface(ctx, request)
}

//RESTDeleteLogicalInterface delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteLogicalInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLogicalInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLogicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetLogicalInterface a REST Get request.
// nolint
func (service *ContrailService) RESTGetLogicalInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLogicalInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLogicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListLogicalInterface handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListLogicalInterface(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListLogicalInterfaceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLogicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
