package models

// AlarmOperand2

import "encoding/json"

// AlarmOperand2
type AlarmOperand2 struct {
	UveAttribute string `json:"uve_attribute"`
	JSONValue    string `json:"json_value"`
}

// String returns json representation of the object
func (model *AlarmOperand2) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAlarmOperand2 makes AlarmOperand2
func MakeAlarmOperand2() *AlarmOperand2 {
	return &AlarmOperand2{
		//TODO(nati): Apply default
		UveAttribute: "",
		JSONValue:    "",
	}
}

// InterfaceToAlarmOperand2 makes AlarmOperand2 from interface
func InterfaceToAlarmOperand2(iData interface{}) *AlarmOperand2 {
	data := iData.(map[string]interface{})
	return &AlarmOperand2{
		UveAttribute: data["uve_attribute"].(string),

		//{"description":"UVE attribute specified in the dotted format. Example: NodeStatus.process_info.process_state","type":"string"}
		JSONValue: data["json_value"].(string),

		//{"description":"json value as string","type":"string"}

	}
}

// InterfaceToAlarmOperand2Slice makes a slice of AlarmOperand2 from interface
func InterfaceToAlarmOperand2Slice(data interface{}) []*AlarmOperand2 {
	list := data.([]interface{})
	result := MakeAlarmOperand2Slice()
	for _, item := range list {
		result = append(result, InterfaceToAlarmOperand2(item))
	}
	return result
}

// MakeAlarmOperand2Slice() makes a slice of AlarmOperand2
func MakeAlarmOperand2Slice() []*AlarmOperand2 {
	return []*AlarmOperand2{}
}
