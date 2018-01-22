package models

// AlarmExpression

import "encoding/json"

// AlarmExpression
type AlarmExpression struct {
	Operation AlarmOperation `json:"operation,omitempty"`
	Operand1  string         `json:"operand1,omitempty"`
	Variables []string       `json:"variables,omitempty"`
	Operand2  *AlarmOperand2 `json:"operand2,omitempty"`
}

// String returns json representation of the object
func (model *AlarmExpression) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAlarmExpression makes AlarmExpression
func MakeAlarmExpression() *AlarmExpression {
	return &AlarmExpression{
		//TODO(nati): Apply default
		Operand1:  "",
		Variables: []string{},
		Operand2:  MakeAlarmOperand2(),
		Operation: MakeAlarmOperation(),
	}
}

// MakeAlarmExpressionSlice() makes a slice of AlarmExpression
func MakeAlarmExpressionSlice() []*AlarmExpression {
	return []*AlarmExpression{}
}
