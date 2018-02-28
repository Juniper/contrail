package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeSequenceType makes SequenceType
// nolint
func MakeSequenceType() *SequenceType {
	return &SequenceType{
		//TODO(nati): Apply default
		Major: 0,
		Minor: 0,
	}
}

// MakeSequenceType makes SequenceType
// nolint
func InterfaceToSequenceType(i interface{}) *SequenceType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SequenceType{
		//TODO(nati): Apply default
		Major: common.InterfaceToInt64(m["major"]),
		Minor: common.InterfaceToInt64(m["minor"]),
	}
}

// MakeSequenceTypeSlice() makes a slice of SequenceType
// nolint
func MakeSequenceTypeSlice() []*SequenceType {
	return []*SequenceType{}
}

// InterfaceToSequenceTypeSlice() makes a slice of SequenceType
// nolint
func InterfaceToSequenceTypeSlice(i interface{}) []*SequenceType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SequenceType{}
	for _, item := range list {
		result = append(result, InterfaceToSequenceType(item))
	}
	return result
}
