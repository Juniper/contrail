package models

// TimerType

import "encoding/json"

type TimerType struct {
	OnInterval  string `json:"on_interval"`
	EndTime     string `json:"end_time"`
	StartTime   string `json:"start_time"`
	OffInterval string `json:"off_interval"`
}

func (model *TimerType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeTimerType() *TimerType {
	return &TimerType{
		//TODO(nati): Apply default
		OffInterval: "",
		OnInterval:  "",
		EndTime:     "",
		StartTime:   "",
	}
}

func InterfaceToTimerType(iData interface{}) *TimerType {
	data := iData.(map[string]interface{})
	return &TimerType{
		OffInterval: data["off_interval"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"OffInterval","GoType":"string"}
		OnInterval: data["on_interval"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"OnInterval","GoType":"string"}
		EndTime: data["end_time"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EndTime","GoType":"string"}
		StartTime: data["start_time"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StartTime","GoType":"string"}

	}
}

func InterfaceToTimerTypeSlice(data interface{}) []*TimerType {
	list := data.([]interface{})
	result := MakeTimerTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToTimerType(item))
	}
	return result
}

func MakeTimerTypeSlice() []*TimerType {
	return []*TimerType{}
}
