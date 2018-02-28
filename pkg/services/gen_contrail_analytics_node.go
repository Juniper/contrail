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

//RESTCreateContrailAnalyticsNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateContrailAnalyticsNode(c echo.Context) error {
	requestData := &models.CreateContrailAnalyticsNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailAnalyticsNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailAnalyticsNode handle a Create API
// nolint
func (service *ContrailService) CreateContrailAnalyticsNode(
	ctx context.Context,
	request *models.CreateContrailAnalyticsNodeRequest) (*models.CreateContrailAnalyticsNodeResponse, error) {
	model := request.ContrailAnalyticsNode
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

	return service.Next().CreateContrailAnalyticsNode(ctx, request)
}

//RESTUpdateContrailAnalyticsNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateContrailAnalyticsNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailAnalyticsNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailAnalyticsNode handles a Update request.
// nolint
func (service *ContrailService) UpdateContrailAnalyticsNode(
	ctx context.Context,
	request *models.UpdateContrailAnalyticsNodeRequest) (*models.UpdateContrailAnalyticsNodeResponse, error) {
	model := request.ContrailAnalyticsNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailAnalyticsNode(ctx, request)
}

//RESTDeleteContrailAnalyticsNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteContrailAnalyticsNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailAnalyticsNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailAnalyticsNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetContrailAnalyticsNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailAnalyticsNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailAnalyticsNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListContrailAnalyticsNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListContrailAnalyticsNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
