package models


// MakePortMap makes PortMap
func MakePortMap() *PortMap{
    return &PortMap{
    //TODO(nati): Apply default
    SRCPort: 0,
        Protocol: "",
        DSTPort: 0,
        
    }
}

// MakePortMapSlice() makes a slice of PortMap
func MakePortMapSlice() []*PortMap {
    return []*PortMap{}
}


