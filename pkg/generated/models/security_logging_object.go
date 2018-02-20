package models

// SecurityLoggingObject

// SecurityLoggingObject
//proteus:generate
type SecurityLoggingObject struct {
	UUID                       string                             `json:"uuid,omitempty"`
	ParentUUID                 string                             `json:"parent_uuid,omitempty"`
	ParentType                 string                             `json:"parent_type,omitempty"`
	FQName                     []string                           `json:"fq_name,omitempty"`
	IDPerms                    *IdPermsType                       `json:"id_perms,omitempty"`
	DisplayName                string                             `json:"display_name,omitempty"`
	Annotations                *KeyValuePairs                     `json:"annotations,omitempty"`
	Perms2                     *PermType2                         `json:"perms2,omitempty"`
	SecurityLoggingObjectRules *SecurityLoggingObjectRuleListType `json:"security_logging_object_rules,omitempty"`
	SecurityLoggingObjectRate  int                                `json:"security_logging_object_rate,omitempty"`

	NetworkPolicyRefs []*SecurityLoggingObjectNetworkPolicyRef `json:"network_policy_refs,omitempty"`
	SecurityGroupRefs []*SecurityLoggingObjectSecurityGroupRef `json:"security_group_refs,omitempty"`
}

// SecurityLoggingObjectSecurityGroupRef references each other
type SecurityLoggingObjectSecurityGroupRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *SecurityLoggingObjectRuleListType
}

// SecurityLoggingObjectNetworkPolicyRef references each other
type SecurityLoggingObjectNetworkPolicyRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *SecurityLoggingObjectRuleListType
}

// MakeSecurityLoggingObject makes SecurityLoggingObject
func MakeSecurityLoggingObject() *SecurityLoggingObject {
	return &SecurityLoggingObject{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		SecurityLoggingObjectRules: MakeSecurityLoggingObjectRuleListType(),
		SecurityLoggingObjectRate:  0,
	}
}

// MakeSecurityLoggingObjectSlice() makes a slice of SecurityLoggingObject
func MakeSecurityLoggingObjectSlice() []*SecurityLoggingObject {
	return []*SecurityLoggingObject{}
}
