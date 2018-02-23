package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeVirtualMachineInterfacePropertiesType makes VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesType() *VirtualMachineInterfacePropertiesType {
	return &VirtualMachineInterfacePropertiesType{
		//TODO(nati): Apply default
		SubInterfaceVlanTag:  0,
		LocalPreference:      0,
		InterfaceMirror:      MakeInterfaceMirrorType(),
		ServiceInterfaceType: "",
	}
}

// MakeVirtualMachineInterfacePropertiesType makes VirtualMachineInterfacePropertiesType
func InterfaceToVirtualMachineInterfacePropertiesType(i interface{}) *VirtualMachineInterfacePropertiesType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualMachineInterfacePropertiesType{
		//TODO(nati): Apply default
		SubInterfaceVlanTag:  schema.InterfaceToInt64(m["sub_interface_vlan_tag"]),
		LocalPreference:      schema.InterfaceToInt64(m["local_preference"]),
		InterfaceMirror:      InterfaceToInterfaceMirrorType(m["interface_mirror"]),
		ServiceInterfaceType: schema.InterfaceToString(m["service_interface_type"]),
	}
}

// MakeVirtualMachineInterfacePropertiesTypeSlice() makes a slice of VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesTypeSlice() []*VirtualMachineInterfacePropertiesType {
	return []*VirtualMachineInterfacePropertiesType{}
}

// InterfaceToVirtualMachineInterfacePropertiesTypeSlice() makes a slice of VirtualMachineInterfacePropertiesType
func InterfaceToVirtualMachineInterfacePropertiesTypeSlice(i interface{}) []*VirtualMachineInterfacePropertiesType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualMachineInterfacePropertiesType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualMachineInterfacePropertiesType(item))
	}
	return result
}
