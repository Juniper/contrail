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

//RESTCreateContrailControlNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailControlNode(c echo.Context) error {
	requestData := &models.CreateContrailControlNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_control_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailControlNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailControlNode handle a Create API
func (service *ContrailService) CreateContrailControlNode(
	ctx context.Context,
	request *models.CreateContrailControlNodeRequest) (*models.CreateContrailControlNodeResponse, error) {
	model := request.ContrailControlNode
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

	return service.Next().CreateContrailControlNode(ctx, request)
}

//RESTUpdateContrailControlNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailControlNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailControlNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_control_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailControlNode handles a Update request.
func (service *ContrailService) UpdateContrailControlNode(
	ctx context.Context,
	request *models.UpdateContrailControlNodeRequest) (*models.UpdateContrailControlNodeResponse, error) {
	model := request.ContrailControlNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailControlNode(ctx, request)
}

//RESTDeleteContrailControlNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailControlNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailControlNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailControlNode a REST Get request.
func (service *ContrailService) RESTGetContrailControlNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailControlNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailControlNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailControlNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailControlNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
