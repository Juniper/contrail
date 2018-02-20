package models

// DhcpOptionsListType

// DhcpOptionsListType
//proteus:generate
type DhcpOptionsListType struct {
	DHCPOption []*DhcpOptionType `json:"dhcp_option,omitempty"`
}

// MakeDhcpOptionsListType makes DhcpOptionsListType
func MakeDhcpOptionsListType() *DhcpOptionsListType {
	return &DhcpOptionsListType{
		//TODO(nati): Apply default

		DHCPOption: MakeDhcpOptionTypeSlice(),
	}
}

// MakeDhcpOptionsListTypeSlice() makes a slice of DhcpOptionsListType
func MakeDhcpOptionsListTypeSlice() []*DhcpOptionsListType {
	return []*DhcpOptionsListType{}
}
