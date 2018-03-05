package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeSubnet makes Subnet
// nolint
func MakeSubnet() *Subnet {
	return &Subnet{
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
		SubnetIPPrefix:       MakeSubnetType(),
	}
}

// MakeSubnet makes Subnet
// nolint
func InterfaceToSubnet(i interface{}) *Subnet {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Subnet{
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
		SubnetIPPrefix:       InterfaceToSubnetType(m["subnet_ip_prefix"]),
	}
}

// MakeSubnetSlice() makes a slice of Subnet
// nolint
func MakeSubnetSlice() []*Subnet {
	return []*Subnet{}
}

// InterfaceToSubnetSlice() makes a slice of Subnet
// nolint
func InterfaceToSubnetSlice(i interface{}) []*Subnet {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Subnet{}
	for _, item := range list {
		result = append(result, InterfaceToSubnet(item))
	}
	return result
}
