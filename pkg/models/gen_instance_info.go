package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeInstanceInfo makes InstanceInfo
// nolint
func MakeInstanceInfo() *InstanceInfo {
	return &InstanceInfo{
		//TODO(nati): Apply default
		DisplayName:  "",
		ImageSource:  "",
		LocalGB:      "",
		MemoryMB:     "",
		NovaHostID:   "",
		RootGB:       "",
		SwapMB:       "",
		Vcpus:        "",
		Capabilities: "",
	}
}

// MakeInstanceInfo makes InstanceInfo
// nolint
func InterfaceToInstanceInfo(i interface{}) *InstanceInfo {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &InstanceInfo{
		//TODO(nati): Apply default
		DisplayName:  common.InterfaceToString(m["display_name"]),
		ImageSource:  common.InterfaceToString(m["image_source"]),
		LocalGB:      common.InterfaceToString(m["local_gb"]),
		MemoryMB:     common.InterfaceToString(m["memory_mb"]),
		NovaHostID:   common.InterfaceToString(m["nova_host_id"]),
		RootGB:       common.InterfaceToString(m["root_gb"]),
		SwapMB:       common.InterfaceToString(m["swap_mb"]),
		Vcpus:        common.InterfaceToString(m["vcpus"]),
		Capabilities: common.InterfaceToString(m["capabilities"]),
	}
}

// MakeInstanceInfoSlice() makes a slice of InstanceInfo
// nolint
func MakeInstanceInfoSlice() []*InstanceInfo {
	return []*InstanceInfo{}
}

// InterfaceToInstanceInfoSlice() makes a slice of InstanceInfo
// nolint
func InterfaceToInstanceInfoSlice(i interface{}) []*InstanceInfo {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*InstanceInfo{}
	for _, item := range list {
		result = append(result, InterfaceToInstanceInfo(item))
	}
	return result
}
