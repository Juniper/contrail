package models

// KeyValuePairs

import "encoding/json"

type KeyValuePairs struct {
	KeyValuePair []*KeyValuePair `json:"key_value_pair"`
}

func (model *KeyValuePairs) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeKeyValuePairs() *KeyValuePairs {
	return &KeyValuePairs{
		//TODO(nati): Apply default

		KeyValuePair: MakeKeyValuePairSlice(),
	}
}

func InterfaceToKeyValuePairs(iData interface{}) *KeyValuePairs {
	data := iData.(map[string]interface{})
	return &KeyValuePairs{

		KeyValuePair: InterfaceToKeyValuePairSlice(data["key_value_pair"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string"},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair"},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair"}

	}
}

func InterfaceToKeyValuePairsSlice(data interface{}) []*KeyValuePairs {
	list := data.([]interface{})
	result := MakeKeyValuePairsSlice()
	for _, item := range list {
		result = append(result, InterfaceToKeyValuePairs(item))
	}
	return result
}

func MakeKeyValuePairsSlice() []*KeyValuePairs {
	return []*KeyValuePairs{}
}
