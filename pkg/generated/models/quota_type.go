package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


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

// MakeQuotaType makes QuotaType
func InterfaceToQuotaType(i interface{}) *QuotaType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &QuotaType{
    //TODO(nati): Apply default
    VirtualRouter: schema.InterfaceToInt64(m["virtual_router"]),
        NetworkPolicy: schema.InterfaceToInt64(m["network_policy"]),
        LoadbalancerPool: schema.InterfaceToInt64(m["loadbalancer_pool"]),
        RouteTable: schema.InterfaceToInt64(m["route_table"]),
        Subnet: schema.InterfaceToInt64(m["subnet"]),
        NetworkIpam: schema.InterfaceToInt64(m["network_ipam"]),
        VirtualDNSRecord: schema.InterfaceToInt64(m["virtual_DNS_record"]),
        LogicalRouter: schema.InterfaceToInt64(m["logical_router"]),
        SecurityGroupRule: schema.InterfaceToInt64(m["security_group_rule"]),
        VirtualDNS: schema.InterfaceToInt64(m["virtual_DNS"]),
        ServiceInstance: schema.InterfaceToInt64(m["service_instance"]),
        ServiceTemplate: schema.InterfaceToInt64(m["service_template"]),
        BGPRouter: schema.InterfaceToInt64(m["bgp_router"]),
        FloatingIP: schema.InterfaceToInt64(m["floating_ip"]),
        FloatingIPPool: schema.InterfaceToInt64(m["floating_ip_pool"]),
        LoadbalancerMember: schema.InterfaceToInt64(m["loadbalancer_member"]),
        AccessControlList: schema.InterfaceToInt64(m["access_control_list"]),
        VirtualMachineInterface: schema.InterfaceToInt64(m["virtual_machine_interface"]),
        InstanceIP: schema.InterfaceToInt64(m["instance_ip"]),
        GlobalVrouterConfig: schema.InterfaceToInt64(m["global_vrouter_config"]),
        SecurityLoggingObject: schema.InterfaceToInt64(m["security_logging_object"]),
        LoadbalancerHealthmonitor: schema.InterfaceToInt64(m["loadbalancer_healthmonitor"]),
        VirtualIP: schema.InterfaceToInt64(m["virtual_ip"]),
        Defaults: schema.InterfaceToInt64(m["defaults"]),
        SecurityGroup: schema.InterfaceToInt64(m["security_group"]),
        VirtualNetwork: schema.InterfaceToInt64(m["virtual_network"]),
        
    }
}

// MakeQuotaTypeSlice() makes a slice of QuotaType
func MakeQuotaTypeSlice() []*QuotaType {
    return []*QuotaType{}
}

// InterfaceToQuotaTypeSlice() makes a slice of QuotaType
func InterfaceToQuotaTypeSlice(i interface{}) []*QuotaType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*QuotaType{}
    for _, item := range list {
        result = append(result, InterfaceToQuotaType(item) )
    }
    return result
}



