package models

// QuotaType

import "encoding/json"

// QuotaType
type QuotaType struct {
	VirtualRouter             int `json:"virtual_router,omitempty"`
	NetworkPolicy             int `json:"network_policy,omitempty"`
	LogicalRouter             int `json:"logical_router,omitempty"`
	FloatingIP                int `json:"floating_ip,omitempty"`
	LoadbalancerMember        int `json:"loadbalancer_member,omitempty"`
	AccessControlList         int `json:"access_control_list,omitempty"`
	SecurityGroup             int `json:"security_group,omitempty"`
	RouteTable                int `json:"route_table,omitempty"`
	VirtualMachineInterface   int `json:"virtual_machine_interface,omitempty"`
	InstanceIP                int `json:"instance_ip,omitempty"`
	GlobalVrouterConfig       int `json:"global_vrouter_config,omitempty"`
	VirtualIP                 int `json:"virtual_ip,omitempty"`
	Defaults                  int `json:"defaults,omitempty"`
	VirtualNetwork            int `json:"virtual_network,omitempty"`
	NetworkIpam               int `json:"network_ipam,omitempty"`
	VirtualDNS                int `json:"virtual_DNS,omitempty"`
	FloatingIPPool            int `json:"floating_ip_pool,omitempty"`
	SecurityLoggingObject     int `json:"security_logging_object,omitempty"`
	LoadbalancerHealthmonitor int `json:"loadbalancer_healthmonitor,omitempty"`
	LoadbalancerPool          int `json:"loadbalancer_pool,omitempty"`
	Subnet                    int `json:"subnet,omitempty"`
	VirtualDNSRecord          int `json:"virtual_DNS_record,omitempty"`
	SecurityGroupRule         int `json:"security_group_rule,omitempty"`
	ServiceInstance           int `json:"service_instance,omitempty"`
	ServiceTemplate           int `json:"service_template,omitempty"`
	BGPRouter                 int `json:"bgp_router,omitempty"`
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
		VirtualDNS:                0,
		FloatingIPPool:            0,
		SecurityLoggingObject:     0,
		NetworkIpam:               0,
		Subnet:                    0,
		VirtualDNSRecord:          0,
		SecurityGroupRule:         0,
		ServiceInstance:           0,
		ServiceTemplate:           0,
		BGPRouter:                 0,
		LoadbalancerHealthmonitor: 0,
		LoadbalancerPool:          0,
		NetworkPolicy:             0,
		LogicalRouter:             0,
		FloatingIP:                0,
		LoadbalancerMember:        0,
		AccessControlList:         0,
		SecurityGroup:             0,
		VirtualRouter:             0,
		VirtualMachineInterface:   0,
		InstanceIP:                0,
		GlobalVrouterConfig:       0,
		VirtualIP:                 0,
		Defaults:                  0,
		VirtualNetwork:            0,
		RouteTable:                0,
	}
}

// MakeQuotaTypeSlice() makes a slice of QuotaType
func MakeQuotaTypeSlice() []*QuotaType {
	return []*QuotaType{}
}
