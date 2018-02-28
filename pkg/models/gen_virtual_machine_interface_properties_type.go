package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualMachineInterfacePropertiesType makes VirtualMachineInterfacePropertiesType
// nolint
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
// nolint
func InterfaceToVirtualMachineInterfacePropertiesType(i interface{}) *VirtualMachineInterfacePropertiesType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualMachineInterfacePropertiesType{
		//TODO(nati): Apply default
		SubInterfaceVlanTag:  common.InterfaceToInt64(m["sub_interface_vlan_tag"]),
		LocalPreference:      common.InterfaceToInt64(m["local_preference"]),
		InterfaceMirror:      InterfaceToInterfaceMirrorType(m["interface_mirror"]),
		ServiceInterfaceType: common.InterfaceToString(m["service_interface_type"]),
	}
}

// MakeVirtualMachineInterfacePropertiesTypeSlice() makes a slice of VirtualMachineInterfacePropertiesType
// nolint
func MakeVirtualMachineInterfacePropertiesTypeSlice() []*VirtualMachineInterfacePropertiesType {
	return []*VirtualMachineInterfacePropertiesType{}
}

// InterfaceToVirtualMachineInterfacePropertiesTypeSlice() makes a slice of VirtualMachineInterfacePropertiesType
// nolint
func InterfaceToVirtualMachineInterfacePropertiesTypeSlice(i interface{}) []*VirtualMachineInterfacePropertiesType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualMachineInterfacePropertiesType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualMachineInterfacePropertiesType(item))
	}
	return result
}
