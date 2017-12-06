package models

// KeyValuePair

import "encoding/json"

// KeyValuePair
type KeyValuePair struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

//  parents relation object

// String returns json representation of the object
func (model *KeyValuePair) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeKeyValuePair makes KeyValuePair
func MakeKeyValuePair() *KeyValuePair {
	return &KeyValuePair{
		//TODO(nati): Apply default
		Value: "",
		Key:   "",
	}
}

// InterfaceToKeyValuePair makes KeyValuePair from interface
func InterfaceToKeyValuePair(iData interface{}) *KeyValuePair {
	data := iData.(map[string]interface{})
	return &KeyValuePair{
		Value: data["value"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}
		Key: data["key"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToKeyValuePairSlice makes a slice of KeyValuePair from interface
func InterfaceToKeyValuePairSlice(data interface{}) []*KeyValuePair {
	list := data.([]interface{})
	result := MakeKeyValuePairSlice()
	for _, item := range list {
		result = append(result, InterfaceToKeyValuePair(item))
	}
	return result
}

// MakeKeyValuePairSlice() makes a slice of KeyValuePair
func MakeKeyValuePairSlice() []*KeyValuePair {
	return []*KeyValuePair{}
}
