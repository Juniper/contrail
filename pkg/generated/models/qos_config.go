package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeQosConfig makes QosConfig
func MakeQosConfig() *QosConfig {
	return &QosConfig{
		//TODO(nati): Apply default
		UUID:                     "",
		ParentUUID:               "",
		ParentType:               "",
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		QosConfigType:            "",
		MPLSExpEntries:           MakeQosIdForwardingClassPairs(),
		VlanPriorityEntries:      MakeQosIdForwardingClassPairs(),
		DefaultForwardingClassID: 0,
		DSCPEntries:              MakeQosIdForwardingClassPairs(),
	}
}

// MakeQosConfig makes QosConfig
func InterfaceToQosConfig(i interface{}) *QosConfig {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &QosConfig{
		//TODO(nati): Apply default
		UUID:                     schema.InterfaceToString(m["uuid"]),
		ParentUUID:               schema.InterfaceToString(m["parent_uuid"]),
		ParentType:               schema.InterfaceToString(m["parent_type"]),
		FQName:                   schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                  InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:              schema.InterfaceToString(m["display_name"]),
		Annotations:              InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                   InterfaceToPermType2(m["perms2"]),
		QosConfigType:            schema.InterfaceToString(m["qos_config_type"]),
		MPLSExpEntries:           InterfaceToQosIdForwardingClassPairs(m["mpls_exp_entries"]),
		VlanPriorityEntries:      InterfaceToQosIdForwardingClassPairs(m["vlan_priority_entries"]),
		DefaultForwardingClassID: schema.InterfaceToInt64(m["default_forwarding_class_id"]),
		DSCPEntries:              InterfaceToQosIdForwardingClassPairs(m["dscp_entries"]),
	}
}

// MakeQosConfigSlice() makes a slice of QosConfig
func MakeQosConfigSlice() []*QosConfig {
	return []*QosConfig{}
}

// InterfaceToQosConfigSlice() makes a slice of QosConfig
func InterfaceToQosConfigSlice(i interface{}) []*QosConfig {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*QosConfig{}
	for _, item := range list {
		result = append(result, InterfaceToQosConfig(item))
	}
	return result
}
