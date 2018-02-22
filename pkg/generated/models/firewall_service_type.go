package models


// MakeFirewallServiceType makes FirewallServiceType
func MakeFirewallServiceType() *FirewallServiceType{
    return &FirewallServiceType{
    //TODO(nati): Apply default
    Protocol: "",
        DSTPorts: MakePortType(),
        SRCPorts: MakePortType(),
        ProtocolID: 0,
        
    }
}

// MakeFirewallServiceTypeSlice() makes a slice of FirewallServiceType
func MakeFirewallServiceTypeSlice() []*FirewallServiceType {
    return []*FirewallServiceType{}
}


