package models

// DsaRule

import "encoding/json"

// DsaRule
type DsaRule struct {
	DisplayName  string                          `json:"display_name"`
	Annotations  *KeyValuePairs                  `json:"annotations"`
	ParentType   string                          `json:"parent_type"`
	FQName       []string                        `json:"fq_name"`
	IDPerms      *IdPermsType                    `json:"id_perms"`
	DsaRuleEntry *DiscoveryServiceAssignmentType `json:"dsa_rule_entry"`
	Perms2       *PermType2                      `json:"perms2"`
	UUID         string                          `json:"uuid"`
	ParentUUID   string                          `json:"parent_uuid"`
}

// String returns json representation of the object
func (model *DsaRule) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDsaRule makes DsaRule
func MakeDsaRule() *DsaRule {
	return &DsaRule{
		//TODO(nati): Apply default
		Annotations:  MakeKeyValuePairs(),
		DisplayName:  "",
		Perms2:       MakePermType2(),
		UUID:         "",
		ParentUUID:   "",
		ParentType:   "",
		FQName:       []string{},
		IDPerms:      MakeIdPermsType(),
		DsaRuleEntry: MakeDiscoveryServiceAssignmentType(),
	}
}

// MakeDsaRuleSlice() makes a slice of DsaRule
func MakeDsaRuleSlice() []*DsaRule {
	return []*DsaRule{}
}
