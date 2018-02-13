package models

// AlarmAndList

import "encoding/json"

// AlarmAndList
//proteus:generate
type AlarmAndList struct {
	AndList []*AlarmExpression `json:"and_list,omitempty"`
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

// MakeAlarmAndListSlice() makes a slice of AlarmAndList
func MakeAlarmAndListSlice() []*AlarmAndList {
	return []*AlarmAndList{}
}
