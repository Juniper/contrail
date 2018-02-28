package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAlarmOperand2 makes AlarmOperand2
// nolint
func MakeAlarmOperand2() *AlarmOperand2 {
	return &AlarmOperand2{
		//TODO(nati): Apply default
		UveAttribute: "",
		JSONValue:    "",
	}
}

// MakeAlarmOperand2 makes AlarmOperand2
// nolint
func InterfaceToAlarmOperand2(i interface{}) *AlarmOperand2 {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AlarmOperand2{
		//TODO(nati): Apply default
		UveAttribute: common.InterfaceToString(m["uve_attribute"]),
		JSONValue:    common.InterfaceToString(m["json_value"]),
	}
}

// MakeAlarmOperand2Slice() makes a slice of AlarmOperand2
// nolint
func MakeAlarmOperand2Slice() []*AlarmOperand2 {
	return []*AlarmOperand2{}
}

// InterfaceToAlarmOperand2Slice() makes a slice of AlarmOperand2
// nolint
func InterfaceToAlarmOperand2Slice(i interface{}) []*AlarmOperand2 {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AlarmOperand2{}
	for _, item := range list {
		result = append(result, InterfaceToAlarmOperand2(item))
	}
	return result
}
