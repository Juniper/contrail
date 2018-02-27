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

//RESTCreateContrailConfigDatabaseNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailConfigDatabaseNode(c echo.Context) error {
	requestData := &models.CreateContrailConfigDatabaseNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_config_database_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailConfigDatabaseNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailConfigDatabaseNode handle a Create API
func (service *ContrailService) CreateContrailConfigDatabaseNode(
	ctx context.Context,
	request *models.CreateContrailConfigDatabaseNodeRequest) (*models.CreateContrailConfigDatabaseNodeResponse, error) {
	model := request.ContrailConfigDatabaseNode
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

	return service.Next().CreateContrailConfigDatabaseNode(ctx, request)
}

//RESTUpdateContrailConfigDatabaseNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailConfigDatabaseNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailConfigDatabaseNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_config_database_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailConfigDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailConfigDatabaseNode handles a Update request.
func (service *ContrailService) UpdateContrailConfigDatabaseNode(
	ctx context.Context,
	request *models.UpdateContrailConfigDatabaseNodeRequest) (*models.UpdateContrailConfigDatabaseNodeResponse, error) {
	model := request.ContrailConfigDatabaseNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailConfigDatabaseNode(ctx, request)
}

//RESTDeleteContrailConfigDatabaseNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailConfigDatabaseNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailConfigDatabaseNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailConfigDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailConfigDatabaseNode a REST Get request.
func (service *ContrailService) RESTGetContrailConfigDatabaseNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailConfigDatabaseNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailConfigDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailConfigDatabaseNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailConfigDatabaseNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailConfigDatabaseNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailConfigDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
