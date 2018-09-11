package types

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

func (sv *ContrailTypeLogicService) validateVrouterAllocationPools(
	ctx context.Context,
	vr *models.VirtualRouter,
) error {
	if len(vr.GetNetworkIpamRefs()) == 0 {
		return nil
	}

	ipamRefUUIDs := vr.GetNetworkIpamRefUUIDs()

	networkIpamsRes, err := sv.ReadService.ListNetworkIpam(ctx, &services.ListNetworkIpamRequest{
		Spec: &baseservices.ListSpec{
			ObjectUUIDs: ipamRefUUIDs,
		},
	})

	if err != nil {
		return common.ErrorBadRequestf("Error in dbe_list: %v", err)
	}

	for _, netIpam := range networkIpamsRes.GetNetworkIpams() {
		if netIpam.GetIpamSubnetMethod() != models.FlatSubnet {
			return errors.Errorf(
				"only flat-subnet ipam can be attached to vrouter: NetworkIpam %s has subnet method %s",
				netIpam.GetUUID(), netIpam.GetIpamSubnetMethod())
		}
	}
	//TODO: Validate vrouter allocation pools

	return nil
}

// CreateVirtualRouter virtual-router create specific logic.
func (sv *ContrailTypeLogicService) CreateVirtualRouter(
	ctx context.Context, request *services.CreateVirtualRouterRequest,
) (*services.CreateVirtualRouterResponse, error) {
	var response *services.CreateVirtualRouterResponse
	vr := request.VirtualRouter
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			err := sv.validateVrouterAllocationPools(ctx, vr)
			if err != nil {
				return common.ErrorBadRequestf("virtual-router %s allocation pools validations failed: %v",
					vr.GetUUID(), err)
			}
			response, err = sv.BaseService.CreateVirtualRouter(ctx, request)
			return err
		})

	return response, err
}

// UpdateVirtualRouter virtual-router update specific logic.
func (sv *ContrailTypeLogicService) UpdateVirtualRouter(
	ctx context.Context, request *services.UpdateVirtualRouterRequest,
) (*services.UpdateVirtualRouterResponse, error) {
	var response *services.UpdateVirtualRouterResponse
	vr := request.GetVirtualRouter()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			err := sv.validateVrouterAllocationPools(ctx, vr)
			if err != nil {
				return common.ErrorBadRequestf("virtual-router %s allocation pools validations failed: %v",
					vr.GetUUID(), err)
			}
			response, err = sv.BaseService.UpdateVirtualRouter(ctx, request)
			return err
		})

	return response, err
}

// DeleteVirtualRouter virtual-router delete specific logic.
func (sv *ContrailTypeLogicService) DeleteVirtualRouter(
	ctx context.Context, request *services.DeleteVirtualRouterRequest,
) (*services.DeleteVirtualRouterResponse, error) {
	return sv.BaseService.DeleteVirtualRouter(ctx, request)
}
