package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeVirtualNetworkPolicyType makes VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyType() *VirtualNetworkPolicyType {
	return &VirtualNetworkPolicyType{
		//TODO(nati): Apply default
		Timer:    MakeTimerType(),
		Sequence: MakeSequenceType(),
	}
}

// MakeVirtualNetworkPolicyType makes VirtualNetworkPolicyType
func InterfaceToVirtualNetworkPolicyType(i interface{}) *VirtualNetworkPolicyType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualNetworkPolicyType{
		//TODO(nati): Apply default
		Timer:    InterfaceToTimerType(m["timer"]),
		Sequence: InterfaceToSequenceType(m["sequence"]),
	}
}

// MakeVirtualNetworkPolicyTypeSlice() makes a slice of VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyTypeSlice() []*VirtualNetworkPolicyType {
	return []*VirtualNetworkPolicyType{}
}

// InterfaceToVirtualNetworkPolicyTypeSlice() makes a slice of VirtualNetworkPolicyType
func InterfaceToVirtualNetworkPolicyTypeSlice(i interface{}) []*VirtualNetworkPolicyType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualNetworkPolicyType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetworkPolicyType(item))
	}
	return result
}
