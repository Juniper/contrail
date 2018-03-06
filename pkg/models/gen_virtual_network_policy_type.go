package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualNetworkPolicyType makes VirtualNetworkPolicyType
// nolint
func MakeVirtualNetworkPolicyType() *VirtualNetworkPolicyType {
	return &VirtualNetworkPolicyType{
		//TODO(nati): Apply default
		Timer:    MakeTimerType(),
		Sequence: MakeSequenceType(),
	}
}

// MakeVirtualNetworkPolicyType makes VirtualNetworkPolicyType
// nolint
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
// nolint
func MakeVirtualNetworkPolicyTypeSlice() []*VirtualNetworkPolicyType {
	return []*VirtualNetworkPolicyType{}
}

// InterfaceToVirtualNetworkPolicyTypeSlice() makes a slice of VirtualNetworkPolicyType
// nolint
func InterfaceToVirtualNetworkPolicyTypeSlice(i interface{}) []*VirtualNetworkPolicyType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualNetworkPolicyType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetworkPolicyType(item))
	}
	return result
}
