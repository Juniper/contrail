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

//SecurityGroupRESTAPI
type SecurityGroupRESTAPI struct {
	DB *sql.DB
}

type SecurityGroupCreateRequest struct {
	Data *models.SecurityGroup `json:"security-group"`
}

type SecurityGroupUpdateRequest struct {
	Data map[string]interface{} `json:"security-group"`
}

//Path returns api path for collections.
func (api *SecurityGroupRESTAPI) Path() string {
	return "/security-groups"
}

//LongPath returns api path for elements.
func (api *SecurityGroupRESTAPI) LongPath() string {
	return "/security-group/:id"
}

//SetDB sets db object
func (api *SecurityGroupRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *SecurityGroupRESTAPI) Create(c echo.Context) error {
	requestData := &SecurityGroupCreateRequest{
		Data: models.MakeSecurityGroup(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
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
			return db.CreateSecurityGroup(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *SecurityGroupRESTAPI) Update(c echo.Context) error {
	id := c.Param("id")
	requestData := &SecurityGroupUpdateRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
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
			return db.UpdateSecurityGroup(tx, id, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("db update failed")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]map[string]string{
		"security-group": {
			"uuid": id,
			"uri":  "/" + "security-group" + "/" + id},
	})
}

//Delete handles a REST Delete request.
func (api *SecurityGroupRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteSecurityGroup(tx, id, auth)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *SecurityGroupRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	auth := common.GetAuthContext(c)
	var result []*models.SecurityGroup
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListSecurityGroup(tx, &common.ListSpec{
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
		"security_group": result,
	})
}

//List handles a List REST API Request.
func (api *SecurityGroupRESTAPI) List(c echo.Context) error {
	var result []*models.SecurityGroup
	var err error
	auth := common.GetAuthContext(c)
	listSpec := common.GetListSpec(c)
	listSpec.Auth = auth
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListSecurityGroup(tx, listSpec)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"security-groups": result,
	})
}
