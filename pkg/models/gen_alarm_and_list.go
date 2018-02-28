package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAlarmAndList makes AlarmAndList
// nolint
func MakeAlarmAndList() *AlarmAndList {
	return &AlarmAndList{
		//TODO(nati): Apply default

		AndList: MakeAlarmExpressionSlice(),
	}
}

// MakeAlarmAndList makes AlarmAndList
// nolint
func InterfaceToAlarmAndList(i interface{}) *AlarmAndList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AlarmAndList{
		//TODO(nati): Apply default

		AndList: InterfaceToAlarmExpressionSlice(m["and_list"]),
	}
}

// MakeAlarmAndListSlice() makes a slice of AlarmAndList
// nolint
func MakeAlarmAndListSlice() []*AlarmAndList {
	return []*AlarmAndList{}
}

// InterfaceToAlarmAndListSlice() makes a slice of AlarmAndList
// nolint
func InterfaceToAlarmAndListSlice(i interface{}) []*AlarmAndList {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AlarmAndList{}
	for _, item := range list {
		result = append(result, InterfaceToAlarmAndList(item))
	}
	return result
}
