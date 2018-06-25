package types

import (
	"net"

	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
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
	request *services.CreateNetworkIpamRequest) (response *services.CreateNetworkIpamResponse, err error) {

	networkIpam := request.GetNetworkIpam()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			ipamSubnets := networkIpam.GetIpamSubnets()
			if ipamSubnets == nil {
				response, err = sv.BaseService.CreateNetworkIpam(ctx, request)
				return err
			}

			if !networkIpam.IsFlatSubnet() {
				return common.ErrorBadRequest("Ipam subnets are allowed only with flat-subnet")
			}

			err = checkSubnetsOverlap(ipamSubnets.GetSubnets())
			if err != nil {
				return common.ErrorBadRequest(err.Error())
			}

			for _, ipamSubnet := range ipamSubnets.GetSubnets() {
				err = ipamSubnet.CheckIfSubnetParamsAreValid()
				if err != nil {
					return err
				}
			}

			ipamUUID := networkIpam.GetUUID()
			for _, ipamSubnet := range ipamSubnets.GetSubnets() {
				subnetUUID, cErr := sv.createIpamSubnet(ctx, ipamSubnet, ipamUUID)
				if cErr != nil {
					return cErr
				}
				ipamSubnet.SubnetUUID = subnetUUID
			}

			response, err = sv.BaseService.CreateNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

// DeleteNetworkIpam do pre check for network ipam deletion
func (sv *ContrailTypeLogicService) DeleteNetworkIpam(
	ctx context.Context,
	request *services.DeleteNetworkIpamRequest) (response *services.DeleteNetworkIpamResponse, err error) {

	id := request.GetID()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var networkIpam *models.NetworkIpam
			networkIpam, err = sv.getNetworkIpam(ctx, id)
			if err != nil {
				return err
			}

			ipamSubnets := networkIpam.GetIpamSubnets()
			if networkIpam != nil && networkIpam.IsFlatSubnet() && ipamSubnets != nil {
				for _, ipamSubnet := range ipamSubnets.GetSubnets() {
					err = sv.deleteIpamSubnet(ctx, ipamSubnet.GetSubnetUUID())
					if err != nil {
						return err
					}
				}
			}

			response, err = sv.BaseService.DeleteNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

// UpdateNetworkIpam do pre check for network ipam update
func (sv *ContrailTypeLogicService) UpdateNetworkIpam(
	ctx context.Context,
	request *services.UpdateNetworkIpamRequest) (response *services.UpdateNetworkIpamResponse, err error) {

	newNetworkIpam := request.GetNetworkIpam()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			oldNetworkIpam, err := sv.getNetworkIpam(ctx, newNetworkIpam.GetUUID())
			if err != nil {
				return err
			}

			if oldNetworkIpam == nil {
				return errors.Errorf("no NetworkIpam found to update")
			}

			fieldMask := request.GetFieldMask()
			err = sv.checkNetworkIpamMGMT(ctx, oldNetworkIpam, newNetworkIpam, fieldMask.GetPaths())
			if err != nil {
				return common.ErrorBadRequest(err.Error())
			}

			err = sv.checkSubnetMethod(oldNetworkIpam, newNetworkIpam, fieldMask.GetPaths())
			if err != nil {
				return common.ErrorBadRequest(err.Error())
			}

			err = sv.checkIpamSubnets(ctx, oldNetworkIpam, newNetworkIpam, fieldMask.GetPaths())
			if err != nil {
				return err
			}

			for _, ipamSubnet := range newNetworkIpam.GetIpamSubnets().GetSubnets() {
				err = ipamSubnet.CheckIfSubnetParamsAreValid()
				if err != nil {
					return err
				}
			}

			err = sv.checkSubnetDelete(oldNetworkIpam, newNetworkIpam, fieldMask.GetPaths())
			if err != nil {
				return common.ErrorBadRequest(err.Error())
			}

			err = sv.validateSubnetUpdate(oldNetworkIpam, newNetworkIpam)
			if err != nil {
				return common.ErrorBadRequest(err.Error())
			}

			err = sv.processIpamUpdate(ctx, oldNetworkIpam, newNetworkIpam, fieldMask.GetPaths())
			if err != nil {
				return common.ErrorBadRequest(err.Error())
			}

			response, err = sv.BaseService.UpdateNetworkIpam(ctx, request)
			return err
		})
	return response, err
}

func (sv *ContrailTypeLogicService) getNetworkIpam(
	ctx context.Context,
	id string) (*models.NetworkIpam, error) {

	networkIpamRes, err := sv.DataService.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{ID: id})
	if err != nil {
		return nil, err
	}
	return networkIpamRes.GetNetworkIpam(), err
}

