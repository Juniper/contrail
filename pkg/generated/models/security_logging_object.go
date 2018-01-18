package models

// SecurityLoggingObject

import "encoding/json"

// SecurityLoggingObject
type SecurityLoggingObject struct {
	Perms2                     *PermType2                         `json:"perms2,omitempty"`
	UUID                       string                             `json:"uuid,omitempty"`
	ParentUUID                 string                             `json:"parent_uuid,omitempty"`
	IDPerms                    *IdPermsType                       `json:"id_perms,omitempty"`
	Annotations                *KeyValuePairs                     `json:"annotations,omitempty"`
	SecurityLoggingObjectRules *SecurityLoggingObjectRuleListType `json:"security_logging_object_rules,omitempty"`
	SecurityLoggingObjectRate  int                                `json:"security_logging_object_rate,omitempty"`
	ParentType                 string                             `json:"parent_type,omitempty"`
	FQName                     []string                           `json:"fq_name,omitempty"`
	DisplayName                string                             `json:"display_name,omitempty"`

	SecurityGroupRefs []*SecurityLoggingObjectSecurityGroupRef `json:"security_group_refs,omitempty"`
	NetworkPolicyRefs []*SecurityLoggingObjectNetworkPolicyRef `json:"network_policy_refs,omitempty"`
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

// String returns json representation of the object
func (model *SecurityLoggingObject) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSecurityLoggingObject makes SecurityLoggingObject
func MakeSecurityLoggingObject() *SecurityLoggingObject {
	return &SecurityLoggingObject{
		//TODO(nati): Apply default
		Perms2:                     MakePermType2(),
		UUID:                       "",
		ParentUUID:                 "",
		IDPerms:                    MakeIdPermsType(),
		SecurityLoggingObjectRules: MakeSecurityLoggingObjectRuleListType(),
		SecurityLoggingObjectRate:  0,
		ParentType:                 "",
		FQName:                     []string{},
		DisplayName:                "",
		Annotations:                MakeKeyValuePairs(),
	}
}

// MakeSecurityLoggingObjectSlice() makes a slice of SecurityLoggingObject
func MakeSecurityLoggingObjectSlice() []*SecurityLoggingObject {
	return []*SecurityLoggingObject{}
}
