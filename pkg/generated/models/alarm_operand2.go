package models

// AlarmOperand2

import "encoding/json"

type AlarmOperand2 struct {
	UveAttribute string `json:"uve_attribute"`
	JSONValue    string `json:"json_value"`
}

func (model *AlarmOperand2) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeAlarmOperand2() *AlarmOperand2 {
	return &AlarmOperand2{
		//TODO(nati): Apply default
		UveAttribute: "",
		JSONValue:    "",
	}
}

func InterfaceToAlarmOperand2(iData interface{}) *AlarmOperand2 {
	data := iData.(map[string]interface{})
	return &AlarmOperand2{
		UveAttribute: data["uve_attribute"].(string),

		//{"Title":"","Description":"UVE attribute specified in the dotted format. Example: NodeStatus.process_info.process_state","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UveAttribute","GoType":"string"}
		JSONValue: data["json_value"].(string),

		//{"Title":"","Description":"json value as string","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JSONValue","GoType":"string"}

	}
}

func InterfaceToAlarmOperand2Slice(data interface{}) []*AlarmOperand2 {
	list := data.([]interface{})
	result := MakeAlarmOperand2Slice()
	for _, item := range list {
		result = append(result, InterfaceToAlarmOperand2(item))
	}
	return result
}

func MakeAlarmOperand2Slice() []*AlarmOperand2 {
	return []*AlarmOperand2{}
}
