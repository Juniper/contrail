package models

// KeyValuePair

import "encoding/json"

// KeyValuePair
type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

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

		//{"type":"string"}
		Key: data["key"].(string),

		//{"type":"string"}

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
