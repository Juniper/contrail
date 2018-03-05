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

//RESTCreateOpenstackControlNode handle a Create REST service.
func (service *ContrailService) RESTCreateOpenstackControlNode(c echo.Context) error {
	requestData := &models.CreateOpenstackControlNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_control_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateOpenstackControlNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateOpenstackControlNode handle a Create API
func (service *ContrailService) CreateOpenstackControlNode(
	ctx context.Context,
	request *models.CreateOpenstackControlNodeRequest) (*models.CreateOpenstackControlNodeResponse, error) {
	model := request.OpenstackControlNode
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

	return service.Next().CreateOpenstackControlNode(ctx, request)
}

//RESTUpdateOpenstackControlNode handles a REST Update request.
func (service *ContrailService) RESTUpdateOpenstackControlNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateOpenstackControlNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_control_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateOpenstackControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackControlNode handles a Update request.
func (service *ContrailService) UpdateOpenstackControlNode(
	ctx context.Context,
	request *models.UpdateOpenstackControlNodeRequest) (*models.UpdateOpenstackControlNodeResponse, error) {
	model := request.OpenstackControlNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateOpenstackControlNode(ctx, request)
}

//RESTDeleteOpenstackControlNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteOpenstackControlNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteOpenstackControlNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteOpenstackControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetOpenstackControlNode a REST Get request.
func (service *ContrailService) RESTGetOpenstackControlNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetOpenstackControlNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetOpenstackControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListOpenstackControlNode handles a List REST service Request.
func (service *ContrailService) RESTListOpenstackControlNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListOpenstackControlNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListOpenstackControlNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
