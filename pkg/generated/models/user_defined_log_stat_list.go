package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeUserDefinedLogStatList makes UserDefinedLogStatList
func MakeUserDefinedLogStatList() *UserDefinedLogStatList {
	return &UserDefinedLogStatList{
		//TODO(nati): Apply default

		Statlist: MakeUserDefinedLogStatSlice(),
	}
}

// MakeUserDefinedLogStatList makes UserDefinedLogStatList
func InterfaceToUserDefinedLogStatList(i interface{}) *UserDefinedLogStatList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &UserDefinedLogStatList{
		//TODO(nati): Apply default

		Statlist: InterfaceToUserDefinedLogStatSlice(m["statlist"]),
	}
}

// MakeUserDefinedLogStatListSlice() makes a slice of UserDefinedLogStatList
func MakeUserDefinedLogStatListSlice() []*UserDefinedLogStatList {
	return []*UserDefinedLogStatList{}
}

// InterfaceToUserDefinedLogStatListSlice() makes a slice of UserDefinedLogStatList
func InterfaceToUserDefinedLogStatListSlice(i interface{}) []*UserDefinedLogStatList {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*UserDefinedLogStatList{}
	for _, item := range list {
		result = append(result, InterfaceToUserDefinedLogStatList(item))
	}
	return result
}
