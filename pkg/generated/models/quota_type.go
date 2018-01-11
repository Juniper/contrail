package models

// QuotaType

import "encoding/json"

// QuotaType
type QuotaType struct {
	LoadbalancerPool          int `json:"loadbalancer_pool"`
	RouteTable                int `json:"route_table"`
	Subnet                    int `json:"subnet"`
	SecurityGroupRule         int `json:"security_group_rule"`
	LoadbalancerMember        int `json:"loadbalancer_member"`
	VirtualMachineInterface   int `json:"virtual_machine_interface"`
	GlobalVrouterConfig       int `json:"global_vrouter_config"`
	NetworkPolicy             int `json:"network_policy"`
	VirtualDNS                int `json:"virtual_DNS"`
	FloatingIP                int `json:"floating_ip"`
	SecurityGroup             int `json:"security_group"`
	VirtualRouter             int `json:"virtual_router"`
	LogicalRouter             int `json:"logical_router"`
	ServiceInstance           int `json:"service_instance"`
	FloatingIPPool            int `json:"floating_ip_pool"`
	AccessControlList         int `json:"access_control_list"`
	VirtualIP                 int `json:"virtual_ip"`
	Defaults                  int `json:"defaults"`
	VirtualNetwork            int `json:"virtual_network"`
	NetworkIpam               int `json:"network_ipam"`
	VirtualDNSRecord          int `json:"virtual_DNS_record"`
	ServiceTemplate           int `json:"service_template"`
	BGPRouter                 int `json:"bgp_router"`
	InstanceIP                int `json:"instance_ip"`
	SecurityLoggingObject     int `json:"security_logging_object"`
	LoadbalancerHealthmonitor int `json:"loadbalancer_healthmonitor"`
}

// String returns json representation of the object
func (model *QuotaType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeQuotaType makes QuotaType
func MakeQuotaType() *QuotaType {
	return &QuotaType{
		//TODO(nati): Apply default
		ServiceTemplate:           0,
		BGPRouter:                 0,
		InstanceIP:                0,
		SecurityLoggingObject:     0,
		LoadbalancerHealthmonitor: 0,
		NetworkIpam:               0,
		VirtualDNSRecord:          0,
		Subnet:                    0,
		SecurityGroupRule:         0,
		LoadbalancerMember:        0,
		VirtualMachineInterface:   0,
		GlobalVrouterConfig:       0,
		LoadbalancerPool:          0,
		RouteTable:                0,
		FloatingIP:                0,
		SecurityGroup:             0,
		NetworkPolicy:             0,
		VirtualDNS:                0,
		ServiceInstance:           0,
		FloatingIPPool:            0,
		AccessControlList:         0,
		VirtualIP:                 0,
		Defaults:                  0,
		VirtualNetwork:            0,
		VirtualRouter:             0,
		LogicalRouter:             0,
	}
}

// MakeQuotaTypeSlice() makes a slice of QuotaType
func MakeQuotaTypeSlice() []*QuotaType {
	return []*QuotaType{}
}
