package models

// SubnetListType

import "encoding/json"

// SubnetListType
type SubnetListType struct {
	Subnet []*SubnetType `json:"subnet"`
}

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

		//{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}}

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
