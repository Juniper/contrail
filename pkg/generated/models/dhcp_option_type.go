package models

// DhcpOptionType

import "encoding/json"

// DhcpOptionType
type DhcpOptionType struct {
	DHCPOptionName       string `json:"dhcp_option_name"`
	DHCPOptionValue      string `json:"dhcp_option_value"`
	DHCPOptionValueBytes string `json:"dhcp_option_value_bytes"`
}

// String returns json representation of the object
func (model *DhcpOptionType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDhcpOptionType makes DhcpOptionType
func MakeDhcpOptionType() *DhcpOptionType {
	return &DhcpOptionType{
		//TODO(nati): Apply default
		DHCPOptionValue:      "",
		DHCPOptionValueBytes: "",
		DHCPOptionName:       "",
	}
}

// MakeDhcpOptionTypeSlice() makes a slice of DhcpOptionType
func MakeDhcpOptionTypeSlice() []*DhcpOptionType {
	return []*DhcpOptionType{}
}
