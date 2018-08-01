package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
)

type metadataGetter interface {
	GetMetaData(ctx context.Context, uuid string, fqName []string) (*models.MetaData, error)
}

// nolint
type ContrailService struct {
	BaseService

	MetadataGetter    metadataGetter
	TypeValidator     *models.TypeValidator
	InTransactionDoer InTransactionDoer
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
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON format")
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

// PropCollectionUpdateRequest is input request for /prop-collection-update endpoint.
type PropCollectionUpdateRequest struct {
	UUID    string                        `json:"uuid"`
	Updates []models.PropCollectionUpdate `json:"updates"`
}

func (p *PropCollectionUpdateRequest) validate() error {
	if p.UUID == "" {
		return common.ErrorBadRequest("prop_collection_update needs object uuid")
	}
	return nil
}

// RESTPropCollectionUpdate handles a prop-collection-update request.
func (service *ContrailService) RESTPropCollectionUpdate(c echo.Context) error {
	var data PropCollectionUpdateRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON format")
	}

	if err := data.validate(); err != nil {
		return common.ToHTTPError(err)
	}

	if err := service.InTransactionDoer.DoInTransaction(c.Request().Context(), func(ctx context.Context) error {
		m, err := service.MetadataGetter.GetMetaData(ctx, data.UUID, nil)
		if err != nil {
			return errors.Wrap(err, "error getting metadata for provided uuid: %v")
		}

		obj, err := GetObject(ctx, service.Next(), m.Type, data.UUID)
		if err != nil {
			return errors.Wrapf(err, "error getting %v with uuid = %v", m.Type, data.UUID)
		}

		updateMap := map[string]interface{}{}

		for _, update := range data.Updates {
			var updated map[string]interface{}
			updated, err = obj.ApplyPropCollectionUpdate(&update)
			if err != nil {
				return common.ToHTTPError(err)
			}
			for key, value := range updated {
				updateMap[key] = value
			}
		}
		e := NewEvent(&EventOption{
			Data:      updateMap,
			Kind:      m.Type,
			UUID:      data.UUID,
			Operation: OperationUpdate,
		})

		_, err = e.Process(ctx, service)
		return err
	}); err != nil {
		return common.ToHTTPError(err)
	}

	return c.NoContent(http.StatusOK)
}

// RefRelax represents ref-relax-for-delete input data.
type RefRelax struct {
	UUID    string `json:"uuid"`
	RefUUID string `json:"ref-uuid"`
}

func (r *RefRelax) validate() error {
	if r.UUID == "" || r.RefUUID == "" {
		return common.ErrorBadRequestf(
			"bad request: both uuid and ref-uuid should be specified: %s, %s", r.UUID, r.RefUUID)
	}

	return nil
}

// RESTRefRelaxForDelete handles a ref-relax-for-delete request.
func (service *ContrailService) RESTRefRelaxForDelete(c echo.Context) error {
	var data RefRelax
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON format")
	}

	if err := data.validate(); err != nil {
		return common.ToHTTPError(err)
	}

	// TODO (Kamil): implement ref-relax logic

	return c.JSON(http.StatusOK, map[string]interface{}{"uuid": data.UUID})
}
