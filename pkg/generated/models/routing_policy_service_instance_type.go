package models

// RoutingPolicyServiceInstanceType

// RoutingPolicyServiceInstanceType
//proteus:generate
type RoutingPolicyServiceInstanceType struct {
	RightSequence string `json:"right_sequence,omitempty"`
	LeftSequence  string `json:"left_sequence,omitempty"`
}

// MakeRoutingPolicyServiceInstanceType makes RoutingPolicyServiceInstanceType
func MakeRoutingPolicyServiceInstanceType() *RoutingPolicyServiceInstanceType {
	return &RoutingPolicyServiceInstanceType{
		//TODO(nati): Apply default
		RightSequence: "",
		LeftSequence:  "",
	}
}

// MakeRoutingPolicyServiceInstanceTypeSlice() makes a slice of RoutingPolicyServiceInstanceType
func MakeRoutingPolicyServiceInstanceTypeSlice() []*RoutingPolicyServiceInstanceType {
	return []*RoutingPolicyServiceInstanceType{}
}
