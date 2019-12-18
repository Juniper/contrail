package types

import (
	"context"
	"net"
	"strings"

	protobuf "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/db"
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
			if err = sv.reserveVirtualNetworkVxLANID(ctx, virtualNetwork); err != nil {
				return err
			}

			virtualNetwork.VirtualNetworkNetworkID, err = sv.IntPoolAllocator.AllocateInt(
				ctx,
				VirtualNetworkIDPoolKey,
				db.EmptyIntOwner,
			)
			if err != nil {
				return err
			}

			err = sv.validateVirtualNetworkOnCreate(ctx, virtualNetwork)
			if err != nil {
				return err
			}

			err = sv.allocateRefSubnets(ctx, virtualNetwork)
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
				return errutil.ErrorBadRequestf("couldn't get virtual network (%v) for an update: %v", requestedVN.UUID, err)
			}

			currentVN := virtualNetworkResponse.GetVirtualNetwork()

			if err = checkMultiPolicyServiceChainConfig(request, *currentVN); err != nil {
				return err
			}

			//TODO: check VirtualNetworkID

			err = sv.processVxlanIDUpdate(ctx, currentVN, requestedVN, &request.FieldMask)
			if err != nil {
				return err
			}

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
			err = sv.updateVnSubnetsInAddrMgmt(ctx, currentVN, requestedVN, &request.FieldMask)
			if err != nil {
				return err
			}

			response, err = sv.BaseService.UpdateVirtualNetwork(ctx, request)

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
				return errutil.ErrorNotFoundf("couldn't get virtual network (%v) for a delete: %v", uuid, err)
			}
			vn := virtualNetworkResponse.GetVirtualNetwork()

			// Deallocate virtual network ID
			err = sv.IntPoolAllocator.DeallocateInt(ctx,
				VirtualNetworkIDPoolKey, vn.VirtualNetworkNetworkID)
			if err != nil {
				return errutil.ErrorBadRequestf("couldn't deallocate virtual network (%v) id(%v): %v",
					uuid, vn.VirtualNetworkNetworkID, err)
			}

			err = sv.deleteDefaultRoutingInstance(ctx, vn)
			if err != nil {
				return err
			}

			var subnetsToDelete *models.IpamSubnets
			subnetsToDelete, err = sv.getRefSubnets(ctx, vn)
			if err != nil {
				return err
			}

			err = sv.deallocateVnSubnetsInAddrMgmt(ctx, vn, subnetsToDelete)
			if err != nil {
				return errutil.ErrorBadRequestf("couldn't remove virtual network subnet objects: %v", err)
			}

			if err = sv.deallocateVirtualNetworkVxLANID(ctx, vn); err != nil {
				return err
			}

			response, err = sv.BaseService.DeleteVirtualNetwork(ctx, request)

			return err
		})

	return response, err
}

// checkMultiPolicyServiceChainConfig validates if in case of
// enabled Multi Policy Service Chain Support resource still doesn't contain
// the same import and export route target.
func checkMultiPolicyServiceChainConfig(
	request *services.UpdateVirtualNetworkRequest, currentVN models.VirtualNetwork,
) error {
	updateVN := request.GetVirtualNetwork()
	fm := request.GetFieldMask()

	if basemodels.FieldMaskContains(&fm, models.VirtualNetworkFieldMultiPolicyServiceChainsEnabled) {
		currentVN.MultiPolicyServiceChainsEnabled = updateVN.MultiPolicyServiceChainsEnabled
	}
	if !currentVN.MultiPolicyServiceChainsEnabled {
		return nil
	}
	if basemodels.FieldMaskContains(&fm, models.VirtualNetworkFieldImportRouteTargetList) {
		currentVN.ImportRouteTargetList = updateVN.ImportRouteTargetList
	}
	if basemodels.FieldMaskContains(&fm, models.VirtualNetworkFieldExportRouteTargetList) {
		currentVN.ExportRouteTargetList = updateVN.ExportRouteTargetList
	}
	if basemodels.FieldMaskContains(&fm, models.VirtualNetworkFieldRouteTargetList) {
		currentVN.RouteTargetList = updateVN.RouteTargetList
	}

	return currentVN.CheckMultiPolicyServiceChainConfig()
}

