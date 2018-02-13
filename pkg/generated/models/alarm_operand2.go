package models

// AlarmOperand2

import "encoding/json"

// AlarmOperand2
//proteus:generate
type AlarmOperand2 struct {
	UveAttribute string `json:"uve_attribute,omitempty"`
	JSONValue    string `json:"json_value,omitempty"`
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

// MakeAlarmOperand2Slice() makes a slice of AlarmOperand2
func MakeAlarmOperand2Slice() []*AlarmOperand2 {
	return []*AlarmOperand2{}
}
