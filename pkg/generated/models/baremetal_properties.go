package models


// MakeBaremetalProperties makes BaremetalProperties
func MakeBaremetalProperties() *BaremetalProperties{
    return &BaremetalProperties{
    //TODO(nati): Apply default
    CPUCount: 0,
        CPUArch: "",
        DiskGB: 0,
        MemoryMB: 0,
        
    }
}

// MakeBaremetalPropertiesSlice() makes a slice of BaremetalProperties
func MakeBaremetalPropertiesSlice() []*BaremetalProperties {
    return []*BaremetalProperties{}
}


