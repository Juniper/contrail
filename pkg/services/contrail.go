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

// RefUpdate represents ref-update input data.
type RefUpdate struct {
	Operation string          `json:"operation"`
	Type      string          `json:"type"`
	UUID      string          `json:"uuid"`
	RefType   string          `json:"ref-type"`
	RefUUID   string          `json:"ref-uuid"`
	RefFQName []string        `json:"ref-fq-name"`
	Attr      json.RawMessage `json:"attr"`
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
	ctx := c.Request().Context()

	if data.Operation != OperationAdd && data.Operation != OperationDelete {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf("Bad Request: operation should be add or delete: %s", data.Operation),
		)
	}

	if data.RefUUID == "" {
		if len(data.RefFQName) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request: ref-uuid or ref-fq-name must be specified")
		}
		m, err := service.MetadataGetter.GetMetaData(ctx, "", data.RefFQName)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad ref-fq-name provided")
		}
		data.RefUUID = m.UUID
	}

	e := NewEventFromRefUpdate(&data)
	if e == nil {
		log.Debug("event from ref update failed")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	_, err := e.Process(ctx, service)
	if err != nil {
		return common.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"uuid": data.UUID})
}
