package models

// IsParentTypeVirtualNetwork checks if parent's type is virtual network
func (fipp *FloatingIPPool) IsParentTypeVirtualNetwork() bool {
	var m VirtualNetwork
	return fipp.GetParentType() == m.Kind()
}
