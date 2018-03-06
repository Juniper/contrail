package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeBaremetalProperties makes BaremetalProperties
// nolint
func MakeBaremetalProperties() *BaremetalProperties {
	return &BaremetalProperties{
		//TODO(nati): Apply default
		CPUCount: 0,
		CPUArch:  "",
		DiskGB:   0,
		MemoryMB: 0,
	}
}

// MakeBaremetalProperties makes BaremetalProperties
// nolint
func InterfaceToBaremetalProperties(i interface{}) *BaremetalProperties {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BaremetalProperties{
		//TODO(nati): Apply default
		CPUCount: common.InterfaceToInt64(m["cpu_count"]),
		CPUArch:  common.InterfaceToString(m["cpu_arch"]),
		DiskGB:   common.InterfaceToInt64(m["disk_gb"]),
		MemoryMB: common.InterfaceToInt64(m["memory_mb"]),
	}
}

// MakeBaremetalPropertiesSlice() makes a slice of BaremetalProperties
// nolint
func MakeBaremetalPropertiesSlice() []*BaremetalProperties {
	return []*BaremetalProperties{}
}

// InterfaceToBaremetalPropertiesSlice() makes a slice of BaremetalProperties
// nolint
func InterfaceToBaremetalPropertiesSlice(i interface{}) []*BaremetalProperties {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BaremetalProperties{}
	for _, item := range list {
		result = append(result, InterfaceToBaremetalProperties(item))
	}
	return result
}
