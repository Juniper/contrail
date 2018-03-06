package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDsaRule makes DsaRule
// nolint
func MakeDsaRule() *DsaRule {
	return &DsaRule{
		//TODO(nati): Apply default
		UUID:         "",
		ParentUUID:   "",
		ParentType:   "",
		FQName:       []string{},
		IDPerms:      MakeIdPermsType(),
		DisplayName:  "",
		Annotations:  MakeKeyValuePairs(),
		Perms2:       MakePermType2(),
		DsaRuleEntry: MakeDiscoveryServiceAssignmentType(),
	}
}

// MakeDsaRule makes DsaRule
// nolint
func InterfaceToDsaRule(i interface{}) *DsaRule {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DsaRule{
		//TODO(nati): Apply default
		UUID:         common.InterfaceToString(m["uuid"]),
		ParentUUID:   common.InterfaceToString(m["parent_uuid"]),
		ParentType:   common.InterfaceToString(m["parent_type"]),
		FQName:       common.InterfaceToStringList(m["fq_name"]),
		IDPerms:      InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:  common.InterfaceToString(m["display_name"]),
		Annotations:  InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:       InterfaceToPermType2(m["perms2"]),
		DsaRuleEntry: InterfaceToDiscoveryServiceAssignmentType(m["dsa_rule_entry"]),
	}
}

// MakeDsaRuleSlice() makes a slice of DsaRule
// nolint
func MakeDsaRuleSlice() []*DsaRule {
	return []*DsaRule{}
}

// InterfaceToDsaRuleSlice() makes a slice of DsaRule
// nolint
func InterfaceToDsaRuleSlice(i interface{}) []*DsaRule {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DsaRule{}
	for _, item := range list {
		result = append(result, InterfaceToDsaRule(item))
	}
	return result
}
