package api

import (
	"database/sql"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//ContrailControllerNodeRoleRESTAPI
type ContrailControllerNodeRoleRESTAPI struct {
	DB *sql.DB
}

type ContrailControllerNodeRoleCreateRequest struct {
	Data *models.ContrailControllerNodeRole `json:"contrail-controller-node-role"`
}

type ContrailControllerNodeRoleUpdateRequest struct {
	Data map[string]interface{} `json:"contrail-controller-node-role"`
}

//Path returns api path for collections.
func (api *ContrailControllerNodeRoleRESTAPI) Path() string {
	return "/contrail-controller-node-roles"
}

//LongPath returns api path for elements.
func (api *ContrailControllerNodeRoleRESTAPI) LongPath() string {
	return "/contrail-controller-node-role/:id"
}

//SetDB sets db object
func (api *ContrailControllerNodeRoleRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *ContrailControllerNodeRoleRESTAPI) Create(c echo.Context) error {
	requestData := &ContrailControllerNodeRoleCreateRequest{
		Data: models.MakeContrailControllerNodeRole(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node_role",
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
			return db.CreateContrailControllerNodeRole(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node_role",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *ContrailControllerNodeRoleRESTAPI) Update(c echo.Context) error {
	id := c.Param("id")
	requestData := &ContrailControllerNodeRoleUpdateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node_role",
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
			return db.UpdateContrailControllerNodeRole(tx, id, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_controller_node_role",
		}).Debug("db update failed")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]map[string]string{
		"contrail-controller-node-role": {
			"uuid": id,
			"uri":  "/" + "contrail-controller-node-role" + "/" + id},
	})
}

//Delete handles a REST Delete request.
func (api *ContrailControllerNodeRoleRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteContrailControllerNodeRole(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *ContrailControllerNodeRoleRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.ContrailControllerNodeRole
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailControllerNodeRole(tx, &common.ListSpec{
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
		"contrail_controller_node_role": result,
	})
}

//List handles a List REST API Request.
func (api *ContrailControllerNodeRoleRESTAPI) List(c echo.Context) error {
	var result []*models.ContrailControllerNodeRole
	var err error
	auth := common.GetAuthContext(c)
	listSpec := common.GetListSpec(c)
	listSpec.Auth = auth
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailControllerNodeRole(tx, listSpec)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"contrail-controller-node-roles": result,
	})
}
