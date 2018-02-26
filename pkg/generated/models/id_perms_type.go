package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeIdPermsType makes IdPermsType
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
func InterfaceToIdPermsType(i interface{}) *IdPermsType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IdPermsType{
		//TODO(nati): Apply default
		Enable:       schema.InterfaceToBool(m["enable"]),
		Description:  schema.InterfaceToString(m["description"]),
		Created:      schema.InterfaceToString(m["created"]),
		Creator:      schema.InterfaceToString(m["creator"]),
		UserVisible:  schema.InterfaceToBool(m["user_visible"]),
		LastModified: schema.InterfaceToString(m["last_modified"]),
		Permissions:  InterfaceToPermType(m["permissions"]),
	}
}

// MakeIdPermsTypeSlice() makes a slice of IdPermsType
func MakeIdPermsTypeSlice() []*IdPermsType {
	return []*IdPermsType{}
}

// InterfaceToIdPermsTypeSlice() makes a slice of IdPermsType
func InterfaceToIdPermsTypeSlice(i interface{}) []*IdPermsType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IdPermsType{}
	for _, item := range list {
		result = append(result, InterfaceToIdPermsType(item))
	}
	return result
}
