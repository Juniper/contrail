package models

// VirtualMachineInterfacePropertiesType

// VirtualMachineInterfacePropertiesType
//proteus:generate
type VirtualMachineInterfacePropertiesType struct {
	SubInterfaceVlanTag  int                  `json:"sub_interface_vlan_tag,omitempty"`
	LocalPreference      int                  `json:"local_preference,omitempty"`
	InterfaceMirror      *InterfaceMirrorType `json:"interface_mirror,omitempty"`
	ServiceInterfaceType ServiceInterfaceType `json:"service_interface_type,omitempty"`
}

// MakeVirtualMachineInterfacePropertiesType makes VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesType() *VirtualMachineInterfacePropertiesType {
	return &VirtualMachineInterfacePropertiesType{
		//TODO(nati): Apply default
		SubInterfaceVlanTag:  0,
		LocalPreference:      0,
		InterfaceMirror:      MakeInterfaceMirrorType(),
		ServiceInterfaceType: MakeServiceInterfaceType(),
	}
}

// MakeVirtualMachineInterfacePropertiesTypeSlice() makes a slice of VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesTypeSlice() []*VirtualMachineInterfacePropertiesType {
	return []*VirtualMachineInterfacePropertiesType{}
}
