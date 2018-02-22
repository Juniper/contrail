package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeAliasIP makes AliasIP
func MakeAliasIP() *AliasIP {
	return &AliasIP{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		AliasIPAddress:       "",
		AliasIPAddressFamily: "",
	}
}

// MakeAliasIP makes AliasIP
func InterfaceToAliasIP(i interface{}) *AliasIP {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AliasIP{
		//TODO(nati): Apply default
		UUID:                 schema.InterfaceToString(m["uuid"]),
		ParentUUID:           schema.InterfaceToString(m["parent_uuid"]),
		ParentType:           schema.InterfaceToString(m["parent_type"]),
		FQName:               schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          schema.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		AliasIPAddress:       schema.InterfaceToString(m["alias_ip_address"]),
		AliasIPAddressFamily: schema.InterfaceToString(m["alias_ip_address_family"]),
	}
}

// MakeAliasIPSlice() makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
	return []*AliasIP{}
}

// InterfaceToAliasIPSlice() makes a slice of AliasIP
func InterfaceToAliasIPSlice(i interface{}) []*AliasIP {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AliasIP{}
	for _, item := range list {
		result = append(result, InterfaceToAliasIP(item))
	}
	return result
}
