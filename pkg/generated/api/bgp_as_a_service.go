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

//BGPAsAServiceRESTAPI
type BGPAsAServiceRESTAPI struct {
	DB *sql.DB
}

type BGPAsAServiceCreateRequest struct {
	Data *models.BGPAsAService `json:"bgp-as-a-service"`
}

//Path returns api path for collections.
func (api *BGPAsAServiceRESTAPI) Path() string {
	return "/bgp-as-a-services"
}

//LongPath returns api path for elements.
func (api *BGPAsAServiceRESTAPI) LongPath() string {
	return "/bgp-as-a-service/:id"
}

//SetDB sets db object
func (api *BGPAsAServiceRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *BGPAsAServiceRESTAPI) Create(c echo.Context) error {
	requestData := &BGPAsAServiceCreateRequest{
		Data: models.MakeBGPAsAService(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
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
			return db.CreateBGPAsAService(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *BGPAsAServiceRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *BGPAsAServiceRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteBGPAsAService(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *BGPAsAServiceRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.BGPAsAService
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowBGPAsAService(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"bgp_as_a_service": result,
	})
}

//List handles a List REST API Request.
func (api *BGPAsAServiceRESTAPI) List(c echo.Context) error {
	var result []*models.BGPAsAService
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListBGPAsAService(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"bgp-as-a-services": result,
	})
}
