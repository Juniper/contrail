package models


// MakeTelemetryStateInfo makes TelemetryStateInfo
func MakeTelemetryStateInfo() *TelemetryStateInfo{
    return &TelemetryStateInfo{
    //TODO(nati): Apply default
    
            
                Resource:  MakeTelemetryResourceInfoSlice(),
            
        ServerPort: 0,
        ServerIP: "",
        
    }
}

// MakeTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
    return []*TelemetryStateInfo{}
}


