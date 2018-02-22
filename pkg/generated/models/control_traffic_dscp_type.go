package models


// MakeControlTrafficDscpType makes ControlTrafficDscpType
func MakeControlTrafficDscpType() *ControlTrafficDscpType{
    return &ControlTrafficDscpType{
    //TODO(nati): Apply default
    Control: 0,
        Analytics: 0,
        DNS: 0,
        
    }
}

// MakeControlTrafficDscpTypeSlice() makes a slice of ControlTrafficDscpType
func MakeControlTrafficDscpTypeSlice() []*ControlTrafficDscpType {
    return []*ControlTrafficDscpType{}
}


