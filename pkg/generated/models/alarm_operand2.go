package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeAlarmOperand2 makes AlarmOperand2
func MakeAlarmOperand2() *AlarmOperand2 {
	return &AlarmOperand2{
		//TODO(nati): Apply default
		UveAttribute: "",
		JSONValue:    "",
	}
}

// MakeAlarmOperand2 makes AlarmOperand2
func InterfaceToAlarmOperand2(i interface{}) *AlarmOperand2 {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AlarmOperand2{
		//TODO(nati): Apply default
		UveAttribute: schema.InterfaceToString(m["uve_attribute"]),
		JSONValue:    schema.InterfaceToString(m["json_value"]),
	}
}

// MakeAlarmOperand2Slice() makes a slice of AlarmOperand2
func MakeAlarmOperand2Slice() []*AlarmOperand2 {
	return []*AlarmOperand2{}
}

// InterfaceToAlarmOperand2Slice() makes a slice of AlarmOperand2
func InterfaceToAlarmOperand2Slice(i interface{}) []*AlarmOperand2 {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AlarmOperand2{}
	for _, item := range list {
		result = append(result, InterfaceToAlarmOperand2(item))
	}
	return result
}
