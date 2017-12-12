package models

// MacAddressesType

import "encoding/json"

// MacAddressesType
type MacAddressesType struct {
	MacAddress []string `json:"mac_address"`
}

// String returns json representation of the object
func (model *MacAddressesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMacAddressesType makes MacAddressesType
func MakeMacAddressesType() *MacAddressesType {
	return &MacAddressesType{
		//TODO(nati): Apply default
		MacAddress: []string{},
	}
}

// InterfaceToMacAddressesType makes MacAddressesType from interface
func InterfaceToMacAddressesType(iData interface{}) *MacAddressesType {
	data := iData.(map[string]interface{})
	return &MacAddressesType{
		MacAddress: data["mac_address"].([]string),

		//{"type":"array","item":{"type":"string"}}

	}
}

// InterfaceToMacAddressesTypeSlice makes a slice of MacAddressesType from interface
func InterfaceToMacAddressesTypeSlice(data interface{}) []*MacAddressesType {
	list := data.([]interface{})
	result := MakeMacAddressesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMacAddressesType(item))
	}
	return result
}

// MakeMacAddressesTypeSlice() makes a slice of MacAddressesType
func MakeMacAddressesTypeSlice() []*MacAddressesType {
	return []*MacAddressesType{}
}
