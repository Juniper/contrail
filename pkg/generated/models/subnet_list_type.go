package models

// SubnetListType

import "encoding/json"

type SubnetListType struct {
	Subnet []*SubnetType `json:"subnet"`
}

func (model *SubnetListType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeSubnetListType() *SubnetListType {
	return &SubnetListType{
		//TODO(nati): Apply default

		Subnet: MakeSubnetTypeSlice(),
	}
}

func InterfaceToSubnetListType(iData interface{}) *SubnetListType {
	data := iData.(map[string]interface{})
	return &SubnetListType{

		Subnet: InterfaceToSubnetTypeSlice(data["subnet"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"GoName":"Subnet","GoType":"[]*SubnetType"}

	}
}

func InterfaceToSubnetListTypeSlice(data interface{}) []*SubnetListType {
	list := data.([]interface{})
	result := MakeSubnetListTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSubnetListType(item))
	}
	return result
}

func MakeSubnetListTypeSlice() []*SubnetListType {
	return []*SubnetListType{}
}
