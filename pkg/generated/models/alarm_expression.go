package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeAlarmExpression makes AlarmExpression
func MakeAlarmExpression() *AlarmExpression {
	return &AlarmExpression{
		//TODO(nati): Apply default
		Operations: "",
		Operand1:   "",
		Variables:  []string{},
		Operand2:   MakeAlarmOperand2(),
	}
}

// MakeAlarmExpression makes AlarmExpression
func InterfaceToAlarmExpression(i interface{}) *AlarmExpression {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AlarmExpression{
		//TODO(nati): Apply default
		Operations: schema.InterfaceToString(m["operations"]),
		Operand1:   schema.InterfaceToString(m["operand1"]),
		Variables:  schema.InterfaceToStringList(m["variables"]),
		Operand2:   InterfaceToAlarmOperand2(m["operand2"]),
	}
}

// MakeAlarmExpressionSlice() makes a slice of AlarmExpression
func MakeAlarmExpressionSlice() []*AlarmExpression {
	return []*AlarmExpression{}
}

// InterfaceToAlarmExpressionSlice() makes a slice of AlarmExpression
func InterfaceToAlarmExpressionSlice(i interface{}) []*AlarmExpression {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AlarmExpression{}
	for _, item := range list {
		result = append(result, InterfaceToAlarmExpression(item))
	}
	return result
}
