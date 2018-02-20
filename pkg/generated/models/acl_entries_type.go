package models

// AclEntriesType

// AclEntriesType
//proteus:generate
type AclEntriesType struct {
	Dynamic bool           `json:"dynamic"`
	ACLRule []*AclRuleType `json:"acl_rule,omitempty"`
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
