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

//ControllerNodeRoleRESTAPI
type ControllerNodeRoleRESTAPI struct {
	DB *sql.DB
}

type ControllerNodeRoleCreateRequest struct {
	Data *models.ControllerNodeRole `json:"controller-node-role"`
}

//Path returns api path for collections.
func (api *ControllerNodeRoleRESTAPI) Path() string {
	return "/controller-node-roles"
}

//LongPath returns api path for elements.
func (api *ControllerNodeRoleRESTAPI) LongPath() string {
	return "/controller-node-role/:id"
}

//SetDB sets db object
func (api *ControllerNodeRoleRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *ControllerNodeRoleRESTAPI) Create(c echo.Context) error {
	requestData := &ControllerNodeRoleCreateRequest{
		Data: models.MakeControllerNodeRole(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "controller_node_role",
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
			return db.CreateControllerNodeRole(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "controller_node_role",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *ControllerNodeRoleRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *ControllerNodeRoleRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteControllerNodeRole(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *ControllerNodeRoleRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.ControllerNodeRole
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowControllerNodeRole(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"controller_node_role": result,
	})
}

//List handles a List REST API Request.
func (api *ControllerNodeRoleRESTAPI) List(c echo.Context) error {
	var result []*models.ControllerNodeRole
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListControllerNodeRole(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"controller-node-roles": result,
	})
}
