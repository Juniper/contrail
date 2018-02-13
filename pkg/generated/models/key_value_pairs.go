package models

// KeyValuePairs

import "encoding/json"

// KeyValuePairs
//proteus:generate
type KeyValuePairs struct {
	KeyValuePair []*KeyValuePair `json:"key_value_pair,omitempty"`
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

// MakeKeyValuePairsSlice() makes a slice of KeyValuePairs
func MakeKeyValuePairsSlice() []*KeyValuePairs {
	return []*KeyValuePairs{}
}
