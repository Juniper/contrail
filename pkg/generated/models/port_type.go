package models


// MakePortType makes PortType
func MakePortType() *PortType{
    return &PortType{
    //TODO(nati): Apply default
    EndPort: 0,
        StartPort: 0,
        
    }
}

// MakePortTypeSlice() makes a slice of PortType
func MakePortTypeSlice() []*PortType {
    return []*PortType{}
}


