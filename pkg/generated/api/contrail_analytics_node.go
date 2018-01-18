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

//ContrailAnalyticsNodeRESTAPI
type ContrailAnalyticsNodeRESTAPI struct {
	DB *sql.DB
}

type ContrailAnalyticsNodeCreateRequest struct {
	Data *models.ContrailAnalyticsNode `json:"contrail-analytics-node"`
}

type ContrailAnalyticsNodeUpdateRequest struct {
	Data map[string]interface{} `json:"contrail-analytics-node"`
}

//Path returns api path for collections.
func (api *ContrailAnalyticsNodeRESTAPI) Path() string {
	return "/contrail-analytics-nodes"
}

//LongPath returns api path for elements.
func (api *ContrailAnalyticsNodeRESTAPI) LongPath() string {
	return "/contrail-analytics-node/:id"
}

//SetDB sets db object
func (api *ContrailAnalyticsNodeRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *ContrailAnalyticsNodeRESTAPI) Create(c echo.Context) error {
	requestData := &ContrailAnalyticsNodeCreateRequest{
		Data: models.MakeContrailAnalyticsNode(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
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
			return db.CreateContrailAnalyticsNode(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *ContrailAnalyticsNodeRESTAPI) Update(c echo.Context) error {
	id := c.Param("id")
	requestData := &ContrailAnalyticsNodeUpdateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
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
			return db.UpdateContrailAnalyticsNode(tx, id, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("db update failed")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]map[string]string{
		"contrail-analytics-node": {
			"uuid": id,
			"uri":  "/" + "contrail-analytics-node" + "/" + id},
	})
}

//Delete handles a REST Delete request.
func (api *ContrailAnalyticsNodeRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteContrailAnalyticsNode(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *ContrailAnalyticsNodeRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.ContrailAnalyticsNode
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailAnalyticsNode(tx, &common.ListSpec{
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
		"contrail_analytics_node": result,
	})
}

//List handles a List REST API Request.
func (api *ContrailAnalyticsNodeRESTAPI) List(c echo.Context) error {
	var result []*models.ContrailAnalyticsNode
	var err error
	auth := common.GetAuthContext(c)
	listSpec := common.GetListSpec(c)
	listSpec.Auth = auth
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailAnalyticsNode(tx, listSpec)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"contrail-analytics-nodes": result,
	})
}
