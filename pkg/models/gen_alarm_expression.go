package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAlarmExpression makes AlarmExpression
// nolint
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
// nolint
func InterfaceToAlarmExpression(i interface{}) *AlarmExpression {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AlarmExpression{
		//TODO(nati): Apply default
		Operations: common.InterfaceToString(m["operations"]),
		Operand1:   common.InterfaceToString(m["operand1"]),
		Variables:  common.InterfaceToStringList(m["variables"]),
		Operand2:   InterfaceToAlarmOperand2(m["operand2"]),
	}
}

// MakeAlarmExpressionSlice() makes a slice of AlarmExpression
// nolint
func MakeAlarmExpressionSlice() []*AlarmExpression {
	return []*AlarmExpression{}
}

// InterfaceToAlarmExpressionSlice() makes a slice of AlarmExpression
// nolint
func InterfaceToAlarmExpressionSlice(i interface{}) []*AlarmExpression {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AlarmExpression{}
	for _, item := range list {
		result = append(result, InterfaceToAlarmExpression(item))
	}
	return result
}
