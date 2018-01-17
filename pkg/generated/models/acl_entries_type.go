package models

// AclEntriesType

import "encoding/json"

// AclEntriesType
type AclEntriesType struct {
	Dynamic bool           `json:"dynamic,omitempty"`
	ACLRule []*AclRuleType `json:"acl_rule,omitempty"`
}

// String returns json representation of the object
func (model *AclEntriesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAclEntriesType makes AclEntriesType
func MakeAclEntriesType() *AclEntriesType {
	return &AclEntriesType{
		//TODO(nati): Apply default
		Dynamic: false,

		ACLRule: MakeAclRuleTypeSlice(),
	}
}

// MakeAclEntriesTypeSlice() makes a slice of AclEntriesType
func MakeAclEntriesTypeSlice() []*AclEntriesType {
	return []*AclEntriesType{}
}
