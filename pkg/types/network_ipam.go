package types

import (
	"context"
	"net"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func isSubnetOverlap(subnet1, subnet2 *net.IPNet) bool {
	return subnet1.Contains(subnet2.IP) || subnet2.Contains(subnet1.IP)
}

func appendSubnetIfNoOverlap(subnets []*net.IPNet, subnet *net.IPNet) ([]*net.IPNet, error) {
	for _, existingSubnet := range subnets {
		if isSubnetOverlap(subnet, existingSubnet) {
			return nil, errors.Errorf(
				"overlapping addresses: %s, %s", subnet.String(), existingSubnet.String())
		}
	}
	return append(subnets, subnet), nil
}

func checkSubnetsOverlap(ipamSubnets []*models.IpamSubnetType) error {
	subnets := []*net.IPNet{}
	for _, subnet := range ipamSubnets {
		n, err := subnet.Subnet.Net()
		if err != nil {
			return err
		}
		subnets, err = appendSubnetIfNoOverlap(subnets, n)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateNetworkIpam do pre check for network ipam
func (sv *ContrailTypeLogicService) CreateNetworkIpam(
	ctx context.Context,
	request *services.CreateNetworkIpamRequest,
) (response *services.CreateNetworkIpamResponse, err error) {

	networkIpam := request.GetNetworkIpam()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			ipamSubnets := networkIpam.GetIpamSubnets()
			if ipamSubnets == nil || len(ipamSubnets.Subnets) == 0 {
				response, err = sv.BaseService.CreateNetworkIpam(ctx, request)
				return err
			}

			if !networkIpam.IsFlatSubnet() {
				return errutil.ErrorBadRequest("Ipam subnets are allowed only with flat-subnet")
			}

			err = checkSubnetsOverlap(ipamSubnets.GetSubnets())
			if err != nil {
				return errutil.ErrorBadRequest(err.Error())
			}

			for _, ipamSubnet := range ipamSubnets.GetSubnets() {
				err = ipamSubnet.ValidateSubnetParams()
				if err != nil {
					return err
				}
			}

			response, err = sv.BaseService.CreateNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

// UpdateNetworkIpam do pre check for network ipam update
func (sv *ContrailTypeLogicService) UpdateNetworkIpam(
	ctx context.Context,
	request *services.UpdateNetworkIpamRequest,
) (response *services.UpdateNetworkIpamResponse, err error) {

	newNetworkIpam := request.GetNetworkIpam()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var oldNetworkIpam *models.NetworkIpam
			oldNetworkIpam, err = sv.getNetworkIpam(ctx, newNetworkIpam.GetUUID())
			if err != nil {
				return err
			}

			fieldMask := request.GetFieldMask()
			err = sv.checkNetworkIpamMGMT(oldNetworkIpam, newNetworkIpam, &fieldMask)
			if err != nil {
				return errutil.ErrorBadRequestf("check for network ipam mgmt failed with error: %v", err)
			}

			err = sv.checkSubnetMethod(oldNetworkIpam, newNetworkIpam, &fieldMask)
			if err != nil {
				return errutil.ErrorBadRequestf("check for subnet method failed with error: %v", err)
			}

			err = sv.checkIpamSubnets(ctx, oldNetworkIpam, newNetworkIpam, &fieldMask)
			if err != nil {
				return err
			}

			err = sv.checkSubnetDelete(ctx, oldNetworkIpam, newNetworkIpam, &fieldMask)
			if err != nil {
				return grpc.Errorf(codes.Aborted, err.Error())
			}

			err = sv.validateSubnetUpdate(oldNetworkIpam.GetIpamSubnets(), newNetworkIpam.GetIpamSubnets())
			if err != nil {
				return errutil.ErrorBadRequestf("validate subnet update failed with error: %v", err)
			}

			err = sv.processIpamUpdate(oldNetworkIpam, newNetworkIpam, &fieldMask)
			if err != nil {
				return errutil.ErrorBadRequestf("ipam update failed with error: %v", err)
			}

			response, err = sv.BaseService.UpdateNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

func (sv *ContrailTypeLogicService) getNetworkIpam(
	ctx context.Context,
	id string,
) (*models.NetworkIpam, error) {

	networkIpamRes, err := sv.ReadService.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{ID: id})
	if err != nil {
		return nil, err
	}
	return networkIpamRes.GetNetworkIpam(), err
}

func (sv *ContrailTypeLogicService) checkNetworkIpamMGMT(
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
	fieldMask *types.FieldMask,
) error {
	dnsMethodPath := []string{models.NetworkIpamFieldNetworkIpamMGMT, models.IpamTypeFieldIpamDNSMethod}
	if !format.CheckPath(fieldMask, dnsMethodPath) {
		return nil
	}
	isDNSChangeAllowed := sv.isDNSChangeAllowed(oldIpam, newIpam)
	if !isDNSChangeAllowed {
		return errors.Errorf("cannot change DNS method with active VMs referring to the IPAM")
	}
	return nil
}

func (sv *ContrailTypeLogicService) isDNSChangeAllowed(
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
) bool {
	isActiveVMPresent := sv.isActiveVMPresent(oldIpam)
	if !isActiveVMPresent {
		return true
	}
	oldDNSMethod := oldIpam.GetNetworkIpamMGMT().GetIpamDNSMethod()
	newDNSMethod := newIpam.GetNetworkIpamMGMT().GetIpamDNSMethod()
	if oldDNSMethod == "default-dns-server" || oldDNSMethod == "virtual-dns-server" {
		if newDNSMethod == "" || newDNSMethod == "tenant-dns-server" {
			return false
		}
	}
	if oldDNSMethod != newDNSMethod && (oldDNSMethod == "tenant-dns-server" || oldDNSMethod == "") {
		return false
	}
	return true
}

func (sv *ContrailTypeLogicService) isActiveVMPresent(
	networkIpam *models.NetworkIpam) bool {
	for _, vn := range networkIpam.GetVirtualNetworkBackRefs() {
		if len(vn.GetVirtualMachineInterfaceBackRefs()) > 0 {
			return true
		}
	}
	return false
}

func (sv *ContrailTypeLogicService) checkSubnetMethod(
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
	fieldMask *types.FieldMask,
) error {
	if format.CheckPath(fieldMask, []string{models.NetworkIpamFieldIpamSubnetMethod}) {
		if oldIpam.GetIpamSubnetMethod() != newIpam.GetIpamSubnetMethod() {
			return errors.Errorf("Subnet method cannot be changed")
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkIpamSubnets(
	ctx context.Context,
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
	fieldMask *types.FieldMask,
) error {
	ipamSubnetsPath := []string{models.NetworkIpamFieldIpamSubnets, models.IpamSubnetsFieldSubnets}
	if format.CheckPath(fieldMask, ipamSubnetsPath) {
		ipamSubnets := newIpam.GetIpamSubnets().GetSubnets()
		if len(ipamSubnets) == 0 {
			return nil
		}

		if oldIpam.GetIpamSubnetMethod() != "flat-subnet" {
			return errutil.ErrorBadRequest("ipam subnets are only allowed with flat subnet")
		}

		err := checkSubnetsOverlap(ipamSubnets)
		if err != nil {
			return errutil.ErrorBadRequest(err.Error())
		}

		var refIpamUUIDList []string
		refIpamUUIDList, err = sv.findFlatSubnetIpams(ctx, oldIpam)
		if err != nil {
			return err
		}

		var refSubnetsList []*models.IpamSubnetType
		refSubnetsList, err = sv.extractSubnetsFromVNRefs(ctx, oldIpam)
		if err != nil {
			return err
		}
		refSubnetsList, err = sv.extractSubnetsFromFlatSubnetIpams(ctx, refIpamUUIDList, refSubnetsList)
		if err != nil {
			return err
		}

		subnetsList := append(refSubnetsList, ipamSubnets...)
		err = checkSubnetsOverlap(subnetsList)
		if err != nil {
			return errutil.ErrorBadRequest(err.Error())
		}

		err = sv.validateIpamSubnets(ipamSubnets)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) validateIpamSubnets(
	ipamSubnets []*models.IpamSubnetType) error {
	for _, ipamSubnet := range ipamSubnets {
		err := ipamSubnet.ValidateSubnetParams()
		if err != nil {
			return err
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) extractSubnetsFromFlatSubnetIpams(
	ctx context.Context,
	refIpamUUIDList []string,
	refSubnetsList []*models.IpamSubnetType,
) ([]*models.IpamSubnetType, error) {
	for _, ipamUUID := range refIpamUUIDList {
		networkIpam, err := sv.getNetworkIpam(ctx, ipamUUID)
		if err != nil {
			return nil, err
		}
		refIpamSubnets := networkIpam.GetIpamSubnets().GetSubnets()
		refSubnetsList = append(refSubnetsList, refIpamSubnets...)
	}
	return refSubnetsList, nil
}

func (sv *ContrailTypeLogicService) findFlatSubnetIpams(
	ctx context.Context,
	networkIpam *models.NetworkIpam,
) ([]string, error) {
	ipamUUID := networkIpam.GetUUID()
	var refIpamUUIDList []string
	for _, vnRef := range networkIpam.GetVirtualNetworkBackRefs() {
		vn, err := sv.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: vnRef.GetUUID()})
		if err != nil {
			return nil, err
		}
		for _, ipamRef := range vn.GetVirtualNetwork().GetNetworkIpamRefs() {
			ipamRefUUID := ipamRef.GetUUID()
			if ipamRefUUID == ipamUUID || format.ContainsString(refIpamUUIDList, ipamRefUUID) {
				continue
			}
			if !ipamRef.GetAttr().IsFlatSubnet() {
				continue
			}
			refIpamUUIDList = append(refIpamUUIDList, ipamRefUUID)
		}
	}
	return refIpamUUIDList, nil
}

func (sv *ContrailTypeLogicService) extractSubnetsFromVNRefs(
	ctx context.Context,
	networkIpam *models.NetworkIpam,
) ([]*models.IpamSubnetType, error) {
	var refSubnetsList []*models.IpamSubnetType
	for _, vnRef := range networkIpam.GetVirtualNetworkBackRefs() {
		vn, err := sv.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: vnRef.GetUUID()})
		if err != nil {
			return nil, err
		}
		for _, ipamRef := range vn.GetVirtualNetwork().GetNetworkIpamRefs() {
			vnIpamSubnets := ipamRef.GetAttr()
			if vnIpamSubnets.IsFlatSubnet() {
				continue
			}
			refSubnetsList = append(refSubnetsList, vnIpamSubnets.GetIpamSubnets()...)
		}
	}
	return refSubnetsList, nil
}

func (sv *ContrailTypeLogicService) checkSubnetDelete(
	ctx context.Context,
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
	fieldMask *types.FieldMask,
) error {
	ipamSubnetsPath := []string{models.NetworkIpamFieldIpamSubnets, models.IpamSubnetsFieldSubnets}
	if !format.CheckPath(fieldMask, ipamSubnetsPath) {
		return nil
	}

	if len(oldIpam.GetIpamSubnets().GetSubnets()) == 0 {
		return nil
	}
	if oldIpam.GetIpamSubnetMethod() != "flat-subnet" {
		return nil
	}
	subnetsToDelete, err := sv.findSubnetsToDelete(oldIpam, newIpam)
	if err != nil {
		return err
	}
	if len(subnetsToDelete) == 0 {
		return nil
	}

	for _, vnRef := range oldIpam.GetVirtualNetworkBackRefs() {
		vn, err := sv.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: vnRef.GetUUID()})
		if err != nil {
			return err
		}
		err = sv.canSubnetsBeDeleted(ctx, vn.GetVirtualNetwork(), &models.IpamSubnets{Subnets: subnetsToDelete})
		if err != nil {
			return err
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) findSubnetsToDelete(
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
) (subnetsToDelete []*models.IpamSubnetType, err error) {
	oldIpamSubnets := oldIpam.GetIpamSubnets().GetSubnets()
	newIpamSubnets := newIpam.GetIpamSubnets().GetSubnets()
	for _, oldSubnet := range oldIpamSubnets {
		oldSn, err := oldSubnet.GetSubnet().Net()
		if err != nil {
			return nil, err
		}
		for _, newSubnet := range newIpamSubnets {
			newSn, err := newSubnet.GetSubnet().Net()
			if err != nil {
				return nil, err
			}
			if newSn != oldSn {
				subnetsToDelete = append(subnetsToDelete, oldSubnet)
			}
		}
	}
	return subnetsToDelete, nil
}

func (sv *ContrailTypeLogicService) validateSubnetUpdate(
	oldSubnetsSet *models.IpamSubnets,
	newSubnetsSet *models.IpamSubnets,
) error {
	if oldSubnetsSet == nil || newSubnetsSet == nil {
		return nil
	}
	return sv.validateDefaultGWChange(oldSubnetsSet.GetSubnets(), newSubnetsSet.GetSubnets())
}

func (sv *ContrailTypeLogicService) validateDefaultGWChange(
	oldSubnetsSet []*models.IpamSubnetType,
	newSubnetsSet []*models.IpamSubnetType,
) error {
	//TODO handle changes in default gateway
	return nil
}

func (sv *ContrailTypeLogicService) processIpamUpdate(
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
	fieldMask *types.FieldMask,
) error {
	ipamSubnetsPath := []string{models.NetworkIpamFieldIpamSubnets, models.IpamSubnetsFieldSubnets}
	if format.CheckPath(fieldMask, ipamSubnetsPath) {
		newIpamSubnets := newIpam.GetIpamSubnets().GetSubnets()
		oldIpamSubnets := oldIpam.GetIpamSubnets().GetSubnets()
		err := sv.validateSubnetChanges(oldIpamSubnets, newIpamSubnets)
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate changes in default gateways, allocation pools, dns server addresses
func (sv *ContrailTypeLogicService) validateSubnetChanges(
	oldIpamSubnets []*models.IpamSubnetType,
	newIpamSubnets []*models.IpamSubnetType,
) error {
	for _, newIpamSubnet := range newIpamSubnets {
		for _, oldIpamSubnet := range oldIpamSubnets {
			if oldIpamSubnet.GetSubnetName() == newIpamSubnet.GetSubnetName() {
				newDefaultGW := newIpamSubnet.GetDefaultGateway()
				// python code does it that way but not sure if this logic is exactly what we want
				//nolint: lll
				// https://github.com/Juniper/contrail-controller/blob/d9d8fdfb/src/config/api-server/vnc_cfg_api_server/vnc_addr_mgmt.py#L2291
				if newDefaultGW != "" && newDefaultGW != oldIpamSubnet.GetDefaultGateway() {
					return errors.Errorf("Cannot change default gateway")
				}
				newIpamSubnet.DNSServerAddress = oldIpamSubnet.GetDNSServerAddress()
				err := sv.checkSubnetAllocPools(newIpamSubnet, oldIpamSubnet)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkSubnetAllocPools(
	oldSubnet *models.IpamSubnetType,
	newSubnet *models.IpamSubnetType,
) error {
	//TODO check changes in allocPools
	return nil
}
