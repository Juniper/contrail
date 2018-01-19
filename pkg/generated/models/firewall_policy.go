package models

// FirewallPolicy

import "encoding/json"

// FirewallPolicy
type FirewallPolicy struct {
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`

	FirewallRuleRefs          []*FirewallPolicyFirewallRuleRef          `json:"firewall_rule_refs,omitempty"`
	SecurityLoggingObjectRefs []*FirewallPolicySecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
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
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeFirewallPolicySlice() makes a slice of FirewallPolicy
func MakeFirewallPolicySlice() []*FirewallPolicy {
	return []*FirewallPolicy{}
}
