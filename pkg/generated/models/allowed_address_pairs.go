package models

// AllowedAddressPairs

import "encoding/json"

type AllowedAddressPairs struct {
	AllowedAddressPair []*AllowedAddressPair `json:"allowed_address_pair"`
}

func (model *AllowedAddressPairs) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeAllowedAddressPairs() *AllowedAddressPairs {
	return &AllowedAddressPairs{
		//TODO(nati): Apply default

		AllowedAddressPair: MakeAllowedAddressPairSlice(),
	}
}

func InterfaceToAllowedAddressPairs(iData interface{}) *AllowedAddressPairs {
	data := iData.(map[string]interface{})
	return &AllowedAddressPairs{

		AllowedAddressPair: InterfaceToAllowedAddressPairSlice(data["allowed_address_pair"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"address_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["active-active","active-standby"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressMode","CollectionType":"","Column":"","Item":null,"GoName":"AddressMode","GoType":"AddressMode"},"ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"IP","GoType":"SubnetType"},"mac":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Mac","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllowedAddressPair","CollectionType":"","Column":"","Item":null,"GoName":"AllowedAddressPair","GoType":"AllowedAddressPair"},"GoName":"AllowedAddressPair","GoType":"[]*AllowedAddressPair"}

	}
}

func InterfaceToAllowedAddressPairsSlice(data interface{}) []*AllowedAddressPairs {
	list := data.([]interface{})
	result := MakeAllowedAddressPairsSlice()
	for _, item := range list {
		result = append(result, InterfaceToAllowedAddressPairs(item))
	}
	return result
}

func MakeAllowedAddressPairsSlice() []*AllowedAddressPairs {
	return []*AllowedAddressPairs{}
}
