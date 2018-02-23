package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeServer makes Server
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
func InterfaceToServer(i interface{}) *Server {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Server{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		Created:     schema.InterfaceToString(m["created"]),
		HostId:      schema.InterfaceToString(m["hostId"]),
		ID:          schema.InterfaceToString(m["id"]),
		Name:        schema.InterfaceToString(m["name"]),
		Image:       InterfaceToOpenStackImageProperty(m["image"]),
		Flavor:      InterfaceToOpenStackFlavorProperty(m["flavor"]),
		Addresses:   InterfaceToOpenStackAddress(m["addresses"]),
		AccessIPv4:  schema.InterfaceToString(m["accessIPv4"]),
		AccessIPv6:  schema.InterfaceToString(m["accessIPv6"]),
		ConfigDrive: schema.InterfaceToBool(m["config_drive"]),
		Progress:    schema.InterfaceToInt64(m["progress"]),
		Status:      schema.InterfaceToString(m["status"]),
		HostStatus:  schema.InterfaceToString(m["host_status"]),
		TenantID:    schema.InterfaceToString(m["tenant_id"]),
		Updated:     schema.InterfaceToString(m["updated"]),
		UserID:      schema.InterfaceToInt64(m["user_id"]),
		Locked:      schema.InterfaceToBool(m["locked"]),
	}
}

// MakeServerSlice() makes a slice of Server
func MakeServerSlice() []*Server {
	return []*Server{}
}

// InterfaceToServerSlice() makes a slice of Server
func InterfaceToServerSlice(i interface{}) []*Server {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Server{}
	for _, item := range list {
		result = append(result, InterfaceToServer(item))
	}
	return result
}
