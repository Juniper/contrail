package models

import "github.com/Juniper/contrail/pkg/common"

// IsParentTypeVirtualNetwork checks if parent's type is virtual network
func (fipp *FloatingIPPool) IsParentTypeVirtualNetwork() bool {
	var m VirtualNetwork
	return fipp.GetParentType() == m.Kind()
}

// HasSubnets checks if floating-ip-pool has any subnets defined
func (fipp *FloatingIPPool) HasSubnets() bool {
	floatingIPPoolSubnets := fipp.GetFloatingIPPoolSubnets()
	return floatingIPPoolSubnets == nil || len(floatingIPPoolSubnets.GetSubnetUUID()) == 0
}

// AreSubnetsInVirtualNetworkSubnets checks if subnets defined in floating-ip-pool object
// are present in the virtual-network
func (fipp *FloatingIPPool) AreSubnetsInVirtualNetworkSubnets(vn *VirtualNetwork) (bool, error) {
	for _, floatingIPPoolSubnetUUID := range fipp.GetFloatingIPPoolSubnets().GetSubnetUUID() {
		subnetFound := false
		for _, ipam := range vn.GetNetworkIpamRefs() {
			for _, ipamSubnet := range ipam.GetAttr().GetIpamSubnets() {
				if ipamSubnet.GetSubnetUUID() == floatingIPPoolSubnetUUID {
					subnetFound = true
					break
				}
			}
		}

		if !subnetFound {
			return false, common.ErrorBadRequestf("Subnet %s was not found in virtual-network %s",
				floatingIPPoolSubnetUUID, vn.GetUUID())
		}
	}
	return true, nil
}
