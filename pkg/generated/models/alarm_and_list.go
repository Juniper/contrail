package models

// AlarmAndList

import "encoding/json"

// AlarmAndList
type AlarmAndList struct {
	AndList []*AlarmExpression `json:"and_list"`
}

// String returns json representation of the object
func (model *AlarmAndList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAlarmAndList makes AlarmAndList
func MakeAlarmAndList() *AlarmAndList {
	return &AlarmAndList{
		//TODO(nati): Apply default

		AndList: MakeAlarmExpressionSlice(),
	}
}

// InterfaceToAlarmAndList makes AlarmAndList from interface
func InterfaceToAlarmAndList(iData interface{}) *AlarmAndList {
	data := iData.(map[string]interface{})
	return &AlarmAndList{

		AndList: InterfaceToAlarmExpressionSlice(data["and_list"]),

		//{"type":"array","item":{"type":"object","properties":{"operand1":{"type":"string"},"operand2":{"type":"object","properties":{"json_value":{"type":"string"},"uve_attribute":{"type":"string"}}},"operation":{"type":"string","enum":["==","!=","\u003c","\u003c=","\u003e","\u003e=","in","not in","range","size==","size!="]},"variables":{"type":"array","item":{"type":"string"}}}}}

	}
}

// InterfaceToAlarmAndListSlice makes a slice of AlarmAndList from interface
func InterfaceToAlarmAndListSlice(data interface{}) []*AlarmAndList {
	list := data.([]interface{})
	result := MakeAlarmAndListSlice()
	for _, item := range list {
		result = append(result, InterfaceToAlarmAndList(item))
	}
	return result
}

// MakeAlarmAndListSlice() makes a slice of AlarmAndList
func MakeAlarmAndListSlice() []*AlarmAndList {
	return []*AlarmAndList{}
}
