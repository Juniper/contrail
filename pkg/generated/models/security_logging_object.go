package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeSecurityLoggingObject makes SecurityLoggingObject
func MakeSecurityLoggingObject() *SecurityLoggingObject {
	return &SecurityLoggingObject{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		SecurityLoggingObjectRules: MakeSecurityLoggingObjectRuleListType(),
		SecurityLoggingObjectRate:  0,
	}
}

// MakeSecurityLoggingObject makes SecurityLoggingObject
func InterfaceToSecurityLoggingObject(i interface{}) *SecurityLoggingObject {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SecurityLoggingObject{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		SecurityLoggingObjectRules: InterfaceToSecurityLoggingObjectRuleListType(m["security_logging_object_rules"]),
		SecurityLoggingObjectRate:  schema.InterfaceToInt64(m["security_logging_object_rate"]),
	}
}

// MakeSecurityLoggingObjectSlice() makes a slice of SecurityLoggingObject
func MakeSecurityLoggingObjectSlice() []*SecurityLoggingObject {
	return []*SecurityLoggingObject{}
}

// InterfaceToSecurityLoggingObjectSlice() makes a slice of SecurityLoggingObject
func InterfaceToSecurityLoggingObjectSlice(i interface{}) []*SecurityLoggingObject {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SecurityLoggingObject{}
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObject(item))
	}
	return result
}
