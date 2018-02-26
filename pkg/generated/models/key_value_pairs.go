package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeKeyValuePairs makes KeyValuePairs
func MakeKeyValuePairs() *KeyValuePairs {
	return &KeyValuePairs{
		//TODO(nati): Apply default

		KeyValuePair: MakeKeyValuePairSlice(),
	}
}

// MakeKeyValuePairs makes KeyValuePairs
func InterfaceToKeyValuePairs(i interface{}) *KeyValuePairs {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &KeyValuePairs{
		//TODO(nati): Apply default

		KeyValuePair: InterfaceToKeyValuePairSlice(m["key_value_pair"]),
	}
}

// MakeKeyValuePairsSlice() makes a slice of KeyValuePairs
func MakeKeyValuePairsSlice() []*KeyValuePairs {
	return []*KeyValuePairs{}
}

// InterfaceToKeyValuePairsSlice() makes a slice of KeyValuePairs
func InterfaceToKeyValuePairsSlice(i interface{}) []*KeyValuePairs {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*KeyValuePairs{}
	for _, item := range list {
		result = append(result, InterfaceToKeyValuePairs(item))
	}
	return result
}
