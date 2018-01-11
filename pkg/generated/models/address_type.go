package models

// AddressType

import "encoding/json"

// AddressType
type AddressType struct {
	SecurityGroup  string        `json:"security_group"`
	Subnet         *SubnetType   `json:"subnet"`
	NetworkPolicy  string        `json:"network_policy"`
	SubnetList     []*SubnetType `json:"subnet_list"`
	VirtualNetwork string        `json:"virtual_network"`
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

		SubnetList: MakeSubnetTypeSlice(),

		VirtualNetwork: "",
		SecurityGroup:  "",
		Subnet:         MakeSubnetType(),
		NetworkPolicy:  "",
	}
}

// MakeAddressTypeSlice() makes a slice of AddressType
func MakeAddressTypeSlice() []*AddressType {
	return []*AddressType{}
}
