package models

// AllowedAddressPairs

import "encoding/json"

// AllowedAddressPairs
type AllowedAddressPairs struct {
	AllowedAddressPair []*AllowedAddressPair `json:"allowed_address_pair"`
}

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

		//{"type":"array","item":{"type":"object","properties":{"address_mode":{"type":"string","enum":["active-active","active-standby"]},"ip":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"mac":{"type":"string"}}}}

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
