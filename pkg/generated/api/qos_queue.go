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

//QosQueueRESTAPI
type QosQueueRESTAPI struct {
	DB *sql.DB
}

type QosQueueCreateRequest struct {
	Data *models.QosQueue `json:"qos-queue"`
}

type QosQueueUpdateRequest struct {
	Data map[string]interface{} `json:"qos-queue"`
}

//Path returns api path for collections.
func (api *QosQueueRESTAPI) Path() string {
	return "/qos-queues"
}

//LongPath returns api path for elements.
func (api *QosQueueRESTAPI) LongPath() string {
	return "/qos-queue/:id"
}

//SetDB sets db object
func (api *QosQueueRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *QosQueueRESTAPI) Create(c echo.Context) error {
	requestData := &QosQueueCreateRequest{
		Data: models.MakeQosQueue(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
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
			return db.CreateQosQueue(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *QosQueueRESTAPI) Update(c echo.Context) error {
	id := c.Param("id")
	requestData := &QosQueueUpdateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
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
			return db.UpdateQosQueue(tx, id, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
		}).Debug("db update failed")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]map[string]string{
		"qos-queue": {
			"uuid": id,
			"uri":  "/" + "qos-queue" + "/" + id},
	})
}

//Delete handles a REST Delete request.
func (api *QosQueueRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteQosQueue(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *QosQueueRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.QosQueue
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListQosQueue(tx, &common.ListSpec{
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
		"qos_queue": result,
	})
}

//List handles a List REST API Request.
func (api *QosQueueRESTAPI) List(c echo.Context) error {
	var result []*models.QosQueue
	var err error
	auth := common.GetAuthContext(c)
	listSpec := common.GetListSpec(c)
	listSpec.Auth = auth
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListQosQueue(tx, listSpec)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"qos-queues": result,
	})
}
