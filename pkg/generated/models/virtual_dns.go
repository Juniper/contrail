package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeVirtualDNS makes VirtualDNS
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
func InterfaceToVirtualDNS(i interface{}) *VirtualDNS {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualDNS{
		//TODO(nati): Apply default
		UUID:           schema.InterfaceToString(m["uuid"]),
		ParentUUID:     schema.InterfaceToString(m["parent_uuid"]),
		ParentType:     schema.InterfaceToString(m["parent_type"]),
		FQName:         schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:        InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:    schema.InterfaceToString(m["display_name"]),
		Annotations:    InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:         InterfaceToPermType2(m["perms2"]),
		VirtualDNSData: InterfaceToVirtualDnsType(m["virtual_DNS_data"]),
	}
}

// MakeVirtualDNSSlice() makes a slice of VirtualDNS
func MakeVirtualDNSSlice() []*VirtualDNS {
	return []*VirtualDNS{}
}

// InterfaceToVirtualDNSSlice() makes a slice of VirtualDNS
func InterfaceToVirtualDNSSlice(i interface{}) []*VirtualDNS {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualDNS{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualDNS(item))
	}
	return result
}
