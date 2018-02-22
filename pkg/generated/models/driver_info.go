package models


// MakeDriverInfo makes DriverInfo
func MakeDriverInfo() *DriverInfo{
    return &DriverInfo{
    //TODO(nati): Apply default
    IpmiAddress: "",
        IpmiUsername: "",
        IpmiPassword: "",
        DeployKernel: "",
        DeployRamdisk: "",
        
    }
}

// MakeDriverInfoSlice() makes a slice of DriverInfo
func MakeDriverInfoSlice() []*DriverInfo {
    return []*DriverInfo{}
}


