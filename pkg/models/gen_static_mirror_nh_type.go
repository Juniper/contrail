package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeStaticMirrorNhType makes StaticMirrorNhType
// nolint
func MakeStaticMirrorNhType() *StaticMirrorNhType {
	return &StaticMirrorNhType{
		//TODO(nati): Apply default
		VtepDSTIPAddress:  "",
		VtepDSTMacAddress: "",
		Vni:               0,
	}
}

// MakeStaticMirrorNhType makes StaticMirrorNhType
// nolint
func InterfaceToStaticMirrorNhType(i interface{}) *StaticMirrorNhType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &StaticMirrorNhType{
		//TODO(nati): Apply default
		VtepDSTIPAddress:  common.InterfaceToString(m["vtep_dst_ip_address"]),
		VtepDSTMacAddress: common.InterfaceToString(m["vtep_dst_mac_address"]),
		Vni:               common.InterfaceToInt64(m["vni"]),
	}
}

// MakeStaticMirrorNhTypeSlice() makes a slice of StaticMirrorNhType
// nolint
func MakeStaticMirrorNhTypeSlice() []*StaticMirrorNhType {
	return []*StaticMirrorNhType{}
}

// InterfaceToStaticMirrorNhTypeSlice() makes a slice of StaticMirrorNhType
// nolint
func InterfaceToStaticMirrorNhTypeSlice(i interface{}) []*StaticMirrorNhType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*StaticMirrorNhType{}
	for _, item := range list {
		result = append(result, InterfaceToStaticMirrorNhType(item))
	}
	return result
}
