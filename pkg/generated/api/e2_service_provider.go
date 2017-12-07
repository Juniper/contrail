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

//E2ServiceProviderRESTAPI
type E2ServiceProviderRESTAPI struct {
	DB *sql.DB
}

type E2ServiceProviderCreateRequest struct {
	Data *models.E2ServiceProvider `json:"e2-service-provider"`
}

//Path returns api path for collections.
func (api *E2ServiceProviderRESTAPI) Path() string {
	return "/e2-service-providers"
}

//LongPath returns api path for elements.
func (api *E2ServiceProviderRESTAPI) LongPath() string {
	return "/e2-service-provider/:id"
}

//SetDB sets db object
func (api *E2ServiceProviderRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *E2ServiceProviderRESTAPI) Create(c echo.Context) error {
	requestData := &E2ServiceProviderCreateRequest{
		Data: models.MakeE2ServiceProvider(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
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
			return db.CreateE2ServiceProvider(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *E2ServiceProviderRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *E2ServiceProviderRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteE2ServiceProvider(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *E2ServiceProviderRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.E2ServiceProvider
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowE2ServiceProvider(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"e2_service_provider": result,
	})
}

//List handles a List REST API Request.
func (api *E2ServiceProviderRESTAPI) List(c echo.Context) error {
	var result []*models.E2ServiceProvider
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListE2ServiceProvider(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"e2-service-providers": result,
	})
}
