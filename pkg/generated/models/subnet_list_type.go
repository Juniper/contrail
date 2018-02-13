package models

// SubnetListType

import "encoding/json"

// SubnetListType
//proteus:generate
type SubnetListType struct {
	Subnet []*SubnetType `json:"subnet,omitempty"`
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

// MakeSubnetListTypeSlice() makes a slice of SubnetListType
func MakeSubnetListTypeSlice() []*SubnetListType {
	return []*SubnetListType{}
}
