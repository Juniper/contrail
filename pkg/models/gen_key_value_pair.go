package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeKeyValuePair makes KeyValuePair
// nolint
func MakeKeyValuePair() *KeyValuePair {
	return &KeyValuePair{
		//TODO(nati): Apply default
		Value: "",
		Key:   "",
	}
}

// MakeKeyValuePair makes KeyValuePair
// nolint
func InterfaceToKeyValuePair(i interface{}) *KeyValuePair {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &KeyValuePair{
		//TODO(nati): Apply default
		Value: common.InterfaceToString(m["value"]),
		Key:   common.InterfaceToString(m["key"]),
	}
}

// MakeKeyValuePairSlice() makes a slice of KeyValuePair
// nolint
func MakeKeyValuePairSlice() []*KeyValuePair {
	return []*KeyValuePair{}
}

// InterfaceToKeyValuePairSlice() makes a slice of KeyValuePair
// nolint
func InterfaceToKeyValuePairSlice(i interface{}) []*KeyValuePair {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*KeyValuePair{}
	for _, item := range list {
		result = append(result, InterfaceToKeyValuePair(item))
	}
	return result
}
