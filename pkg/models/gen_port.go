package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePort makes Port
// nolint
func MakePort() *Port {
	return &Port{
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
		MacAddress:           "",
		NodeUUID:             "",
		PxeEnabled:           false,
		LocalLinkConnection:  MakeLocalLinkConnection(),
	}
}

// MakePort makes Port
// nolint
func InterfaceToPort(i interface{}) *Port {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Port{
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
		MacAddress:           common.InterfaceToString(m["mac_address"]),
		NodeUUID:             common.InterfaceToString(m["node_uuid"]),
		PxeEnabled:           common.InterfaceToBool(m["pxe_enabled"]),
		LocalLinkConnection:  InterfaceToLocalLinkConnection(m["local_link_connection"]),
	}
}

// MakePortSlice() makes a slice of Port
// nolint
func MakePortSlice() []*Port {
	return []*Port{}
}

// InterfaceToPortSlice() makes a slice of Port
// nolint
func InterfaceToPortSlice(i interface{}) []*Port {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Port{}
	for _, item := range list {
		result = append(result, InterfaceToPort(item))
	}
	return result
}