func (sv *ContrailTypeLogicService) prevalidateVirtualNetwork(vn *models.VirtualNetwork) error {
	// check if multiple policy service chain supported
	if err := vn.CheckMultiPolicyServiceChainConfig(); err != nil {
		return err
	}

	// Does not authorize to set the virtual network ID as it's allocated
	// by the vnc server
	return sv.checkVirtualNetworkID(nil, vn, nil)
}

func (sv *ContrailTypeLogicService) validateVirtualNetworkOnCreate(
	ctx context.Context, virtualNetwork *models.VirtualNetwork,
) error {
	if err := sv.processIpamNetworkSubnets(ctx, virtualNetwork); err != nil {
		return err
	}
	//TODO: check route target (depends on global system config)
	//TODO: check provider details
	if err := sv.checkIsProviderNetwork(nil, virtualNetwork, nil); err != nil {
		return err
	}
	if err := sv.checkProviderNetwork(ctx, nil, virtualNetwork, nil); err != nil {
		return err
	}
	if err := sv.checkNetworkSupportBGPTypes(ctx, virtualNetwork); err != nil {
		return err
	}
	if err := sv.checkBGPVPNRefs(ctx, virtualNetwork); err != nil {
		return err
	}
	return nil
}

func (sv *ContrailTypeLogicService) createDefaultRoutingInstance(
	ctx context.Context, vn *models.VirtualNetwork,
) error {
	_, err := sv.WriteService.CreateRoutingInstance(
		ctx, &services.CreateRoutingInstanceRequest{
			RoutingInstance: vn.MakeDefaultRoutingInstance(),
		},
	)

	if err != nil {
		return errors.Wrap(err, "could not create default routing instance for VN")
	}

	return nil
}

