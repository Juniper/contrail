package logic

import (
	"context"

	"github.com/gogo/protobuf/types"

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
		// TODO Wrap.
		return nil, err
	}

	// TODO Wrap err.
	return r.vncToNeutron(lrResponse.GetLogicalRouter())
}

// Read logic
func (r *Router) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	// TODO: If fields == ["tenant_id"], return {"id": id, "tenant_id": None}

	lrResponse, err := rp.ReadService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{
		ID: id,
	})
	if err != nil {
		// TODO Return a special error if IsNotFound(err) ?
		// TODO Wrap.
		return nil, err
	}

	return r.vncToNeutron(lrResponse.GetLogicalRouter())
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
	if err != nil {
		// TODO Return a special error if IsNotFound(err) ?
		// TODO Wrap.
		return nil, err
	}

	return r.Read(ctx, rp, id)
}

// Delete logic
func (r *Router) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	// TODO: Check VMI refs.

	if _, err := rp.WriteService.DeleteLogicalRouter(ctx, &services.DeleteLogicalRouterRequest{
		ID: id,
	}); err != nil {
		// TODO Wrap.
		return nil, err
	}

	return &RouterResponse{}, nil
}

func (r *Router) neutronToVnc(ctx context.Context, rp RequestParameters) (*models.LogicalRouter, error) {
	projectUUID, err := neutronIDToVncUUID(rp.RequestContext.TenantID)
	if err != nil {
		// TODO Wrap.
		return nil, err
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

func (r *Router) vncToNeutron(lr *models.LogicalRouter) (*RouterResponse, error) {
	response := &RouterResponse{
		ID:                  lr.GetUUID(),
		Name:                lr.GetDisplayName(), // TODO or Name if it's empty.
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

	if contrailExtensionsEnabled {
		response.FQName = lr.GetFQName()
	}

	return response, nil
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
