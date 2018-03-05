package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualDNSRecord makes VirtualDNSRecord
// nolint
func MakeVirtualDNSRecord() *VirtualDNSRecord {
	return &VirtualDNSRecord{
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
		VirtualDNSRecordData: MakeVirtualDnsRecordType(),
	}
}

// MakeVirtualDNSRecord makes VirtualDNSRecord
// nolint
func InterfaceToVirtualDNSRecord(i interface{}) *VirtualDNSRecord {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualDNSRecord{
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
		VirtualDNSRecordData: InterfaceToVirtualDnsRecordType(m["virtual_DNS_record_data"]),
	}
}

// MakeVirtualDNSRecordSlice() makes a slice of VirtualDNSRecord
// nolint
func MakeVirtualDNSRecordSlice() []*VirtualDNSRecord {
	return []*VirtualDNSRecord{}
}

// InterfaceToVirtualDNSRecordSlice() makes a slice of VirtualDNSRecord
// nolint
func InterfaceToVirtualDNSRecordSlice(i interface{}) []*VirtualDNSRecord {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualDNSRecord{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualDNSRecord(item))
	}
	return result
}
