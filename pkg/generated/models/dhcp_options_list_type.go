package models

// DhcpOptionsListType

import "encoding/json"

// DhcpOptionsListType
type DhcpOptionsListType struct {
	DHCPOption []*DhcpOptionType `json:"dhcp_option"`
}

//  parents relation object

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

		//{"Title":"","Description":"List of DHCP options","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"dhcp_option_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionName","GoType":"string","GoPremitive":true},"dhcp_option_value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValue","GoType":"string","GoPremitive":true},"dhcp_option_value_bytes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValueBytes","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DhcpOptionType","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOption","GoType":"DhcpOptionType","GoPremitive":false},"GoName":"DHCPOption","GoType":"[]*DhcpOptionType","GoPremitive":true}

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
