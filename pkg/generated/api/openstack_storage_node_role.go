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

//OpenstackStorageNodeRoleRESTAPI
type OpenstackStorageNodeRoleRESTAPI struct {
	DB *sql.DB
}

type OpenstackStorageNodeRoleCreateRequest struct {
	Data *models.OpenstackStorageNodeRole `json:"openstack-storage-node-role"`
}

type OpenstackStorageNodeRoleUpdateRequest struct {
	Data map[string]interface{} `json:"openstack-storage-node-role"`
}

//Path returns api path for collections.
func (api *OpenstackStorageNodeRoleRESTAPI) Path() string {
	return "/openstack-storage-node-roles"
}

//LongPath returns api path for elements.
func (api *OpenstackStorageNodeRoleRESTAPI) LongPath() string {
	return "/openstack-storage-node-role/:id"
}

//SetDB sets db object
func (api *OpenstackStorageNodeRoleRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *OpenstackStorageNodeRoleRESTAPI) Create(c echo.Context) error {
	requestData := &OpenstackStorageNodeRoleCreateRequest{
		Data: models.MakeOpenstackStorageNodeRole(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_storage_node_role",
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
			return db.CreateOpenstackStorageNodeRole(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_storage_node_role",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *OpenstackStorageNodeRoleRESTAPI) Update(c echo.Context) error {
	id := c.Param("id")
	requestData := &OpenstackStorageNodeRoleUpdateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_storage_node_role",
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
			return db.UpdateOpenstackStorageNodeRole(tx, id, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_storage_node_role",
		}).Debug("db update failed")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]map[string]string{
		"openstack-storage-node-role": {
			"uuid": id,
			"uri":  "/" + "openstack-storage-node-role" + "/" + id},
	})
}

//Delete handles a REST Delete request.
func (api *OpenstackStorageNodeRoleRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteOpenstackStorageNodeRole(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *OpenstackStorageNodeRoleRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.OpenstackStorageNodeRole
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListOpenstackStorageNodeRole(tx, &common.ListSpec{
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
		"openstack_storage_node_role": result,
	})
}

//List handles a List REST API Request.
func (api *OpenstackStorageNodeRoleRESTAPI) List(c echo.Context) error {
	var result []*models.OpenstackStorageNodeRole
	var err error
	auth := common.GetAuthContext(c)
	listSpec := common.GetListSpec(c)
	listSpec.Auth = auth
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListOpenstackStorageNodeRole(tx, listSpec)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"openstack-storage-node-roles": result,
	})
}
