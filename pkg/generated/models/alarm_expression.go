package models


// MakeAlarmExpression makes AlarmExpression
func MakeAlarmExpression() *AlarmExpression{
    return &AlarmExpression{
    //TODO(nati): Apply default
    Operations: "",
        Operand1: "",
        Variables: []string{},
        Operand2: MakeAlarmOperand2(),
        
    }
}

// MakeAlarmExpressionSlice() makes a slice of AlarmExpression
func MakeAlarmExpressionSlice() []*AlarmExpression {
    return []*AlarmExpression{}
}


