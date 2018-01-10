package models

// DhcpOptionType

import "encoding/json"

// DhcpOptionType
type DhcpOptionType struct {
	DHCPOptionValue      string `json:"dhcp_option_value"`
	DHCPOptionValueBytes string `json:"dhcp_option_value_bytes"`
	DHCPOptionName       string `json:"dhcp_option_name"`
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
		DHCPOptionName:       "",
		DHCPOptionValue:      "",
		DHCPOptionValueBytes: "",
	}
}

// InterfaceToDhcpOptionType makes DhcpOptionType from interface
func InterfaceToDhcpOptionType(iData interface{}) *DhcpOptionType {
	data := iData.(map[string]interface{})
	return &DhcpOptionType{
		DHCPOptionValue: data["dhcp_option_value"].(string),

		//{"description":"Encoded DHCP option value (decimal)","type":"string"}
		DHCPOptionValueBytes: data["dhcp_option_value_bytes"].(string),

		//{"description":"Value of the DHCP option to be copied byte by byte","type":"string"}
		DHCPOptionName: data["dhcp_option_name"].(string),

		//{"description":"Name of the DHCP option","type":"string"}

	}
}

// InterfaceToDhcpOptionTypeSlice makes a slice of DhcpOptionType from interface
func InterfaceToDhcpOptionTypeSlice(data interface{}) []*DhcpOptionType {
	list := data.([]interface{})
	result := MakeDhcpOptionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDhcpOptionType(item))
	}
	return result
}

// MakeDhcpOptionTypeSlice() makes a slice of DhcpOptionType
func MakeDhcpOptionTypeSlice() []*DhcpOptionType {
	return []*DhcpOptionType{}
}
