package models

// AlarmExpression

// AlarmExpression
//proteus:generate
type AlarmExpression struct {
	Operations AlarmOperation `json:"operations,omitempty"`
	Operand1   string         `json:"operand1,omitempty"`
	Variables  []string       `json:"variables,omitempty"`
	Operand2   *AlarmOperand2 `json:"operand2,omitempty"`
}

// MakeAlarmExpression makes AlarmExpression
func MakeAlarmExpression() *AlarmExpression {
	return &AlarmExpression{
		//TODO(nati): Apply default
		Operations: MakeAlarmOperation(),
		Operand1:   "",
		Variables:  []string{},
		Operand2:   MakeAlarmOperand2(),
	}
}

// MakeAlarmExpressionSlice() makes a slice of AlarmExpression
func MakeAlarmExpressionSlice() []*AlarmExpression {
	return []*AlarmExpression{}
}
