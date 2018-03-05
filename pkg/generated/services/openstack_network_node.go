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

//RESTCreateOpenstackNetworkNode handle a Create REST service.
func (service *ContrailService) RESTCreateOpenstackNetworkNode(c echo.Context) error {
	requestData := &models.CreateOpenstackNetworkNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_network_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateOpenstackNetworkNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateOpenstackNetworkNode handle a Create API
func (service *ContrailService) CreateOpenstackNetworkNode(
	ctx context.Context,
	request *models.CreateOpenstackNetworkNodeRequest) (*models.CreateOpenstackNetworkNodeResponse, error) {
	model := request.OpenstackNetworkNode
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

	return service.Next().CreateOpenstackNetworkNode(ctx, request)
}

//RESTUpdateOpenstackNetworkNode handles a REST Update request.
func (service *ContrailService) RESTUpdateOpenstackNetworkNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateOpenstackNetworkNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_network_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateOpenstackNetworkNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackNetworkNode handles a Update request.
func (service *ContrailService) UpdateOpenstackNetworkNode(
	ctx context.Context,
	request *models.UpdateOpenstackNetworkNodeRequest) (*models.UpdateOpenstackNetworkNodeResponse, error) {
	model := request.OpenstackNetworkNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateOpenstackNetworkNode(ctx, request)
}

//RESTDeleteOpenstackNetworkNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteOpenstackNetworkNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteOpenstackNetworkNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteOpenstackNetworkNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetOpenstackNetworkNode a REST Get request.
func (service *ContrailService) RESTGetOpenstackNetworkNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetOpenstackNetworkNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetOpenstackNetworkNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListOpenstackNetworkNode handles a List REST service Request.
func (service *ContrailService) RESTListOpenstackNetworkNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListOpenstackNetworkNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListOpenstackNetworkNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
