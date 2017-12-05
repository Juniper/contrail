package models

// DhcpOptionType

import "encoding/json"

type DhcpOptionType struct {
	DHCPOptionValue      string `json:"dhcp_option_value"`
	DHCPOptionValueBytes string `json:"dhcp_option_value_bytes"`
	DHCPOptionName       string `json:"dhcp_option_name"`
}

func (model *DhcpOptionType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeDhcpOptionType() *DhcpOptionType {
	return &DhcpOptionType{
		//TODO(nati): Apply default
		DHCPOptionValue:      "",
		DHCPOptionValueBytes: "",
		DHCPOptionName:       "",
	}
}

func InterfaceToDhcpOptionType(iData interface{}) *DhcpOptionType {
	data := iData.(map[string]interface{})
	return &DhcpOptionType{
		DHCPOptionValue: data["dhcp_option_value"].(string),

		//{"Title":"","Description":"Encoded DHCP option value (decimal)","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValue","GoType":"string"}
		DHCPOptionValueBytes: data["dhcp_option_value_bytes"].(string),

		//{"Title":"","Description":"Value of the DHCP option to be copied byte by byte","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValueBytes","GoType":"string"}
		DHCPOptionName: data["dhcp_option_name"].(string),

		//{"Title":"","Description":"Name of the DHCP option","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionName","GoType":"string"}

	}
}

func InterfaceToDhcpOptionTypeSlice(data interface{}) []*DhcpOptionType {
	list := data.([]interface{})
	result := MakeDhcpOptionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDhcpOptionType(item))
	}
	return result
}

func MakeDhcpOptionTypeSlice() []*DhcpOptionType {
	return []*DhcpOptionType{}
}
