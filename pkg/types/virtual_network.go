package types

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
)

var errorMultiPolicyServiceChain = common.ErrorBadRequest("Multi policy service chains are not supported, with both import export external route targets")

func checkMultiPolicyServiceChainSupported(virtualNetwork *models.VirtualNetwork) error {
	if !virtualNetwork.MultiPolicyServiceChainsEnabled {
		return nil
	}
	if len(virtualNetwork.GetRouteTargetList().GetRouteTarget()) != 0 {
		return errorMultiPolicyServiceChain
	}
	for _, importRouteTarget := range virtualNetwork.GetImportRouteTargetList().GetRouteTarget() {
		for _, exportRouteTarget := range virtualNetwork.GetExportRouteTargetList().GetRouteTarget() {
			if importRouteTarget == exportRouteTarget {
				return errorMultiPolicyServiceChain
			}
		}
	}
	return nil
}

//CreateVirtualNetwork do pre check for virtual network.
func (service *ContrailTypeLogicService) CreateVirtualNetwork(
	ctx context.Context,
	request *models.CreateVirtualNetworkRequest) (response *models.CreateVirtualNetworkResponse, err error) {
	virtualNetwork := request.VirtualNetwork
	// check if multiple policy service chain supported
	err = checkMultiPolicyServiceChainSupported(virtualNetwork)
	if err != nil {
		return nil, err
	}
	//  neutorn <-> vnc sharing
	if virtualNetwork.Perms2.GlobalAccess == PERMS_RWX {
		virtualNetwork.IsShared = true
	}
	if virtualNetwork.IsShared == true {
		virtualNetwork.Perms2.GlobalAccess = PERMS_RWX
	}
	// Does not authorize to set the virtual network ID as it's allocated
	// by the vnc server
	if virtualNetwork.VirtualNetworkNetworkID != 0 {
		return nil, common.ErrorBadRequest("Cannot set the virtual network ID")
	}

	db.DoInTransaction(
		ctx,
		service.DB.DB,
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

//DeleteVirtualNetwork do pre check for delete network.
func (service *ContrailTypeLogicService) DeleteVirtualNetwork(
	ctx context.Context,
	request *models.DeleteVirtualNetworkRequest) (response *models.DeleteVirtualNetworkResponse, err error) {
	id := request.ID

	db.DoInTransaction(
		ctx,
		service.DB.DB,
		func(ctx context.Context) error {
			// deallocate virtual network ID
			getResponse, err := service.DB.GetVirtualNetwork(ctx, &models.GetVirtualNetworkRequest{
				ID: id,
			})
			if err != nil {
				return err
			}
			virtualNetworkID := getResponse.VirtualNetwork.VirtualNetworkNetworkID
			err = service.DB.DeallocateInt(ctx, VirtualNetworkIDPoolKey, virtualNetworkID)
			if err != nil {
				return err
			}

			response, err = service.Next().DeleteVirtualNetwork(ctx, request)
			return err
		})

	return response, err
}
