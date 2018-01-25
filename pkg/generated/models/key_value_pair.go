package models

// KeyValuePair

// KeyValuePair
//proteus:generate
type KeyValuePair struct {
	Value string `json:"value,omitempty"`
	Key   string `json:"key,omitempty"`
}

// MakeKeyValuePair makes KeyValuePair
func MakeKeyValuePair() *KeyValuePair {
	return &KeyValuePair{
		//TODO(nati): Apply default
		Value: "",
		Key:   "",
	}
}

// MakeKeyValuePairSlice() makes a slice of KeyValuePair
func MakeKeyValuePairSlice() []*KeyValuePair {
	return []*KeyValuePair{}
}
