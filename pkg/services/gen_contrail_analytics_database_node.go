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

//RESTCreateContrailAnalyticsDatabaseNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateContrailAnalyticsDatabaseNode(c echo.Context) error {
	requestData := &models.CreateContrailAnalyticsDatabaseNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_database_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailAnalyticsDatabaseNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailAnalyticsDatabaseNode handle a Create API
// nolint
func (service *ContrailService) CreateContrailAnalyticsDatabaseNode(
	ctx context.Context,
	request *models.CreateContrailAnalyticsDatabaseNodeRequest) (*models.CreateContrailAnalyticsDatabaseNodeResponse, error) {
	model := request.ContrailAnalyticsDatabaseNode
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

	return service.Next().CreateContrailAnalyticsDatabaseNode(ctx, request)
}

//RESTUpdateContrailAnalyticsDatabaseNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateContrailAnalyticsDatabaseNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailAnalyticsDatabaseNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_database_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailAnalyticsDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailAnalyticsDatabaseNode handles a Update request.
// nolint
func (service *ContrailService) UpdateContrailAnalyticsDatabaseNode(
	ctx context.Context,
	request *models.UpdateContrailAnalyticsDatabaseNodeRequest) (*models.UpdateContrailAnalyticsDatabaseNodeResponse, error) {
	model := request.ContrailAnalyticsDatabaseNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailAnalyticsDatabaseNode(ctx, request)
}

//RESTDeleteContrailAnalyticsDatabaseNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteContrailAnalyticsDatabaseNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailAnalyticsDatabaseNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailAnalyticsDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailAnalyticsDatabaseNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetContrailAnalyticsDatabaseNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailAnalyticsDatabaseNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailAnalyticsDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailAnalyticsDatabaseNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListContrailAnalyticsDatabaseNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListContrailAnalyticsDatabaseNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailAnalyticsDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
