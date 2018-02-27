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

//RESTCreateDashboard handle a Create REST service.
func (service *ContrailService) RESTCreateDashboard(c echo.Context) error {
	requestData := &models.CreateDashboardRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dashboard",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateDashboard(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateDashboard handle a Create API
func (service *ContrailService) CreateDashboard(
	ctx context.Context,
	request *models.CreateDashboardRequest) (*models.CreateDashboardResponse, error) {
	model := request.Dashboard
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

	return service.Next().CreateDashboard(ctx, request)
}

//RESTUpdateDashboard handles a REST Update request.
func (service *ContrailService) RESTUpdateDashboard(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateDashboardRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dashboard",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateDashboard(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateDashboard handles a Update request.
func (service *ContrailService) UpdateDashboard(
	ctx context.Context,
	request *models.UpdateDashboardRequest) (*models.UpdateDashboardResponse, error) {
	model := request.Dashboard
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateDashboard(ctx, request)
}

//RESTDeleteDashboard delete a resource using REST service.
func (service *ContrailService) RESTDeleteDashboard(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteDashboardRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteDashboard(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetDashboard a REST Get request.
func (service *ContrailService) RESTGetDashboard(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetDashboardRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetDashboard(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListDashboard handles a List REST service Request.
func (service *ContrailService) RESTListDashboard(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListDashboardRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListDashboard(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
