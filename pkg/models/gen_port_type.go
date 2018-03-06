package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePortType makes PortType
// nolint
func MakePortType() *PortType {
	return &PortType{
		//TODO(nati): Apply default
		EndPort:   0,
		StartPort: 0,
	}
}

// MakePortType makes PortType
// nolint
func InterfaceToPortType(i interface{}) *PortType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PortType{
		//TODO(nati): Apply default
		EndPort:   common.InterfaceToInt64(m["end_port"]),
		StartPort: common.InterfaceToInt64(m["start_port"]),
	}
}

// MakePortTypeSlice() makes a slice of PortType
// nolint
func MakePortTypeSlice() []*PortType {
	return []*PortType{}
}

// InterfaceToPortTypeSlice() makes a slice of PortType
// nolint
func InterfaceToPortTypeSlice(i interface{}) []*PortType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PortType{}
	for _, item := range list {
		result = append(result, InterfaceToPortType(item))
	}
	return result
}
