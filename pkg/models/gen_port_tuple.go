package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePortTuple makes PortTuple
// nolint
func MakePortTuple() *PortTuple {
	return &PortTuple{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ConfigurationVersion: 0,
	}
}

// MakePortTuple makes PortTuple
// nolint
func InterfaceToPortTuple(i interface{}) *PortTuple {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PortTuple{
		//TODO(nati): Apply default
		UUID:                 common.InterfaceToString(m["uuid"]),
		ParentUUID:           common.InterfaceToString(m["parent_uuid"]),
		ParentType:           common.InterfaceToString(m["parent_type"]),
		FQName:               common.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          common.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion: common.InterfaceToInt64(m["configuration_version"]),
	}
}

// MakePortTupleSlice() makes a slice of PortTuple
// nolint
func MakePortTupleSlice() []*PortTuple {
	return []*PortTuple{}
}

// InterfaceToPortTupleSlice() makes a slice of PortTuple
// nolint
func InterfaceToPortTupleSlice(i interface{}) []*PortTuple {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PortTuple{}
	for _, item := range list {
		result = append(result, InterfaceToPortTuple(item))
	}
	return result
}
