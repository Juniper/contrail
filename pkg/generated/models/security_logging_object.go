package models

// SecurityLoggingObject

import "encoding/json"

// SecurityLoggingObject
type SecurityLoggingObject struct {
	SecurityLoggingObjectRate  int                                `json:"security_logging_object_rate"`
	UUID                       string                             `json:"uuid"`
	ParentUUID                 string                             `json:"parent_uuid"`
	ParentType                 string                             `json:"parent_type"`
	FQName                     []string                           `json:"fq_name"`
	Annotations                *KeyValuePairs                     `json:"annotations"`
	SecurityLoggingObjectRules *SecurityLoggingObjectRuleListType `json:"security_logging_object_rules"`
	Perms2                     *PermType2                         `json:"perms2"`
	IDPerms                    *IdPermsType                       `json:"id_perms"`
	DisplayName                string                             `json:"display_name"`

	SecurityGroupRefs []*SecurityLoggingObjectSecurityGroupRef `json:"security_group_refs"`
	NetworkPolicyRefs []*SecurityLoggingObjectNetworkPolicyRef `json:"network_policy_refs"`
}

// SecurityLoggingObjectNetworkPolicyRef references each other
type SecurityLoggingObjectNetworkPolicyRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *SecurityLoggingObjectRuleListType
}

// SecurityLoggingObjectSecurityGroupRef references each other
type SecurityLoggingObjectSecurityGroupRef struct {
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
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		Annotations:               MakeKeyValuePairs(),
		SecurityLoggingObjectRate: 0,
		UUID:                       "",
		IDPerms:                    MakeIdPermsType(),
		DisplayName:                "",
		SecurityLoggingObjectRules: MakeSecurityLoggingObjectRuleListType(),
		Perms2: MakePermType2(),
	}
}

// MakeSecurityLoggingObjectSlice() makes a slice of SecurityLoggingObject
func MakeSecurityLoggingObjectSlice() []*SecurityLoggingObject {
	return []*SecurityLoggingObject{}
}
