package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
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

	// TODO: Update gateway.

	// TODO Wrap err.
	return r.vncToNeutron(lrResponse.GetLogicalRouter())
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
	}, nil
}

func (r *Router) vncToNeutron(lr *models.LogicalRouter) (*RouterResponse, error) {
	response := &RouterResponse{
		ID:           lr.GetUUID(),
		Name:         lr.GetDisplayName(), // TODO or Name if it's empty.
		TenantID:     VncUUIDToNeutronID(lr.GetParentUUID()),
		AdminStateUp: lr.GetIDPerms().GetEnable(),
		Shared:       false,
		Status:       netStatusActive,
		GWPortID:     "",
		// TODO ExternalGatewayInfo.
		Description: lr.GetIDPerms().GetDescription(),
		CreatedAt:   lr.GetIDPerms().GetCreated(),
		UpdatedAt:   lr.GetIDPerms().GetLastModified(),
	}

	if contrailExtensionsEnabled {
		response.FQName = lr.GetFQName()
	}

	return response, nil
}
