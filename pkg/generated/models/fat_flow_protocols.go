package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeFatFlowProtocols makes FatFlowProtocols
func MakeFatFlowProtocols() *FatFlowProtocols {
	return &FatFlowProtocols{
		//TODO(nati): Apply default

		FatFlowProtocol: MakeProtocolTypeSlice(),
	}
}

// MakeFatFlowProtocols makes FatFlowProtocols
func InterfaceToFatFlowProtocols(i interface{}) *FatFlowProtocols {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FatFlowProtocols{
		//TODO(nati): Apply default

		FatFlowProtocol: InterfaceToProtocolTypeSlice(m["fat_flow_protocol"]),
	}
}

// MakeFatFlowProtocolsSlice() makes a slice of FatFlowProtocols
func MakeFatFlowProtocolsSlice() []*FatFlowProtocols {
	return []*FatFlowProtocols{}
}

// InterfaceToFatFlowProtocolsSlice() makes a slice of FatFlowProtocols
func InterfaceToFatFlowProtocolsSlice(i interface{}) []*FatFlowProtocols {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FatFlowProtocols{}
	for _, item := range list {
		result = append(result, InterfaceToFatFlowProtocols(item))
	}
	return result
}
