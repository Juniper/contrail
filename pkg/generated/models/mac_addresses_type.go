package models

// MacAddressesType

import "encoding/json"

type MacAddressesType struct {
	MacAddress []string `json:"mac_address"`
}

func (model *MacAddressesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeMacAddressesType() *MacAddressesType {
	return &MacAddressesType{
		//TODO(nati): Apply default
		MacAddress: []string{},
	}
}

func InterfaceToMacAddressesType(iData interface{}) *MacAddressesType {
	data := iData.(map[string]interface{})
	return &MacAddressesType{
		MacAddress: data["mac_address"].([]string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MacAddress","GoType":"string"},"GoName":"MacAddress","GoType":"[]string"}

	}
}

func InterfaceToMacAddressesTypeSlice(data interface{}) []*MacAddressesType {
	list := data.([]interface{})
	result := MakeMacAddressesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMacAddressesType(item))
	}
	return result
}

func MakeMacAddressesTypeSlice() []*MacAddressesType {
	return []*MacAddressesType{}
}
