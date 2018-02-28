package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeQuotaType makes QuotaType
// nolint
func MakeQuotaType() *QuotaType {
	return &QuotaType{
		//TODO(nati): Apply default
		VirtualRouter:             0,
		NetworkPolicy:             0,
		LoadbalancerPool:          0,
		RouteTable:                0,
		Subnet:                    0,
		NetworkIpam:               0,
		VirtualDNSRecord:          0,
		LogicalRouter:             0,
		SecurityGroupRule:         0,
		VirtualDNS:                0,
		ServiceInstance:           0,
		ServiceTemplate:           0,
		BGPRouter:                 0,
		FloatingIP:                0,
		FloatingIPPool:            0,
		LoadbalancerMember:        0,
		AccessControlList:         0,
		VirtualMachineInterface:   0,
		InstanceIP:                0,
		GlobalVrouterConfig:       0,
		SecurityLoggingObject:     0,
		LoadbalancerHealthmonitor: 0,
		VirtualIP:                 0,
		Defaults:                  0,
		SecurityGroup:             0,
		VirtualNetwork:            0,
	}
}

// MakeQuotaType makes QuotaType
// nolint
func InterfaceToQuotaType(i interface{}) *QuotaType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &QuotaType{
		//TODO(nati): Apply default
		VirtualRouter:             common.InterfaceToInt64(m["virtual_router"]),
		NetworkPolicy:             common.InterfaceToInt64(m["network_policy"]),
		LoadbalancerPool:          common.InterfaceToInt64(m["loadbalancer_pool"]),
		RouteTable:                common.InterfaceToInt64(m["route_table"]),
		Subnet:                    common.InterfaceToInt64(m["subnet"]),
		NetworkIpam:               common.InterfaceToInt64(m["network_ipam"]),
		VirtualDNSRecord:          common.InterfaceToInt64(m["virtual_DNS_record"]),
		LogicalRouter:             common.InterfaceToInt64(m["logical_router"]),
		SecurityGroupRule:         common.InterfaceToInt64(m["security_group_rule"]),
		VirtualDNS:                common.InterfaceToInt64(m["virtual_DNS"]),
		ServiceInstance:           common.InterfaceToInt64(m["service_instance"]),
		ServiceTemplate:           common.InterfaceToInt64(m["service_template"]),
		BGPRouter:                 common.InterfaceToInt64(m["bgp_router"]),
		FloatingIP:                common.InterfaceToInt64(m["floating_ip"]),
		FloatingIPPool:            common.InterfaceToInt64(m["floating_ip_pool"]),
		LoadbalancerMember:        common.InterfaceToInt64(m["loadbalancer_member"]),
		AccessControlList:         common.InterfaceToInt64(m["access_control_list"]),
		VirtualMachineInterface:   common.InterfaceToInt64(m["virtual_machine_interface"]),
		InstanceIP:                common.InterfaceToInt64(m["instance_ip"]),
		GlobalVrouterConfig:       common.InterfaceToInt64(m["global_vrouter_config"]),
		SecurityLoggingObject:     common.InterfaceToInt64(m["security_logging_object"]),
		LoadbalancerHealthmonitor: common.InterfaceToInt64(m["loadbalancer_healthmonitor"]),
		VirtualIP:                 common.InterfaceToInt64(m["virtual_ip"]),
		Defaults:                  common.InterfaceToInt64(m["defaults"]),
		SecurityGroup:             common.InterfaceToInt64(m["security_group"]),
		VirtualNetwork:            common.InterfaceToInt64(m["virtual_network"]),
	}
}

// MakeQuotaTypeSlice() makes a slice of QuotaType
// nolint
func MakeQuotaTypeSlice() []*QuotaType {
	return []*QuotaType{}
}

// InterfaceToQuotaTypeSlice() makes a slice of QuotaType
// nolint
func InterfaceToQuotaTypeSlice(i interface{}) []*QuotaType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*QuotaType{}
	for _, item := range list {
		result = append(result, InterfaceToQuotaType(item))
	}
	return result
}
