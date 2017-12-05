package models

// KeyValuePair

import "encoding/json"

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (model *KeyValuePair) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeKeyValuePair() *KeyValuePair {
	return &KeyValuePair{
		//TODO(nati): Apply default
		Value: "",
		Key:   "",
	}
}

func InterfaceToKeyValuePair(iData interface{}) *KeyValuePair {
	data := iData.(map[string]interface{})
	return &KeyValuePair{
		Key: data["key"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string"}
		Value: data["value"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}

	}
}

func InterfaceToKeyValuePairSlice(data interface{}) []*KeyValuePair {
	list := data.([]interface{})
	result := MakeKeyValuePairSlice()
	for _, item := range list {
		result = append(result, InterfaceToKeyValuePair(item))
	}
	return result
}

func MakeKeyValuePairSlice() []*KeyValuePair {
	return []*KeyValuePair{}
}