func (sv *ContrailTypeLogicService) createIpamSubnet(
	ctx context.Context,
	ipamSubnet *models.IpamSubnetType,
	ipamUUID string) (subnetUUID string, err error) {
	createIpamSubnetParams := &ipam.CreateIpamSubnetRequest{
		IpamSubnet:      ipamSubnet,
		NetworkIpamUUID: ipamUUID,
	}
	return sv.AddressManager.CreateIpamSubnet(ctx, createIpamSubnetParams)
}

func (sv *ContrailTypeLogicService) deleteIpamSubnet(
	ctx context.Context,
	subnetUUID string) error {
	deleteIpamSubnetParams := &ipam.DeleteIpamSubnetRequest{
		SubnetUUID: subnetUUID,
	}
	return sv.AddressManager.DeleteIpamSubnet(ctx, deleteIpamSubnetParams)
}

func (sv *ContrailTypeLogicService) checkNetworkIpamMGMT(
	ctx context.Context,
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
	fieldMask []string) error {
	if common.ContainsString(fieldMask, models.NetworkIpamPropertyIDNetworkIpamMGMT) {
		if oldIpam.GetNetworkIpamMGMT() == nil || newIpam.GetNetworkIpamMGMT() == nil {
			return nil
		}
		isChangeAllowed := sv.isChangeAllowed(ctx, oldIpam, newIpam)
		if !isChangeAllowed {
			return errors.Errorf("Cannot change DNS method with active VMs referring to the IPAM")
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) isChangeAllowed(
	ctx context.Context,
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam) bool {
	isActiveVMPresent := sv.isActiveVMPresent(ctx, oldIpam)
	if isActiveVMPresent {
		oldDNSMethod := oldIpam.GetNetworkIpamMGMT().GetIpamDNSMethod()
		newDNSMethod := newIpam.GetNetworkIpamMGMT().GetIpamDNSMethod()
		if oldDNSMethod == "default-dns-server" || oldDNSMethod == "virtual-dns-server" {
			if newDNSMethod == "" || newDNSMethod == "tenant-dns-server" {
				return false
			}
		}
		if oldDNSMethod == "tenant-dns-server" && oldDNSMethod != newDNSMethod {
			return false
		}
		if oldDNSMethod != "" && oldDNSMethod != newDNSMethod {
			return false
		}
	}
	return true
}

// check old ipam here
func (sv *ContrailTypeLogicService) isActiveVMPresent(
	ctx context.Context,
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
	fieldMask []string) error {
	if common.ContainsString(fieldMask, models.NetworkIpamPropertyIDIpamSubnetMethod) {
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
	fieldMask []string) error {
	if common.ContainsString(fieldMask, models.NetworkIpamPropertyIDIpamSubnets) {
		if oldIpam.GetIpamSubnetMethod() != "flat-subnet" {
			return common.ErrorBadRequest("Ipam subnets are only allowed with flat subnet")
		}

		ipamSubnets := newIpam.GetIpamSubnets().GetSubnets()
		err := checkSubnetsOverlap(ipamSubnets)
		if err != nil {
			return common.ErrorBadRequest(err.Error())
		}

		var refIpamUUIDList []string
		var refSubnetsList []*models.IpamSubnetType
		for _, vn := range newIpam.GetVirtualNetworkBackRefs() {
			for _, ipamRef := range vn.GetNetworkIpamRefs() {
				if ipamRef.GetUUID() == oldIpam.GetUUID() || common.ContainsString(refIpamUUIDList, ipamRef.GetUUID()) {
					continue
				}

				vnIpamSubnets := ipamRef.GetAttr().GetIpamSubnets()
				if len(vnIpamSubnets) == 1 {
					refIpamSubnet := vnIpamSubnets[0]
					if refIpamSubnet.GetSubnet().IPPrefix != "" {
						refIpamUUIDList = append(refIpamUUIDList, ipamRef.GetUUID())
					}
				}
				//processedVNSubnets := sv.vnToSubnets(vnIpamSubnets)
				//refSubnetsList = append(refSubnetsList, processedVNSubnets...)
				refSubnetsList = append(refSubnetsList, vnIpamSubnets...)
			}
		}
		for _, ipamUUID := range refIpamUUIDList {
			networkIpam, err := sv.DataService.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{ID: ipamUUID})
			if err != nil {
				return err
			}
			refIpamSubnets := networkIpam.GetNetworkIpam().GetIpamSubnets().GetSubnets()
			refSubnetsList = append(refSubnetsList, refIpamSubnets...)
		}
		// refs + exact from ipam
		err = checkSubnetsOverlap(refSubnetsList)
		if err != nil {
			return common.ErrorBadRequest(err.Error())
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) checkSubnetDelete(
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
	fieldMask []string) error {
	if common.ContainsString(fieldMask, models.NetworkIpamPropertyIDIpamSubnets) {
		if len(oldIpam.GetIpamSubnets().GetSubnets()) == 0 {
			return nil
		}
		if oldIpam.GetIpamSubnetMethod() != "flat-subnet" {
			return nil
		}
		subnetsToDelete, err := sv.findIpamSubnetsToDelete(oldIpam, newIpam)
		if err != nil {
			return err
		}
		if len(subnetsToDelete) == 0 {
			return nil
		}

		for _, vn := range oldIpam.GetVirtualNetworkBackRefs() {
			err := sv.checkSubnetToDelete(subnetsToDelete, vn)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (sv *ContrailTypeLogicService) findIpamSubnetsToDelete(
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam) (subnetsToDelete []*models.IpamSubnetType, err error) {
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

//check subnet delete
func (sv *ContrailTypeLogicService) checkSubnetToDelete(
	subnetsSet []*models.IpamSubnetType,
	vn *models.VirtualNetwork) error {
	for _, instanceIP := range vn.GetInstanceIPBackRefs() {
		err := checkIfSubnetsSetIncludeIP(subnetsSet, instanceIP.GetInstanceIPAddress())
		if err != nil {
			return err
		}
	}

	for _, floatingIPPool := range vn.GetFloatingIPPools() {
		for _, floatingIP := range floatingIPPool.GetFloatingIPs() {
			err := checkIfSubnetsSetIncludeIP(subnetsSet, floatingIP.GetFloatingIPAddress())
			if err != nil {
				return err
			}
		}
	}

	//FLOATING IP CHECK
	for _, aliasIPPools := range vn.GetAliasIPPools() {
		for _, aliasIP := range aliasIPPools.GetAliasIPs() {
			err := checkIfSubnetsSetIncludeIP(subnetsSet, aliasIP.GetAliasIPAddress())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func checkIfSubnetsSetIncludeIP(
	subnetsSet []*models.IpamSubnetType,
	ipString string) error {
	for _, ipamSubnet := range subnetsSet {
		subnet, err := ipamSubnet.GetSubnet().Net()
		if err != nil {
			return err
		}
		err = models.IsIpInSubnet(subnet, ipString)
		if err != nil {
			return err //IP in use
		}
	}
	return nil
}

func (sv *ContrailTypeLogicService) validateSubnetUpdate(
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam) error {
	oldSubnetsSet := oldIpam.GetIpamSubnets()
	newSubnetsSet := newIpam.GetIpamSubnets()
	if oldSubnetsSet == nil || newSubnetsSet == nil {
		return nil
	}
	err := sv.validateSubnetChanges(oldSubnetsSet.GetSubnets(), newSubnetsSet.GetSubnets())
	if err != nil {
		return err
	}
	return nil
}

func (sv *ContrailTypeLogicService) validateSubnetChanges(
	oldSubnetsSet []*models.IpamSubnetType,
	newSubnetsSet []*models.IpamSubnetType) error {

	//TODO handle changes in default gateway
	//for _, newSubnet := range newSubnetsSet {
	//	newCidr := newSubnet.GetSubnet()
	//	if newCidr == nil {
	//		continue
	//	}
	//	newDFGateway := newSubnet.GetDefaultGateway()
	//	for _, oldSubnet := range oldSubnetsSet {
	//		oldCidr := oldSubnet.GetSubnet()
	//		if oldCidr == nil || oldCidr != newCidr {
	//			continue
	//		}
	//		oldDFGateway := oldSubnet.GetDefaultGateway()
	//		var defaultGateway string
	//		if oldSubnet.AddrFromStart {
	//			m := oldSubnet.GetSubnet()
	//			cidr := m.IPPrefix + "/" + strconv.Itoa(int(m.IPPrefixLen))
	//			first, _, _ := net.ParseCIDR(cidr)
	//			first[4]++
	//			defaultGateway = first.String()
	//		} //else {
	//		//	defaultGateway :=
	//		//}
	//		if newDFGateway != oldDFGateway {
	//			if newDFGateway == "" && oldDFGateway != defaultGateway {
	//				return errors.Errorf("default gateway change is not allowed")
	//			}
	//			if newDFGateway != "" && newDFGateway != defaultGateway {
	//				return errors.Errorf("default gateway change is not allowed")
	//			}
	//		}
	//	}
	//}
	return nil
}

func (sv *ContrailTypeLogicService) processIpamUpdate(
	ctx context.Context,
	oldIpam *models.NetworkIpam,
	newIpam *models.NetworkIpam,
	fieldMask []string) error {
	if common.ContainsString(fieldMask, models.NetworkIpamPropertyIDIpamSubnets) {
		subnetsToDelete, err := sv.findIpamSubnetsToDelete(oldIpam, newIpam)
		if err != nil {
			return err
		}
		for _, subnet := range subnetsToDelete {
			err := sv.deleteIpamSubnet(ctx, subnet.GetSubnetUUID())
			if err != nil {
				return err
			}
		}
		//TODO remove stale entries in subnet_obj
		//TODO validate changes in default gateway, dns server, allocation pool
		//TODO create new ipam subnets
	}
	return nil
}
