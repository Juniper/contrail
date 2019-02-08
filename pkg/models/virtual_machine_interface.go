package models

// FindInterfaceRouteTableRef finds Interface Route Table Reference.
func (m *VirtualMachineInterface) FindInterfaceRouteTableRef(
	predicate func(*VirtualMachineInterfaceInterfaceRouteTableRef) bool,
) *VirtualMachineInterfaceInterfaceRouteTableRef {
	for _, ref := range m.InterfaceRouteTableRefs {
		if predicate(ref) {
			return ref
		}
	}
	return nil
}
