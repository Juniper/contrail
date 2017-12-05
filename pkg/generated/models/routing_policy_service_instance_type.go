package models

// RoutingPolicyServiceInstanceType

import "encoding/json"

type RoutingPolicyServiceInstanceType struct {
	RightSequence string `json:"right_sequence"`
	LeftSequence  string `json:"left_sequence"`
}

func (model *RoutingPolicyServiceInstanceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeRoutingPolicyServiceInstanceType() *RoutingPolicyServiceInstanceType {
	return &RoutingPolicyServiceInstanceType{
		//TODO(nati): Apply default
		RightSequence: "",
		LeftSequence:  "",
	}
}

func InterfaceToRoutingPolicyServiceInstanceType(iData interface{}) *RoutingPolicyServiceInstanceType {
	data := iData.(map[string]interface{})
	return &RoutingPolicyServiceInstanceType{
		RightSequence: data["right_sequence"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"right_sequence","Item":null,"GoName":"RightSequence","GoType":"string"}
		LeftSequence: data["left_sequence"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"left_sequence","Item":null,"GoName":"LeftSequence","GoType":"string"}

	}
}

func InterfaceToRoutingPolicyServiceInstanceTypeSlice(data interface{}) []*RoutingPolicyServiceInstanceType {
	list := data.([]interface{})
	result := MakeRoutingPolicyServiceInstanceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRoutingPolicyServiceInstanceType(item))
	}
	return result
}

func MakeRoutingPolicyServiceInstanceTypeSlice() []*RoutingPolicyServiceInstanceType {
	return []*RoutingPolicyServiceInstanceType{}
}
