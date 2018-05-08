package types

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
)

var errorMultiPolicyServiceChain = common.ErrorBadRequest("Multi policy service chains are not supported, with both import export external route targets")

//CreateVirtualNetwork do pre check for virtual network.
func (service *ContrailTypeLogicService) CreateVirtualNetwork(
	ctx context.Context,
	request *models.CreateVirtualNetworkRequest) (response *models.CreateVirtualNetworkResponse, err error) {
	virtualNetwork := request.VirtualNetwork
	// check if multiple policy service chain supported
	if !virtualNetwork.isValidMultiPolicyServiceChainConfig(virtualNetwork) {
		return nil, errorMultiPolicyServiceChain
	}
	//  neutorn <-> vnc sharing
	virtualNetwork.MakeNeutronCompatible()
	// Does not authorize to set the virtual network ID as it's allocated
	// by the vnc server
	if virtualNetwork.HasVirtualNetworkNetworkID() {
		return nil, common.ErrorBadRequest("Cannot set the virtual network ID")
	}

	err = db.DoInTransaction(
		ctx,
		service.DB.GetDB(),
		func(ctx context.Context) error {
			// allocate virtual network ID
			virtualNetwork.VirtualNetworkNetworkID, err = service.DB.AllocateInt(ctx, VirtualNetworkIDPoolKey)
			if err != nil {
				return err
			}
			// check ipam network subnets
			// check route target
			// check provider network property
			// check network support bgp vpn types
			// check if we can reference the BGP VPNs
			// process network ipam refs references
			response, err = service.Next().CreateVirtualNetwork(ctx, request)

			// create native/vn-default routing instance

			return err
		})

	return response, err
}

func (service *ContrailTypeLogicService) getVirtualNetworkID(ctx context.Context, id string) (string, error) {
	var getResponse *models.GetVirtualNetworkResponse
	getResponse, err = service.DB.GetVirtualNetwork(ctx, &models.GetVirtualNetworkRequest{
		ID: id,
	})
	if err != nil {
		return "", err
	}
	return getResponse.VirtualNetwork.VirtualNetworkNetworkID, nil
}

//DeleteVirtualNetwork do pre check for delete network.
func (service *ContrailTypeLogicService) DeleteVirtualNetwork(
	ctx context.Context,
	request *models.DeleteVirtualNetworkRequest) (response *models.DeleteVirtualNetworkResponse, err error) {
	id := request.ID

	err = db.DoInTransaction(
		ctx,
		service.DB.DB,
		func(ctx context.Context) error {
			// deallocate virtual network ID
			virtualNetworkID, err := service.getVirtualNetworkID(ctx, id)
			if err != nil {
				return err
			}
			err = service.DB.DeallocateInt(ctx, VirtualNetworkIDPoolKey, virtualNetworkID)
			if err != nil {
				return err
			}

			response, err = service.Next().DeleteVirtualNetwork(ctx, request)
			return err
		})

	return response, err
}
