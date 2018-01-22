package models

// AllowedAddressPair

import "encoding/json"

// AllowedAddressPair
type AllowedAddressPair struct {
	AddressMode AddressMode `json:"address_mode,omitempty"`
	IP          *SubnetType `json:"ip,omitempty"`
	Mac         string      `json:"mac,omitempty"`
}

// String returns json representation of the object
func (model *AllowedAddressPair) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAllowedAddressPair makes AllowedAddressPair
func MakeAllowedAddressPair() *AllowedAddressPair {
	return &AllowedAddressPair{
		//TODO(nati): Apply default
		IP:          MakeSubnetType(),
		Mac:         "",
		AddressMode: MakeAddressMode(),
	}
}

// MakeAllowedAddressPairSlice() makes a slice of AllowedAddressPair
func MakeAllowedAddressPairSlice() []*AllowedAddressPair {
	return []*AllowedAddressPair{}
}
