package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceObject makes ServiceObject
// nolint
func MakeServiceObject() *ServiceObject {
	return &ServiceObject{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeServiceObject makes ServiceObject
// nolint
func InterfaceToServiceObject(i interface{}) *ServiceObject {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceObject{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
	}
}

// MakeServiceObjectSlice() makes a slice of ServiceObject
// nolint
func MakeServiceObjectSlice() []*ServiceObject {
	return []*ServiceObject{}
}

// InterfaceToServiceObjectSlice() makes a slice of ServiceObject
// nolint
func InterfaceToServiceObjectSlice(i interface{}) []*ServiceObject {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceObject{}
	for _, item := range list {
		result = append(result, InterfaceToServiceObject(item))
	}
	return result
}
