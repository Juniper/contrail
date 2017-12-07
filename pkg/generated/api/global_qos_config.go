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

//GlobalQosConfigRESTAPI
type GlobalQosConfigRESTAPI struct {
	DB *sql.DB
}

type GlobalQosConfigCreateRequest struct {
	Data *models.GlobalQosConfig `json:"global-qos-config"`
}

//Path returns api path for collections.
func (api *GlobalQosConfigRESTAPI) Path() string {
	return "/global-qos-configs"
}

//LongPath returns api path for elements.
func (api *GlobalQosConfigRESTAPI) LongPath() string {
	return "/global-qos-config/:id"
}

//SetDB sets db object
func (api *GlobalQosConfigRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *GlobalQosConfigRESTAPI) Create(c echo.Context) error {
	requestData := &GlobalQosConfigCreateRequest{
		Data: models.MakeGlobalQosConfig(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_qos_config",
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
			return db.CreateGlobalQosConfig(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "global_qos_config",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *GlobalQosConfigRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *GlobalQosConfigRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteGlobalQosConfig(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *GlobalQosConfigRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.GlobalQosConfig
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowGlobalQosConfig(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"global_qos_config": result,
	})
}

//List handles a List REST API Request.
func (api *GlobalQosConfigRESTAPI) List(c echo.Context) error {
	var result []*models.GlobalQosConfig
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListGlobalQosConfig(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"global-qos-configs": result,
	})
}
