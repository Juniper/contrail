package models

// DhcpOptionsListType

import "encoding/json"

// DhcpOptionsListType
type DhcpOptionsListType struct {
	DHCPOption []*DhcpOptionType `json:"dhcp_option"`
}

// String returns json representation of the object
func (model *DhcpOptionsListType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
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
