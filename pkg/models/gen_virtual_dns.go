package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualDNS makes VirtualDNS
// nolint
func MakeVirtualDNS() *VirtualDNS {
	return &VirtualDNS{
		//TODO(nati): Apply default
		UUID:           "",
		ParentUUID:     "",
		ParentType:     "",
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		VirtualDNSData: MakeVirtualDnsType(),
	}
}

// MakeVirtualDNS makes VirtualDNS
// nolint
func InterfaceToVirtualDNS(i interface{}) *VirtualDNS {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualDNS{
		//TODO(nati): Apply default
		UUID:           common.InterfaceToString(m["uuid"]),
		ParentUUID:     common.InterfaceToString(m["parent_uuid"]),
		ParentType:     common.InterfaceToString(m["parent_type"]),
		FQName:         common.InterfaceToStringList(m["fq_name"]),
		IDPerms:        InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:    common.InterfaceToString(m["display_name"]),
		Annotations:    InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:         InterfaceToPermType2(m["perms2"]),
		VirtualDNSData: InterfaceToVirtualDnsType(m["virtual_DNS_data"]),
	}
}

// MakeVirtualDNSSlice() makes a slice of VirtualDNS
// nolint
func MakeVirtualDNSSlice() []*VirtualDNS {
	return []*VirtualDNS{}
}

// InterfaceToVirtualDNSSlice() makes a slice of VirtualDNS
// nolint
func InterfaceToVirtualDNSSlice(i interface{}) []*VirtualDNS {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualDNS{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualDNS(item))
	}
	return result
}
