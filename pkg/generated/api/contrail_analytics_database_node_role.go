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

//ContrailAnalyticsDatabaseNodeRoleRESTAPI
type ContrailAnalyticsDatabaseNodeRoleRESTAPI struct {
	DB *sql.DB
}

type ContrailAnalyticsDatabaseNodeRoleCreateRequest struct {
	Data *models.ContrailAnalyticsDatabaseNodeRole `json:"contrail-analytics-database-node-role"`
}

//Path returns api path for collections.
func (api *ContrailAnalyticsDatabaseNodeRoleRESTAPI) Path() string {
	return "/contrail-analytics-database-node-roles"
}

//LongPath returns api path for elements.
func (api *ContrailAnalyticsDatabaseNodeRoleRESTAPI) LongPath() string {
	return "/contrail-analytics-database-node-role/:id"
}

//SetDB sets db object
func (api *ContrailAnalyticsDatabaseNodeRoleRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *ContrailAnalyticsDatabaseNodeRoleRESTAPI) Create(c echo.Context) error {
	requestData := &ContrailAnalyticsDatabaseNodeRoleCreateRequest{
		Data: models.MakeContrailAnalyticsDatabaseNodeRole(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_database_node_role",
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
			return db.CreateContrailAnalyticsDatabaseNodeRole(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_database_node_role",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *ContrailAnalyticsDatabaseNodeRoleRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *ContrailAnalyticsDatabaseNodeRoleRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteContrailAnalyticsDatabaseNodeRole(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *ContrailAnalyticsDatabaseNodeRoleRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.ContrailAnalyticsDatabaseNodeRole
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailAnalyticsDatabaseNodeRole(tx, &common.ListSpec{
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
		"contrail_analytics_database_node_role": result,
	})
}

//List handles a List REST API Request.
func (api *ContrailAnalyticsDatabaseNodeRoleRESTAPI) List(c echo.Context) error {
	var result []*models.ContrailAnalyticsDatabaseNodeRole
	var err error
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailAnalyticsDatabaseNodeRole(tx, &common.ListSpec{
				Limit: 1000,
				Auth:  auth,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"contrail-analytics-database-node-roles": result,
	})
}
