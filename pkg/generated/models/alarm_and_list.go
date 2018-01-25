package models

// AlarmAndList

// AlarmAndList
//proteus:generate
type AlarmAndList struct {
	AndList []*AlarmExpression `json:"and_list,omitempty"`
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
