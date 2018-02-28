package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeBaremetalPort makes BaremetalPort
// nolint
func MakeBaremetalPort() *BaremetalPort {
	return &BaremetalPort{
		//TODO(nati): Apply default
		UUID:                "",
		ParentUUID:          "",
		ParentType:          "",
		FQName:              []string{},
		IDPerms:             MakeIdPermsType(),
		DisplayName:         "",
		Annotations:         MakeKeyValuePairs(),
		Perms2:              MakePermType2(),
		MacAddress:          "",
		CreatedAt:           "",
		UpdatedAt:           "",
		Node:                "",
		PxeEnabled:          false,
		LocalLinkConnection: MakeLocalLinkConnection(),
	}
}

// MakeBaremetalPort makes BaremetalPort
// nolint
func InterfaceToBaremetalPort(i interface{}) *BaremetalPort {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BaremetalPort{
		//TODO(nati): Apply default
		UUID:                common.InterfaceToString(m["uuid"]),
		ParentUUID:          common.InterfaceToString(m["parent_uuid"]),
		ParentType:          common.InterfaceToString(m["parent_type"]),
		FQName:              common.InterfaceToStringList(m["fq_name"]),
		IDPerms:             InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:         common.InterfaceToString(m["display_name"]),
		Annotations:         InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:              InterfaceToPermType2(m["perms2"]),
		MacAddress:          common.InterfaceToString(m["mac_address"]),
		CreatedAt:           common.InterfaceToString(m["created_at"]),
		UpdatedAt:           common.InterfaceToString(m["updated_at"]),
		Node:                common.InterfaceToString(m["node"]),
		PxeEnabled:          common.InterfaceToBool(m["pxe_enabled"]),
		LocalLinkConnection: InterfaceToLocalLinkConnection(m["local_link_connection"]),
	}
}

// MakeBaremetalPortSlice() makes a slice of BaremetalPort
// nolint
func MakeBaremetalPortSlice() []*BaremetalPort {
	return []*BaremetalPort{}
}

// InterfaceToBaremetalPortSlice() makes a slice of BaremetalPort
// nolint
func InterfaceToBaremetalPortSlice(i interface{}) []*BaremetalPort {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BaremetalPort{}
	for _, item := range list {
		result = append(result, InterfaceToBaremetalPort(item))
	}
	return result
}
