package models


// MakeUserDefinedLogStat makes UserDefinedLogStat
func MakeUserDefinedLogStat() *UserDefinedLogStat{
    return &UserDefinedLogStat{
    //TODO(nati): Apply default
    Pattern: "",
        Name: "",
        
    }
}

// MakeUserDefinedLogStatSlice() makes a slice of UserDefinedLogStat
func MakeUserDefinedLogStatSlice() []*UserDefinedLogStat {
    return []*UserDefinedLogStat{}
}


