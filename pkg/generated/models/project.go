package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeProject makes Project
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
func InterfaceToProject(i interface{}) *Project {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Project{
		//TODO(nati): Apply default
		UUID:         schema.InterfaceToString(m["uuid"]),
		ParentUUID:   schema.InterfaceToString(m["parent_uuid"]),
		ParentType:   schema.InterfaceToString(m["parent_type"]),
		FQName:       schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:      InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:  schema.InterfaceToString(m["display_name"]),
		Annotations:  InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:       InterfaceToPermType2(m["perms2"]),
		VxlanRouting: schema.InterfaceToBool(m["vxlan_routing"]),
		AlarmEnable:  schema.InterfaceToBool(m["alarm_enable"]),
		Quota:        InterfaceToQuotaType(m["quota"]),
	}
}

// MakeProjectSlice() makes a slice of Project
func MakeProjectSlice() []*Project {
	return []*Project{}
}

// InterfaceToProjectSlice() makes a slice of Project
func InterfaceToProjectSlice(i interface{}) []*Project {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Project{}
	for _, item := range list {
		result = append(result, InterfaceToProject(item))
	}
	return result
}
