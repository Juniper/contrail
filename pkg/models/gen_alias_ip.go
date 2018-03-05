package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAliasIP makes AliasIP
// nolint
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
		ConfigurationVersion: 0,
		AliasIPAddress:       "",
		AliasIPAddressFamily: "",
	}
}

// MakeAliasIP makes AliasIP
// nolint
func InterfaceToAliasIP(i interface{}) *AliasIP {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AliasIP{
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
		AliasIPAddress:       common.InterfaceToString(m["alias_ip_address"]),
		AliasIPAddressFamily: common.InterfaceToString(m["alias_ip_address_family"]),
	}
}

// MakeAliasIPSlice() makes a slice of AliasIP
// nolint
func MakeAliasIPSlice() []*AliasIP {
	return []*AliasIP{}
}

// InterfaceToAliasIPSlice() makes a slice of AliasIP
// nolint
func InterfaceToAliasIPSlice(i interface{}) []*AliasIP {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AliasIP{}
	for _, item := range list {
		result = append(result, InterfaceToAliasIP(item))
	}
	return result
}
