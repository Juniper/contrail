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

//RESTCreateAppformixNode handle a Create REST service.
func (service *ContrailService) RESTCreateAppformixNode(c echo.Context) error {
	requestData := &models.CreateAppformixNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "appformix_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAppformixNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAppformixNode handle a Create API
func (service *ContrailService) CreateAppformixNode(
	ctx context.Context,
	request *models.CreateAppformixNodeRequest) (*models.CreateAppformixNodeResponse, error) {
	model := request.AppformixNode
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

	return service.Next().CreateAppformixNode(ctx, request)
}

//RESTUpdateAppformixNode handles a REST Update request.
func (service *ContrailService) RESTUpdateAppformixNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAppformixNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "appformix_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAppformixNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAppformixNode handles a Update request.
func (service *ContrailService) UpdateAppformixNode(
	ctx context.Context,
	request *models.UpdateAppformixNodeRequest) (*models.UpdateAppformixNodeResponse, error) {
	model := request.AppformixNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateAppformixNode(ctx, request)
}

//RESTDeleteAppformixNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteAppformixNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAppformixNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAppformixNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetAppformixNode a REST Get request.
func (service *ContrailService) RESTGetAppformixNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAppformixNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAppformixNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListAppformixNode handles a List REST service Request.
func (service *ContrailService) RESTListAppformixNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAppformixNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAppformixNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
