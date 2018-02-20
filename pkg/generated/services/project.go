package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//RESTProjectUpdateRequest for update request for REST.
type RESTProjectUpdateRequest struct {
	Data map[string]interface{} `json:"project"`
}

//RESTCreateProject handle a Create REST service.
func (service *ContrailService) RESTCreateProject(c echo.Context) error {
	requestData := &models.CreateProjectRequest{
		Project: models.MakeProject(),
	}
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
	if model.FQName == nil {
		return nil, common.ErrorBadRequest("Missing fq_name")
	}

	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateProject(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "project",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateProjectResponse{
		Project: request.Project,
	}, nil
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
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateProject(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "project",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateProjectResponse{
		Project: model,
	}, nil
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

//DeleteProject delete a resource.
func (service *ContrailService) DeleteProject(ctx context.Context, request *models.DeleteProjectRequest) (*models.DeleteProjectResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteProject(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteProjectResponse{
		ID: request.ID,
	}, nil
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

//GetProject a Get request.
func (service *ContrailService) GetProject(ctx context.Context, request *models.GetProjectRequest) (response *models.GetProjectResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListProjectRequest{
		Spec: spec,
	}
	var result *models.ListProjectResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListProject(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Projects) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetProjectResponse{
		Project: result.Projects[0],
	}
	return response, nil
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

//ListProject handles a List service Request.
func (service *ContrailService) ListProject(
	ctx context.Context,
	request *models.ListProjectRequest) (response *models.ListProjectResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListProject(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
