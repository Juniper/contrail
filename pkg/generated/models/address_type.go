package models

// AddressType

import "encoding/json"

// AddressType
type AddressType struct {
	SubnetList     []*SubnetType `json:"subnet_list,omitempty"`
	VirtualNetwork string        `json:"virtual_network,omitempty"`
	SecurityGroup  string        `json:"security_group,omitempty"`
	Subnet         *SubnetType   `json:"subnet,omitempty"`
	NetworkPolicy  string        `json:"network_policy,omitempty"`
}

// String returns json representation of the object
func (model *AddressType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAddressType makes AddressType
func MakeAddressType() *AddressType {
	return &AddressType{
		//TODO(nati): Apply default
		SecurityGroup: "",
		Subnet:        MakeSubnetType(),
		NetworkPolicy: "",

		SubnetList: MakeSubnetTypeSlice(),

		VirtualNetwork: "",
	}
}

// MakeAddressTypeSlice() makes a slice of AddressType
func MakeAddressTypeSlice() []*AddressType {
	return []*AddressType{}
}
