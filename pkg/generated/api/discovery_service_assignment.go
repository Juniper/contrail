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

//DiscoveryServiceAssignmentRESTAPI
type DiscoveryServiceAssignmentRESTAPI struct {
	DB *sql.DB
}

type DiscoveryServiceAssignmentCreateRequest struct {
	Data *models.DiscoveryServiceAssignment `json:"discovery-service-assignment"`
}

//Path returns api path for collections.
func (api *DiscoveryServiceAssignmentRESTAPI) Path() string {
	return "/discovery-service-assignments"
}

//LongPath returns api path for elements.
func (api *DiscoveryServiceAssignmentRESTAPI) LongPath() string {
	return "/discovery-service-assignment/:id"
}

//SetDB sets db object
func (api *DiscoveryServiceAssignmentRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *DiscoveryServiceAssignmentRESTAPI) Create(c echo.Context) error {
	requestData := &DiscoveryServiceAssignmentCreateRequest{
		Data: models.MakeDiscoveryServiceAssignment(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "discovery_service_assignment",
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
			return db.CreateDiscoveryServiceAssignment(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "discovery_service_assignment",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *DiscoveryServiceAssignmentRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *DiscoveryServiceAssignmentRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteDiscoveryServiceAssignment(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *DiscoveryServiceAssignmentRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.DiscoveryServiceAssignment
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListDiscoveryServiceAssignment(tx, &common.ListSpec{
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
		"discovery_service_assignment": result,
	})
}

//List handles a List REST API Request.
func (api *DiscoveryServiceAssignmentRESTAPI) List(c echo.Context) error {
	var result []*models.DiscoveryServiceAssignment
	var err error
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListDiscoveryServiceAssignment(tx, &common.ListSpec{
				Limit: 1000,
				Auth:  auth,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"discovery-service-assignments": result,
	})
}
