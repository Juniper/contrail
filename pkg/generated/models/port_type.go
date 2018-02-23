package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakePortType makes PortType
func MakePortType() *PortType {
	return &PortType{
		//TODO(nati): Apply default
		EndPort:   0,
		StartPort: 0,
	}
}

// MakePortType makes PortType
func InterfaceToPortType(i interface{}) *PortType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PortType{
		//TODO(nati): Apply default
		EndPort:   schema.InterfaceToInt64(m["end_port"]),
		StartPort: schema.InterfaceToInt64(m["start_port"]),
	}
}

// MakePortTypeSlice() makes a slice of PortType
func MakePortTypeSlice() []*PortType {
	return []*PortType{}
}

// InterfaceToPortTypeSlice() makes a slice of PortType
func InterfaceToPortTypeSlice(i interface{}) []*PortType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PortType{}
	for _, item := range list {
		result = append(result, InterfaceToPortType(item))
	}
	return result
}
