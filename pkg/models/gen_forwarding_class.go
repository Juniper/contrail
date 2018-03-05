package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeForwardingClass makes ForwardingClass
// nolint
func MakeForwardingClass() *ForwardingClass {
	return &ForwardingClass{
		//TODO(nati): Apply default
		UUID:                        "",
		ParentUUID:                  "",
		ParentType:                  "",
		FQName:                      []string{},
		IDPerms:                     MakeIdPermsType(),
		DisplayName:                 "",
		Annotations:                 MakeKeyValuePairs(),
		Perms2:                      MakePermType2(),
		ConfigurationVersion:        0,
		ForwardingClassDSCP:         0,
		ForwardingClassVlanPriority: 0,
		ForwardingClassMPLSExp:      0,
		ForwardingClassID:           0,
	}
}

// MakeForwardingClass makes ForwardingClass
// nolint
func InterfaceToForwardingClass(i interface{}) *ForwardingClass {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ForwardingClass{
		//TODO(nati): Apply default
		UUID:                        common.InterfaceToString(m["uuid"]),
		ParentUUID:                  common.InterfaceToString(m["parent_uuid"]),
		ParentType:                  common.InterfaceToString(m["parent_type"]),
		FQName:                      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                 common.InterfaceToString(m["display_name"]),
		Annotations:                 InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                      InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:        common.InterfaceToInt64(m["configuration_version"]),
		ForwardingClassDSCP:         common.InterfaceToInt64(m["forwarding_class_dscp"]),
		ForwardingClassVlanPriority: common.InterfaceToInt64(m["forwarding_class_vlan_priority"]),
		ForwardingClassMPLSExp:      common.InterfaceToInt64(m["forwarding_class_mpls_exp"]),
		ForwardingClassID:           common.InterfaceToInt64(m["forwarding_class_id"]),
	}
}

// MakeForwardingClassSlice() makes a slice of ForwardingClass
// nolint
func MakeForwardingClassSlice() []*ForwardingClass {
	return []*ForwardingClass{}
}

// InterfaceToForwardingClassSlice() makes a slice of ForwardingClass
// nolint
func InterfaceToForwardingClassSlice(i interface{}) []*ForwardingClass {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ForwardingClass{}
	for _, item := range list {
		result = append(result, InterfaceToForwardingClass(item))
	}
	return result
}
