package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeInterfaceMirrorType makes InterfaceMirrorType
// nolint
func MakeInterfaceMirrorType() *InterfaceMirrorType {
	return &InterfaceMirrorType{
		//TODO(nati): Apply default
		TrafficDirection: "",
		MirrorTo:         MakeMirrorActionType(),
	}
}

// MakeInterfaceMirrorType makes InterfaceMirrorType
// nolint
func InterfaceToInterfaceMirrorType(i interface{}) *InterfaceMirrorType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &InterfaceMirrorType{
		//TODO(nati): Apply default
		TrafficDirection: common.InterfaceToString(m["traffic_direction"]),
		MirrorTo:         InterfaceToMirrorActionType(m["mirror_to"]),
	}
}

// MakeInterfaceMirrorTypeSlice() makes a slice of InterfaceMirrorType
// nolint
func MakeInterfaceMirrorTypeSlice() []*InterfaceMirrorType {
	return []*InterfaceMirrorType{}
}

// InterfaceToInterfaceMirrorTypeSlice() makes a slice of InterfaceMirrorType
// nolint
func InterfaceToInterfaceMirrorTypeSlice(i interface{}) []*InterfaceMirrorType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*InterfaceMirrorType{}
	for _, item := range list {
		result = append(result, InterfaceToInterfaceMirrorType(item))
	}
	return result
}
