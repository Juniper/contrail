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

//VirtualMachineInterfaceRESTAPI
type VirtualMachineInterfaceRESTAPI struct {
	DB *sql.DB
}

type VirtualMachineInterfaceCreateRequest struct {
	Data *models.VirtualMachineInterface `json:"virtual-machine-interface"`
}

//Path returns api path for collections.
func (api *VirtualMachineInterfaceRESTAPI) Path() string {
	return "/virtual-machine-interfaces"
}

//LongPath returns api path for elements.
func (api *VirtualMachineInterfaceRESTAPI) LongPath() string {
	return "/virtual-machine-interface/:id"
}

//SetDB sets db object
func (api *VirtualMachineInterfaceRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *VirtualMachineInterfaceRESTAPI) Create(c echo.Context) error {
	requestData := &VirtualMachineInterfaceCreateRequest{
		Data: models.MakeVirtualMachineInterface(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
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
			return db.CreateVirtualMachineInterface(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *VirtualMachineInterfaceRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *VirtualMachineInterfaceRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteVirtualMachineInterface(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *VirtualMachineInterfaceRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.VirtualMachineInterface
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowVirtualMachineInterface(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"virtual_machine_interface": result,
	})
}

//List handles a List REST API Request.
func (api *VirtualMachineInterfaceRESTAPI) List(c echo.Context) error {
	var result []*models.VirtualMachineInterface
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListVirtualMachineInterface(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"virtual-machine-interfaces": result,
	})
}
