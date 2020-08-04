package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models"
	"github.com/labstack/echo"
)

// Endpoint paths.
const (
	RefUpdatePath         = "/ref-update"
	RefRelaxForDeletePath = "/ref-relax-for-delete"
)

type RefUpdater interface {
	UpdateRef(context.Context, *RefUpdate) error
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
	Operation RefOperation           `json:"operation"`
	Type      string                 `json:"type"`
	UUID      string                 `json:"uuid"`
	RefType   string                 `json:"ref-type"`
	RefUUID   string                 `json:"ref-uuid"`
	RefFQName []string               `json:"ref-fq-name"`
	Attr      map[string]interface{} `json:"attr,omitempty"`
}

// Validate checks validity of request data.
func (r *RefUpdate) Validate() error {
	if r.UUID == "" || r.Type == "" || r.RefType == "" || r.Operation == "" {
		return errutil.ErrorBadRequestf(
			"uuid/type/ref-type/operation is null: %s, %s, %s, %s",
			r.UUID, r.Type, r.RefType, r.Operation,
		)
	}

	if r.Operation != RefOperationAdd && r.Operation != RefOperationDelete {
		return errutil.ErrorBadRequestf("operation should be ADD or DELETE, was %s", r.Operation)
	}

	return nil
}

// RefUpdatePlugin creates a RefUpdate endpoint that uses provided RefUpdater.
type RefUpdatePlugin struct {
	MetadataGetter MetadataGetter
	RefUpdater     RefUpdater
}

// RegisterHTTPAPI registers HTTP endpoints.
func (p *RefUpdatePlugin) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.POST(RefUpdatePath, p.RESTRefUpdate)
}

// RegisterGRPCAPI does nothing.
func (p *RefUpdatePlugin) RegisterGRPCAPI(r apiserver.GRPCRouter) {}

// RESTRefUpdate handles a ref-update request.
func (p *RefUpdatePlugin) RESTRefUpdate(c echo.Context) error {
	var data RefUpdate
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if err := data.Validate(); err != nil {
		return errutil.ToHTTPError(err)
	}

	ctx := c.Request().Context()
	if data.RefUUID == "" {
		m, err := p.MetadataGetter.GetMetadata(ctx, models.FQNameMetadata(data.RefFQName, data.RefType))
		if err != nil {
			return errutil.ToHTTPError(errutil.ErrorBadRequestf("error resolving ref-uuid using ref-fq-name: %v", err))
		}
		data.RefUUID = m.UUID
	}

	if err := p.RefUpdater.UpdateRef(ctx, &data); err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"uuid": data.UUID})
}

// RefRelaxer makes references not prevent the referenced resource from being deleted.
type RefRelaxer interface {
	RelaxRef(ctx context.Context, fromUUID, toUUID string) error
}

// RefRelaxPlugin is a plugin that adds endpoints allowing to relax references.
type RefRelaxPlugin struct {
	RefRelaxer RefRelaxer
}

// RegisterHTTPAPI registers HTTP endpoints.
func (p *RefRelaxPlugin) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.POST(RefRelaxForDeletePath, p.RESTRefRelaxForDelete)
}

// RegisterGRPCAPI registers GRPC endpoints.
func (p *RefRelaxPlugin) RegisterGRPCAPI(r apiserver.GRPCRouter) {
	r.RegisterService(&_RefRelax_serviceDesc, p)
}

// RESTRefRelaxForDelete handles a ref-relax-for-delete request.
func (p *RefRelaxPlugin) RESTRefRelaxForDelete(c echo.Context) error {
	var data RelaxRefRequest

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if err := validateRelaxRefRequest(&data); err != nil {
		return errutil.ToHTTPError(err)
	}

	response, err := p.RelaxRef(c.Request().Context(), &data)
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, response)
}

func validateRelaxRefRequest(r *RelaxRefRequest) error {
	if r.UUID == "" || r.RefUUID == "" {
		return errutil.ErrorBadRequestf(
			"bad request: both uuid and ref-uuid should be specified: %s, %s", r.UUID, r.RefUUID)
	}

	return nil
}

// RelaxRef makes a reference not prevent the referenced resource from being deleted.
func (p *RefRelaxPlugin) RelaxRef(ctx context.Context, request *RelaxRefRequest) (*RelaxRefResponse, error) {
	if err := p.RefRelaxer.RelaxRef(ctx, request.GetUUID(), request.GetRefUUID()); err != nil {
		return nil, err
	}
	return &RelaxRefResponse{UUID: request.UUID}, nil
}
