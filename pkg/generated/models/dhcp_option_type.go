package models

// DhcpOptionType

// DhcpOptionType
//proteus:generate
type DhcpOptionType struct {
	DHCPOptionValue      string `json:"dhcp_option_value,omitempty"`
	DHCPOptionValueBytes string `json:"dhcp_option_value_bytes,omitempty"`
	DHCPOptionName       string `json:"dhcp_option_name,omitempty"`
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
