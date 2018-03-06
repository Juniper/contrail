package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualDnsType makes VirtualDnsType
// nolint
func MakeVirtualDnsType() *VirtualDnsType {
	return &VirtualDnsType{
		//TODO(nati): Apply default
		FloatingIPRecord:         "",
		DomainName:               "",
		ExternalVisible:          false,
		NextVirtualDNS:           "",
		DynamicRecordsFromClient: false,
		ReverseResolution:        false,
		DefaultTTLSeconds:        0,
		RecordOrder:              "",
	}
}

// MakeVirtualDnsType makes VirtualDnsType
// nolint
func InterfaceToVirtualDnsType(i interface{}) *VirtualDnsType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualDnsType{
		//TODO(nati): Apply default
		FloatingIPRecord:         common.InterfaceToString(m["floating_ip_record"]),
		DomainName:               common.InterfaceToString(m["domain_name"]),
		ExternalVisible:          common.InterfaceToBool(m["external_visible"]),
		NextVirtualDNS:           common.InterfaceToString(m["next_virtual_DNS"]),
		DynamicRecordsFromClient: common.InterfaceToBool(m["dynamic_records_from_client"]),
		ReverseResolution:        common.InterfaceToBool(m["reverse_resolution"]),
		DefaultTTLSeconds:        common.InterfaceToInt64(m["default_ttl_seconds"]),
		RecordOrder:              common.InterfaceToString(m["record_order"]),
	}
}

// MakeVirtualDnsTypeSlice() makes a slice of VirtualDnsType
// nolint
func MakeVirtualDnsTypeSlice() []*VirtualDnsType {
	return []*VirtualDnsType{}
}

// InterfaceToVirtualDnsTypeSlice() makes a slice of VirtualDnsType
// nolint
func InterfaceToVirtualDnsTypeSlice(i interface{}) []*VirtualDnsType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualDnsType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualDnsType(item))
	}
	return result
}
