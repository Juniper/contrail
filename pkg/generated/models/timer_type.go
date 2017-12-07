package models

// TimerType

import "encoding/json"

// TimerType
type TimerType struct {
	StartTime   string `json:"start_time"`
	OffInterval string `json:"off_interval"`
	OnInterval  string `json:"on_interval"`
	EndTime     string `json:"end_time"`
}

//  parents relation object

// String returns json representation of the object
func (model *TimerType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeTimerType makes TimerType
func MakeTimerType() *TimerType {
	return &TimerType{
		//TODO(nati): Apply default
		StartTime:   "",
		OffInterval: "",
		OnInterval:  "",
		EndTime:     "",
	}
}

// InterfaceToTimerType makes TimerType from interface
func InterfaceToTimerType(iData interface{}) *TimerType {
	data := iData.(map[string]interface{})
	return &TimerType{
		StartTime: data["start_time"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StartTime","GoType":"string","GoPremitive":true}
		OffInterval: data["off_interval"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"OffInterval","GoType":"string","GoPremitive":true}
		OnInterval: data["on_interval"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"OnInterval","GoType":"string","GoPremitive":true}
		EndTime: data["end_time"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EndTime","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToTimerTypeSlice makes a slice of TimerType from interface
func InterfaceToTimerTypeSlice(data interface{}) []*TimerType {
	list := data.([]interface{})
	result := MakeTimerTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToTimerType(item))
	}
	return result
}

// MakeTimerTypeSlice() makes a slice of TimerType
func MakeTimerTypeSlice() []*TimerType {
	return []*TimerType{}
}
