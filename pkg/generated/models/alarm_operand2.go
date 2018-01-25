package models

// AlarmOperand2

// AlarmOperand2
//proteus:generate
type AlarmOperand2 struct {
	UveAttribute string `json:"uve_attribute,omitempty"`
	JSONValue    string `json:"json_value,omitempty"`
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
