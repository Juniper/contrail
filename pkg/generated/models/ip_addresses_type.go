package models

// IpAddressesType

import "encoding/json"

// IpAddressesType
type IpAddressesType struct {
	IPAddress IpAddressType `json:"ip_address,omitempty"`
}

// String returns json representation of the object
func (model *IpAddressesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIpAddressesType makes IpAddressesType
func MakeIpAddressesType() *IpAddressesType {
	return &IpAddressesType{
		//TODO(nati): Apply default
		IPAddress: MakeIpAddressType(),
	}
}

// MakeIpAddressesTypeSlice() makes a slice of IpAddressesType
func MakeIpAddressesTypeSlice() []*IpAddressesType {
	return []*IpAddressesType{}
}
