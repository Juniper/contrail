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

//RESTCreateLocation handle a Create REST service.
func (service *ContrailService) RESTCreateLocation(c echo.Context) error {
	requestData := &models.CreateLocationRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "location",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLocation(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLocation handle a Create API
func (service *ContrailService) CreateLocation(
	ctx context.Context,
	request *models.CreateLocationRequest) (*models.CreateLocationResponse, error) {
	model := request.Location
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

	return service.Next().CreateLocation(ctx, request)
}

//RESTUpdateLocation handles a REST Update request.
func (service *ContrailService) RESTUpdateLocation(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLocationRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "location",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLocation(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLocation handles a Update request.
func (service *ContrailService) UpdateLocation(
	ctx context.Context,
	request *models.UpdateLocationRequest) (*models.UpdateLocationResponse, error) {
	model := request.Location
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateLocation(ctx, request)
}

//RESTDeleteLocation delete a resource using REST service.
func (service *ContrailService) RESTDeleteLocation(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLocationRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLocation(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetLocation a REST Get request.
func (service *ContrailService) RESTGetLocation(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLocationRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLocation(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListLocation handles a List REST service Request.
func (service *ContrailService) RESTListLocation(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLocationRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLocation(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
