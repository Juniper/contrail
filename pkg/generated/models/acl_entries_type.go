package models


// MakeAclEntriesType makes AclEntriesType
func MakeAclEntriesType() *AclEntriesType{
    return &AclEntriesType{
    //TODO(nati): Apply default
    Dynamic: false,
        
            
                ACLRule:  MakeAclRuleTypeSlice(),
            
        
    }
}

// MakeAclEntriesTypeSlice() makes a slice of AclEntriesType
func MakeAclEntriesTypeSlice() []*AclEntriesType {
    return []*AclEntriesType{}
}


