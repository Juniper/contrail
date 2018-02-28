package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeUserDefinedLogStatList makes UserDefinedLogStatList
// nolint
func MakeUserDefinedLogStatList() *UserDefinedLogStatList {
	return &UserDefinedLogStatList{
		//TODO(nati): Apply default

		Statlist: MakeUserDefinedLogStatSlice(),
	}
}

// MakeUserDefinedLogStatList makes UserDefinedLogStatList
// nolint
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
// nolint
func MakeUserDefinedLogStatListSlice() []*UserDefinedLogStatList {
	return []*UserDefinedLogStatList{}
}

// InterfaceToUserDefinedLogStatListSlice() makes a slice of UserDefinedLogStatList
// nolint
func InterfaceToUserDefinedLogStatListSlice(i interface{}) []*UserDefinedLogStatList {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*UserDefinedLogStatList{}
	for _, item := range list {
		result = append(result, InterfaceToUserDefinedLogStatList(item))
	}
	return result
}
