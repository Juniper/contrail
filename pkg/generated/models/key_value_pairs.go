package models

// KeyValuePairs

import "encoding/json"

// KeyValuePairs
type KeyValuePairs struct {
	KeyValuePair []*KeyValuePair `json:"key_value_pair"`
}

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

		//{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}

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
