package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePortMappings makes PortMappings
// nolint
func MakePortMappings() *PortMappings {
	return &PortMappings{
		//TODO(nati): Apply default

		PortMappings: MakePortMapSlice(),
	}
}

// MakePortMappings makes PortMappings
// nolint
func InterfaceToPortMappings(i interface{}) *PortMappings {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PortMappings{
		//TODO(nati): Apply default

		PortMappings: InterfaceToPortMapSlice(m["port_mappings"]),
	}
}

// MakePortMappingsSlice() makes a slice of PortMappings
// nolint
func MakePortMappingsSlice() []*PortMappings {
	return []*PortMappings{}
}

// InterfaceToPortMappingsSlice() makes a slice of PortMappings
// nolint
func InterfaceToPortMappingsSlice(i interface{}) []*PortMappings {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PortMappings{}
	for _, item := range list {
		result = append(result, InterfaceToPortMappings(item))
	}
	return result
}
