package models

// AlarmExpression

import "encoding/json"

// AlarmExpression
type AlarmExpression struct {
	Operand1  string         `json:"operand1"`
	Variables []string       `json:"variables"`
	Operand2  *AlarmOperand2 `json:"operand2"`
	Operation AlarmOperation `json:"operation"`
}

//  parents relation object

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

// InterfaceToAlarmExpression makes AlarmExpression from interface
func InterfaceToAlarmExpression(iData interface{}) *AlarmExpression {
	data := iData.(map[string]interface{})
	return &AlarmExpression{
		Variables: data["variables"].([]string),

		//{"Title":"","Description":"List of UVE attributes that would be useful when the alarm is raised. For example, user may want to raise an alarm if the NodeStatus.process_info.process_state != PROCESS_STATE_RUNNING. But, it would be useful to know the process_name whose state != PROCESS_STATE_RUNNING. This UVE attribute which is neither part of operand1 nor operand2 may be specified in variables","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Variables","GoType":"string","GoPremitive":true},"GoName":"Variables","GoType":"[]string","GoPremitive":true}
		Operand2: InterfaceToAlarmOperand2(data["operand2"]),

		//{"Title":"","Description":"UVE attribute or a json value to compare with the UVE attribute in operand1","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"json_value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JSONValue","GoType":"string","GoPremitive":true},"uve_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UveAttribute","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AlarmOperand2","CollectionType":"","Column":"","Item":null,"GoName":"Operand2","GoType":"AlarmOperand2","GoPremitive":false}
		Operation: InterfaceToAlarmOperation(data["operation"]),

		//{"Title":"","Description":"operation to compare operand1 and operand2","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["==","!=","\u003c","\u003c=","\u003e","\u003e=","in","not in","range","size==","size!="],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AlarmOperation","CollectionType":"","Column":"","Item":null,"GoName":"Operation","GoType":"AlarmOperation","GoPremitive":false}
		Operand1: data["operand1"].(string),

		//{"Title":"","Description":"UVE attribute specified in the dotted format. Example: NodeStatus.process_info.process_state","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Operand1","GoType":"string","GoPremitive":true}

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
