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

// InterfaceToDhcpOptionsListType makes DhcpOptionsListType from interface
func InterfaceToDhcpOptionsListType(iData interface{}) *DhcpOptionsListType {
	data := iData.(map[string]interface{})
	return &DhcpOptionsListType{

		DHCPOption: InterfaceToDhcpOptionTypeSlice(data["dhcp_option"]),

		//{"description":"List of DHCP options","type":"array","item":{"type":"object","properties":{"dhcp_option_name":{"type":"string"},"dhcp_option_value":{"type":"string"},"dhcp_option_value_bytes":{"type":"string"}}}}

	}
}

// InterfaceToDhcpOptionsListTypeSlice makes a slice of DhcpOptionsListType from interface
func InterfaceToDhcpOptionsListTypeSlice(data interface{}) []*DhcpOptionsListType {
	list := data.([]interface{})
	result := MakeDhcpOptionsListTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDhcpOptionsListType(item))
	}
	return result
}

// MakeDhcpOptionsListTypeSlice() makes a slice of DhcpOptionsListType
func MakeDhcpOptionsListTypeSlice() []*DhcpOptionsListType {
	return []*DhcpOptionsListType{}
}
