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

//RESTCreateConfigNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateConfigNode(c echo.Context) error {
	requestData := &models.CreateConfigNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateConfigNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateConfigNode handle a Create API
// nolint
func (service *ContrailService) CreateConfigNode(
	ctx context.Context,
	request *models.CreateConfigNodeRequest) (*models.CreateConfigNodeResponse, error) {
	model := request.ConfigNode
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

	return service.Next().CreateConfigNode(ctx, request)
}

//RESTUpdateConfigNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateConfigNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateConfigNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "config_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateConfigNode handles a Update request.
// nolint
func (service *ContrailService) UpdateConfigNode(
	ctx context.Context,
	request *models.UpdateConfigNodeRequest) (*models.UpdateConfigNodeResponse, error) {
	model := request.ConfigNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateConfigNode(ctx, request)
}

//RESTDeleteConfigNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteConfigNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteConfigNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetConfigNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetConfigNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetConfigNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListConfigNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListConfigNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListConfigNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
