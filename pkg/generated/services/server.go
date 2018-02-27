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

//RESTCreateServer handle a Create REST service.
func (service *ContrailService) RESTCreateServer(c echo.Context) error {
	requestData := &models.CreateServerRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "server",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServer(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServer handle a Create API
func (service *ContrailService) CreateServer(
	ctx context.Context,
	request *models.CreateServerRequest) (*models.CreateServerResponse, error) {
	model := request.Server
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

	return service.Next().CreateServer(ctx, request)
}

//RESTUpdateServer handles a REST Update request.
func (service *ContrailService) RESTUpdateServer(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServerRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "server",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServer handles a Update request.
func (service *ContrailService) UpdateServer(
	ctx context.Context,
	request *models.UpdateServerRequest) (*models.UpdateServerResponse, error) {
	model := request.Server
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateServer(ctx, request)
}

//RESTDeleteServer delete a resource using REST service.
func (service *ContrailService) RESTDeleteServer(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetServer a REST Get request.
func (service *ContrailService) RESTGetServer(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListServer handles a List REST service Request.
func (service *ContrailService) RESTListServer(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServerRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
