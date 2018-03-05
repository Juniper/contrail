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

//RESTCreateOpenstackComputeNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateOpenstackComputeNode(c echo.Context) error {
	requestData := &models.CreateOpenstackComputeNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_compute_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateOpenstackComputeNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateOpenstackComputeNode handle a Create API
// nolint
func (service *ContrailService) CreateOpenstackComputeNode(
	ctx context.Context,
	request *models.CreateOpenstackComputeNodeRequest) (*models.CreateOpenstackComputeNodeResponse, error) {
	model := request.OpenstackComputeNode
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

	return service.Next().CreateOpenstackComputeNode(ctx, request)
}

//RESTUpdateOpenstackComputeNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateOpenstackComputeNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateOpenstackComputeNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_compute_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateOpenstackComputeNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackComputeNode handles a Update request.
// nolint
func (service *ContrailService) UpdateOpenstackComputeNode(
	ctx context.Context,
	request *models.UpdateOpenstackComputeNodeRequest) (*models.UpdateOpenstackComputeNodeResponse, error) {
	model := request.OpenstackComputeNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateOpenstackComputeNode(ctx, request)
}

//RESTDeleteOpenstackComputeNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteOpenstackComputeNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteOpenstackComputeNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteOpenstackComputeNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetOpenstackComputeNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetOpenstackComputeNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetOpenstackComputeNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetOpenstackComputeNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListOpenstackComputeNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListOpenstackComputeNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListOpenstackComputeNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListOpenstackComputeNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
