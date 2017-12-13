package models

// UveKeysType

import "encoding/json"

// UveKeysType
type UveKeysType struct {
	UveKey []string `json:"uve_key"`
}

// String returns json representation of the object
func (model *UveKeysType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeUveKeysType makes UveKeysType
func MakeUveKeysType() *UveKeysType {
	return &UveKeysType{
		//TODO(nati): Apply default
		UveKey: []string{},
	}
}

// InterfaceToUveKeysType makes UveKeysType from interface
func InterfaceToUveKeysType(iData interface{}) *UveKeysType {
	data := iData.(map[string]interface{})
	return &UveKeysType{
		UveKey: data["uve_key"].([]string),

		//{"description":"List of UVE tables where this alarm config should be applied","type":"array","item":{"type":"string"}}

	}
}

// InterfaceToUveKeysTypeSlice makes a slice of UveKeysType from interface
func InterfaceToUveKeysTypeSlice(data interface{}) []*UveKeysType {
	list := data.([]interface{})
	result := MakeUveKeysTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToUveKeysType(item))
	}
	return result
}

// MakeUveKeysTypeSlice() makes a slice of UveKeysType
func MakeUveKeysTypeSlice() []*UveKeysType {
	return []*UveKeysType{}
}
