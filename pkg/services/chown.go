package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	models "github.com/Juniper/contrail/pkg/models"
	"github.com/gogo/protobuf/types"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// RESTChown handles chown request.
// TODO(dfurman): it should be Contrail plugin
func (p *ContrailEndpointPlugin) RESTChown(c echo.Context) error {
	var data ChownRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	ctx := c.Request().Context()
	if _, err := p.Chown(ctx, &data); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// Chown handles chown request.
func (p *ContrailEndpointPlugin) Chown(ctx context.Context, request *ChownRequest) (*types.Empty, error) {
	if err := validateChownRequest(request); err != nil {
		return nil, err
	}

	if err := p.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		metadata, err := p.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{UUID: request.GetUUID()})
		if err != nil {
			return errors.Wrapf(err, "failed to change the owner of the resource with UUID '%v'", request.GetUUID())
		}

		// nolint: lll
		// TODO: check permissions, see https://github.com/Juniper/contrail-controller/blob/137e2a08025e1ae7084621c0f081f7b99d1b04cd/src/config/api-server/vnc_cfg_api_server/vnc_cfg_api_server.py#L2409

		var fm types.FieldMask
		basemodels.FieldMaskAppend(&fm, basemodels.CommonFieldPerms2, models.PermType2FieldOwner)

		event, err := NewEvent(EventOption{
			UUID:      request.GetUUID(),
			Kind:      metadata.Type,
			Operation: OperationUpdate,
			Data: map[string]interface{}{
				"perms2": map[string]interface{}{
					"owner": request.GetOwner(),
				},
			},
			FieldMask: &fm,
		})
		if err != nil {
			return errors.Wrapf(err, "failed to change the owner of '%v' with UUID '%v'", metadata.Type, request.GetUUID())
		}

		_, err = event.Process(ctx, p.Service)
		return errors.Wrapf(err, "failed to change the owner of '%v' with UUID '%v'", metadata.Type, request.GetUUID())
	}); err != nil {
		return nil, err
	}

	return &types.Empty{}, nil
}

func validateChownRequest(r *ChownRequest) error {
	if r == nil || r.UUID == "" || r.Owner == "" {
		return errutil.ErrorBadRequestf(
			"bad request: both uuid and owner should be specified: %s, %s", r.GetUUID(), r.GetOwner())
	}

	if _, err := uuid.FromString(r.GetUUID()); err != nil {
		return errutil.ErrorBadRequestf(
			"bad request: invalid uuid format (not UUID): %s", r.GetUUID())
	}
	if _, err := uuid.FromString(r.GetOwner()); err != nil {
		return errutil.ErrorBadRequestf(
			"bad request: invalid owner format (not UUID): %s", r.GetOwner())
	}

	return nil
}
