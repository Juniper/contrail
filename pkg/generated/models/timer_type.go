package models

// TimerType

import "encoding/json"

// TimerType
type TimerType struct {
	OffInterval string `json:"off_interval,omitempty"`
	OnInterval  string `json:"on_interval,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
}

// String returns json representation of the object
func (model *TimerType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeTimerType makes TimerType
func MakeTimerType() *TimerType {
	return &TimerType{
		//TODO(nati): Apply default
		EndTime:     "",
		StartTime:   "",
		OffInterval: "",
		OnInterval:  "",
	}
}

// MakeTimerTypeSlice() makes a slice of TimerType
func MakeTimerTypeSlice() []*TimerType {
	return []*TimerType{}
}
