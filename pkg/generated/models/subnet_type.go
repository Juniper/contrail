package models

// SubnetType

import "encoding/json"

// SubnetType
type SubnetType struct {
	IPPrefix    string `json:"ip_prefix"`
	IPPrefixLen int    `json:"ip_prefix_len"`
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

// MakeSubnetTypeSlice() makes a slice of SubnetType
func MakeSubnetTypeSlice() []*SubnetType {
	return []*SubnetType{}
}
