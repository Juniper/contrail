package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeVirtualDNSRecord makes VirtualDNSRecord
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
		VirtualDNSRecordData: MakeVirtualDnsRecordType(),
	}
}

// MakeVirtualDNSRecord makes VirtualDNSRecord
func InterfaceToVirtualDNSRecord(i interface{}) *VirtualDNSRecord {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualDNSRecord{
		//TODO(nati): Apply default
		UUID:                 schema.InterfaceToString(m["uuid"]),
		ParentUUID:           schema.InterfaceToString(m["parent_uuid"]),
		ParentType:           schema.InterfaceToString(m["parent_type"]),
		FQName:               schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          schema.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		VirtualDNSRecordData: InterfaceToVirtualDnsRecordType(m["virtual_DNS_record_data"]),
	}
}

// MakeVirtualDNSRecordSlice() makes a slice of VirtualDNSRecord
func MakeVirtualDNSRecordSlice() []*VirtualDNSRecord {
	return []*VirtualDNSRecord{}
}

// InterfaceToVirtualDNSRecordSlice() makes a slice of VirtualDNSRecord
func InterfaceToVirtualDNSRecordSlice(i interface{}) []*VirtualDNSRecord {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualDNSRecord{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualDNSRecord(item))
	}
	return result
}
