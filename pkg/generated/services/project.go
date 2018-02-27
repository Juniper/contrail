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

//RESTCreateProject handle a Create REST service.
func (service *ContrailService) RESTCreateProject(c echo.Context) error {
	requestData := &models.CreateProjectRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "project",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateProject(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateProject handle a Create API
func (service *ContrailService) CreateProject(
	ctx context.Context,
	request *models.CreateProjectRequest) (*models.CreateProjectResponse, error) {
	model := request.Project
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

	return service.Next().CreateProject(ctx, request)
}

//RESTUpdateProject handles a REST Update request.
func (service *ContrailService) RESTUpdateProject(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateProjectRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "project",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateProject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateProject handles a Update request.
func (service *ContrailService) UpdateProject(
	ctx context.Context,
	request *models.UpdateProjectRequest) (*models.UpdateProjectResponse, error) {
	model := request.Project
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateProject(ctx, request)
}

//RESTDeleteProject delete a resource using REST service.
func (service *ContrailService) RESTDeleteProject(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteProjectRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteProject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetProject a REST Get request.
func (service *ContrailService) RESTGetProject(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetProjectRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetProject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListProject handles a List REST service Request.
func (service *ContrailService) RESTListProject(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListProjectRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListProject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
