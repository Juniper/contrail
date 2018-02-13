package models

// DhcpOptionType

import "encoding/json"

// DhcpOptionType
//proteus:generate
type DhcpOptionType struct {
	DHCPOptionValue      string `json:"dhcp_option_value,omitempty"`
	DHCPOptionValueBytes string `json:"dhcp_option_value_bytes,omitempty"`
	DHCPOptionName       string `json:"dhcp_option_name,omitempty"`
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
