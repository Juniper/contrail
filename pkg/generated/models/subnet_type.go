package models

// SubnetType

import "encoding/json"

// SubnetType
type SubnetType struct {
	IPPrefixLen int    `json:"ip_prefix_len"`
	IPPrefix    string `json:"ip_prefix"`
}

// String returns json representation of the object
func (model *SubnetType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSubnetType makes SubnetType
func MakeSubnetType() *SubnetType {
	return &SubnetType{
		//TODO(nati): Apply default
		IPPrefix:    "",
		IPPrefixLen: 0,
	}
}

// InterfaceToSubnetType makes SubnetType from interface
func InterfaceToSubnetType(iData interface{}) *SubnetType {
	data := iData.(map[string]interface{})
	return &SubnetType{
		IPPrefix: data["ip_prefix"].(string),

		//{"type":"string"}
		IPPrefixLen: data["ip_prefix_len"].(int),

		//{"type":"integer"}

	}
}

// InterfaceToSubnetTypeSlice makes a slice of SubnetType from interface
func InterfaceToSubnetTypeSlice(data interface{}) []*SubnetType {
	list := data.([]interface{})
	result := MakeSubnetTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSubnetType(item))
	}
	return result
}

// MakeSubnetTypeSlice() makes a slice of SubnetType
func MakeSubnetTypeSlice() []*SubnetType {
	return []*SubnetType{}
}
