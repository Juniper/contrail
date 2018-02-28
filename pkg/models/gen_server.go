package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServer makes Server
// nolint
func MakeServer() *Server {
	return &Server{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		Created:     "",
		HostId:      "",
		ID:          "",
		Name:        "",
		Image:       MakeOpenStackImageProperty(),
		Flavor:      MakeOpenStackFlavorProperty(),
		Addresses:   MakeOpenStackAddress(),
		AccessIPv4:  "",
		AccessIPv6:  "",
		ConfigDrive: false,
		Progress:    0,
		Status:      "",
		HostStatus:  "",
		TenantID:    "",
		Updated:     "",
		UserID:      0,
		Locked:      false,
	}
}

// MakeServer makes Server
// nolint
func InterfaceToServer(i interface{}) *Server {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Server{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		Created:     common.InterfaceToString(m["created"]),
		HostId:      common.InterfaceToString(m["hostId"]),
		ID:          common.InterfaceToString(m["id"]),
		Name:        common.InterfaceToString(m["name"]),
		Image:       InterfaceToOpenStackImageProperty(m["image"]),
		Flavor:      InterfaceToOpenStackFlavorProperty(m["flavor"]),
		Addresses:   InterfaceToOpenStackAddress(m["addresses"]),
		AccessIPv4:  common.InterfaceToString(m["accessIPv4"]),
		AccessIPv6:  common.InterfaceToString(m["accessIPv6"]),
		ConfigDrive: common.InterfaceToBool(m["config_drive"]),
		Progress:    common.InterfaceToInt64(m["progress"]),
		Status:      common.InterfaceToString(m["status"]),
		HostStatus:  common.InterfaceToString(m["host_status"]),
		TenantID:    common.InterfaceToString(m["tenant_id"]),
		Updated:     common.InterfaceToString(m["updated"]),
		UserID:      common.InterfaceToInt64(m["user_id"]),
		Locked:      common.InterfaceToBool(m["locked"]),
	}
}

// MakeServerSlice() makes a slice of Server
// nolint
func MakeServerSlice() []*Server {
	return []*Server{}
}

// InterfaceToServerSlice() makes a slice of Server
// nolint
func InterfaceToServerSlice(i interface{}) []*Server {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Server{}
	for _, item := range list {
		result = append(result, InterfaceToServer(item))
	}
	return result
}
