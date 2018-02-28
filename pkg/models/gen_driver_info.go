package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDriverInfo makes DriverInfo
// nolint
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

// MakeDriverInfo makes DriverInfo
// nolint
func InterfaceToDriverInfo(i interface{}) *DriverInfo {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DriverInfo{
		//TODO(nati): Apply default
		IpmiAddress:   common.InterfaceToString(m["ipmi_address"]),
		IpmiUsername:  common.InterfaceToString(m["ipmi_username"]),
		IpmiPassword:  common.InterfaceToString(m["ipmi_password"]),
		DeployKernel:  common.InterfaceToString(m["deploy_kernel"]),
		DeployRamdisk: common.InterfaceToString(m["deploy_ramdisk"]),
	}
}

// MakeDriverInfoSlice() makes a slice of DriverInfo
// nolint
func MakeDriverInfoSlice() []*DriverInfo {
	return []*DriverInfo{}
}

// InterfaceToDriverInfoSlice() makes a slice of DriverInfo
// nolint
func InterfaceToDriverInfoSlice(i interface{}) []*DriverInfo {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DriverInfo{}
	for _, item := range list {
		result = append(result, InterfaceToDriverInfo(item))
	}
	return result
}
