package models


// MakeUserDefinedLogStatList makes UserDefinedLogStatList
func MakeUserDefinedLogStatList() *UserDefinedLogStatList{
    return &UserDefinedLogStatList{
    //TODO(nati): Apply default
    
            
                Statlist:  MakeUserDefinedLogStatSlice(),
            
        
    }
}

// MakeUserDefinedLogStatListSlice() makes a slice of UserDefinedLogStatList
func MakeUserDefinedLogStatListSlice() []*UserDefinedLogStatList {
    return []*UserDefinedLogStatList{}
}


