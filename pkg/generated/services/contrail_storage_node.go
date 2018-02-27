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

//RESTCreateContrailStorageNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailStorageNode(c echo.Context) error {
	requestData := &models.CreateContrailStorageNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_storage_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailStorageNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailStorageNode handle a Create API
func (service *ContrailService) CreateContrailStorageNode(
	ctx context.Context,
	request *models.CreateContrailStorageNodeRequest) (*models.CreateContrailStorageNodeResponse, error) {
	model := request.ContrailStorageNode
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

	return service.Next().CreateContrailStorageNode(ctx, request)
}

//RESTUpdateContrailStorageNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailStorageNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailStorageNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_storage_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailStorageNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailStorageNode handles a Update request.
func (service *ContrailService) UpdateContrailStorageNode(
	ctx context.Context,
	request *models.UpdateContrailStorageNodeRequest) (*models.UpdateContrailStorageNodeResponse, error) {
	model := request.ContrailStorageNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailStorageNode(ctx, request)
}

//RESTDeleteContrailStorageNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailStorageNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailStorageNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailStorageNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailStorageNode a REST Get request.
func (service *ContrailService) RESTGetContrailStorageNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailStorageNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailStorageNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailStorageNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailStorageNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailStorageNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailStorageNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
