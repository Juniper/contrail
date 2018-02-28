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

//RESTCreateDatabaseNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateDatabaseNode(c echo.Context) error {
	requestData := &models.CreateDatabaseNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "database_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateDatabaseNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateDatabaseNode handle a Create API
// nolint
func (service *ContrailService) CreateDatabaseNode(
	ctx context.Context,
	request *models.CreateDatabaseNodeRequest) (*models.CreateDatabaseNodeResponse, error) {
	model := request.DatabaseNode
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

	return service.Next().CreateDatabaseNode(ctx, request)
}

//RESTUpdateDatabaseNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateDatabaseNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateDatabaseNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "database_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateDatabaseNode handles a Update request.
// nolint
func (service *ContrailService) UpdateDatabaseNode(
	ctx context.Context,
	request *models.UpdateDatabaseNodeRequest) (*models.UpdateDatabaseNodeResponse, error) {
	model := request.DatabaseNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateDatabaseNode(ctx, request)
}

//RESTDeleteDatabaseNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteDatabaseNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteDatabaseNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetDatabaseNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetDatabaseNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetDatabaseNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListDatabaseNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListDatabaseNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListDatabaseNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListDatabaseNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
