package models

// UserDefinedLogStat

import "encoding/json"

// UserDefinedLogStat
//proteus:generate
type UserDefinedLogStat struct {
	Pattern string `json:"pattern,omitempty"`
	Name    string `json:"name,omitempty"`
}

// String returns json representation of the object
func (model *UserDefinedLogStat) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeUserDefinedLogStat makes UserDefinedLogStat
func MakeUserDefinedLogStat() *UserDefinedLogStat {
	return &UserDefinedLogStat{
		//TODO(nati): Apply default
		Pattern: "",
		Name:    "",
	}
}

// MakeUserDefinedLogStatSlice() makes a slice of UserDefinedLogStat
func MakeUserDefinedLogStatSlice() []*UserDefinedLogStat {
	return []*UserDefinedLogStat{}
}
