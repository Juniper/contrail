package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeSubnet makes Subnet
func MakeSubnet() *Subnet {
	return &Subnet{
		//TODO(nati): Apply default
		UUID:           "",
		ParentUUID:     "",
		ParentType:     "",
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		SubnetIPPrefix: MakeSubnetType(),
	}
}

// MakeSubnet makes Subnet
func InterfaceToSubnet(i interface{}) *Subnet {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Subnet{
		//TODO(nati): Apply default
		UUID:           schema.InterfaceToString(m["uuid"]),
		ParentUUID:     schema.InterfaceToString(m["parent_uuid"]),
		ParentType:     schema.InterfaceToString(m["parent_type"]),
		FQName:         schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:        InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:    schema.InterfaceToString(m["display_name"]),
		Annotations:    InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:         InterfaceToPermType2(m["perms2"]),
		SubnetIPPrefix: InterfaceToSubnetType(m["subnet_ip_prefix"]),
	}
}

// MakeSubnetSlice() makes a slice of Subnet
func MakeSubnetSlice() []*Subnet {
	return []*Subnet{}
}

// InterfaceToSubnetSlice() makes a slice of Subnet
func InterfaceToSubnetSlice(i interface{}) []*Subnet {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Subnet{}
	for _, item := range list {
		result = append(result, InterfaceToSubnet(item))
	}
	return result
}
