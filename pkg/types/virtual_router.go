package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/services"
)

// CreateVirtualRouter virtual-router create specific logic.
func (sv *ContrailTypeLogicService) CreateVirtualRouter(
	ctx context.Context, request *services.CreateVirtualRouterRequest,
) (*services.CreateVirtualRouterResponse, error) {
	return sv.BaseService.CreateVirtualRouter(ctx, request)
}

// UpdateVirtualRouter virtual-router update specific logic.
func (sv *ContrailTypeLogicService) UpdateVirtualRouter(
	ctx context.Context, request *services.UpdateVirtualRouterRequest,
) (*services.UpdateVirtualRouterResponse, error) {
	return sv.BaseService.UpdateVirtualRouter(ctx, request)
}

// DeleteVirtualRouter virtual-router delete specific logic.
func (sv *ContrailTypeLogicService) DeleteVirtualRouter(
	ctx context.Context, request *services.DeleteVirtualRouterRequest,
) (*services.DeleteVirtualRouterResponse, error) {
	return sv.BaseService.DeleteVirtualRouter(ctx, request)
}
