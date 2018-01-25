package models

// KeyValuePairs

// KeyValuePairs
//proteus:generate
type KeyValuePairs struct {
	KeyValuePair []*KeyValuePair `json:"key_value_pair,omitempty"`
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
