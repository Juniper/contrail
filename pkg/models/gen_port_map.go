package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePortMap makes PortMap
// nolint
func MakePortMap() *PortMap {
	return &PortMap{
		//TODO(nati): Apply default
		SRCPort:  0,
		Protocol: "",
		DSTPort:  0,
	}
}

// MakePortMap makes PortMap
// nolint
func InterfaceToPortMap(i interface{}) *PortMap {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PortMap{
		//TODO(nati): Apply default
		SRCPort:  common.InterfaceToInt64(m["src_port"]),
		Protocol: common.InterfaceToString(m["protocol"]),
		DSTPort:  common.InterfaceToInt64(m["dst_port"]),
	}
}

// MakePortMapSlice() makes a slice of PortMap
// nolint
func MakePortMapSlice() []*PortMap {
	return []*PortMap{}
}

// InterfaceToPortMapSlice() makes a slice of PortMap
// nolint
func InterfaceToPortMapSlice(i interface{}) []*PortMap {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PortMap{}
	for _, item := range list {
		result = append(result, InterfaceToPortMap(item))
	}
	return result
}
