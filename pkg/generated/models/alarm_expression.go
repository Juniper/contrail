package models

// AlarmExpression

import "encoding/json"

// AlarmExpression
type AlarmExpression struct {
	Operation AlarmOperation `json:"operation"`
	Operand1  string         `json:"operand1"`
	Variables []string       `json:"variables"`
	Operand2  *AlarmOperand2 `json:"operand2"`
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
		Operation: MakeAlarmOperation(),
		Operand1:  "",
		Variables: []string{},
		Operand2:  MakeAlarmOperand2(),
	}
}

// MakeAlarmExpressionSlice() makes a slice of AlarmExpression
func MakeAlarmExpressionSlice() []*AlarmExpression {
	return []*AlarmExpression{}
}
