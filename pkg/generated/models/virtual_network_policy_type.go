package models

// VirtualNetworkPolicyType

import "encoding/json"

// VirtualNetworkPolicyType
type VirtualNetworkPolicyType struct {
	Timer    *TimerType    `json:"timer"`
	Sequence *SequenceType `json:"sequence"`
}

// String returns json representation of the object
func (model *VirtualNetworkPolicyType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualNetworkPolicyType makes VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyType() *VirtualNetworkPolicyType {
	return &VirtualNetworkPolicyType{
		//TODO(nati): Apply default
		Timer:    MakeTimerType(),
		Sequence: MakeSequenceType(),
	}
}

// MakeVirtualNetworkPolicyTypeSlice() makes a slice of VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyTypeSlice() []*VirtualNetworkPolicyType {
	return []*VirtualNetworkPolicyType{}
}
