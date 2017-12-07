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

//ServiceConnectionModuleRESTAPI
type ServiceConnectionModuleRESTAPI struct {
	DB *sql.DB
}

type ServiceConnectionModuleCreateRequest struct {
	Data *models.ServiceConnectionModule `json:"service-connection-module"`
}

//Path returns api path for collections.
func (api *ServiceConnectionModuleRESTAPI) Path() string {
	return "/service-connection-modules"
}

//LongPath returns api path for elements.
func (api *ServiceConnectionModuleRESTAPI) LongPath() string {
	return "/service-connection-module/:id"
}

//SetDB sets db object
func (api *ServiceConnectionModuleRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *ServiceConnectionModuleRESTAPI) Create(c echo.Context) error {
	requestData := &ServiceConnectionModuleCreateRequest{
		Data: models.MakeServiceConnectionModule(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
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
			return db.CreateServiceConnectionModule(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_connection_module",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *ServiceConnectionModuleRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *ServiceConnectionModuleRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServiceConnectionModule(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *ServiceConnectionModuleRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.ServiceConnectionModule
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowServiceConnectionModule(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"service_connection_module": result,
	})
}

//List handles a List REST API Request.
func (api *ServiceConnectionModuleRESTAPI) List(c echo.Context) error {
	var result []*models.ServiceConnectionModule
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceConnectionModule(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"service-connection-modules": result,
	})
}
