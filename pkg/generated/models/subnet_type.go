package models

// SubnetType

import "encoding/json"

type SubnetType struct {
	IPPrefix    string `json:"ip_prefix"`
	IPPrefixLen int    `json:"ip_prefix_len"`
}

func (model *SubnetType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeSubnetType() *SubnetType {
	return &SubnetType{
		//TODO(nati): Apply default
		IPPrefix:    "",
		IPPrefixLen: 0,
	}
}

func InterfaceToSubnetType(iData interface{}) *SubnetType {
	data := iData.(map[string]interface{})
	return &SubnetType{
		IPPrefix: data["ip_prefix"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix","Item":null,"GoName":"IPPrefix","GoType":"string"}
		IPPrefixLen: data["ip_prefix_len"].(int),

		//{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix_len","Item":null,"GoName":"IPPrefixLen","GoType":"int"}

	}
}

func InterfaceToSubnetTypeSlice(data interface{}) []*SubnetType {
	list := data.([]interface{})
	result := MakeSubnetTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSubnetType(item))
	}
	return result
}

func MakeSubnetTypeSlice() []*SubnetType {
	return []*SubnetType{}
}
