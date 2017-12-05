package models

// AllowedAddressPair

import "encoding/json"

type AllowedAddressPair struct {
	IP          *SubnetType `json:"ip"`
	Mac         string      `json:"mac"`
	AddressMode AddressMode `json:"address_mode"`
}

func (model *AllowedAddressPair) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeAllowedAddressPair() *AllowedAddressPair {
	return &AllowedAddressPair{
		//TODO(nati): Apply default
		Mac:         "",
		AddressMode: MakeAddressMode(),
		IP:          MakeSubnetType(),
	}
}

func InterfaceToAllowedAddressPair(iData interface{}) *AllowedAddressPair {
	data := iData.(map[string]interface{})
	return &AllowedAddressPair{
		IP: InterfaceToSubnetType(data["ip"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"IP","GoType":"SubnetType"}
		Mac: data["mac"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Mac","GoType":"string"}
		AddressMode: InterfaceToAddressMode(data["address_mode"]),

		//{"Title":"","Description":"Address-mode active-backup is used for VRRP address.                           Address-mode active-active is used for ECMP.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["active-active","active-standby"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressMode","CollectionType":"","Column":"","Item":null,"GoName":"AddressMode","GoType":"AddressMode"}

	}
}

func InterfaceToAllowedAddressPairSlice(data interface{}) []*AllowedAddressPair {
	list := data.([]interface{})
	result := MakeAllowedAddressPairSlice()
	for _, item := range list {
		result = append(result, InterfaceToAllowedAddressPair(item))
	}
	return result
}

func MakeAllowedAddressPairSlice() []*AllowedAddressPair {
	return []*AllowedAddressPair{}
}
