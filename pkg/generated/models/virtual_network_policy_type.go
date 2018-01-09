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

// InterfaceToVirtualNetworkPolicyType makes VirtualNetworkPolicyType from interface
func InterfaceToVirtualNetworkPolicyType(iData interface{}) *VirtualNetworkPolicyType {
	data := iData.(map[string]interface{})
	return &VirtualNetworkPolicyType{
		Sequence: InterfaceToSequenceType(data["sequence"]),

		//{"description":"Sequence number to specify order of policy attachment to network","type":"object","properties":{"major":{"type":"integer"},"minor":{"type":"integer"}}}
		Timer: InterfaceToTimerType(data["timer"]),

		//{"description":"Timer to specify when the policy can be active","type":"object","properties":{"end_time":{"type":"string"},"off_interval":{"type":"string"},"on_interval":{"type":"string"},"start_time":{"type":"string"}}}

	}
}

// InterfaceToVirtualNetworkPolicyTypeSlice makes a slice of VirtualNetworkPolicyType from interface
func InterfaceToVirtualNetworkPolicyTypeSlice(data interface{}) []*VirtualNetworkPolicyType {
	list := data.([]interface{})
	result := MakeVirtualNetworkPolicyTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetworkPolicyType(item))
	}
	return result
}

// MakeVirtualNetworkPolicyTypeSlice() makes a slice of VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyTypeSlice() []*VirtualNetworkPolicyType {
	return []*VirtualNetworkPolicyType{}
}
