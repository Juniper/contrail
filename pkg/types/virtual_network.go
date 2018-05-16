package types

import (
	"net"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
)

var errorMultiPolicyServiceChain = common.ErrorBadRequest("Multi policy service chains are not supported, with both import export external route targets")

func isSubnetOverlap(subnet1, subnet2 *net.IPNet) bool {
	return subnet1.Contains(subnet2.IP) || subnet2.Contains(subnet1.IP)
}

func appendSubnetIfNoOverlap(subnets []*net.IPNet, subnet *net.IPNet) ([]*net.IPNet, error) {
	for _, existingSubnet := range subnets {
		if isSubnetOverlap(subnet, existingSubnet) {
			return nil, common.ErrorBadRequest(
				"Overlapping addresses: " + subnet.String() + "," + existingSubnet.String())
		}
	}
	return append(subnets, subnet), nil
}

func mergeSubnetIfNoOverlap(subnets []*net.IPNet, ipamSubnets []*models.IpamSubnetType) ([]*net.IPNet, error) {
	for _, subnet := range ipamSubnets {
		n, err := subnet.Subnet.Net()
		if err != nil {
			return nil, err
		}
		subnets, err = appendSubnetIfNoOverlap(subnets, n)
		if err != nil {
			return nil, err
		}
	}
	return subnets, nil
}

func (service *ContrailTypeLogicService) checkIpamNetworkSubnets(ctx context.Context, virtualNetwork *models.VirtualNetwork) error {
	// check for each ipam references
	ipamReferences := virtualNetwork.NetworkIpamRefs
	// ip subnet must not overlap in a same network.
	// so we create a list of subnet for overlap check.
	subnets := []*net.IPNet{}
	for _, ipamReference := range ipamReferences {
		ipamResponse, err := service.DB.GetNetworkIpam(ctx, &models.GetNetworkIpamRequest{
			ID: ipamReference.UUID,
		})
		if err != nil {
			return err
		}
		var ipamSubnets []*models.IpamSubnetType
		ipam := ipamResponse.NetworkIpam
		vnSubnet := ipamReference.GetAttr()
		if ipam.IsFlatSubnet() {
			// network mode must be L3
			if !virtualNetwork.IsL3Mode() {
				return common.ErrorBadRequest("flat-subnet is allowed only with l3 network")
			}
			err := vnSubnet.ValidateFlatSubnet()
			if err != nil {
				return err
			}
			ipamSubnets = ipam.GetIpamSubnets().GetSubnets()
		} else {
			err := vnSubnet.ValidateUserDefined()
			if err != nil {
				return err
			}
			ipamSubnets = vnSubnet.GetIpamSubnets()
		}
		subnets, err = mergeSubnetIfNoOverlap(subnets, ipamSubnets)
		if err != nil {
			return err
		}
		//TODO: check network subnet quota
	}
	return nil
}

//CreateVirtualNetwork do pre check for virtual network.
func (service *ContrailTypeLogicService) CreateVirtualNetwork(
	ctx context.Context,
	request *models.CreateVirtualNetworkRequest) (response *models.CreateVirtualNetworkResponse, err error) {
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

	err = db.DoInTransaction(
		ctx,
		service.DB.DB(),
		func(ctx context.Context) error {
			// allocate virtual network ID
			virtualNetwork.VirtualNetworkNetworkID, err = service.DB.AllocateInt(ctx, VirtualNetworkIDPoolKey)
			if err != nil {
				return err
			}
			//check ipam network subnets
			err = service.checkIpamNetworkSubnets(ctx, virtualNetwork)
			if err != nil {
				return err
			}
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
	var getResponse *models.GetVirtualNetworkResponse
	getResponse, err := service.DB.GetVirtualNetwork(ctx, &models.GetVirtualNetworkRequest{
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
	request *models.DeleteVirtualNetworkRequest) (response *models.DeleteVirtualNetworkResponse, err error) {
	id := request.ID

	err = db.DoInTransaction(
		ctx,
		service.DB.DB(),
		func(ctx context.Context) error {
			// deallocate virtual network ID
			var virtualNetworkID int64
			virtualNetworkID, err = service.getVirtualNetworkID(ctx, id)
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
