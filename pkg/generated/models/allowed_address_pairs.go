package models

// AllowedAddressPairs

import "encoding/json"

// AllowedAddressPairs
type AllowedAddressPairs struct {
	AllowedAddressPair []*AllowedAddressPair `json:"allowed_address_pair"`
}

//  parents relation object

// String returns json representation of the object
func (model *AllowedAddressPairs) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAllowedAddressPairs makes AllowedAddressPairs
func MakeAllowedAddressPairs() *AllowedAddressPairs {
	return &AllowedAddressPairs{
		//TODO(nati): Apply default

		AllowedAddressPair: MakeAllowedAddressPairSlice(),
	}
}

// InterfaceToAllowedAddressPairs makes AllowedAddressPairs from interface
func InterfaceToAllowedAddressPairs(iData interface{}) *AllowedAddressPairs {
	data := iData.(map[string]interface{})
	return &AllowedAddressPairs{

		AllowedAddressPair: InterfaceToAllowedAddressPairSlice(data["allowed_address_pair"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"address_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["active-active","active-standby"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressMode","CollectionType":"","Column":"","Item":null,"GoName":"AddressMode","GoType":"AddressMode","GoPremitive":false},"ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"IP","GoType":"SubnetType","GoPremitive":false},"mac":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Mac","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllowedAddressPair","CollectionType":"","Column":"","Item":null,"GoName":"AllowedAddressPair","GoType":"AllowedAddressPair","GoPremitive":false},"GoName":"AllowedAddressPair","GoType":"[]*AllowedAddressPair","GoPremitive":true}

	}
}

// InterfaceToAllowedAddressPairsSlice makes a slice of AllowedAddressPairs from interface
func InterfaceToAllowedAddressPairsSlice(data interface{}) []*AllowedAddressPairs {
	list := data.([]interface{})
	result := MakeAllowedAddressPairsSlice()
	for _, item := range list {
		result = append(result, InterfaceToAllowedAddressPairs(item))
	}
	return result
}

// MakeAllowedAddressPairsSlice() makes a slice of AllowedAddressPairs
func MakeAllowedAddressPairsSlice() []*AllowedAddressPairs {
	return []*AllowedAddressPairs{}
}
