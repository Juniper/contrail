package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
)

type metadataGetter interface {
	GetMetaData(ctx context.Context, uuid string, fqName []string) (*models.MetaData, error)
}

// nolint
type ContrailService struct {
	BaseService

	MetadataGetter metadataGetter
	TypeValidator  *models.TypeValidator
}

// RESTSync handles Sync API request.
func (service *ContrailService) RESTSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	// TODO: Call events.Sort()

	responses, err := events.Process(c.Request().Context(), service)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, responses.Events)
}

// RefOperation is enum type for ref-update operation.
type RefOperation string

// RefOperation values.
const (
	RefOperationAdd    RefOperation = "ADD"
	RefOperationDelete RefOperation = "DELETE"
)

// RefUpdate represents ref-update input data.
type RefUpdate struct {
	Operation RefOperation    `json:"operation"`
	Type      string          `json:"type"`
	UUID      string          `json:"uuid"`
	RefType   string          `json:"ref-type"`
	RefUUID   string          `json:"ref-uuid"`
	RefFQName []string        `json:"ref-fq-name"`
	Attr      json.RawMessage `json:"attr"`
}

func (r *RefUpdate) validate() error {
	if r.UUID == "" || r.Type == "" || r.RefType == "" || r.Operation == "" {
		return common.ErrorBadRequestf(
			"uuid/type/ref-type/operation is null: %s, %s, %s, %s",
			r.UUID, r.Type, r.RefType, r.Operation,
		)
	}

	if r.Operation != RefOperationAdd && r.Operation != RefOperationDelete {
		return common.ErrorBadRequestf("operation should be ADD or DELETE, was %s", r.Operation)
	}

	return nil
}

// RESTRefUpdate handles a ref-update request.
func (service *ContrailService) RESTRefUpdate(c echo.Context) error {
	var data RefUpdate
	if err := c.Bind(&data); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("bind failed on ref-update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	if err := data.validate(); err != nil {
		return common.ToHTTPError(err)
	}

	ctx := c.Request().Context()
	if data.RefUUID == "" {
		m, err := service.MetadataGetter.GetMetaData(ctx, "", data.RefFQName)
		if err != nil {
			return common.ToHTTPError(common.ErrorBadRequestf("error resolving ref-uuid using ref-fq-name: %v", err))
		}
		data.RefUUID = m.UUID
	}

	e, err := NewEventFromRefUpdate(&data)
	if err != nil {
		return common.ToHTTPError(common.ErrorBadRequest(err.Error()))
	}
	if _, err = e.Process(ctx, service); err != nil {
		return common.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"uuid": data.UUID})
}

// RefRelax represents ref-relax-for-delete input data.
type RefRelax struct {
	UUID    string `json:"uuid"`
	RefUUID string `json:"ref-uuid"`
}

func (r *RefRelax) validate() error {
	if r.UUID == "" || r.RefUUID == "" {
		return common.ErrorBadRequestf(
			"Bad Request: Both uuid and ref-uuid should be specified: %s, %s", r.UUID, r.RefUUID)
	}

	return nil
}

// RESTRefRelaxForDelete handles a ref-relax-for-delete request.
func (service *ContrailService) RESTRefRelaxForDelete(c echo.Context) error {
	var data RefRelax
	if err := c.Bind(&data); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("bind failed on ref-relax-for-delete")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	if err := data.validate(); err != nil {
		return common.ToHTTPError(err)
	}

	// TODO (Kamil): implement ref-relax logic

	return c.JSON(http.StatusOK, map[string]interface{}{"uuid": data.UUID})
}

// Application is a part of set-tag input data. TODO: Investigate it
type Application struct {
	IsGlobal bool   `json:"is_global"`
	Value    string `json:"value"`
}

// SetTag represents set-tag input data.
type SetTag struct {
	App     Application `json:"application"`
	ObjUUID string      `json:"obj_uuid"`
	ObjType string      `json:"obj_type"`
}

func (a *Application) validate() bool {
	return !(a.Value == "")
}

func (t *SetTag) validate() error {
	if t.ObjUUID == "" || t.ObjType == "" {
		return common.ErrorBadRequestf(
			"Bad Request: Both obj_uuid and obj_type should be specified but got uuid: '%s' and type: '%s",
			t.ObjUUID, t.ObjType,
		)
	}
	if !t.checkIfObjTypeFits() {
		return common.ErrorBadRequestf(
			"Bad Request: obj_type should be project, virtual-network, virtual-machine, ",
			"virtual-machine-interface or application-policy-set",
		)
	}
	if !t.App.validate() {
		return common.ErrorBadRequestf(
			"Bad Request: set-tag requires specified application value (uuid: %s type: %s)",
			t.ObjUUID, t.ObjType,
		)
	}
	return nil
}

// TODO: This method is during investigation
func (t *SetTag) checkIfObjTypeFits() bool {
	switch t.ObjType {
	case "project":
		return true
	case "virtual-network":
		return true
	case "virtual-machine":
		return true
	case "virtual-machine-interface":
		return true
	case "application-policy-set":
		return true
	}
	return false
}

// RESTSetTag handles set-tag request.
func (service *ContrailService) RESTSetTag(c echo.Context) error {
	var data SetTag

	if err := c.Bind(&data); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("bind failed on set-tag")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	if err := data.validate(); err != nil {
		return common.ToHTTPError(err)
	}

	// TODO (Ignacy): implement set-tag logic

	return c.JSON(http.StatusOK, map[string]interface{}{})
}
