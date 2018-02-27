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

//RESTCreateAnalyticsNode handle a Create REST service.
func (service *ContrailService) RESTCreateAnalyticsNode(c echo.Context) error {
	requestData := &models.CreateAnalyticsNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "analytics_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAnalyticsNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAnalyticsNode handle a Create API
func (service *ContrailService) CreateAnalyticsNode(
	ctx context.Context,
	request *models.CreateAnalyticsNodeRequest) (*models.CreateAnalyticsNodeResponse, error) {
	model := request.AnalyticsNode
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

	return service.Next().CreateAnalyticsNode(ctx, request)
}

//RESTUpdateAnalyticsNode handles a REST Update request.
func (service *ContrailService) RESTUpdateAnalyticsNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAnalyticsNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "analytics_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAnalyticsNode handles a Update request.
func (service *ContrailService) UpdateAnalyticsNode(
	ctx context.Context,
	request *models.UpdateAnalyticsNodeRequest) (*models.UpdateAnalyticsNodeResponse, error) {
	model := request.AnalyticsNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateAnalyticsNode(ctx, request)
}

//RESTDeleteAnalyticsNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteAnalyticsNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAnalyticsNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetAnalyticsNode a REST Get request.
func (service *ContrailService) RESTGetAnalyticsNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAnalyticsNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListAnalyticsNode handles a List REST service Request.
func (service *ContrailService) RESTListAnalyticsNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAnalyticsNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
