package models

// VirtualNetworkPolicyType

// VirtualNetworkPolicyType
//proteus:generate
type VirtualNetworkPolicyType struct {
	Timer    *TimerType    `json:"timer,omitempty"`
	Sequence *SequenceType `json:"sequence,omitempty"`
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
