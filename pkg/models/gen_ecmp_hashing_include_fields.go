package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeEcmpHashingIncludeFields makes EcmpHashingIncludeFields
// nolint
func MakeEcmpHashingIncludeFields() *EcmpHashingIncludeFields {
	return &EcmpHashingIncludeFields{
		//TODO(nati): Apply default
		DestinationIP:     false,
		IPProtocol:        false,
		SourceIP:          false,
		HashingConfigured: false,
		SourcePort:        false,
		DestinationPort:   false,
	}
}

// MakeEcmpHashingIncludeFields makes EcmpHashingIncludeFields
// nolint
func InterfaceToEcmpHashingIncludeFields(i interface{}) *EcmpHashingIncludeFields {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &EcmpHashingIncludeFields{
		//TODO(nati): Apply default
		DestinationIP:     common.InterfaceToBool(m["destination_ip"]),
		IPProtocol:        common.InterfaceToBool(m["ip_protocol"]),
		SourceIP:          common.InterfaceToBool(m["source_ip"]),
		HashingConfigured: common.InterfaceToBool(m["hashing_configured"]),
		SourcePort:        common.InterfaceToBool(m["source_port"]),
		DestinationPort:   common.InterfaceToBool(m["destination_port"]),
	}
}

// MakeEcmpHashingIncludeFieldsSlice() makes a slice of EcmpHashingIncludeFields
// nolint
func MakeEcmpHashingIncludeFieldsSlice() []*EcmpHashingIncludeFields {
	return []*EcmpHashingIncludeFields{}
}

// InterfaceToEcmpHashingIncludeFieldsSlice() makes a slice of EcmpHashingIncludeFields
// nolint
func InterfaceToEcmpHashingIncludeFieldsSlice(i interface{}) []*EcmpHashingIncludeFields {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*EcmpHashingIncludeFields{}
	for _, item := range list {
		result = append(result, InterfaceToEcmpHashingIncludeFields(item))
	}
	return result
}
