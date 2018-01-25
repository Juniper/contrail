package models

// DsaRule

// DsaRule
//proteus:generate
type DsaRule struct {
	UUID         string                          `json:"uuid,omitempty"`
	ParentUUID   string                          `json:"parent_uuid,omitempty"`
	ParentType   string                          `json:"parent_type,omitempty"`
	FQName       []string                        `json:"fq_name,omitempty"`
	IDPerms      *IdPermsType                    `json:"id_perms,omitempty"`
	DisplayName  string                          `json:"display_name,omitempty"`
	Annotations  *KeyValuePairs                  `json:"annotations,omitempty"`
	Perms2       *PermType2                      `json:"perms2,omitempty"`
	DsaRuleEntry *DiscoveryServiceAssignmentType `json:"dsa_rule_entry,omitempty"`
}

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

// MakeDsaRuleSlice() makes a slice of DsaRule
func MakeDsaRuleSlice() []*DsaRule {
	return []*DsaRule{}
}
