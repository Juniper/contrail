package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAddressGroup makes AddressGroup
// nolint
func MakeAddressGroup() *AddressGroup {
	return &AddressGroup{
		//TODO(nati): Apply default
		UUID:               "",
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		AddressGroupPrefix: MakeSubnetListType(),
	}
}

// MakeAddressGroup makes AddressGroup
// nolint
func InterfaceToAddressGroup(i interface{}) *AddressGroup {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AddressGroup{
		//TODO(nati): Apply default
		UUID:               common.InterfaceToString(m["uuid"]),
		ParentUUID:         common.InterfaceToString(m["parent_uuid"]),
		ParentType:         common.InterfaceToString(m["parent_type"]),
		FQName:             common.InterfaceToStringList(m["fq_name"]),
		IDPerms:            InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:        common.InterfaceToString(m["display_name"]),
		Annotations:        InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:             InterfaceToPermType2(m["perms2"]),
		AddressGroupPrefix: InterfaceToSubnetListType(m["address_group_prefix"]),
	}
}

// MakeAddressGroupSlice() makes a slice of AddressGroup
// nolint
func MakeAddressGroupSlice() []*AddressGroup {
	return []*AddressGroup{}
}

// InterfaceToAddressGroupSlice() makes a slice of AddressGroup
// nolint
func InterfaceToAddressGroupSlice(i interface{}) []*AddressGroup {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AddressGroup{}
	for _, item := range list {
		result = append(result, InterfaceToAddressGroup(item))
	}
	return result
}
