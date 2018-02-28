package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDashboard makes Dashboard
// nolint
func MakeDashboard() *Dashboard {
	return &Dashboard{
		//TODO(nati): Apply default
		UUID:            "",
		ParentUUID:      "",
		ParentType:      "",
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		ContainerConfig: "",
	}
}

// MakeDashboard makes Dashboard
// nolint
func InterfaceToDashboard(i interface{}) *Dashboard {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Dashboard{
		//TODO(nati): Apply default
		UUID:            common.InterfaceToString(m["uuid"]),
		ParentUUID:      common.InterfaceToString(m["parent_uuid"]),
		ParentType:      common.InterfaceToString(m["parent_type"]),
		FQName:          common.InterfaceToStringList(m["fq_name"]),
		IDPerms:         InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:     common.InterfaceToString(m["display_name"]),
		Annotations:     InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:          InterfaceToPermType2(m["perms2"]),
		ContainerConfig: common.InterfaceToString(m["container_config"]),
	}
}

// MakeDashboardSlice() makes a slice of Dashboard
// nolint
func MakeDashboardSlice() []*Dashboard {
	return []*Dashboard{}
}

// InterfaceToDashboardSlice() makes a slice of Dashboard
// nolint
func InterfaceToDashboardSlice(i interface{}) []*Dashboard {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Dashboard{}
	for _, item := range list {
		result = append(result, InterfaceToDashboard(item))
	}
	return result
}
