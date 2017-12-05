package models

// UveKeysType

import "encoding/json"

type UveKeysType struct {
	UveKey []string `json:"uve_key"`
}

func (model *UveKeysType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeUveKeysType() *UveKeysType {
	return &UveKeysType{
		//TODO(nati): Apply default
		UveKey: []string{},
	}
}

func InterfaceToUveKeysType(iData interface{}) *UveKeysType {
	data := iData.(map[string]interface{})
	return &UveKeysType{
		UveKey: data["uve_key"].([]string),

		//{"Title":"","Description":"List of UVE tables where this alarm config should be applied","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UveKey","GoType":"string"},"GoName":"UveKey","GoType":"[]string"}

	}
}

func InterfaceToUveKeysTypeSlice(data interface{}) []*UveKeysType {
	list := data.([]interface{})
	result := MakeUveKeysTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToUveKeysType(item))
	}
	return result
}

func MakeUveKeysTypeSlice() []*UveKeysType {
	return []*UveKeysType{}
}
