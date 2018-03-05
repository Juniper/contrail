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

//RESTCreateOpenstackStorageNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateOpenstackStorageNode(c echo.Context) error {
	requestData := &models.CreateOpenstackStorageNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_storage_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateOpenstackStorageNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateOpenstackStorageNode handle a Create API
// nolint
func (service *ContrailService) CreateOpenstackStorageNode(
	ctx context.Context,
	request *models.CreateOpenstackStorageNodeRequest) (*models.CreateOpenstackStorageNodeResponse, error) {
	model := request.OpenstackStorageNode
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

	return service.Next().CreateOpenstackStorageNode(ctx, request)
}

//RESTUpdateOpenstackStorageNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateOpenstackStorageNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateOpenstackStorageNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_storage_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateOpenstackStorageNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackStorageNode handles a Update request.
// nolint
func (service *ContrailService) UpdateOpenstackStorageNode(
	ctx context.Context,
	request *models.UpdateOpenstackStorageNodeRequest) (*models.UpdateOpenstackStorageNodeResponse, error) {
	model := request.OpenstackStorageNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateOpenstackStorageNode(ctx, request)
}

//RESTDeleteOpenstackStorageNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteOpenstackStorageNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteOpenstackStorageNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteOpenstackStorageNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetOpenstackStorageNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetOpenstackStorageNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetOpenstackStorageNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetOpenstackStorageNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListOpenstackStorageNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListOpenstackStorageNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListOpenstackStorageNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListOpenstackStorageNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
