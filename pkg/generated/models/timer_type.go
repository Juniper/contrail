package models


// MakeTimerType makes TimerType
func MakeTimerType() *TimerType{
    return &TimerType{
    //TODO(nati): Apply default
    StartTime: "",
        OffInterval: "",
        OnInterval: "",
        EndTime: "",
        
    }
}

// MakeTimerTypeSlice() makes a slice of TimerType
func MakeTimerTypeSlice() []*TimerType {
    return []*TimerType{}
}


