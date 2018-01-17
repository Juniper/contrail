package models

// VirtualMachineInterfacePropertiesType

import "encoding/json"

// VirtualMachineInterfacePropertiesType
type VirtualMachineInterfacePropertiesType struct {
	ServiceInterfaceType ServiceInterfaceType `json:"service_interface_type,omitempty"`
	SubInterfaceVlanTag  int                  `json:"sub_interface_vlan_tag,omitempty"`
	LocalPreference      int                  `json:"local_preference,omitempty"`
	InterfaceMirror      *InterfaceMirrorType `json:"interface_mirror,omitempty"`
}

// String returns json representation of the object
func (model *VirtualMachineInterfacePropertiesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualMachineInterfacePropertiesType makes VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesType() *VirtualMachineInterfacePropertiesType {
	return &VirtualMachineInterfacePropertiesType{
		//TODO(nati): Apply default
		ServiceInterfaceType: MakeServiceInterfaceType(),
		SubInterfaceVlanTag:  0,
		LocalPreference:      0,
		InterfaceMirror:      MakeInterfaceMirrorType(),
	}
}

// MakeVirtualMachineInterfacePropertiesTypeSlice() makes a slice of VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesTypeSlice() []*VirtualMachineInterfacePropertiesType {
	return []*VirtualMachineInterfacePropertiesType{}
}
