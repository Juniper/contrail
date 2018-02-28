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

//RESTCreateContrailControllerNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateContrailControllerNode(c echo.Context) error {
	requestData := &models.CreateContrailControllerNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailControllerNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailControllerNode handle a Create API
// nolint
func (service *ContrailService) CreateContrailControllerNode(
	ctx context.Context,
	request *models.CreateContrailControllerNodeRequest) (*models.CreateContrailControllerNodeResponse, error) {
	model := request.ContrailControllerNode
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

	return service.Next().CreateContrailControllerNode(ctx, request)
}

//RESTUpdateContrailControllerNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateContrailControllerNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailControllerNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailControllerNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailControllerNode handles a Update request.
// nolint
func (service *ContrailService) UpdateContrailControllerNode(
	ctx context.Context,
	request *models.UpdateContrailControllerNodeRequest) (*models.UpdateContrailControllerNodeResponse, error) {
	model := request.ContrailControllerNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailControllerNode(ctx, request)
}

//RESTDeleteContrailControllerNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteContrailControllerNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailControllerNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailControllerNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailControllerNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetContrailControllerNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailControllerNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailControllerNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailControllerNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListContrailControllerNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListContrailControllerNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailControllerNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
