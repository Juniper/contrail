package models

// RoutingPolicyServiceInstanceType

import "encoding/json"

// RoutingPolicyServiceInstanceType
type RoutingPolicyServiceInstanceType struct {
	LeftSequence  string `json:"left_sequence"`
	RightSequence string `json:"right_sequence"`
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

// InterfaceToRoutingPolicyServiceInstanceType makes RoutingPolicyServiceInstanceType from interface
func InterfaceToRoutingPolicyServiceInstanceType(iData interface{}) *RoutingPolicyServiceInstanceType {
	data := iData.(map[string]interface{})
	return &RoutingPolicyServiceInstanceType{
		RightSequence: data["right_sequence"].(string),

		//{"type":"string"}
		LeftSequence: data["left_sequence"].(string),

		//{"type":"string"}

	}
}

// InterfaceToRoutingPolicyServiceInstanceTypeSlice makes a slice of RoutingPolicyServiceInstanceType from interface
func InterfaceToRoutingPolicyServiceInstanceTypeSlice(data interface{}) []*RoutingPolicyServiceInstanceType {
	list := data.([]interface{})
	result := MakeRoutingPolicyServiceInstanceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRoutingPolicyServiceInstanceType(item))
	}
	return result
}

// MakeRoutingPolicyServiceInstanceTypeSlice() makes a slice of RoutingPolicyServiceInstanceType
func MakeRoutingPolicyServiceInstanceTypeSlice() []*RoutingPolicyServiceInstanceType {
	return []*RoutingPolicyServiceInstanceType{}
}
