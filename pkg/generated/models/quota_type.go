package models


// MakeQuotaType makes QuotaType
func MakeQuotaType() *QuotaType{
    return &QuotaType{
    //TODO(nati): Apply default
    VirtualRouter: 0,
        NetworkPolicy: 0,
        LoadbalancerPool: 0,
        RouteTable: 0,
        Subnet: 0,
        NetworkIpam: 0,
        VirtualDNSRecord: 0,
        LogicalRouter: 0,
        SecurityGroupRule: 0,
        VirtualDNS: 0,
        ServiceInstance: 0,
        ServiceTemplate: 0,
        BGPRouter: 0,
        FloatingIP: 0,
        FloatingIPPool: 0,
        LoadbalancerMember: 0,
        AccessControlList: 0,
        VirtualMachineInterface: 0,
        InstanceIP: 0,
        GlobalVrouterConfig: 0,
        SecurityLoggingObject: 0,
        LoadbalancerHealthmonitor: 0,
        VirtualIP: 0,
        Defaults: 0,
        SecurityGroup: 0,
        VirtualNetwork: 0,
        
    }
}

// MakeQuotaTypeSlice() makes a slice of QuotaType
func MakeQuotaTypeSlice() []*QuotaType {
    return []*QuotaType{}
}


