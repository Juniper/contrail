package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAlarmOrList makes AlarmOrList
// nolint
func MakeAlarmOrList() *AlarmOrList {
	return &AlarmOrList{
		//TODO(nati): Apply default

		OrList: MakeAlarmAndListSlice(),
	}
}

// MakeAlarmOrList makes AlarmOrList
// nolint
func InterfaceToAlarmOrList(i interface{}) *AlarmOrList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AlarmOrList{
		//TODO(nati): Apply default

		OrList: InterfaceToAlarmAndListSlice(m["or_list"]),
	}
}

// MakeAlarmOrListSlice() makes a slice of AlarmOrList
// nolint
func MakeAlarmOrListSlice() []*AlarmOrList {
	return []*AlarmOrList{}
}

// InterfaceToAlarmOrListSlice() makes a slice of AlarmOrList
// nolint
func InterfaceToAlarmOrListSlice(i interface{}) []*AlarmOrList {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AlarmOrList{}
	for _, item := range list {
		result = append(result, InterfaceToAlarmOrList(item))
	}
	return result
}
