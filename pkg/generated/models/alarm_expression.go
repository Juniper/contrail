package models

// AlarmExpression

import "encoding/json"

// AlarmExpression
type AlarmExpression struct {
	Operand2  *AlarmOperand2 `json:"operand2"`
	Operation AlarmOperation `json:"operation"`
	Operand1  string         `json:"operand1"`
	Variables []string       `json:"variables"`
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
		Operand2:  MakeAlarmOperand2(),
		Operation: MakeAlarmOperation(),
		Operand1:  "",
		Variables: []string{},
	}
}

// InterfaceToAlarmExpression makes AlarmExpression from interface
func InterfaceToAlarmExpression(iData interface{}) *AlarmExpression {
	data := iData.(map[string]interface{})
	return &AlarmExpression{
		Operation: InterfaceToAlarmOperation(data["operation"]),

		//{"description":"operation to compare operand1 and operand2","type":"string","enum":["==","!=","\u003c","\u003c=","\u003e","\u003e=","in","not in","range","size==","size!="]}
		Operand1: data["operand1"].(string),

		//{"description":"UVE attribute specified in the dotted format. Example: NodeStatus.process_info.process_state","type":"string"}
		Variables: data["variables"].([]string),

		//{"description":"List of UVE attributes that would be useful when the alarm is raised. For example, user may want to raise an alarm if the NodeStatus.process_info.process_state != PROCESS_STATE_RUNNING. But, it would be useful to know the process_name whose state != PROCESS_STATE_RUNNING. This UVE attribute which is neither part of operand1 nor operand2 may be specified in variables","type":"array","item":{"type":"string"}}
		Operand2: InterfaceToAlarmOperand2(data["operand2"]),

		//{"description":"UVE attribute or a json value to compare with the UVE attribute in operand1","type":"object","properties":{"json_value":{"type":"string"},"uve_attribute":{"type":"string"}}}

	}
}

// InterfaceToAlarmExpressionSlice makes a slice of AlarmExpression from interface
func InterfaceToAlarmExpressionSlice(data interface{}) []*AlarmExpression {
	list := data.([]interface{})
	result := MakeAlarmExpressionSlice()
	for _, item := range list {
		result = append(result, InterfaceToAlarmExpression(item))
	}
	return result
}

// MakeAlarmExpressionSlice() makes a slice of AlarmExpression
func MakeAlarmExpressionSlice() []*AlarmExpression {
	return []*AlarmExpression{}
}
