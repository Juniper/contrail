package models

// AlarmOrList

import "encoding/json"

// AlarmOrList
type AlarmOrList struct {
	OrList []*AlarmAndList `json:"or_list"`
}

// String returns json representation of the object
func (model *AlarmOrList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAlarmOrList makes AlarmOrList
func MakeAlarmOrList() *AlarmOrList {
	return &AlarmOrList{
		//TODO(nati): Apply default

		OrList: MakeAlarmAndListSlice(),
	}
}

// InterfaceToAlarmOrList makes AlarmOrList from interface
func InterfaceToAlarmOrList(iData interface{}) *AlarmOrList {
	data := iData.(map[string]interface{})
	return &AlarmOrList{

		OrList: InterfaceToAlarmAndListSlice(data["or_list"]),

		//{"type":"array","item":{"type":"object","properties":{"and_list":{"type":"array","item":{"type":"object","properties":{"operand1":{"type":"string"},"operand2":{"type":"object","properties":{"json_value":{"type":"string"},"uve_attribute":{"type":"string"}}},"operation":{"type":"string","enum":["==","!=","\u003c","\u003c=","\u003e","\u003e=","in","not in","range","size==","size!="]},"variables":{"type":"array","item":{"type":"string"}}}}}}}}

	}
}

// InterfaceToAlarmOrListSlice makes a slice of AlarmOrList from interface
func InterfaceToAlarmOrListSlice(data interface{}) []*AlarmOrList {
	list := data.([]interface{})
	result := MakeAlarmOrListSlice()
	for _, item := range list {
		result = append(result, InterfaceToAlarmOrList(item))
	}
	return result
}

// MakeAlarmOrListSlice() makes a slice of AlarmOrList
func MakeAlarmOrListSlice() []*AlarmOrList {
	return []*AlarmOrList{}
}
