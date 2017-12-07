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

//GlobalSystemConfigRESTAPI
type GlobalSystemConfigRESTAPI struct {
	DB *sql.DB
}

type GlobalSystemConfigCreateRequest struct {
	Data *models.GlobalSystemConfig `json:"global-system-config"`
}

//Path returns api path for collections.
func (api *GlobalSystemConfigRESTAPI) Path() string {
	return "/global-system-configs"
}

//LongPath returns api path for elements.
func (api *GlobalSystemConfigRESTAPI) LongPath() string {
	return "/global-system-config/:id"
}

//SetDB sets db object
func (api *GlobalSystemConfigRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *GlobalSystemConfigRESTAPI) Create(c echo.Context) error {
	requestData := &GlobalSystemConfigCreateRequest{
		Data: models.MakeGlobalSystemConfig(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_system_config",
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
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.CreateGlobalSystemConfig(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_system_config",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *GlobalSystemConfigRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *GlobalSystemConfigRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteGlobalSystemConfig(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *GlobalSystemConfigRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.GlobalSystemConfig
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowGlobalSystemConfig(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"global_system_config": result,
	})
}

//List handles a List REST API Request.
func (api *GlobalSystemConfigRESTAPI) List(c echo.Context) error {
	var result []*models.GlobalSystemConfig
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListGlobalSystemConfig(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"global-system-configs": result,
	})
}
