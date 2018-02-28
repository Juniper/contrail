package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeBGPaaServiceParametersType makes BGPaaServiceParametersType
// nolint
func MakeBGPaaServiceParametersType() *BGPaaServiceParametersType {
	return &BGPaaServiceParametersType{
		//TODO(nati): Apply default
		PortStart: 0,
		PortEnd:   0,
	}
}

// MakeBGPaaServiceParametersType makes BGPaaServiceParametersType
// nolint
func InterfaceToBGPaaServiceParametersType(i interface{}) *BGPaaServiceParametersType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BGPaaServiceParametersType{
		//TODO(nati): Apply default
		PortStart: common.InterfaceToInt64(m["port_start"]),
		PortEnd:   common.InterfaceToInt64(m["port_end"]),
	}
}

// MakeBGPaaServiceParametersTypeSlice() makes a slice of BGPaaServiceParametersType
// nolint
func MakeBGPaaServiceParametersTypeSlice() []*BGPaaServiceParametersType {
	return []*BGPaaServiceParametersType{}
}

// InterfaceToBGPaaServiceParametersTypeSlice() makes a slice of BGPaaServiceParametersType
// nolint
func InterfaceToBGPaaServiceParametersTypeSlice(i interface{}) []*BGPaaServiceParametersType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BGPaaServiceParametersType{}
	for _, item := range list {
		result = append(result, InterfaceToBGPaaServiceParametersType(item))
	}
	return result
}
