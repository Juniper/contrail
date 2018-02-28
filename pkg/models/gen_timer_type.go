package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeTimerType makes TimerType
// nolint
func MakeTimerType() *TimerType {
	return &TimerType{
		//TODO(nati): Apply default
		StartTime:   "",
		OffInterval: "",
		OnInterval:  "",
		EndTime:     "",
	}
}

// MakeTimerType makes TimerType
// nolint
func InterfaceToTimerType(i interface{}) *TimerType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &TimerType{
		//TODO(nati): Apply default
		StartTime:   common.InterfaceToString(m["start_time"]),
		OffInterval: common.InterfaceToString(m["off_interval"]),
		OnInterval:  common.InterfaceToString(m["on_interval"]),
		EndTime:     common.InterfaceToString(m["end_time"]),
	}
}

// MakeTimerTypeSlice() makes a slice of TimerType
// nolint
func MakeTimerTypeSlice() []*TimerType {
	return []*TimerType{}
}

// InterfaceToTimerTypeSlice() makes a slice of TimerType
// nolint
func InterfaceToTimerTypeSlice(i interface{}) []*TimerType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*TimerType{}
	for _, item := range list {
		result = append(result, InterfaceToTimerType(item))
	}
	return result
}
