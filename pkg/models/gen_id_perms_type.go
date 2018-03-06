package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeIdPermsType makes IdPermsType
// nolint
func MakeIdPermsType() *IdPermsType {
	return &IdPermsType{
		//TODO(nati): Apply default
		Enable:       false,
		Description:  "",
		Created:      "",
		Creator:      "",
		UserVisible:  false,
		LastModified: "",
		Permissions:  MakePermType(),
	}
}

// MakeIdPermsType makes IdPermsType
// nolint
func InterfaceToIdPermsType(i interface{}) *IdPermsType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IdPermsType{
		//TODO(nati): Apply default
		Enable:       common.InterfaceToBool(m["enable"]),
		Description:  common.InterfaceToString(m["description"]),
		Created:      common.InterfaceToString(m["created"]),
		Creator:      common.InterfaceToString(m["creator"]),
		UserVisible:  common.InterfaceToBool(m["user_visible"]),
		LastModified: common.InterfaceToString(m["last_modified"]),
		Permissions:  InterfaceToPermType(m["permissions"]),
	}
}

// MakeIdPermsTypeSlice() makes a slice of IdPermsType
// nolint
func MakeIdPermsTypeSlice() []*IdPermsType {
	return []*IdPermsType{}
}

// InterfaceToIdPermsTypeSlice() makes a slice of IdPermsType
// nolint
func InterfaceToIdPermsTypeSlice(i interface{}) []*IdPermsType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IdPermsType{}
	for _, item := range list {
		result = append(result, InterfaceToIdPermsType(item))
	}
	return result
}
