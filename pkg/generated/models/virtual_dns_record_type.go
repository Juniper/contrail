package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeVirtualDnsRecordType makes VirtualDnsRecordType
func MakeVirtualDnsRecordType() *VirtualDnsRecordType {
	return &VirtualDnsRecordType{
		//TODO(nati): Apply default
		RecordName:         "",
		RecordClass:        "",
		RecordData:         "",
		RecordType:         "",
		RecordTTLSeconds:   0,
		RecordMXPreference: 0,
	}
}

// MakeVirtualDnsRecordType makes VirtualDnsRecordType
func InterfaceToVirtualDnsRecordType(i interface{}) *VirtualDnsRecordType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualDnsRecordType{
		//TODO(nati): Apply default
		RecordName:         schema.InterfaceToString(m["record_name"]),
		RecordClass:        schema.InterfaceToString(m["record_class"]),
		RecordData:         schema.InterfaceToString(m["record_data"]),
		RecordType:         schema.InterfaceToString(m["record_type"]),
		RecordTTLSeconds:   schema.InterfaceToInt64(m["record_ttl_seconds"]),
		RecordMXPreference: schema.InterfaceToInt64(m["record_mx_preference"]),
	}
}

// MakeVirtualDnsRecordTypeSlice() makes a slice of VirtualDnsRecordType
func MakeVirtualDnsRecordTypeSlice() []*VirtualDnsRecordType {
	return []*VirtualDnsRecordType{}
}

// InterfaceToVirtualDnsRecordTypeSlice() makes a slice of VirtualDnsRecordType
func InterfaceToVirtualDnsRecordTypeSlice(i interface{}) []*VirtualDnsRecordType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualDnsRecordType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualDnsRecordType(item))
	}
	return result
}
