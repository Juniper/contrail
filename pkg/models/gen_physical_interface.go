package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePhysicalInterface makes PhysicalInterface
// nolint
func MakePhysicalInterface() *PhysicalInterface {
	return &PhysicalInterface{
		//TODO(nati): Apply default
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		ConfigurationVersion:      0,
		EthernetSegmentIdentifier: "",
	}
}

// MakePhysicalInterface makes PhysicalInterface
// nolint
func InterfaceToPhysicalInterface(i interface{}) *PhysicalInterface {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PhysicalInterface{
		//TODO(nati): Apply default
		UUID:                      common.InterfaceToString(m["uuid"]),
		ParentUUID:                common.InterfaceToString(m["parent_uuid"]),
		ParentType:                common.InterfaceToString(m["parent_type"]),
		FQName:                    common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               common.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:      common.InterfaceToInt64(m["configuration_version"]),
		EthernetSegmentIdentifier: common.InterfaceToString(m["ethernet_segment_identifier"]),
	}
}

// MakePhysicalInterfaceSlice() makes a slice of PhysicalInterface
// nolint
func MakePhysicalInterfaceSlice() []*PhysicalInterface {
	return []*PhysicalInterface{}
}

// InterfaceToPhysicalInterfaceSlice() makes a slice of PhysicalInterface
// nolint
func InterfaceToPhysicalInterfaceSlice(i interface{}) []*PhysicalInterface {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PhysicalInterface{}
	for _, item := range list {
		result = append(result, InterfaceToPhysicalInterface(item))
	}
	return result
}
