package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeKeyValuePairs makes KeyValuePairs
// nolint
func MakeKeyValuePairs() *KeyValuePairs {
	return &KeyValuePairs{
		//TODO(nati): Apply default

		KeyValuePair: MakeKeyValuePairSlice(),
	}
}

// MakeKeyValuePairs makes KeyValuePairs
// nolint
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
// nolint
func MakeKeyValuePairsSlice() []*KeyValuePairs {
	return []*KeyValuePairs{}
}

// InterfaceToKeyValuePairsSlice() makes a slice of KeyValuePairs
// nolint
func InterfaceToKeyValuePairsSlice(i interface{}) []*KeyValuePairs {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*KeyValuePairs{}
	for _, item := range list {
		result = append(result, InterfaceToKeyValuePairs(item))
	}
	return result
}
