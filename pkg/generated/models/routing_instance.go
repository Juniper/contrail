package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeRoutingInstance makes RoutingInstance
func MakeRoutingInstance() *RoutingInstance {
	return &RoutingInstance{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeRoutingInstance makes RoutingInstance
func InterfaceToRoutingInstance(i interface{}) *RoutingInstance {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RoutingInstance{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
	}
}

// MakeRoutingInstanceSlice() makes a slice of RoutingInstance
func MakeRoutingInstanceSlice() []*RoutingInstance {
	return []*RoutingInstance{}
}

// InterfaceToRoutingInstanceSlice() makes a slice of RoutingInstance
func InterfaceToRoutingInstanceSlice(i interface{}) []*RoutingInstance {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RoutingInstance{}
	for _, item := range list {
		result = append(result, InterfaceToRoutingInstance(item))
	}
	return result
}
