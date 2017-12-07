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

//CustomerAttachmentRESTAPI
type CustomerAttachmentRESTAPI struct {
	DB *sql.DB
}

type CustomerAttachmentCreateRequest struct {
	Data *models.CustomerAttachment `json:"customer-attachment"`
}

//Path returns api path for collections.
func (api *CustomerAttachmentRESTAPI) Path() string {
	return "/customer-attachments"
}

//LongPath returns api path for elements.
func (api *CustomerAttachmentRESTAPI) LongPath() string {
	return "/customer-attachment/:id"
}

//SetDB sets db object
func (api *CustomerAttachmentRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *CustomerAttachmentRESTAPI) Create(c echo.Context) error {
	requestData := &CustomerAttachmentCreateRequest{
		Data: models.MakeCustomerAttachment(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "customer_attachment",
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
			return db.CreateCustomerAttachment(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "customer_attachment",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *CustomerAttachmentRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *CustomerAttachmentRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteCustomerAttachment(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *CustomerAttachmentRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.CustomerAttachment
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowCustomerAttachment(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"customer_attachment": result,
	})
}

//List handles a List REST API Request.
func (api *CustomerAttachmentRESTAPI) List(c echo.Context) error {
	var result []*models.CustomerAttachment
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListCustomerAttachment(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"customer-attachments": result,
	})
}
