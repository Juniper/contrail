package models

// DhcpOptionType

import "encoding/json"

// DhcpOptionType
type DhcpOptionType struct {
	DHCPOptionValue      string `json:"dhcp_option_value"`
	DHCPOptionValueBytes string `json:"dhcp_option_value_bytes"`
	DHCPOptionName       string `json:"dhcp_option_name"`
}

//  parents relation object

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

// InterfaceToDhcpOptionType makes DhcpOptionType from interface
func InterfaceToDhcpOptionType(iData interface{}) *DhcpOptionType {
	data := iData.(map[string]interface{})
	return &DhcpOptionType{
		DHCPOptionValue: data["dhcp_option_value"].(string),

		//{"Title":"","Description":"Encoded DHCP option value (decimal)","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValue","GoType":"string","GoPremitive":true}
		DHCPOptionValueBytes: data["dhcp_option_value_bytes"].(string),

		//{"Title":"","Description":"Value of the DHCP option to be copied byte by byte","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValueBytes","GoType":"string","GoPremitive":true}
		DHCPOptionName: data["dhcp_option_name"].(string),

		//{"Title":"","Description":"Name of the DHCP option","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionName","GoType":"string","GoPremitive":true}

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
