package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDomain makes Domain
// nolint
func MakeDomain() *Domain {
	return &Domain{
		//TODO(nati): Apply default
		UUID:         "",
		ParentUUID:   "",
		ParentType:   "",
		FQName:       []string{},
		IDPerms:      MakeIdPermsType(),
		DisplayName:  "",
		Annotations:  MakeKeyValuePairs(),
		Perms2:       MakePermType2(),
		DomainLimits: MakeDomainLimitsType(),
	}
}

// MakeDomain makes Domain
// nolint
func InterfaceToDomain(i interface{}) *Domain {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Domain{
		//TODO(nati): Apply default
		UUID:         common.InterfaceToString(m["uuid"]),
		ParentUUID:   common.InterfaceToString(m["parent_uuid"]),
		ParentType:   common.InterfaceToString(m["parent_type"]),
		FQName:       common.InterfaceToStringList(m["fq_name"]),
		IDPerms:      InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:  common.InterfaceToString(m["display_name"]),
		Annotations:  InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:       InterfaceToPermType2(m["perms2"]),
		DomainLimits: InterfaceToDomainLimitsType(m["domain_limits"]),
	}
}

// MakeDomainSlice() makes a slice of Domain
// nolint
func MakeDomainSlice() []*Domain {
	return []*Domain{}
}

// InterfaceToDomainSlice() makes a slice of Domain
// nolint
func InterfaceToDomainSlice(i interface{}) []*Domain {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Domain{}
	for _, item := range list {
		result = append(result, InterfaceToDomain(item))
	}
	return result
}
