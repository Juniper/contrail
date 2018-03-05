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

//RESTCreateContrailWebuiNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailWebuiNode(c echo.Context) error {
	requestData := &models.CreateContrailWebuiNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_webui_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailWebuiNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailWebuiNode handle a Create API
func (service *ContrailService) CreateContrailWebuiNode(
	ctx context.Context,
	request *models.CreateContrailWebuiNodeRequest) (*models.CreateContrailWebuiNodeResponse, error) {
	model := request.ContrailWebuiNode
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

	return service.Next().CreateContrailWebuiNode(ctx, request)
}

//RESTUpdateContrailWebuiNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailWebuiNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailWebuiNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_webui_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailWebuiNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailWebuiNode handles a Update request.
func (service *ContrailService) UpdateContrailWebuiNode(
	ctx context.Context,
	request *models.UpdateContrailWebuiNodeRequest) (*models.UpdateContrailWebuiNodeResponse, error) {
	model := request.ContrailWebuiNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailWebuiNode(ctx, request)
}

//RESTDeleteContrailWebuiNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailWebuiNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailWebuiNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailWebuiNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailWebuiNode a REST Get request.
func (service *ContrailService) RESTGetContrailWebuiNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailWebuiNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailWebuiNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailWebuiNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailWebuiNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailWebuiNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailWebuiNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
