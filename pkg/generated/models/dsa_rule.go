package models

// DsaRule

import "encoding/json"

// DsaRule
type DsaRule struct {
	ParentType   string                          `json:"parent_type,omitempty"`
	IDPerms      *IdPermsType                    `json:"id_perms,omitempty"`
	DisplayName  string                          `json:"display_name,omitempty"`
	ParentUUID   string                          `json:"parent_uuid,omitempty"`
	Perms2       *PermType2                      `json:"perms2,omitempty"`
	UUID         string                          `json:"uuid,omitempty"`
	DsaRuleEntry *DiscoveryServiceAssignmentType `json:"dsa_rule_entry,omitempty"`
	FQName       []string                        `json:"fq_name,omitempty"`
	Annotations  *KeyValuePairs                  `json:"annotations,omitempty"`
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
		Perms2:       MakePermType2(),
		UUID:         "",
		DsaRuleEntry: MakeDiscoveryServiceAssignmentType(),
		FQName:       []string{},
		ParentUUID:   "",
		ParentType:   "",
		IDPerms:      MakeIdPermsType(),
		DisplayName:  "",
	}
}

// MakeDsaRuleSlice() makes a slice of DsaRule
func MakeDsaRuleSlice() []*DsaRule {
	return []*DsaRule{}
}
