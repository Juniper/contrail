package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualDnsRecordType makes VirtualDnsRecordType
// nolint
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
// nolint
func InterfaceToVirtualDnsRecordType(i interface{}) *VirtualDnsRecordType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualDnsRecordType{
		//TODO(nati): Apply default
		RecordName:         common.InterfaceToString(m["record_name"]),
		RecordClass:        common.InterfaceToString(m["record_class"]),
		RecordData:         common.InterfaceToString(m["record_data"]),
		RecordType:         common.InterfaceToString(m["record_type"]),
		RecordTTLSeconds:   common.InterfaceToInt64(m["record_ttl_seconds"]),
		RecordMXPreference: common.InterfaceToInt64(m["record_mx_preference"]),
	}
}

// MakeVirtualDnsRecordTypeSlice() makes a slice of VirtualDnsRecordType
// nolint
func MakeVirtualDnsRecordTypeSlice() []*VirtualDnsRecordType {
	return []*VirtualDnsRecordType{}
}

// InterfaceToVirtualDnsRecordTypeSlice() makes a slice of VirtualDnsRecordType
// nolint
func InterfaceToVirtualDnsRecordTypeSlice(i interface{}) []*VirtualDnsRecordType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualDnsRecordType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualDnsRecordType(item))
	}
	return result
}
