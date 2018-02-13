package models

// UserDefinedLogStatList

import "encoding/json"

// UserDefinedLogStatList
//proteus:generate
type UserDefinedLogStatList struct {
	Statlist []*UserDefinedLogStat `json:"statlist,omitempty"`
}

// String returns json representation of the object
func (model *UserDefinedLogStatList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeUserDefinedLogStatList makes UserDefinedLogStatList
func MakeUserDefinedLogStatList() *UserDefinedLogStatList {
	return &UserDefinedLogStatList{
		//TODO(nati): Apply default

		Statlist: MakeUserDefinedLogStatSlice(),
	}
}

// MakeUserDefinedLogStatListSlice() makes a slice of UserDefinedLogStatList
func MakeUserDefinedLogStatListSlice() []*UserDefinedLogStatList {
	return []*UserDefinedLogStatList{}
}
