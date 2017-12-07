package models

// RoutingPolicyServiceInstanceType

import "encoding/json"

// RoutingPolicyServiceInstanceType
type RoutingPolicyServiceInstanceType struct {
	RightSequence string `json:"right_sequence"`
	LeftSequence  string `json:"left_sequence"`
}

//  parents relation object

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

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"right_sequence","Item":null,"GoName":"RightSequence","GoType":"string","GoPremitive":true}
		LeftSequence: data["left_sequence"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"left_sequence","Item":null,"GoName":"LeftSequence","GoType":"string","GoPremitive":true}

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
