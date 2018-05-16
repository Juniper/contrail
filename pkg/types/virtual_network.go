package types

import (
	"net"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

//CreateVirtualNetwork do pre check for virtual network.
func (sv *ContrailTypeLogicService) CreateVirtualNetwork(
	ctx context.Context,
	request *services.CreateVirtualNetworkRequest) (response *services.CreateVirtualNetworkResponse, err error) {
	virtualNetwork := request.VirtualNetwork
	// check if multiple policy service chain supported
	if !virtualNetwork.IsValidMultiPolicyServiceChainConfig() {
		return nil, common.ErrorBadRequest("Multi policy service chains are not supported, with both import export external route targets")
	}
	//  neutron <-> vnc sharing
	virtualNetwork.MakeNeutronCompatible()
	// Does not authorize to set the virtual network ID as it's allocated
	// by the vnc server
	if virtualNetwork.HasVirtualNetworkNetworkID() {

		return nil, common.ErrorForbiden("Cannot set the virtual network ID, it's allocated by the server")
	}

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			// Allocate virtual network ID
			virtualNetwork.VirtualNetworkNetworkID, err = sv.IntPoolAllocator.AllocateInt(ctx, VirtualNetworkIDPoolKey)
			if err != nil {
				return err
			}
			err = sv.processIpamNetworkSubnets(ctx, virtualNetwork)
			if err != nil {
				return err
			}
			//TODO: check route target (depends on global system config)
			//TODO: check provider details
			err = sv.checkProviderNetwork(ctx, virtualNetwork)
			if err != nil {
				return err
			}
			err = sv.checkNetworkSupportBGPTypes(ctx, virtualNetwork)
			if err != nil {
				return err
			}
			err = sv.checkBGPVPNRefs(ctx, virtualNetwork)
			if err != nil {
				return err
			}
			//TODO: process network ipam refs references
			response, err = sv.BaseService.CreateVirtualNetwork(ctx, request)

			//TODO: create native/vn-default routing instance

			return err
		})

	return response, err
}

//DeleteVirtualNetwork do pre check for delete network.
func (sv *ContrailTypeLogicService) DeleteVirtualNetwork(
	ctx context.Context,
	request *services.DeleteVirtualNetworkRequest) (response *services.DeleteVirtualNetworkResponse, err error) {
	id := request.ID

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			// Deallocate virtual network ID
			var virtualNetworkID int64
			virtualNetworkID, err = sv.getVirtualNetworkID(ctx, id)
			if err != nil {
				return err
			}

			err = sv.IntPoolAllocator.DeallocateInt(ctx, VirtualNetworkIDPoolKey, virtualNetworkID)
			if err != nil {
				return err
			}

			response, err = sv.BaseService.DeleteVirtualNetwork(ctx, request)

			// TODO: Delete native/vn-default routing instance
			return err
		})

	return response, err
}

func isSubnetOverlap(subnet1, subnet2 *net.IPNet) bool {
	return subnet1.Contains(subnet2.IP) || subnet2.Contains(subnet1.IP)
}

func appendSubnetIfNoOverlap(subnets []*net.IPNet, subnet *net.IPNet) ([]*net.IPNet, error) {
	for _, existingSubnet := range subnets {
		if isSubnetOverlap(subnet, existingSubnet) {
			return nil, errors.New(
				"overlapping addresses: " + subnet.String() + "," + existingSubnet.String())
		}
	}
	return append(subnets, subnet), nil
}

func mergeSubnetIfNoOverlap(subnets []*net.IPNet, ipamSubnets []*models.IpamSubnetType) ([]*net.IPNet, error) {
	for _, subnet := range ipamSubnets {
		n, err := subnet.Subnet.Net()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse subnet("+subnet.SubnetUUID+")")
		}
		subnets, err = appendSubnetIfNoOverlap(subnets, n)
		if err != nil {
			return nil, err
		}
	}
	return subnets, nil
}

