package models

// AlarmOrList

import "encoding/json"

// AlarmOrList
type AlarmOrList struct {
	OrList []*AlarmAndList `json:"or_list,omitempty"`
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

// MakeAlarmOrListSlice() makes a slice of AlarmOrList
func MakeAlarmOrListSlice() []*AlarmOrList {
	return []*AlarmOrList{}
}
