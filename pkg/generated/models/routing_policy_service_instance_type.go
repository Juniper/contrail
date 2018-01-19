package models

// RoutingPolicyServiceInstanceType

import "encoding/json"

// RoutingPolicyServiceInstanceType
type RoutingPolicyServiceInstanceType struct {
	RightSequence string `json:"right_sequence,omitempty"`
	LeftSequence  string `json:"left_sequence,omitempty"`
}

// String returns json representation of the object
func (model *RoutingPolicyServiceInstanceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
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
