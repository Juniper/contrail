package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeDsaRule makes DsaRule
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
func InterfaceToDsaRule(i interface{}) *DsaRule {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DsaRule{
		//TODO(nati): Apply default
		UUID:         schema.InterfaceToString(m["uuid"]),
		ParentUUID:   schema.InterfaceToString(m["parent_uuid"]),
		ParentType:   schema.InterfaceToString(m["parent_type"]),
		FQName:       schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:      InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:  schema.InterfaceToString(m["display_name"]),
		Annotations:  InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:       InterfaceToPermType2(m["perms2"]),
		DsaRuleEntry: InterfaceToDiscoveryServiceAssignmentType(m["dsa_rule_entry"]),
	}
}

// MakeDsaRuleSlice() makes a slice of DsaRule
func MakeDsaRuleSlice() []*DsaRule {
	return []*DsaRule{}
}

// InterfaceToDsaRuleSlice() makes a slice of DsaRule
func InterfaceToDsaRuleSlice(i interface{}) []*DsaRule {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DsaRule{}
	for _, item := range list {
		result = append(result, InterfaceToDsaRule(item))
	}
	return result
}