func (sv *ContrailTypeLogicService) processIpamNetworkSubnets(ctx context.Context, virtualNetwork *models.VirtualNetwork) error {
	// check for each ipam references
	ipamReferences := virtualNetwork.NetworkIpamRefs
	// ip subnet must not overlap in a same network.
	// so we create a list of subnet for overlap check.
	subnets := []*net.IPNet{}
	for _, ipamReference := range ipamReferences {
		ipamResponse, err := sv.DataService.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{
			ID: ipamReference.UUID,
		})
		if err != nil {
			return common.ErrorBadRequest("Couldn't process ipam subnets: " + err.Error())
		}
		var ipamSubnets []*models.IpamSubnetType
		ipam := ipamResponse.NetworkIpam
		vnSubnet := ipamReference.GetAttr()
		if ipam.IsFlatSubnet() {
			if !virtualNetwork.IsSupportingL3VPNType() {
				return common.ErrorBadRequest("Couldn't process ipam subnets: flat-subnet is allowed only with l3 network")
			}
			err = vnSubnet.ValidateFlatSubnet()
			if err != nil {
				return common.ErrorBadRequest("Couldn't process ipam subnets: " + err.Error())
			}
			ipamSubnets = ipam.GetIpamSubnets().GetSubnets()
		} else {
			err = vnSubnet.ValidateUserDefined()
			if err != nil {
				return common.ErrorBadRequest("Couldn't process ipam subnets: " + err.Error())
			}
			ipamSubnets = vnSubnet.GetIpamSubnets()
		}
		subnets, err = mergeSubnetIfNoOverlap(subnets, ipamSubnets)
		if err != nil {
			return common.ErrorBadRequest("Couldn't process ipam subnets: " + err.Error())
		}
		//TODO: check network subnet quota
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkProviderNetwork(
	ctx context.Context, virtualNetwork *models.VirtualNetwork) error {

	if virtualNetwork.IsProviderNetwork {
		return common.ErrorBadRequest("Non provider VN (" + virtualNetwork.UUID + ") can not be configured with is_provider_network = True")
	}

	// no further checks if not linked to a provider network
	if len(virtualNetwork.VirtualNetworkRefs) == 0 {
		return nil
	}

	// non provider network can connect to only one provider network.
	if len(virtualNetwork.VirtualNetworkRefs) > 1 {
		return common.ErrorBadRequest("Non Provider VN (" + virtualNetwork.UUID + ") can connect to one provider VN but trying to connect to multiple VN")
	}
	refUUID := virtualNetwork.VirtualNetworkRefs[0].UUID
	ok, err := sv.isVirtualNetworkProviderNetwork(ctx, refUUID)
	if err != nil {
		return err
	}
	if !ok {
		return common.ErrorBadRequest("Non Provider VN (" + virtualNetwork.UUID + ") can connect only " +
			"to one provider VN but not (" + refUUID + ")")
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkNetworkSupportBGPTypes(
	ctx context.Context, virtualNetwork *models.VirtualNetwork) error {
	if virtualNetwork.IsSupportingAnyVPNType() {
		return nil
	}
	for _, ref := range virtualNetwork.BGPVPNRefs {
		vpnType, err := sv.getBGPVPNType(ctx, ref.UUID)
		if err != nil {
			return err
		}
		if vpnType != virtualNetwork.VirtualNetworkProperties.ForwardingMode {
			return common.ErrorBadRequest("BGP types check failed: cannot associate bgpvpn type '" +
				vpnType + "' with a virtual network in forwarding mode " +
				virtualNetwork.VirtualNetworkProperties.ForwardingMode)
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkBGPVPNRefs(
	ctx context.Context, virtualNetwork *models.VirtualNetwork) error {
	if len(virtualNetwork.BGPVPNRefs) == 0 {
		return nil
	}

	if len(virtualNetwork.LogicalRouterRefs) == 0 {
		return nil
	}

	for _, logicalRouterUUID := range virtualNetwork.LogicalRouterRefs {
		refUUID := logicalRouterUUID.UUID
		logicalRouterResponse, err := sv.DataService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{ID: refUUID})
		if err != nil {
			return err
		}
		logicalRouter := logicalRouterResponse.GetLogicalRouter()
		vpnUUIDs := logicalRouter.BGPVPNRefs
		if len(vpnUUIDs) > 0 {
			vpnUUIDStrings := []string{}
			for _, vpnUUID := range vpnUUIDs {
				vpnUUIDStrings = append(vpnUUIDStrings, vpnUUID.UUID)
			}
			return common.ErrorBadRequest("BGP VPN check failed: network " +
				virtualNetwork.UUID +
				"is linked to a logical router which is associated to bgpvpn(s):" + strings.Join(vpnUUIDStrings, ","))
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) getVirtualNetworkID(ctx context.Context, id string) (int64, error) {
	getResponse, err := sv.DataService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: id,
	})
	if err != nil {
		return 0, err
	}
	return getResponse.VirtualNetwork.VirtualNetworkNetworkID, nil
}

func (sv *ContrailTypeLogicService) getBGPVPNType(ctx context.Context, id string) (string, error) {
	getResponse, err := sv.DataService.GetBGPVPN(ctx, &services.GetBGPVPNRequest{
		ID: id,
	})
	if err != nil {
		return "", err
	}
	return getResponse.BGPVPN.BGPVPNType, nil
}

func (sv *ContrailTypeLogicService) isVirtualNetworkProviderNetwork(ctx context.Context, id string) (bool, error) {
	getResponse, err := sv.DataService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: id,
	})
	if err != nil {
		return false, err
	}
	return getResponse.VirtualNetwork.IsProviderNetwork, nil
}
