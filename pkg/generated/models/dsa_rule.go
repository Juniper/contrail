package models

// DsaRule

import "encoding/json"

// DsaRule
type DsaRule struct {
	IDPerms      *IdPermsType                    `json:"id_perms,omitempty"`
	DisplayName  string                          `json:"display_name,omitempty"`
	Annotations  *KeyValuePairs                  `json:"annotations,omitempty"`
	UUID         string                          `json:"uuid,omitempty"`
	FQName       []string                        `json:"fq_name,omitempty"`
	ParentType   string                          `json:"parent_type,omitempty"`
	Perms2       *PermType2                      `json:"perms2,omitempty"`
	ParentUUID   string                          `json:"parent_uuid,omitempty"`
	DsaRuleEntry *DiscoveryServiceAssignmentType `json:"dsa_rule_entry,omitempty"`
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
		FQName:       []string{},
		IDPerms:      MakeIdPermsType(),
		DisplayName:  "",
		Annotations:  MakeKeyValuePairs(),
		UUID:         "",
		DsaRuleEntry: MakeDiscoveryServiceAssignmentType(),
		ParentType:   "",
		Perms2:       MakePermType2(),
		ParentUUID:   "",
	}
}

// MakeDsaRuleSlice() makes a slice of DsaRule
func MakeDsaRuleSlice() []*DsaRule {
	return []*DsaRule{}
}
