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

//RoutingPolicyRESTAPI
type RoutingPolicyRESTAPI struct {
	DB *sql.DB
}

type RoutingPolicyCreateRequest struct {
	Data *models.RoutingPolicy `json:"routing-policy"`
}

//Path returns api path for collections.
func (api *RoutingPolicyRESTAPI) Path() string {
	return "/routing-policys"
}

//LongPath returns api path for elements.
func (api *RoutingPolicyRESTAPI) LongPath() string {
	return "/routing-policy/:id"
}

//SetDB sets db object
func (api *RoutingPolicyRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *RoutingPolicyRESTAPI) Create(c echo.Context) error {
	requestData := &RoutingPolicyCreateRequest{
		Data: models.MakeRoutingPolicy(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
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
			return db.CreateRoutingPolicy(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *RoutingPolicyRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *RoutingPolicyRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteRoutingPolicy(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *RoutingPolicyRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.RoutingPolicy
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowRoutingPolicy(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"routing_policy": result,
	})
}

//List handles a List REST API Request.
func (api *RoutingPolicyRESTAPI) List(c echo.Context) error {
	var result []*models.RoutingPolicy
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListRoutingPolicy(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"routing-policys": result,
	})
}
