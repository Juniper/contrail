package models

// TimerType

// TimerType
//proteus:generate
type TimerType struct {
	StartTime   string `json:"start_time,omitempty"`
	OffInterval string `json:"off_interval,omitempty"`
	OnInterval  string `json:"on_interval,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
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

// MakeTimerTypeSlice() makes a slice of TimerType
func MakeTimerTypeSlice() []*TimerType {
	return []*TimerType{}
}
