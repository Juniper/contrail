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

//RESTCreateWidget handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateWidget(c echo.Context) error {
	requestData := &models.CreateWidgetRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "widget",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateWidget(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateWidget handle a Create API
// nolint
func (service *ContrailService) CreateWidget(
	ctx context.Context,
	request *models.CreateWidgetRequest) (*models.CreateWidgetResponse, error) {
	model := request.Widget
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

	return service.Next().CreateWidget(ctx, request)
}

//RESTUpdateWidget handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateWidget(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateWidgetRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "widget",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateWidget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateWidget handles a Update request.
// nolint
func (service *ContrailService) UpdateWidget(
	ctx context.Context,
	request *models.UpdateWidgetRequest) (*models.UpdateWidgetResponse, error) {
	model := request.Widget
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateWidget(ctx, request)
}

//RESTDeleteWidget delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteWidget(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteWidgetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteWidget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetWidget a REST Get request.
// nolint
func (service *ContrailService) RESTGetWidget(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetWidgetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetWidget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListWidget handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListWidget(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListWidgetRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListWidget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
