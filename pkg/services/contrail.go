package services

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
)

// MetaData represents resource meta data.
type MetaData struct {
	UUID   string
	FQName []string
	Type   string
}

type metadataGetter interface {
	GetMetaData(ctx context.Context, uuid string, fqName []string) (*MetaData, error)
}

// nolint
type ContrailService struct {
	BaseService

	MetadataGetter metadataGetter
	TypeValidator  *models.TypeValidator
}

//RESTSync handles a bulk create request.
func (service *ContrailService) RESTSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("bind failed on sync")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	responses, err := events.Process(ctx, service)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, responses.Events)
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

func (r *RefUpdate) loadUUIDfromFQName(ctx context.Context, mg metadataGetter) error {
	if len(r.RefFQName) == 0 {
		return common.ErrorBadRequest("ref-uuid or ref-fq-name must be specified")
	}
	m, err := mg.GetMetaData(ctx, "", r.RefFQName)
	if err != nil {
		return common.ErrorBadRequestf("bad ref-fq-name provided")
	}
	r.RefUUID = m.UUID
	return nil
}
func (r *RefUpdate) validate() error {
	if r.UUID == "" || r.Type == "" || r.RefType == "" || r.Operation == "" {
		return common.ErrorBadRequestf(
			"uuid/type/ref-type/operation is null: %s, %s, %s, %s",
			r.UUID, r.Type, r.RefType, r.Operation,
		)
	}

	if r.Operation != RefOperationAdd && r.Operation != RefOperationDelete {
		return common.ErrorBadRequestf("operation should be add or delete, was %s", r.Operation)
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
		if err := data.loadUUIDfromFQName(ctx, service.MetadataGetter); err != nil {
			return common.ToHTTPError(err)
		}
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
