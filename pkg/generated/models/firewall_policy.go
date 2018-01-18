package models

// FirewallPolicy

import "encoding/json"

// FirewallPolicy
type FirewallPolicy struct {
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`

	FirewallRuleRefs          []*FirewallPolicyFirewallRuleRef          `json:"firewall_rule_refs,omitempty"`
	SecurityLoggingObjectRefs []*FirewallPolicySecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
}

// FirewallPolicyFirewallRuleRef references each other
type FirewallPolicyFirewallRuleRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *FirewallSequence
}

// FirewallPolicySecurityLoggingObjectRef references each other
type FirewallPolicySecurityLoggingObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

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
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
	}
}

// MakeFirewallPolicySlice() makes a slice of FirewallPolicy
func MakeFirewallPolicySlice() []*FirewallPolicy {
	return []*FirewallPolicy{}
}
