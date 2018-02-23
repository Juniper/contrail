package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeKeyValuePair makes KeyValuePair
func MakeKeyValuePair() *KeyValuePair {
	return &KeyValuePair{
		//TODO(nati): Apply default
		Value: "",
		Key:   "",
	}
}

// MakeKeyValuePair makes KeyValuePair
func InterfaceToKeyValuePair(i interface{}) *KeyValuePair {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &KeyValuePair{
		//TODO(nati): Apply default
		Value: schema.InterfaceToString(m["value"]),
		Key:   schema.InterfaceToString(m["key"]),
	}
}

// MakeKeyValuePairSlice() makes a slice of KeyValuePair
func MakeKeyValuePairSlice() []*KeyValuePair {
	return []*KeyValuePair{}
}

// InterfaceToKeyValuePairSlice() makes a slice of KeyValuePair
func InterfaceToKeyValuePairSlice(i interface{}) []*KeyValuePair {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*KeyValuePair{}
	for _, item := range list {
		result = append(result, InterfaceToKeyValuePair(item))
	}
	return result
}
