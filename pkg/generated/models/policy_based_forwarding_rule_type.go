package models


// MakePolicyBasedForwardingRuleType makes PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleType() *PolicyBasedForwardingRuleType{
    return &PolicyBasedForwardingRuleType{
    //TODO(nati): Apply default
    DSTMac: "",
        Protocol: "",
        Ipv6ServiceChainAddress: "",
        Direction: "",
        MPLSLabel: 0,
        VlanTag: 0,
        SRCMac: "",
        ServiceChainAddress: "",
        
    }
}

// MakePolicyBasedForwardingRuleTypeSlice() makes a slice of PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleTypeSlice() []*PolicyBasedForwardingRuleType {
    return []*PolicyBasedForwardingRuleType{}
}


