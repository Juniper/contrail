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
	"github.com/Juniper/contrail/pkg/types/ipam"
)

//CreateVirtualNetwork do pre check and post setup for virtual network.
func (sv *ContrailTypeLogicService) CreateVirtualNetwork(
	ctx context.Context,
	request *services.CreateVirtualNetworkRequest) (response *services.CreateVirtualNetworkResponse, err error) {
	virtualNetwork := request.VirtualNetwork

	if err = sv.prevalidateVirtualNetwork(virtualNetwork); err != nil {
		return nil, err
	}

	// neutron <-> vnc sharing
	virtualNetwork.MakeNeutronCompatible()

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
			err = sv.allocateVnSubnetsInAddrMgmt(ctx, virtualNetwork.GetSubnets())
			if err != nil {
				return err
			}

			response, err = sv.BaseService.CreateVirtualNetwork(ctx, request)
			if err != nil {
				return err
			}

			return sv.createDefaultRoutingInstance(ctx, response.VirtualNetwork)
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
			virtualNetworkResponse, err = sv.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
				ID: requestedVN.UUID,
			})
			if err != nil {
				return common.ErrorBadRequestf("couldn't get virtual network (%v) for an update: %v", requestedVN.UUID, err)
			}

			currentVN := virtualNetworkResponse.GetVirtualNetwork()

			//TODO: check VirtualNetworkID
			//TODO: check changes in virtual_network_properties
			//      we need to read ipam_refs from db and for any ipam
			//      if subnet_method is flat-subnet, network_mode should be l3

			err = sv.processIpamNetworkSubnets(ctx, requestedVN)
			if err != nil {
				return err
			}
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
			err = sv.updateVnSubnetsInAddrMgmt(ctx, currentVN, requestedVN)
			if err != nil {
				return err
			}

			response, err = sv.Next().UpdateVirtualNetwork(ctx, request)

			//TODO: update native/default-VN routing instance
			return err
		})

	return response, err
}

// DeleteVirtualNetwork do pre/post check/teardown for delete network.
func (sv *ContrailTypeLogicService) DeleteVirtualNetwork(
	ctx context.Context,
	request *services.DeleteVirtualNetworkRequest) (response *services.DeleteVirtualNetworkResponse, err error) {
	uuid := request.ID

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var virtualNetworkResponse *services.GetVirtualNetworkResponse
			virtualNetworkResponse, err = sv.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
				ID: uuid,
			})
			if err != nil {
				return common.ErrorBadRequestf("couldn't get virtual network (%v) for a delete: %v", uuid, err)
			}
			vn := virtualNetworkResponse.GetVirtualNetwork()

			// Deallocate virtual network ID
			err = sv.IntPoolAllocator.DeallocateInt(ctx,
				VirtualNetworkIDPoolKey, vn.VirtualNetworkNetworkID)
			if err != nil {
				return common.ErrorBadRequestf("couldn't deallocate virtual network (%v) id(%v): %v",
					uuid, vn.VirtualNetworkNetworkID, err)
			}

			err = sv.deleteDefaultRoutingInstance(ctx, vn)
			if err != nil {
				return err
			}

			err = sv.deallocateVnSubnetsInAddrMgmt(ctx, vn, vn.GetSubnets())
			if err != nil {
				return common.ErrorBadRequestf("couldn't remove virtual network subnet objects: %v", err)
			}

			response, err = sv.BaseService.DeleteVirtualNetwork(ctx, request)

			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) prevalidateVirtualNetwork(vn *models.VirtualNetwork) error {
	// check if multiple policy service chain supported
	if !vn.IsValidMultiPolicyServiceChainConfig() {
		return common.ErrorBadRequest(
			"multi policy service chains are not supported, with both import export external route targets")
	}

	// Does not authorize to set the virtual network ID as it's allocated
	// by the vnc server
	return sv.checkVirtualNetworkID(nil, vn, nil)
}

func (sv *ContrailTypeLogicService) createDefaultRoutingInstance(
	ctx context.Context, vn *models.VirtualNetwork,
) error {
	_, err := sv.WriteService.CreateRoutingInstance(ctx, &services.CreateRoutingInstanceRequest{
		RoutingInstance: &models.RoutingInstance{
			Name:                      vn.Name,
			ParentUUID:                vn.UUID,
			RoutingInstanceIsDefault:  true,
			RoutingInstanceFabricSnat: vn.FabricSnat,
		},
	})

	if err != nil {
		return errors.Wrap(err, "could not create default routing instance for VN")
	}

	return nil
}

