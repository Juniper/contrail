package models

// KeyValuePairs

import "encoding/json"

// KeyValuePairs
type KeyValuePairs struct {
	KeyValuePair []*KeyValuePair `json:"key_value_pair"`
}

//  parents relation object

// String returns json representation of the object
func (model *KeyValuePairs) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeKeyValuePairs makes KeyValuePairs
func MakeKeyValuePairs() *KeyValuePairs {
	return &KeyValuePairs{
		//TODO(nati): Apply default

		KeyValuePair: MakeKeyValuePairSlice(),
	}
}

// InterfaceToKeyValuePairs makes KeyValuePairs from interface
func InterfaceToKeyValuePairs(iData interface{}) *KeyValuePairs {
	data := iData.(map[string]interface{})
	return &KeyValuePairs{

		KeyValuePair: InterfaceToKeyValuePairSlice(data["key_value_pair"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair","GoPremitive":false},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair","GoPremitive":true}

	}
}

// InterfaceToKeyValuePairsSlice makes a slice of KeyValuePairs from interface
func InterfaceToKeyValuePairsSlice(data interface{}) []*KeyValuePairs {
	list := data.([]interface{})
	result := MakeKeyValuePairsSlice()
	for _, item := range list {
		result = append(result, InterfaceToKeyValuePairs(item))
	}
	return result
}

// MakeKeyValuePairsSlice() makes a slice of KeyValuePairs
func MakeKeyValuePairsSlice() []*KeyValuePairs {
	return []*KeyValuePairs{}
}
