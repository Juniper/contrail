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

//RESTCreateNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateNode(c echo.Context) error {
	requestData := &models.CreateNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateNode handle a Create API
// nolint
func (service *ContrailService) CreateNode(
	ctx context.Context,
	request *models.CreateNodeRequest) (*models.CreateNodeResponse, error) {
	model := request.Node
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

	return service.Next().CreateNode(ctx, request)
}

//RESTUpdateNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateNode handles a Update request.
// nolint
func (service *ContrailService) UpdateNode(
	ctx context.Context,
	request *models.UpdateNodeRequest) (*models.UpdateNodeResponse, error) {
	model := request.Node
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateNode(ctx, request)
}

//RESTDeleteNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
