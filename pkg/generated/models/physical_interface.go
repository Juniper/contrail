package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakePhysicalInterface makes PhysicalInterface
func MakePhysicalInterface() *PhysicalInterface {
	return &PhysicalInterface{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		EthernetSegmentIdentifier: "",
	}
}

// MakePhysicalInterface makes PhysicalInterface
func InterfaceToPhysicalInterface(i interface{}) *PhysicalInterface {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PhysicalInterface{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		EthernetSegmentIdentifier: schema.InterfaceToString(m["ethernet_segment_identifier"]),
	}
}

// MakePhysicalInterfaceSlice() makes a slice of PhysicalInterface
func MakePhysicalInterfaceSlice() []*PhysicalInterface {
	return []*PhysicalInterface{}
}

// InterfaceToPhysicalInterfaceSlice() makes a slice of PhysicalInterface
func InterfaceToPhysicalInterfaceSlice(i interface{}) []*PhysicalInterface {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PhysicalInterface{}
	for _, item := range list {
		result = append(result, InterfaceToPhysicalInterface(item))
	}
	return result
}
