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

//InterfaceRouteTableRESTAPI
type InterfaceRouteTableRESTAPI struct {
	DB *sql.DB
}

type InterfaceRouteTableCreateRequest struct {
	Data *models.InterfaceRouteTable `json:"interface-route-table"`
}

//Path returns api path for collections.
func (api *InterfaceRouteTableRESTAPI) Path() string {
	return "/interface-route-tables"
}

//LongPath returns api path for elements.
func (api *InterfaceRouteTableRESTAPI) LongPath() string {
	return "/interface-route-table/:id"
}

//SetDB sets db object
func (api *InterfaceRouteTableRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *InterfaceRouteTableRESTAPI) Create(c echo.Context) error {
	requestData := &InterfaceRouteTableCreateRequest{
		Data: models.MakeInterfaceRouteTable(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
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
			return db.CreateInterfaceRouteTable(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *InterfaceRouteTableRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *InterfaceRouteTableRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteInterfaceRouteTable(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *InterfaceRouteTableRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.InterfaceRouteTable
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowInterfaceRouteTable(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"interface_route_table": result,
	})
}

//List handles a List REST API Request.
func (api *InterfaceRouteTableRESTAPI) List(c echo.Context) error {
	var result []*models.InterfaceRouteTable
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListInterfaceRouteTable(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"interface-route-tables": result,
	})
}
