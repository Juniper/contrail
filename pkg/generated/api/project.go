package api

import (
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//ProjectRESTAPI
type ProjectRESTAPI struct {
	DB *sql.DB
}

type ProjectCreateRequest struct {
	Data *models.Project `json:"project"`
}

type ProjectUpdateRequest struct {
	Data map[string]interface{} `json:"project"`
}

//Path returns api path for collections.
func (api *ProjectRESTAPI) Path() string {
	return "/projects"
}

//LongPath returns api path for elements.
func (api *ProjectRESTAPI) LongPath() string {
	return "/project/:id"
}

//SetDB sets db object
func (api *ProjectRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *ProjectRESTAPI) Create(c echo.Context) error {
	requestData := &ProjectCreateRequest{
		Data: models.MakeProject(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "project",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	model := requestData.Data
	if model == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}

	if model.FQName == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing FQName")
	}

	auth := common.GetAuthContext(c)
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.CreateProject(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "project",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *ProjectRESTAPI) Update(c echo.Context) error {
	id := c.Param("id")
	requestData := &ProjectUpdateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "project",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	model := requestData.Data
	if model == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	auth := common.GetAuthContext(c)
	ok := common.SetValueByPath(model, "Perms2.Owner", ".", auth.ProjectID())
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.UpdateProject(tx, id, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "project",
		}).Debug("db update failed")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]map[string]string{
		"project": {
			"uuid": id,
			"uri":  "/" + "project" + "/" + id},
	})
}

//Delete handles a REST Delete request.
func (api *ProjectRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteProject(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *ProjectRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.Project
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListProject(tx, &common.ListSpec{
				Limit: 1000,
				Auth:  auth,
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

//List handles a List REST API Request.
func (api *ProjectRESTAPI) List(c echo.Context) error {
	var result []*models.Project
	var err error
	auth := common.GetAuthContext(c)
	listSpec := common.GetListSpec(c)
	listSpec.Auth = auth
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListProject(tx, listSpec)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"projects": result,
	})
}
