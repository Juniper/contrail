package models

// DsaRule

import "encoding/json"

// DsaRule
type DsaRule struct {
	ParentUUID   string                          `json:"parent_uuid"`
	ParentType   string                          `json:"parent_type"`
	IDPerms      *IdPermsType                    `json:"id_perms"`
	Annotations  *KeyValuePairs                  `json:"annotations"`
	DsaRuleEntry *DiscoveryServiceAssignmentType `json:"dsa_rule_entry"`
	FQName       []string                        `json:"fq_name"`
	DisplayName  string                          `json:"display_name"`
	Perms2       *PermType2                      `json:"perms2"`
	UUID         string                          `json:"uuid"`
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
		ParentUUID:   "",
		ParentType:   "",
		IDPerms:      MakeIdPermsType(),
		Annotations:  MakeKeyValuePairs(),
		DsaRuleEntry: MakeDiscoveryServiceAssignmentType(),
		FQName:       []string{},
		DisplayName:  "",
		Perms2:       MakePermType2(),
		UUID:         "",
	}
}

// InterfaceToDsaRule makes DsaRule from interface
func InterfaceToDsaRule(iData interface{}) *DsaRule {
	data := iData.(map[string]interface{})
	return &DsaRule{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		DsaRuleEntry: InterfaceToDiscoveryServiceAssignmentType(data["dsa_rule_entry"]),

		//{"description":"rule entry defining publisher set and subscriber set.","type":"object","properties":{"publisher":{"type":"object","properties":{"ep_id":{"type":"string"},"ep_prefix":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"ep_type":{"type":"string"},"ep_version":{"type":"string"}}},"subscriber":{"type":"array","item":{"type":"object","properties":{"ep_id":{"type":"string"},"ep_prefix":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"ep_type":{"type":"string"},"ep_version":{"type":"string"}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}

	}
}

// InterfaceToDsaRuleSlice makes a slice of DsaRule from interface
func InterfaceToDsaRuleSlice(data interface{}) []*DsaRule {
	list := data.([]interface{})
	result := MakeDsaRuleSlice()
	for _, item := range list {
		result = append(result, InterfaceToDsaRule(item))
	}
	return result
}

// MakeDsaRuleSlice() makes a slice of DsaRule
func MakeDsaRuleSlice() []*DsaRule {
	return []*DsaRule{}
}
