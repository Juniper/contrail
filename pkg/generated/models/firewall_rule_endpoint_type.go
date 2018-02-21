package models


// MakeFirewallRuleEndpointType makes FirewallRuleEndpointType
func MakeFirewallRuleEndpointType() *FirewallRuleEndpointType{
    return &FirewallRuleEndpointType{
    //TODO(nati): Apply default
    AddressGroup: "",
        Subnet: MakeSubnetType(),
        Tags: []string{},
        
            
                TagIds: []int64{},
            
        VirtualNetwork: "",
        Any: false,
        
    }
}

// MakeFirewallRuleEndpointTypeSlice() makes a slice of FirewallRuleEndpointType
func MakeFirewallRuleEndpointTypeSlice() []*FirewallRuleEndpointType {
    return []*FirewallRuleEndpointType{}
}


