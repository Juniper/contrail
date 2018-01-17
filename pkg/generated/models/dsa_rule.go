package models

// DsaRule

import "encoding/json"

// DsaRule
type DsaRule struct {
	Annotations  *KeyValuePairs                  `json:"annotations,omitempty"`
	Perms2       *PermType2                      `json:"perms2,omitempty"`
	UUID         string                          `json:"uuid,omitempty"`
	ParentUUID   string                          `json:"parent_uuid,omitempty"`
	IDPerms      *IdPermsType                    `json:"id_perms,omitempty"`
	DsaRuleEntry *DiscoveryServiceAssignmentType `json:"dsa_rule_entry,omitempty"`
	ParentType   string                          `json:"parent_type,omitempty"`
	FQName       []string                        `json:"fq_name,omitempty"`
	DisplayName  string                          `json:"display_name,omitempty"`
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
		DsaRuleEntry: MakeDiscoveryServiceAssignmentType(),
		ParentType:   "",
		FQName:       []string{},
		DisplayName:  "",
		Annotations:  MakeKeyValuePairs(),
		Perms2:       MakePermType2(),
		UUID:         "",
		ParentUUID:   "",
		IDPerms:      MakeIdPermsType(),
	}
}

// MakeDsaRuleSlice() makes a slice of DsaRule
func MakeDsaRuleSlice() []*DsaRule {
	return []*DsaRule{}
}
