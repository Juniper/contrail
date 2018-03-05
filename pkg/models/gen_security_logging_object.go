package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeSecurityLoggingObject makes SecurityLoggingObject
// nolint
func MakeSecurityLoggingObject() *SecurityLoggingObject {
	return &SecurityLoggingObject{
		//TODO(nati): Apply default
		UUID:                       "",
		ParentUUID:                 "",
		ParentType:                 "",
		FQName:                     []string{},
		IDPerms:                    MakeIdPermsType(),
		DisplayName:                "",
		Annotations:                MakeKeyValuePairs(),
		Perms2:                     MakePermType2(),
		ConfigurationVersion:       0,
		SecurityLoggingObjectRules: MakeSecurityLoggingObjectRuleListType(),
		SecurityLoggingObjectRate:  0,
	}
}

// MakeSecurityLoggingObject makes SecurityLoggingObject
// nolint
func InterfaceToSecurityLoggingObject(i interface{}) *SecurityLoggingObject {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SecurityLoggingObject{
		//TODO(nati): Apply default
		UUID:                       common.InterfaceToString(m["uuid"]),
		ParentUUID:                 common.InterfaceToString(m["parent_uuid"]),
		ParentType:                 common.InterfaceToString(m["parent_type"]),
		FQName:                     common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                    InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                common.InterfaceToString(m["display_name"]),
		Annotations:                InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                     InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:       common.InterfaceToInt64(m["configuration_version"]),
		SecurityLoggingObjectRules: InterfaceToSecurityLoggingObjectRuleListType(m["security_logging_object_rules"]),
		SecurityLoggingObjectRate:  common.InterfaceToInt64(m["security_logging_object_rate"]),
	}
}

// MakeSecurityLoggingObjectSlice() makes a slice of SecurityLoggingObject
// nolint
func MakeSecurityLoggingObjectSlice() []*SecurityLoggingObject {
	return []*SecurityLoggingObject{}
}

// InterfaceToSecurityLoggingObjectSlice() makes a slice of SecurityLoggingObject
// nolint
func InterfaceToSecurityLoggingObjectSlice(i interface{}) []*SecurityLoggingObject {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SecurityLoggingObject{}
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObject(item))
	}
	return result
}
