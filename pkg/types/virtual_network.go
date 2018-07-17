package types

import (
	"context"
	"net"
	"strings"

	protobuf "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

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
	err = sv.checkVirtualNetworkID(nil, virtualNetwork, nil)
	if err != nil {
		return nil, err
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
			err = sv.checkIsProviderNetwork(nil, virtualNetwork, nil)
			if err != nil {
				return err
			}
			err = sv.checkProviderNetwork(ctx, nil, virtualNetwork, nil)
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

// UpdateVirtualNetwork do pre check for virtual network update.
func (sv *ContrailTypeLogicService) UpdateVirtualNetwork(
	ctx context.Context,
	request *services.UpdateVirtualNetworkRequest) (response *services.UpdateVirtualNetworkResponse, err error) {

	requestedVN := request.VirtualNetwork

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var virtualNetworkResponse *services.GetVirtualNetworkResponse
			virtualNetworkResponse, err = sv.DataService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
				ID: requestedVN.UUID,
			})
			if err != nil {
				return common.ErrorBadRequestf("couldn't get virtual network (%v) for an update: %v", requestedVN.UUID, err)
			}

			currentVN := virtualNetworkResponse.GetVirtualNetwork()

			//TODO: check VirtualNetworkID
			//TODO: process ipam network subnets
			//TODO: check route target (depends on global system config)
			//TODO: check provider details

			err = sv.checkIsProviderNetwork(currentVN, requestedVN, &request.FieldMask)
			if err != nil {
				return err
			}
			err = sv.checkProviderNetwork(ctx, currentVN, requestedVN, &request.FieldMask)
			if err != nil {
				return err
			}
			//TODO: check network support BGP types
			//TODO: check BGPVPN Refs

			response, err = sv.Next().UpdateVirtualNetwork(ctx, request)
			return err
		})

	return response, err
}

// DeleteVirtualNetwork do pre check for delete network.
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

func (sv *ContrailTypeLogicService) checkVirtualNetworkID(
	currentVN *models.VirtualNetwork, requestedVN *models.VirtualNetwork, fieldMask *protobuf.FieldMask) error {

	if currentVN == nil {
		if requestedVN.HasVirtualNetworkNetworkID() {
			return common.ErrorForbidden("cannot set the virtual network ID, it's allocated by the server")
		}
		return nil
	}

	if !common.ContainsString(fieldMask.GetPaths(), models.VirtualNetworkPropertyIDVirtualNetworkNetworkID) {
		return nil
	}

	if currentVN.GetVirtualNetworkNetworkID() != requestedVN.GetVirtualNetworkNetworkID() {
		return common.ErrorForbidden("cannot update the virtual network ID, it's allocated by the server")
	}

	return nil
}

func (sv *ContrailTypeLogicService) checkIsProviderNetwork(
	currentVN *models.VirtualNetwork, requestedVN *models.VirtualNetwork,
	fieldMask *protobuf.FieldMask) error {

	if currentVN == nil {
		if requestedVN.IsProviderNetwork {
			return common.ErrorBadRequestf("non-provider VN (%v) can not be configured with %v = True",
				requestedVN.UUID, models.VirtualNetworkPropertyIDIsProviderNetwork)
		}
		return nil
	}

	if !common.ContainsString(fieldMask.GetPaths(), models.VirtualNetworkPropertyIDIsProviderNetwork) {
		return nil
	}

	if currentVN.IsProviderNetwork != requestedVN.IsProviderNetwork {
		return common.ErrorBadRequestf("update %v property of VN (%v) is not allowed",
			models.VirtualNetworkPropertyIDIsProviderNetwork, requestedVN.UUID)
	}

	return nil
}

func (sv *ContrailTypeLogicService) isProviderNetwork(
	currentVN *models.VirtualNetwork, requestedVN *models.VirtualNetwork,
	fieldMask *protobuf.FieldMask) bool {
	isProviderNetwork := requestedVN.IsProviderNetwork
	if currentVN != nil && !common.ContainsString(fieldMask.GetPaths(), models.VirtualNetworkPropertyIDIsProviderNetwork) {
		isProviderNetwork = currentVN.IsProviderNetwork
	}
	return isProviderNetwork
}

func (sv *ContrailTypeLogicService) checkProviderNetwork(
	ctx context.Context, currentVN *models.VirtualNetwork, requestedVN *models.VirtualNetwork,
	fieldMask *protobuf.FieldMask) error {

	// No further checks if not linked to a provider network.
	if len(requestedVN.VirtualNetworkRefs) == 0 {
		return nil
	}

	isProviderNetwork := sv.isProviderNetwork(currentVN, requestedVN, fieldMask)

	// Non-provider network can connect to only one provider network.
	if !isProviderNetwork && len(requestedVN.VirtualNetworkRefs) > 1 {
		return common.ErrorBadRequestf(
			"non-provider VN (%v) can be connected to one provider VN but trying to connect to multiple VN: %v",
			requestedVN.UUID, func() (vnUUIDs []string) {
				for _, vnRef := range requestedVN.VirtualNetworkRefs {
					vnUUIDs = append(vnUUIDs, vnRef.GetUUID())
				}
				return vnUUIDs
			}())
	}

	linkedProviderVirtualNetworkUUIDs, err := sv.getLinkedProviderVirtualNetworks(ctx, requestedVN)
	if err != nil {
		return err
	}

	// Provider VN can not connect to another provider VN.
	if isProviderNetwork && len(linkedProviderVirtualNetworkUUIDs) > 0 {
		return common.ErrorBadRequestf("provider VN (%v) cannot be connected to another provider VN (%v)",
			requestedVN.UUID, linkedProviderVirtualNetworkUUIDs)
	}

	// Non-provider network can connect to only one provider network.
	if !isProviderNetwork && len(linkedProviderVirtualNetworkUUIDs) != 1 {
		return common.ErrorBadRequestf("non-provider VN (%v) can be connected to one provider VN but not to (%v)",
			requestedVN.UUID, linkedProviderVirtualNetworkUUIDs)
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
				"bgp types check failed: cannot associate bgpvpn type '%v' "+
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
				"bgp VPN check failed: network %v is linked to a logical router "+
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

func (sv *ContrailTypeLogicService) getLinkedProviderVirtualNetworks(
	ctx context.Context, virtualNetwork *models.VirtualNetwork) ([]string, error) {

	var linkedProviderVirtualNetworkUUIDs []string
	for _, vnRef := range virtualNetwork.GetVirtualNetworkRefs() {
		getResponse, err := sv.DataService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
			ID: vnRef.GetUUID(),
		})
		if err != nil {
			return nil, err
		}
		linkedVirtualNetwork := getResponse.GetVirtualNetwork()
		if linkedVirtualNetwork.GetIsProviderNetwork() {
			linkedProviderVirtualNetworkUUIDs = append(linkedProviderVirtualNetworkUUIDs, linkedVirtualNetwork.GetUUID())
		}
	}
	return linkedProviderVirtualNetworkUUIDs, nil
}
