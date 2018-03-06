package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAliasIPPool makes AliasIPPool
// nolint
func MakeAliasIPPool() *AliasIPPool {
	return &AliasIPPool{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeAliasIPPool makes AliasIPPool
// nolint
func InterfaceToAliasIPPool(i interface{}) *AliasIPPool {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AliasIPPool{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
	}
}

// MakeAliasIPPoolSlice() makes a slice of AliasIPPool
// nolint
func MakeAliasIPPoolSlice() []*AliasIPPool {
	return []*AliasIPPool{}
}

// InterfaceToAliasIPPoolSlice() makes a slice of AliasIPPool
// nolint
func InterfaceToAliasIPPoolSlice(i interface{}) []*AliasIPPool {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AliasIPPool{}
	for _, item := range list {
		result = append(result, InterfaceToAliasIPPool(item))
	}
	return result
}
