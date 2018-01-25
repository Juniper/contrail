package models

// AlarmOrList

// AlarmOrList
//proteus:generate
type AlarmOrList struct {
	OrList []*AlarmAndList `json:"or_list,omitempty"`
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
