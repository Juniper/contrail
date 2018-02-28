package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeUser makes User
// nolint
func MakeUser() *User {
	return &User{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		Password:    "",
	}
}

// MakeUser makes User
// nolint
func InterfaceToUser(i interface{}) *User {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &User{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		Password:    common.InterfaceToString(m["password"]),
	}
}

// MakeUserSlice() makes a slice of User
// nolint
func MakeUserSlice() []*User {
	return []*User{}
}

// InterfaceToUserSlice() makes a slice of User
// nolint
func InterfaceToUserSlice(i interface{}) []*User {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*User{}
	for _, item := range list {
		result = append(result, InterfaceToUser(item))
	}
	return result
}