func (sv *ContrailTypeLogicService) deleteDefaultRoutingInstance(
	ctx context.Context, vn *models.VirtualNetwork,
) error {
	// Delete native/VN-default routing instance if it hasn't been deleted by the user
	for _, ri := range vn.RoutingInstances {
		if !ri.GetRoutingInstanceIsDefault() {
			continue
		}

		// TODO: delete children of the default routing instance
		_, err := sv.WriteService.DeleteRoutingInstance(ctx, &services.DeleteRoutingInstanceRequest{
			ID: ri.UUID,
		})

		if err != nil {
			return errors.Wrapf(err,
				"could not delete default routing instance (uuid: %v) for VN (%v)", ri.UUID, vn.UUID)
		}

		return nil
	}

	return nil
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
		ipamResponse, err := sv.ReadService.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{
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

func (sv *ContrailTypeLogicService) allocateVnSubnet(
	ctx context.Context, vnSubnet *models.IpamSubnetType) error {

	subnetAlreadyCreated, err := sv.AddressManager.CheckIfIpamSubnetExists(ctx, vnSubnet.SubnetUUID)
	if err != nil {
		return common.ErrorBadRequestf("couldn't allocate ipam subnet %v: %v", vnSubnet.SubnetUUID, err)
	}

	if subnetAlreadyCreated {
		return nil
	}

	vnSubnet.SubnetUUID, err = sv.AddressManager.CreateIpamSubnet(ctx, &ipam.CreateIpamSubnetRequest{
		IpamSubnet: vnSubnet,
	})

	if err != nil {
		return common.ErrorBadRequestf("couldn't allocate ipam subnet %v: %v", vnSubnet.SubnetUUID, err)
	}

	return nil
}

func (sv *ContrailTypeLogicService) deallocateVnSubnet(
	ctx context.Context, vnSubnet *models.IpamSubnetType) (err error) {
	err = sv.AddressManager.DeleteIpamSubnet(ctx, &ipam.DeleteIpamSubnetRequest{
		SubnetUUID: vnSubnet.GetSubnetUUID(),
	})

	if err != nil {
		return common.ErrorBadRequestf("couldn't deallocate ipam subnet %v: %v", vnSubnet.SubnetUUID, err)
	}

	return nil
}

func (sv *ContrailTypeLogicService) updateVnSubnetsInAddrMgmt(
	ctx context.Context, currentVN *models.VirtualNetwork, requestedVN *models.VirtualNetwork,
) error {
	vnSubnetsToDelete := models.IpamSubnetsSubtract(currentVN.GetSubnets(), requestedVN.GetSubnets())
	err := sv.deallocateVnSubnetsInAddrMgmt(ctx, currentVN, vnSubnetsToDelete)
	if err != nil {
		return err
	}

	// TODO: update existing subnets
	return sv.allocateVnSubnetsInAddrMgmt(ctx, requestedVN.GetSubnets())
}

func (sv *ContrailTypeLogicService) allocateVnSubnetsInAddrMgmt(
	ctx context.Context, vnSubnets []*models.IpamSubnetType,
) error {
	for _, vnSubnet := range vnSubnets {
		err := sv.allocateVnSubnet(ctx, vnSubnet)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) deallocateVnSubnetsInAddrMgmt(
	ctx context.Context, vn *models.VirtualNetwork, vnSubnets []*models.IpamSubnetType,
) error {

	err := sv.canSubnetsBeDeleted(ctx, vn, vnSubnets)
	if err != nil {
		return common.ErrorConflictf("subnets cannot be deleted: %v", err)
	}

	for _, vnSubnet := range vnSubnets {
		err := sv.deallocateVnSubnet(ctx, vnSubnet)
		if err != nil {
			return err
		}
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

	if !common.ContainsString(fieldMask.GetPaths(), models.VirtualNetworkFieldVirtualNetworkNetworkID) {
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
				requestedVN.UUID, models.VirtualNetworkFieldIsProviderNetwork)
		}
		return nil
	}

	if !common.ContainsString(fieldMask.GetPaths(), models.VirtualNetworkFieldIsProviderNetwork) {
		return nil
	}

	if currentVN.IsProviderNetwork != requestedVN.IsProviderNetwork {
		return common.ErrorBadRequestf("update %v property of VN (%v) is not allowed",
			models.VirtualNetworkFieldIsProviderNetwork, requestedVN.UUID)
	}

	return nil
}

func (sv *ContrailTypeLogicService) isProviderNetwork(
	currentVN *models.VirtualNetwork, requestedVN *models.VirtualNetwork,
	fieldMask *protobuf.FieldMask) bool {
	isProviderNetwork := requestedVN.IsProviderNetwork
	if currentVN != nil && !common.ContainsString(fieldMask.GetPaths(), models.VirtualNetworkFieldIsProviderNetwork) {
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
		logicalRouterResponse, err := sv.ReadService.GetLogicalRouter(
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

func (sv *ContrailTypeLogicService) getBGPVPNType(ctx context.Context, id string) (string, error) {
	getResponse, err := sv.ReadService.GetBGPVPN(ctx, &services.GetBGPVPNRequest{
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
		getResponse, err := sv.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
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

func (sv *ContrailTypeLogicService) canSubnetsBeDeleted(
	ctx context.Context,
	vn *models.VirtualNetwork,
	subnetsSet []*models.IpamSubnetType,
) error {
	for _, instanceIPBackRef := range vn.GetInstanceIPBackRefs() {
		err := sv.checkInstanceIP(ctx, instanceIPBackRef, subnetsSet)
		if err != nil {
			return err
		}
	}

	for _, floatingIPPoolRef := range vn.GetFloatingIPPools() {
		err := sv.checkFloatingIPPool(ctx, floatingIPPoolRef, subnetsSet)
		if err != nil {
			return err
		}
	}

	for _, aliasIPPoolRef := range vn.GetAliasIPPools() {
		err := sv.checkAliasIPPool(ctx, aliasIPPoolRef, subnetsSet)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkInstanceIP(
	ctx context.Context,
	instanceIPBackRef *models.InstanceIP,
	subnetsSet []*models.IpamSubnetType,
) error {
	instanceIP, err := sv.ReadService.GetInstanceIP(ctx, &services.GetInstanceIPRequest{
		ID: instanceIPBackRef.GetUUID()})
	if err != nil {
		return err
	}
	return checkIfSubnetsSetIncludeIP(subnetsSet, instanceIP.GetInstanceIP().GetInstanceIPAddress())
}

func (sv *ContrailTypeLogicService) checkFloatingIPPool(
	ctx context.Context,
	floatingIPPoolRef *models.FloatingIPPool,
	subnetsSet []*models.IpamSubnetType,
) error {
	floatingIPPool, err := sv.ReadService.GetFloatingIPPool(ctx, &services.GetFloatingIPPoolRequest{
		ID: floatingIPPoolRef.GetUUID()})
	if err != nil {
		return err
	}
	for _, floatingIP := range floatingIPPool.GetFloatingIPPool().GetFloatingIPs() {
		err = checkIfSubnetsSetIncludeIP(subnetsSet, floatingIP.GetFloatingIPAddress())
		if err != nil {
			return err
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkAliasIPPool(
	ctx context.Context,
	aliasIPPoolRef *models.AliasIPPool,
	subnetsSet []*models.IpamSubnetType,
) error {
	aliasIPPool, err := sv.ReadService.GetAliasIPPool(ctx, &services.GetAliasIPPoolRequest{
		ID: aliasIPPoolRef.GetUUID()})
	if err != nil {
		return err
	}
	for _, aliasIP := range aliasIPPool.GetAliasIPPool().GetAliasIPs() {
		err = checkIfSubnetsSetIncludeIP(subnetsSet, aliasIP.GetAliasIPAddress())
		if err != nil {
			return err
		}
	}
	return nil
}

func checkIfSubnetsSetIncludeIP(
	subnetsSet []*models.IpamSubnetType,
	ipString string,
) error {
	for _, ipamSubnet := range subnetsSet {
		ip := net.ParseIP(ipString)
		if ip == nil {
			return errors.Errorf("invalid address: " + ipString)
		}

		contains, err := ipamSubnet.Contains(ip)
		if err != nil {
			return err
		}

		if contains {
			return errors.Errorf("subnet %s contains address %s", ipamSubnet.SubnetUUID, ipString)
		}
	}
	return nil
}
