package types

import (
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/services"
)

var errorMultiPolicyServiceChain = common.ErrorBadRequest(
	"Multi policy service chains are not supported, with both import export external route targets")

//CreateVirtualNetwork do pre check for virtual network.
func (service *ContrailTypeLogicService) CreateVirtualNetwork(
	ctx context.Context,
	request *services.CreateVirtualNetworkRequest) (response *services.CreateVirtualNetworkResponse, err error) {
	virtualNetwork := request.VirtualNetwork
	// check if multiple policy service chain supported
	if !virtualNetwork.IsValidMultiPolicyServiceChainConfig() {
		return nil, errorMultiPolicyServiceChain
	}
	//  neutorn <-> vnc sharing
	virtualNetwork.MakeNeutronCompatible()
	// Does not authorize to set the virtual network ID as it's allocated
	// by the vnc server
	if virtualNetwork.HasVirtualNetworkNetworkID() {
		return nil, common.ErrorBadRequest("Cannot set the virtual network ID")
	}

	err = service.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			// allocate virtual network ID
			virtualNetwork.VirtualNetworkNetworkID, err = service.IntPoolAllocator.AllocateInt(ctx, VirtualNetworkIDPoolKey)
			if err != nil {
				return err
			}
			//TODO: check ipam network subnets
			//TODO: check route target
			//TODO: check provider network property
			//TODO: check network support bgp vpn types
			//TODO: check if we can reference the BGP VPNs
			//TODO: process network ipam refs references
			response, err = service.Next().CreateVirtualNetwork(ctx, request)

			//TODO: create native/vn-default routing instance

			return err
		})

	return response, err
}

func (service *ContrailTypeLogicService) getVirtualNetworkID(ctx context.Context, id string) (int64, error) {
	var getResponse *services.GetVirtualNetworkResponse
	getResponse, err := service.DataService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: id,
	})
	if err != nil {
		return 0, err
	}
	return getResponse.VirtualNetwork.VirtualNetworkNetworkID, nil
}

//DeleteVirtualNetwork do pre check for delete network.
func (service *ContrailTypeLogicService) DeleteVirtualNetwork(
	ctx context.Context,
	request *services.DeleteVirtualNetworkRequest) (response *services.DeleteVirtualNetworkResponse, err error) {
	id := request.ID

	err = service.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			// deallocate virtual network ID
			var virtualNetworkID int64
			virtualNetworkID, err = service.getVirtualNetworkID(ctx, id)
			if err != nil {
				return err
			}

			err = service.IntPoolAllocator.DeallocateInt(ctx, VirtualNetworkIDPoolKey, virtualNetworkID)
			if err != nil {
				return err
			}

			response, err = service.Next().DeleteVirtualNetwork(ctx, request)
			return err
		})

	return response, err
}
