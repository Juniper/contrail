package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeSequenceType makes SequenceType
func MakeSequenceType() *SequenceType {
	return &SequenceType{
		//TODO(nati): Apply default
		Major: 0,
		Minor: 0,
	}
}

// MakeSequenceType makes SequenceType
func InterfaceToSequenceType(i interface{}) *SequenceType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SequenceType{
		//TODO(nati): Apply default
		Major: schema.InterfaceToInt64(m["major"]),
		Minor: schema.InterfaceToInt64(m["minor"]),
	}
}

// MakeSequenceTypeSlice() makes a slice of SequenceType
func MakeSequenceTypeSlice() []*SequenceType {
	return []*SequenceType{}
}

// InterfaceToSequenceTypeSlice() makes a slice of SequenceType
func InterfaceToSequenceTypeSlice(i interface{}) []*SequenceType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SequenceType{}
	for _, item := range list {
		result = append(result, InterfaceToSequenceType(item))
	}
	return result
}
