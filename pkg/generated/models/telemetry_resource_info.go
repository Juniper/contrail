package models


// MakeTelemetryResourceInfo makes TelemetryResourceInfo
func MakeTelemetryResourceInfo() *TelemetryResourceInfo{
    return &TelemetryResourceInfo{
    //TODO(nati): Apply default
    Path: "",
        Rate: "",
        Name: "",
        
    }
}

// MakeTelemetryResourceInfoSlice() makes a slice of TelemetryResourceInfo
func MakeTelemetryResourceInfoSlice() []*TelemetryResourceInfo {
    return []*TelemetryResourceInfo{}
}


