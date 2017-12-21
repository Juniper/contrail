package models

// FirewallPolicy

import "encoding/json"

// FirewallPolicy
type FirewallPolicy struct {
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`

	FirewallRuleRefs          []*FirewallPolicyFirewallRuleRef          `json:"firewall_rule_refs"`
	SecurityLoggingObjectRefs []*FirewallPolicySecurityLoggingObjectRef `json:"security_logging_object_refs"`
}

// FirewallPolicySecurityLoggingObjectRef references each other
type FirewallPolicySecurityLoggingObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// FirewallPolicyFirewallRuleRef references each other
type FirewallPolicyFirewallRuleRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *FirewallSequence
}

// String returns json representation of the object
func (model *FirewallPolicy) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallPolicy makes FirewallPolicy
func MakeFirewallPolicy() *FirewallPolicy {
	return &FirewallPolicy{
		//TODO(nati): Apply default
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
	}
}

// InterfaceToFirewallPolicy makes FirewallPolicy from interface
func InterfaceToFirewallPolicy(iData interface{}) *FirewallPolicy {
	data := iData.(map[string]interface{})
	return &FirewallPolicy{
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToFirewallPolicySlice makes a slice of FirewallPolicy from interface
func InterfaceToFirewallPolicySlice(data interface{}) []*FirewallPolicy {
	list := data.([]interface{})
	result := MakeFirewallPolicySlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallPolicy(item))
	}
	return result
}

// MakeFirewallPolicySlice() makes a slice of FirewallPolicy
func MakeFirewallPolicySlice() []*FirewallPolicy {
	return []*FirewallPolicy{}
}
