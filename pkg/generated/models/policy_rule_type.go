package models


// MakePolicyRuleType makes PolicyRuleType
func MakePolicyRuleType() *PolicyRuleType{
    return &PolicyRuleType{
    //TODO(nati): Apply default
    Direction: "",
        Protocol: "",
        
            
                DSTAddresses:  MakeAddressTypeSlice(),
            
        ActionList: MakeActionListType(),
        Created: "",
        RuleUUID: "",
        
            
                DSTPorts:  MakePortTypeSlice(),
            
        Application: []string{},
        LastModified: "",
        Ethertype: "",
        
            
                SRCAddresses:  MakeAddressTypeSlice(),
            
        RuleSequence: MakeSequenceType(),
        
            
                SRCPorts:  MakePortTypeSlice(),
            
        
    }
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
    return []*PolicyRuleType{}
}


