package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

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

// MakeDriverInfo makes DriverInfo
func InterfaceToDriverInfo(i interface{}) *DriverInfo {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DriverInfo{
		//TODO(nati): Apply default
		IpmiAddress:   schema.InterfaceToString(m["ipmi_address"]),
		IpmiUsername:  schema.InterfaceToString(m["ipmi_username"]),
		IpmiPassword:  schema.InterfaceToString(m["ipmi_password"]),
		DeployKernel:  schema.InterfaceToString(m["deploy_kernel"]),
		DeployRamdisk: schema.InterfaceToString(m["deploy_ramdisk"]),
	}
}

// MakeDriverInfoSlice() makes a slice of DriverInfo
func MakeDriverInfoSlice() []*DriverInfo {
	return []*DriverInfo{}
}

// InterfaceToDriverInfoSlice() makes a slice of DriverInfo
func InterfaceToDriverInfoSlice(i interface{}) []*DriverInfo {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DriverInfo{}
	for _, item := range list {
		result = append(result, InterfaceToDriverInfo(item))
	}
	return result
}
