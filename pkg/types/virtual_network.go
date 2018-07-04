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
		return nil, common.ErrorBadRequest(
			"multi policy service chains are not supported, with both import export external route targets")
	}
	//  neutron <-> vnc sharing
	virtualNetwork.MakeNeutronCompatible()
	// Does not authorize to set the virtual network ID as it's allocated
	// by the vnc server
	if virtualNetwork.HasVirtualNetworkNetworkID() {
		return nil, common.ErrorForbidden(
			"cannot set the virtual network ID, it's allocated by the server")
	}

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
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

func mergeIpamSubnetsIfNoOverlap(subnets []*net.IPNet, ipamSubnets []*models.IpamSubnetType) ([]*net.IPNet, error) {
	for _, subnet := range ipamSubnets {
		n, err := subnet.Subnet.Net()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse subnet(%v)", subnet.SubnetUUID)
		}
		subnets, err = appendSubnetIfNoOverlap(subnets, n)
		if err != nil {
			return nil, err
		}
	}
	return subnets, nil
}

func validateIpamSubnets(
	virtualNetwork *models.VirtualNetwork, ipam *models.NetworkIpam, vnSubnets *models.VnSubnetsType,
) error {
	if !ipam.IsFlatSubnet() {
		return vnSubnets.ValidateUserDefined()
	}

	if !virtualNetwork.IsSupportingL3VPNType() {
		return errors.New("flat-subnet is allowed only with l3 network")
	}

	return vnSubnets.ValidateFlatSubnet()
}

func extractIpamSubnets(ipam *models.NetworkIpam, vnSubnets *models.VnSubnetsType) []*models.IpamSubnetType {
	if ipam.IsFlatSubnet() {
		return ipam.GetIpamSubnets().GetSubnets()
	}

	return vnSubnets.GetIpamSubnets()
}

func (sv *ContrailTypeLogicService) processIpamNetworkSubnets(
	ctx context.Context, virtualNetwork *models.VirtualNetwork,
) error {
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
			return common.ErrorBadRequestf("getting referenced network IPAM with UUID %s failed: %v",
				ipamReference.UUID, err)
		}

		ipam := ipamResponse.NetworkIpam
		vnSubnet := ipamReference.GetAttr()
		err = validateIpamSubnets(virtualNetwork, ipam, vnSubnet)
		if err != nil {
			return common.ErrorBadRequestf(
				"validation of IPAM subnets of referenced network IPAM with UUID %s failed: %v",
				ipamReference.UUID, err)
		}

		ipamSubnets := extractIpamSubnets(ipam, vnSubnet)
		subnets, err = mergeIpamSubnetsIfNoOverlap(subnets, ipamSubnets)
		if err != nil {
			return common.ErrorBadRequestf(
				"merging of IPAM subnets of referenced network IPAM with UUID %s failed: %v",
				ipamReference.UUID, err)
		}
		//TODO: check network subnet quota
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkProviderNetwork(
	ctx context.Context, virtualNetwork *models.VirtualNetwork) error {

	if virtualNetwork.IsProviderNetwork {
		return common.ErrorBadRequestf(
			"non provider VN (%v) can not be configured with is_provider_network = True", virtualNetwork.UUID)
	}

	// no further checks if not linked to a provider network
	if len(virtualNetwork.VirtualNetworkRefs) == 0 {
		return nil
	}

	// non provider network can connect to only one provider network.
	if len(virtualNetwork.VirtualNetworkRefs) > 1 {
		return common.ErrorBadRequestf(
			"non Provider VN (%v) can connect to one provider VN but trying to connect to multiple VN",
			virtualNetwork.UUID,
		)
	}
	refUUID := virtualNetwork.VirtualNetworkRefs[0].UUID
	ok, err := sv.isVirtualNetworkProviderNetwork(ctx, refUUID)
	if err != nil {
		return err
	}
	if !ok {
		return common.ErrorBadRequestf("non Provider VN (%v) can connect only "+
			"to one provider VN but not (%v)", virtualNetwork.UUID, refUUID)
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
			return common.ErrorBadRequestf(
				"BGP types check failed: cannot associate bgpvpn type '%v' "+
					"with a virtual network in forwarding mode '%v'",
				vpnType,
				virtualNetwork.VirtualNetworkProperties.ForwardingMode,
			)
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
		logicalRouterResponse, err := sv.DataService.GetLogicalRouter(
			ctx,
			&services.GetLogicalRouterRequest{ID: refUUID},
		)
		if err != nil {
			return err
		}
		logicalRouter := logicalRouterResponse.GetLogicalRouter()
		vpnRefs := logicalRouter.BGPVPNRefs
		if len(vpnRefs) > 0 {
			vpnUUIDs := []string{}
			for _, vpnRef := range vpnRefs {
				vpnUUIDs = append(vpnUUIDs, vpnRef.UUID)
			}
			return common.ErrorBadRequestf(
				"BGP VPN check failed: network %v is linked to a logical router "+
					"which is associated to bgpvpn(s) %v",
				virtualNetwork.UUID,
				strings.Join(vpnUUIDs, ","),
			)
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
