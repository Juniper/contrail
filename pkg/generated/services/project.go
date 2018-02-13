package services 

import (
    "context"
    "net/http"
    "database/sql"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/Juniper/contrail/pkg/generated/db"
    "github.com/satori/go.uuid"
    "github.com/labstack/echo"
    "github.com/Juniper/contrail/pkg/common"

	log "github.com/sirupsen/logrus"
)

//RESTProjectUpdateRequest for update request for REST.
type RESTProjectUpdateRequest struct {
    Data map[string]interface{} `json:"project"`
}

//RESTCreateProject handle a Create REST service.
func (service *ContrailService) RESTCreateProject(c echo.Context) error {
    requestData := &models.ProjectCreateRequest{
        Project: models.MakeProject(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
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
    request *models.ProjectCreateRequest) (*models.ProjectCreateResponse, error) {
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
    model.Perms2.Owner = auth.ProjectID()
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.CreateProject(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "project",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ProjectCreateResponse{
        Project: request.Project,
    }, nil
}

//RESTUpdateProject handles a REST Update request.
func (service *ContrailService) RESTUpdateProject(c echo.Context) error {
    id := c.Param("id")
    request := &models.ProjectUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "project",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateProject(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateProject handles a Update request.
func (service *ContrailService) UpdateProject(ctx context.Context, request *models.ProjectUpdateRequest) (*models.ProjectUpdateResponse, error) {
    id = request.ID
    model = request.Project
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    auth := common.GetAuthCTX(ctx)
    ok := common.SetValueByPath(model, "Perms2.Owner", ".", auth.ProjectID())
    if !ok {
        return nil, common.ErrorBadRequest("Invalid JSON format")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateProject(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "project",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.Project.UpdateResponse{
        Project: model,
    }, nil
}

//RESTDeleteProject delete a resource using REST service.
func (service *ContrailService) RESTDeleteProject(c echo.Context) error {
    id := c.Param("id")
    request := &models.ProjectDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteProject(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteProject delete a resource.
func (service *ContrailService) DeleteProject(ctx context.Context, request *models.ProjectDeleteRequest) (*models.ProjectDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteProject(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ProjectDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowProject a REST Show request.
func (service *ContrailService) RESTShowProject(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.Project
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListProject(tx, &common.ListSpec{
                Limit: 1,
                Auth: auth,
                Filter: common.Filter{
                    "uuid": []string{id},
                },
            })
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "project": result,
    })
}

//RESTListProject handles a List REST service Request.
func (service *ContrailService) RESTListProject(c echo.Context) (error) {
    var result []*models.Project
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListProject(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "projects": result,
    })
}