func (sv *ContrailTypeLogicService) deleteDefaultRoutingInstance(
	ctx context.Context, vn *models.VirtualNetwork,
) error {
	// Delete native/VN-default routing instance if it hasn't been deleted by the user
	if ri := vn.GetDefaultRoutingInstance(); ri != nil {
		// TODO: delete children of the default routing instance
		_, err := sv.WriteService.DeleteRoutingInstance(
			ctx, &services.DeleteRoutingInstanceRequest{
				ID: ri.UUID,
			},
		)

		if err != nil {
			return errors.Wrapf(err,
				"could not delete default routing instance (uuid: %v) for VN (%v)", ri.UUID, vn.UUID)
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) processVxlanIDUpdate(
	ctx context.Context,
	currentVN *models.VirtualNetwork, requestedVN *models.VirtualNetwork,
	fieldMask *protobuf.FieldMask,
) error {
	if !basemodels.FieldMaskContains(
		fieldMask,
		models.VirtualNetworkFieldVirtualNetworkProperties,
		models.VirtualNetworkTypeFieldVxlanNetworkIdentifier,
	) {
		return nil
	}

	currentVxLANID := currentVN.GetVirtualNetworkProperties().GetVxlanNetworkIdentifier()
	requestedVxLANID := requestedVN.GetVirtualNetworkProperties().GetVxlanNetworkIdentifier()
	if currentVxLANID == requestedVxLANID {
		return nil
	}

	if err := sv.deallocateVirtualNetworkVxLANID(ctx, currentVN); err != nil {
		return err
	}

	requestedVN.FQName = currentVN.FQName
	if err := sv.reserveVirtualNetworkVxLANID(ctx, requestedVN); err != nil {
		return err
	}

	return nil
}

func (sv *ContrailTypeLogicService) reserveVirtualNetworkVxLANID(
	ctx context.Context, virtualNetwork *models.VirtualNetwork,
) error {
	vxLANID := virtualNetwork.GetVirtualNetworkProperties().GetVxlanNetworkIdentifier()
	if vxLANID == 0 {
		return nil
	}

	err := sv.IntPoolAllocator.SetInt(ctx, VirtualNetworkIDPoolKey, vxLANID, virtualNetwork.VxLANIntOwner())
	if err != nil {
		return errutil.ErrorBadRequestf("cannot allocate provided vxlan identifier(%v): %v", vxLANID, err)
	}
	return nil
}

func (sv *ContrailTypeLogicService) deallocateVirtualNetworkVxLANID(
	ctx context.Context, virtualNetwork *models.VirtualNetwork,
) error {
	vxLANID := virtualNetwork.GetVirtualNetworkProperties().GetVxlanNetworkIdentifier()
	if vxLANID == 0 {
		return nil
	}

	err := sv.IntPoolAllocator.DeallocateInt(ctx, VirtualNetworkIDPoolKey, vxLANID)
	if err != nil {
		return errutil.ErrorBadRequestf("couldn't deallocate vxlan network id(%v): %v", vxLANID, err)
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
			return errutil.ErrorBadRequestf("getting referenced network IPAM with UUID %s failed: %v",
				ipamReference.UUID, err)
		}

		ipam := ipamResponse.NetworkIpam
		vnSubnet := ipamReference.GetAttr()
		err = validateIpamSubnets(virtualNetwork, ipam, vnSubnet)
		if err != nil {
			return errutil.ErrorBadRequestf(
				"validation of IPAM subnets of referenced network IPAM with UUID %s failed: %v",
				ipamReference.UUID, err)
		}

		ipamSubnets := extractIpamSubnets(ipam, vnSubnet)
		subnets, err = mergeIpamSubnetsIfNoOverlap(subnets, ipamSubnets)
		if err != nil {
			return errutil.ErrorBadRequestf(
				"merging of IPAM subnets of referenced network IPAM with UUID %s failed: %v",
				ipamReference.UUID, err)
		}
		//TODO: check network subnet quota
	}
	return nil
}

func (sv *ContrailTypeLogicService) allocateVnSubnet(
	ctx context.Context, vnSubnet *models.IpamSubnetType) (subnetUUID string, err error) {
	var subnetAlreadyCreated bool
	subnetAlreadyCreated, err = sv.AddressManager.CheckIfIpamSubnetExists(ctx, vnSubnet.SubnetUUID)
	if err != nil {
		return "", errutil.ErrorBadRequestf("couldn't check if ipam subnet with UUID %v exists: %v", vnSubnet.SubnetUUID, err)
	}

	if subnetAlreadyCreated {
		return vnSubnet.SubnetUUID, nil
	}

	subnetUUID, err = sv.AddressManager.CreateIpamSubnet(ctx, &ipam.CreateIpamSubnetRequest{
		IpamSubnet: vnSubnet,
	})

	if err != nil {
		return "", errutil.ErrorBadRequestf("couldn't allocate ipam subnet %v: %v", vnSubnet.SubnetUUID, err)
	}

	return subnetUUID, nil
}

func (sv *ContrailTypeLogicService) deallocateVnSubnet(
	ctx context.Context, subnetUUID string) (err error) {
	err = sv.AddressManager.DeleteIpamSubnet(ctx, &ipam.DeleteIpamSubnetRequest{
		SubnetUUID: subnetUUID,
	})

	if err != nil {
		return errutil.ErrorBadRequestf("couldn't deallocate ipam subnet %v: %v", subnetUUID, err)
	}

	return nil
}

func (sv *ContrailTypeLogicService) updateVnSubnetsInAddrMgmt(
	ctx context.Context,
	currentVN *models.VirtualNetwork,
	requestedVN *models.VirtualNetwork,
	fieldMask *protobuf.FieldMask,
) error {

	if !basemodels.FieldMaskContains(fieldMask, models.VirtualNetworkFieldNetworkIpamRefs) {
		return nil
	}

	vnSubnetsToDelete := currentVN.GetIpamSubnets().Subtract(requestedVN.GetIpamSubnets())
	err := sv.deallocateVnSubnetsInAddrMgmt(ctx, currentVN, vnSubnetsToDelete)
	if err != nil {
		return err
	}

	// TODO: update existing subnets
	return sv.allocateRefSubnets(ctx, requestedVN)
}

//TODO handle multiple subnet in ipam refs
func (sv *ContrailTypeLogicService) allocateVnSubnetsInAddrMgmt(
	ctx context.Context, ipamRef *models.VirtualNetworkNetworkIpamRef,
) error {
	if ipamRef.GetAttr() != nil {
		if len(ipamRef.GetAttr().GetIpamSubnets()) > 0 {
			linkedIpamSubnets := ipamRef.GetAttr().GetIpamSubnets()
			for _, vnSubnet := range linkedIpamSubnets {
				subnetUUID, err := sv.allocateVnSubnet(ctx, vnSubnet)
				if err != nil {
					return err
				}
				vnSubnet.SubnetUUID = subnetUUID
			}
		} else {
			ipamResponse, err := sv.ReadService.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{
				ID: ipamRef.GetUUID(),
			})
			if err != nil {
				return errutil.ErrorBadRequestf("getting referenced network IPAM with UUID %s failed: %v",
					ipamRef.GetUUID(), err)
			}

			linkedIpamSubnets := ipamResponse.GetNetworkIpam().GetIpamSubnets().GetSubnets()
			for _, vnSubnet := range linkedIpamSubnets {
				subnetUUID, err := sv.allocateVnSubnet(ctx, vnSubnet)
				if err != nil {
					return err
				}
				ipamRef.Attr.IpamSubnets = append(ipamRef.Attr.IpamSubnets, &models.IpamSubnetType{
					SubnetUUID: subnetUUID})
			}
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) allocateRefSubnets(
	ctx context.Context, vn *models.VirtualNetwork,
) error {
	for _, netIpamRef := range vn.GetNetworkIpamRefs() {
		if err := sv.allocateVnSubnetsInAddrMgmt(ctx, netIpamRef); err != nil {
			return err
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) deallocateVnSubnetsInAddrMgmt(
	ctx context.Context, vn *models.VirtualNetwork, vnSubnets *models.IpamSubnets,
) error {
	err := sv.canSubnetsBeDeleted(ctx, vn, vnSubnets)
	if err != nil {
		return errutil.ErrorConflictf("subnets from virtual network %v cannot be deleted: %v", vn.GetUUID(), err)
	}

	for _, subnetUUID := range vn.GetIpamSubnets().UUIDs() {
		err := sv.deallocateVnSubnet(ctx, subnetUUID)
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
			return errutil.ErrorForbidden("cannot set the virtual network ID, it's allocated by the server")
		}
		return nil
	}

	if !format.ContainsString(fieldMask.GetPaths(), models.VirtualNetworkFieldVirtualNetworkNetworkID) {
		return nil
	}

	if currentVN.GetVirtualNetworkNetworkID() != requestedVN.GetVirtualNetworkNetworkID() {
		return errutil.ErrorForbidden("cannot update the virtual network ID, it's allocated by the server")
	}

	return nil
}

func (sv *ContrailTypeLogicService) checkIsProviderNetwork(
	currentVN *models.VirtualNetwork, requestedVN *models.VirtualNetwork,
	fieldMask *protobuf.FieldMask) error {

	if currentVN == nil {
		if requestedVN.IsProviderNetwork {
			return errutil.ErrorBadRequestf("non-provider VN (%v) can not be configured with %v = True",
				requestedVN.UUID, models.VirtualNetworkFieldIsProviderNetwork)
		}
		return nil
	}

	if !format.ContainsString(fieldMask.GetPaths(), models.VirtualNetworkFieldIsProviderNetwork) {
		return nil
	}

	if currentVN.IsProviderNetwork != requestedVN.IsProviderNetwork {
		return errutil.ErrorBadRequestf("update %v property of VN (%v) is not allowed",
			models.VirtualNetworkFieldIsProviderNetwork, requestedVN.UUID)
	}

	return nil
}

func (sv *ContrailTypeLogicService) isProviderNetwork(
	currentVN *models.VirtualNetwork, requestedVN *models.VirtualNetwork,
	fieldMask *protobuf.FieldMask) bool {
	isProviderNetwork := requestedVN.IsProviderNetwork
	if currentVN != nil && !format.ContainsString(fieldMask.GetPaths(), models.VirtualNetworkFieldIsProviderNetwork) {
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
		return errutil.ErrorBadRequestf(
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
		return errutil.ErrorBadRequestf("provider VN (%v) cannot be connected to another provider VN (%v)",
			requestedVN.UUID, linkedProviderVirtualNetworkUUIDs)
	}

	// Non-provider network can connect to only one provider network.
	if !isProviderNetwork && len(linkedProviderVirtualNetworkUUIDs) != 1 {
		return errutil.ErrorBadRequestf("non-provider VN (%v) can be connected to one provider VN but not to (%v)",
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
		if vpnType != virtualNetwork.GetVirtualNetworkProperties().GetForwardingMode() {
			return errutil.ErrorBadRequestf(
				"bgp types check failed: cannot associate bgpvpn type '%v' "+
					"with a virtual network in forwarding mode '%v'",
				vpnType,
				virtualNetwork.GetVirtualNetworkProperties().GetForwardingMode(),
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
			return errutil.ErrorBadRequestf(
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
	vnSubnets *models.IpamSubnets,
) error {

	err := sv.checkIfInstanceIPsAreNotUsedInSubnets(ctx, vn, vnSubnets)
	if err != nil {
		return err
	}

	err = sv.checkIfFloatingIPsAreNotUsedInSubnets(ctx, vn, vnSubnets)
	if err != nil {
		return err
	}

	return sv.checkIfAliasIPsAreNotUsedInSubnets(ctx, vn, vnSubnets)
}

func (sv *ContrailTypeLogicService) checkIfInstanceIPsAreNotUsedInSubnets(
	ctx context.Context,
	vn *models.VirtualNetwork,
	vnSubnets *models.IpamSubnets,
) error {
	for _, instanceIPBackRef := range vn.GetInstanceIPBackRefs() {
		instanceIPRes, err := sv.ReadService.GetInstanceIP(ctx, &services.GetInstanceIPRequest{
			ID: instanceIPBackRef.GetUUID()})
		if err != nil {
			return err
		}
		instanceIP := instanceIPRes.GetInstanceIP()
		contains, err := vnSubnets.Contains(instanceIP.GetInstanceIPAddress())
		if err != nil {
			return err
		}

		if contains {
			return errors.Errorf("cannot delete IP block, ip(%v) is in use", instanceIP.GetInstanceIPAddress())
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) checkIfFloatingIPsAreNotUsedInSubnets(
	ctx context.Context,
	vn *models.VirtualNetwork,
	vnSubnets *models.IpamSubnets,
) error {
	for _, floatingIPPoolRef := range vn.GetFloatingIPPools() {
		floatingIPPoolRes, err := sv.ReadService.GetFloatingIPPool(ctx, &services.GetFloatingIPPoolRequest{
			ID: floatingIPPoolRef.GetUUID()})
		if err != nil {
			return err
		}
		ipAddresses, err := floatingIPPoolRes.GetFloatingIPPool().GetIPsInSubnets(vnSubnets)
		if err != nil {
			return err
		}

		if len(ipAddresses) > 0 {
			return errors.Errorf("cannot delete IP block, floating ip addresses (%v) are in use", ipAddresses)
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) checkIfAliasIPsAreNotUsedInSubnets(
	ctx context.Context,
	vn *models.VirtualNetwork,
	vnSubnets *models.IpamSubnets,
) error {

	for _, aliasIPPoolRef := range vn.GetAliasIPPools() {
		aliasIPPoolRes, err := sv.ReadService.GetAliasIPPool(ctx, &services.GetAliasIPPoolRequest{
			ID: aliasIPPoolRef.GetUUID()})
		if err != nil {
			return err
		}

		ipAddresses, err := aliasIPPoolRes.GetAliasIPPool().GetIPsInSubnets(vnSubnets)
		if err != nil {
			return err
		}

		if len(ipAddresses) > 0 {
			return errors.Errorf("cannot delete IP block, alias ip addresses (%v) are in use", ipAddresses)
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) getRefSubnets(
	ctx context.Context,
	vn *models.VirtualNetwork,
) (*models.IpamSubnets, error) {

	if vn.GetAddressAllocationMethod() == models.UserDefinedSubnetOnly {
		return vn.GetIpamSubnets(), nil
	} else if vn.GetAddressAllocationMethod() == models.FlatSubnetOnly {
		var subnets []*models.IpamSubnetType
		for _, ipamRef := range vn.GetNetworkIpamRefs() {
			ipamResponse, err := sv.ReadService.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{
				ID: ipamRef.GetUUID(),
			})
			if err != nil {
				return nil, errutil.ErrorBadRequestf("getting referenced network IPAM with UUID %s failed: %v",
					ipamRef.GetUUID(), err)
			}

			subnets = append(subnets, ipamResponse.GetNetworkIpam().GetIpamSubnets().GetSubnets()...)
			return &models.IpamSubnets{
				Subnets: subnets,
			}, nil
		}
	}
	return nil, nil
}
