package models

// AlarmOperand2

import "encoding/json"

// AlarmOperand2
type AlarmOperand2 struct {
	UveAttribute string `json:"uve_attribute"`
	JSONValue    string `json:"json_value"`
}

//  parents relation object

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

		//{"Title":"","Description":"UVE attribute specified in the dotted format. Example: NodeStatus.process_info.process_state","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UveAttribute","GoType":"string","GoPremitive":true}
		JSONValue: data["json_value"].(string),

		//{"Title":"","Description":"json value as string","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JSONValue","GoType":"string","GoPremitive":true}

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
