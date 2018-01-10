package models

// AllowedAddressPair

import "encoding/json"

// AllowedAddressPair
type AllowedAddressPair struct {
	Mac         string      `json:"mac"`
	AddressMode AddressMode `json:"address_mode"`
	IP          *SubnetType `json:"ip"`
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

// InterfaceToAllowedAddressPair makes AllowedAddressPair from interface
func InterfaceToAllowedAddressPair(iData interface{}) *AllowedAddressPair {
	data := iData.(map[string]interface{})
	return &AllowedAddressPair{
		IP: InterfaceToSubnetType(data["ip"]),

		//{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}
		Mac: data["mac"].(string),

		//{"type":"string"}
		AddressMode: InterfaceToAddressMode(data["address_mode"]),

		//{"description":"Address-mode active-backup is used for VRRP address.                           Address-mode active-active is used for ECMP.","type":"string","enum":["active-active","active-standby"]}

	}
}

// InterfaceToAllowedAddressPairSlice makes a slice of AllowedAddressPair from interface
func InterfaceToAllowedAddressPairSlice(data interface{}) []*AllowedAddressPair {
	list := data.([]interface{})
	result := MakeAllowedAddressPairSlice()
	for _, item := range list {
		result = append(result, InterfaceToAllowedAddressPair(item))
	}
	return result
}

// MakeAllowedAddressPairSlice() makes a slice of AllowedAddressPair
func MakeAllowedAddressPairSlice() []*AllowedAddressPair {
	return []*AllowedAddressPair{}
}
