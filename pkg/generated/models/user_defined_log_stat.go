package models

// UserDefinedLogStat

// UserDefinedLogStat
//proteus:generate
type UserDefinedLogStat struct {
	Pattern string `json:"pattern,omitempty"`
	Name    string `json:"name,omitempty"`
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
