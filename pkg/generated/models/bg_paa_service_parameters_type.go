package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeBGPaaServiceParametersType makes BGPaaServiceParametersType
func MakeBGPaaServiceParametersType() *BGPaaServiceParametersType {
	return &BGPaaServiceParametersType{
		//TODO(nati): Apply default
		PortStart: 0,
		PortEnd:   0,
	}
}

// MakeBGPaaServiceParametersType makes BGPaaServiceParametersType
func InterfaceToBGPaaServiceParametersType(i interface{}) *BGPaaServiceParametersType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BGPaaServiceParametersType{
		//TODO(nati): Apply default
		PortStart: schema.InterfaceToInt64(m["port_start"]),
		PortEnd:   schema.InterfaceToInt64(m["port_end"]),
	}
}

// MakeBGPaaServiceParametersTypeSlice() makes a slice of BGPaaServiceParametersType
func MakeBGPaaServiceParametersTypeSlice() []*BGPaaServiceParametersType {
	return []*BGPaaServiceParametersType{}
}

// InterfaceToBGPaaServiceParametersTypeSlice() makes a slice of BGPaaServiceParametersType
func InterfaceToBGPaaServiceParametersTypeSlice(i interface{}) []*BGPaaServiceParametersType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BGPaaServiceParametersType{}
	for _, item := range list {
		result = append(result, InterfaceToBGPaaServiceParametersType(item))
	}
	return result
}
