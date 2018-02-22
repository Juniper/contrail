package models


// MakeFirewallServiceGroupType makes FirewallServiceGroupType
func MakeFirewallServiceGroupType() *FirewallServiceGroupType{
    return &FirewallServiceGroupType{
    //TODO(nati): Apply default
    
            
                FirewallService:  MakeFirewallServiceTypeSlice(),
            
        
    }
}

// MakeFirewallServiceGroupTypeSlice() makes a slice of FirewallServiceGroupType
func MakeFirewallServiceGroupTypeSlice() []*FirewallServiceGroupType {
    return []*FirewallServiceGroupType{}
}


