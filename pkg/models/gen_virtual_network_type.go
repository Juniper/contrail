package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualNetworkType makes VirtualNetworkType
// nolint
func MakeVirtualNetworkType() *VirtualNetworkType {
	return &VirtualNetworkType{
		//TODO(nati): Apply default
		ForwardingMode:         "",
		AllowTransit:           false,
		NetworkID:              0,
		MirrorDestination:      false,
		VxlanNetworkIdentifier: 0,
		RPF: "",
	}
}

// MakeVirtualNetworkType makes VirtualNetworkType
// nolint
func InterfaceToVirtualNetworkType(i interface{}) *VirtualNetworkType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualNetworkType{
		//TODO(nati): Apply default
		ForwardingMode:         common.InterfaceToString(m["forwarding_mode"]),
		AllowTransit:           common.InterfaceToBool(m["allow_transit"]),
		NetworkID:              common.InterfaceToInt64(m["network_id"]),
		MirrorDestination:      common.InterfaceToBool(m["mirror_destination"]),
		VxlanNetworkIdentifier: common.InterfaceToInt64(m["vxlan_network_identifier"]),
		RPF: common.InterfaceToString(m["rpf"]),
	}
}

// MakeVirtualNetworkTypeSlice() makes a slice of VirtualNetworkType
// nolint
func MakeVirtualNetworkTypeSlice() []*VirtualNetworkType {
	return []*VirtualNetworkType{}
}

// InterfaceToVirtualNetworkTypeSlice() makes a slice of VirtualNetworkType
// nolint
func InterfaceToVirtualNetworkTypeSlice(i interface{}) []*VirtualNetworkType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualNetworkType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetworkType(item))
	}
	return result
}
