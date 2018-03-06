package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeProject makes Project
// nolint
func MakeProject() *Project {
	return &Project{
		//TODO(nati): Apply default
		UUID:         "",
		ParentUUID:   "",
		ParentType:   "",
		FQName:       []string{},
		IDPerms:      MakeIdPermsType(),
		DisplayName:  "",
		Annotations:  MakeKeyValuePairs(),
		Perms2:       MakePermType2(),
		VxlanRouting: false,
		AlarmEnable:  false,
		Quota:        MakeQuotaType(),
	}
}

// MakeProject makes Project
// nolint
func InterfaceToProject(i interface{}) *Project {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Project{
		//TODO(nati): Apply default
		UUID:         common.InterfaceToString(m["uuid"]),
		ParentUUID:   common.InterfaceToString(m["parent_uuid"]),
		ParentType:   common.InterfaceToString(m["parent_type"]),
		FQName:       common.InterfaceToStringList(m["fq_name"]),
		IDPerms:      InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:  common.InterfaceToString(m["display_name"]),
		Annotations:  InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:       InterfaceToPermType2(m["perms2"]),
		VxlanRouting: common.InterfaceToBool(m["vxlan_routing"]),
		AlarmEnable:  common.InterfaceToBool(m["alarm_enable"]),
		Quota:        InterfaceToQuotaType(m["quota"]),
	}
}

// MakeProjectSlice() makes a slice of Project
// nolint
func MakeProjectSlice() []*Project {
	return []*Project{}
}

// InterfaceToProjectSlice() makes a slice of Project
// nolint
func InterfaceToProjectSlice(i interface{}) []*Project {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Project{}
	for _, item := range list {
		result = append(result, InterfaceToProject(item))
	}
	return result
}
