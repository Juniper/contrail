package models

// DriverInfo

// DriverInfo
//proteus:generate
type DriverInfo struct {
	IpmiAddress   string `json:"ipmi_address,omitempty"`
	IpmiUsername  string `json:"ipmi_username,omitempty"`
	IpmiPassword  string `json:"ipmi_password,omitempty"`
	DeployKernel  string `json:"deploy_kernel,omitempty"`
	DeployRamdisk string `json:"deploy_ramdisk,omitempty"`
}

// MakeDriverInfo makes DriverInfo
func MakeDriverInfo() *DriverInfo {
	return &DriverInfo{
		//TODO(nati): Apply default
		IpmiAddress:   "",
		IpmiUsername:  "",
		IpmiPassword:  "",
		DeployKernel:  "",
		DeployRamdisk: "",
	}
}

// MakeDriverInfoSlice() makes a slice of DriverInfo
func MakeDriverInfoSlice() []*DriverInfo {
	return []*DriverInfo{}
}
