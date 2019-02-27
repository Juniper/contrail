package logic

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// ReadAll logic
func (r *Router) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	// TODO implement ReadAll logic
	return []RouterResponse{}, nil
}

// Create logic
func (r *Router) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	lr, err := r.neutronToVnc(ctx, rp)
	if err != nil {
		return nil, err
	}

	lrResponse, err := rp.WriteService.CreateLogicalRouter(ctx, &services.CreateLogicalRouterRequest{
		LogicalRouter: lr,
	})
	if err != nil {
		return nil, err
	}

	return r.vncToNeutron(lrResponse.GetLogicalRouter()), nil
}

// Read logic
func (r *Router) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	// TODO: If fields == ["tenant_id"], return {"id": id, "tenant_id": None}

	lrResponse, err := rp.ReadService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{
		ID: id,
	})
	if errutil.IsNotFound(err) {
		return nil, newRouterNotFoundError(id)
	} else if err != nil {
		return nil, err
	}

	return r.vncToNeutron(lrResponse.GetLogicalRouter()), nil
}

// Update logic
func (r *Router) Update(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	lr, err := r.neutronToVnc(ctx, rp)
	if err != nil {
		return nil, err
	}
	lr.UUID = id

	// TODO: Check the referred VN's RouterExternal.

	_, err = rp.WriteService.UpdateLogicalRouter(ctx, &services.UpdateLogicalRouterRequest{
		LogicalRouter: lr,
		FieldMask: types.FieldMask{
			Paths: []string{models.LogicalRouterFieldVirtualNetworkRefs},
			// TODO: Update other fields.
		},
	})
	if errutil.IsNotFound(err) {
		return nil, newRouterNotFoundError(id)
	} else if err != nil {
		return nil, err
	}

	return r.Read(ctx, rp, id)
}

// Delete logic
func (r *Router) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	// TODO: Check VMI refs.

	if _, err := rp.WriteService.DeleteLogicalRouter(ctx, &services.DeleteLogicalRouterRequest{
		ID: id,
	}); errutil.IsNotFound(err) {
		return nil, newRouterNotFoundError(id)
	} else if errutil.IsConflict(err) {
		return nil, newRouterInUseError(id)
	} else if err != nil {
		return nil, err
	}

	return &RouterResponse{}, nil
}

func (r *Router) neutronToVnc(ctx context.Context, rp RequestParameters) (*models.LogicalRouter, error) {
	projectUUID, err := neutronIDToVncUUID(rp.RequestContext.TenantID)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid tenant id: %s", rp.RequestContext.TenantID)
	}

	return &models.LogicalRouter{
		Name:        r.Name,
		DisplayName: r.Name,
		UUID:        r.ID,
		ParentUUID:  projectUUID,
		ParentType:  models.KindProject,
		IDPerms: &models.IdPermsType{
			Enable:      r.AdminStateUp,
			Description: r.Description,
		},
		VirtualNetworkRefs: r.makeVNRefs(rp),
	}, nil
}

func (r *Router) vncToNeutron(lr *models.LogicalRouter) *RouterResponse {
	response := &RouterResponse{
		ID:                  lr.GetUUID(),
		TenantID:            VncUUIDToNeutronID(lr.GetParentUUID()),
		AdminStateUp:        lr.GetIDPerms().GetEnable(),
		Shared:              false,
		Status:              netStatusActive,
		GWPortID:            "",
		ExternalGatewayInfo: r.makeExternalGatewayInfo(lr),
		Description:         lr.GetIDPerms().GetDescription(),
		CreatedAt:           lr.GetIDPerms().GetCreated(),
		UpdatedAt:           lr.GetIDPerms().GetLastModified(),
	}

	response.Name = lr.GetDisplayName()
	if response.Name == "" {
		response.Name = lr.GetName()
	}

	if contrailExtensionsEnabled {
		response.FQName = lr.GetFQName()
	}

	return response
}

func (r *Router) makeExternalGatewayInfo(lr *models.LogicalRouter) ExtGatewayInfo {
	vnRefs := lr.GetVirtualNetworkRefs()
	if len(vnRefs) == 0 {
		return ExtGatewayInfo{}
	}
	vnUUID := vnRefs[0].GetUUID()
	if vnUUID == "" {
		return ExtGatewayInfo{}
	}

	return ExtGatewayInfo{
		NetworkID:  vnUUID,
		EnableSnat: true,
	}
}

func (r *Router) makeVNRefs(rp RequestParameters) (refs []*models.LogicalRouterVirtualNetworkRef) {
	// TODO: Make r.ExternalGatewayInfo a pointer and check if it is nil
	// instead of checking r.ExternalGatewayInfo.NetworkID in the fieldmask.
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(
		RouterFieldExternalGatewayInfo,
		ExtGatewayInfoFieldNetworkID,
	)) {
		refs = append(refs, &models.LogicalRouterVirtualNetworkRef{
			UUID: r.ExternalGatewayInfo.NetworkID,
		})
	}

	return refs
}

func newRouterNotFoundError(id string) *Error {
	return newNeutronError(routerNotFound, errorFields{
		"router_id": id,
	})
}

func newRouterInUseError(id string) *Error {
	return newNeutronError(routerInUse, errorFields{
		"router_id": id,
	})
}
