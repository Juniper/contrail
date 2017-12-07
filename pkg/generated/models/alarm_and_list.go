package models

// AlarmAndList

import "encoding/json"

// AlarmAndList
type AlarmAndList struct {
	AndList []*AlarmExpression `json:"and_list"`
}

//  parents relation object

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

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"operand1":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Operand1","GoType":"string","GoPremitive":true},"operand2":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"json_value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JSONValue","GoType":"string","GoPremitive":true},"uve_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UveAttribute","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AlarmOperand2","CollectionType":"","Column":"","Item":null,"GoName":"Operand2","GoType":"AlarmOperand2","GoPremitive":false},"operation":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["==","!=","\u003c","\u003c=","\u003e","\u003e=","in","not in","range","size==","size!="],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AlarmOperation","CollectionType":"","Column":"","Item":null,"GoName":"Operation","GoType":"AlarmOperation","GoPremitive":false},"variables":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Variables","GoType":"string","GoPremitive":true},"GoName":"Variables","GoType":"[]string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AlarmExpression","CollectionType":"","Column":"","Item":null,"GoName":"AndList","GoType":"AlarmExpression","GoPremitive":false},"GoName":"AndList","GoType":"[]*AlarmExpression","GoPremitive":true}

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
