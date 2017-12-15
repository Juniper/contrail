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

//LoadbalancerHealthmonitorRESTAPI
type LoadbalancerHealthmonitorRESTAPI struct {
	DB *sql.DB
}

type LoadbalancerHealthmonitorCreateRequest struct {
	Data *models.LoadbalancerHealthmonitor `json:"loadbalancer-healthmonitor"`
}

//Path returns api path for collections.
func (api *LoadbalancerHealthmonitorRESTAPI) Path() string {
	return "/loadbalancer-healthmonitors"
}

//LongPath returns api path for elements.
func (api *LoadbalancerHealthmonitorRESTAPI) LongPath() string {
	return "/loadbalancer-healthmonitor/:id"
}

//SetDB sets db object
func (api *LoadbalancerHealthmonitorRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *LoadbalancerHealthmonitorRESTAPI) Create(c echo.Context) error {
	requestData := &LoadbalancerHealthmonitorCreateRequest{
		Data: models.MakeLoadbalancerHealthmonitor(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
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
	auth := common.GetAuthContext(c)
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.CreateLoadbalancerHealthmonitor(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *LoadbalancerHealthmonitorRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *LoadbalancerHealthmonitorRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteLoadbalancerHealthmonitor(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *LoadbalancerHealthmonitorRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.LoadbalancerHealthmonitor
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListLoadbalancerHealthmonitor(tx, &common.ListSpec{
				Limit: 1000,
				Auth:  auth,
				Filter: common.Filter{
					"uuid": id,
				},
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"loadbalancer_healthmonitor": result,
	})
}

//List handles a List REST API Request.
func (api *LoadbalancerHealthmonitorRESTAPI) List(c echo.Context) error {
	var result []*models.LoadbalancerHealthmonitor
	var err error
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListLoadbalancerHealthmonitor(tx, &common.ListSpec{
				Limit: 1000,
				Auth:  auth,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"loadbalancer-healthmonitors": result,
	})
}
