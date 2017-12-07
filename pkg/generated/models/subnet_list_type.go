package models

// SubnetListType

import "encoding/json"

// SubnetListType
type SubnetListType struct {
	Subnet []*SubnetType `json:"subnet"`
}

//  parents relation object

// String returns json representation of the object
func (model *SubnetListType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSubnetListType makes SubnetListType
func MakeSubnetListType() *SubnetListType {
	return &SubnetListType{
		//TODO(nati): Apply default

		Subnet: MakeSubnetTypeSlice(),
	}
}

// InterfaceToSubnetListType makes SubnetListType from interface
func InterfaceToSubnetListType(iData interface{}) *SubnetListType {
	data := iData.(map[string]interface{})
	return &SubnetListType{

		Subnet: InterfaceToSubnetTypeSlice(data["subnet"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType","GoPremitive":false},"GoName":"Subnet","GoType":"[]*SubnetType","GoPremitive":true}

	}
}

// InterfaceToSubnetListTypeSlice makes a slice of SubnetListType from interface
func InterfaceToSubnetListTypeSlice(data interface{}) []*SubnetListType {
	list := data.([]interface{})
	result := MakeSubnetListTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSubnetListType(item))
	}
	return result
}

// MakeSubnetListTypeSlice() makes a slice of SubnetListType
func MakeSubnetListTypeSlice() []*SubnetListType {
	return []*SubnetListType{}
}
