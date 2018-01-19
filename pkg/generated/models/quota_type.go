package models

// QuotaType

import "encoding/json"

// QuotaType
type QuotaType struct {
	InstanceIP                int `json:"instance_ip,omitempty"`
	VirtualIP                 int `json:"virtual_ip,omitempty"`
	Defaults                  int `json:"defaults,omitempty"`
	VirtualNetwork            int `json:"virtual_network,omitempty"`
	VirtualDNSRecord          int `json:"virtual_DNS_record,omitempty"`
	VirtualDNS                int `json:"virtual_DNS,omitempty"`
	BGPRouter                 int `json:"bgp_router,omitempty"`
	FloatingIP                int `json:"floating_ip,omitempty"`
	VirtualMachineInterface   int `json:"virtual_machine_interface,omitempty"`
	SecurityLoggingObject     int `json:"security_logging_object,omitempty"`
	RouteTable                int `json:"route_table,omitempty"`
	ServiceInstance           int `json:"service_instance,omitempty"`
	NetworkIpam               int `json:"network_ipam,omitempty"`
	FloatingIPPool            int `json:"floating_ip_pool,omitempty"`
	LoadbalancerMember        int `json:"loadbalancer_member,omitempty"`
	LoadbalancerHealthmonitor int `json:"loadbalancer_healthmonitor,omitempty"`
	LoadbalancerPool          int `json:"loadbalancer_pool,omitempty"`
	Subnet                    int `json:"subnet,omitempty"`
	LogicalRouter             int `json:"logical_router,omitempty"`
	SecurityGroupRule         int `json:"security_group_rule,omitempty"`
	ServiceTemplate           int `json:"service_template,omitempty"`
	AccessControlList         int `json:"access_control_list,omitempty"`
	GlobalVrouterConfig       int `json:"global_vrouter_config,omitempty"`
	SecurityGroup             int `json:"security_group,omitempty"`
	VirtualRouter             int `json:"virtual_router,omitempty"`
	NetworkPolicy             int `json:"network_policy,omitempty"`
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
		VirtualDNSRecord:          0,
		VirtualDNS:                0,
		InstanceIP:                0,
		VirtualIP:                 0,
		Defaults:                  0,
		VirtualNetwork:            0,
		RouteTable:                0,
		ServiceInstance:           0,
		BGPRouter:                 0,
		FloatingIP:                0,
		VirtualMachineInterface:   0,
		SecurityLoggingObject:     0,
		LoadbalancerPool:          0,
		Subnet:                    0,
		NetworkIpam:               0,
		FloatingIPPool:            0,
		LoadbalancerMember:        0,
		LoadbalancerHealthmonitor: 0,
		GlobalVrouterConfig:       0,
		SecurityGroup:             0,
		VirtualRouter:             0,
		NetworkPolicy:             0,
		LogicalRouter:             0,
		SecurityGroupRule:         0,
		ServiceTemplate:           0,
		AccessControlList:         0,
	}
}

// MakeQuotaTypeSlice() makes a slice of QuotaType
func MakeQuotaTypeSlice() []*QuotaType {
	return []*QuotaType{}
}
