package models

// QuotaType

import "encoding/json"

// QuotaType
type QuotaType struct {
	SecurityLoggingObject     int `json:"security_logging_object"`
	LogicalRouter             int `json:"logical_router"`
	SecurityGroupRule         int `json:"security_group_rule"`
	ServiceTemplate           int `json:"service_template"`
	BGPRouter                 int `json:"bgp_router"`
	FloatingIPPool            int `json:"floating_ip_pool"`
	AccessControlList         int `json:"access_control_list"`
	LoadbalancerHealthmonitor int `json:"loadbalancer_healthmonitor"`
	VirtualNetwork            int `json:"virtual_network"`
	Subnet                    int `json:"subnet"`
	NetworkIpam               int `json:"network_ipam"`
	VirtualDNS                int `json:"virtual_DNS"`
	ServiceInstance           int `json:"service_instance"`
	LoadbalancerMember        int `json:"loadbalancer_member"`
	InstanceIP                int `json:"instance_ip"`
	LoadbalancerPool          int `json:"loadbalancer_pool"`
	VirtualDNSRecord          int `json:"virtual_DNS_record"`
	FloatingIP                int `json:"floating_ip"`
	VirtualMachineInterface   int `json:"virtual_machine_interface"`
	GlobalVrouterConfig       int `json:"global_vrouter_config"`
	SecurityGroup             int `json:"security_group"`
	VirtualRouter             int `json:"virtual_router"`
	NetworkPolicy             int `json:"network_policy"`
	RouteTable                int `json:"route_table"`
	VirtualIP                 int `json:"virtual_ip"`
	Defaults                  int `json:"defaults"`
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
		VirtualMachineInterface:   0,
		GlobalVrouterConfig:       0,
		SecurityGroup:             0,
		LoadbalancerPool:          0,
		VirtualDNSRecord:          0,
		FloatingIP:                0,
		VirtualIP:                 0,
		Defaults:                  0,
		VirtualRouter:             0,
		NetworkPolicy:             0,
		RouteTable:                0,
		BGPRouter:                 0,
		FloatingIPPool:            0,
		AccessControlList:         0,
		SecurityLoggingObject:     0,
		LogicalRouter:             0,
		SecurityGroupRule:         0,
		ServiceTemplate:           0,
		ServiceInstance:           0,
		LoadbalancerMember:        0,
		InstanceIP:                0,
		LoadbalancerHealthmonitor: 0,
		VirtualNetwork:            0,
		Subnet:                    0,
		NetworkIpam:               0,
		VirtualDNS:                0,
	}
}

// InterfaceToQuotaType makes QuotaType from interface
func InterfaceToQuotaType(iData interface{}) *QuotaType {
	data := iData.(map[string]interface{})
	return &QuotaType{
		NetworkIpam: data["network_ipam"].(int),

		//{"description":"Maximum number of network IPAMs","type":"integer"}
		VirtualDNS: data["virtual_DNS"].(int),

		//{"description":"Maximum number of virtual DNS servers","type":"integer"}
		ServiceInstance: data["service_instance"].(int),

		//{"description":"Maximum number of service instances","type":"integer"}
		LoadbalancerMember: data["loadbalancer_member"].(int),

		//{"description":"Maximum number of loadbalancer member","type":"integer"}
		InstanceIP: data["instance_ip"].(int),

		//{"description":"Maximum number of instance ips","type":"integer"}
		LoadbalancerHealthmonitor: data["loadbalancer_healthmonitor"].(int),

		//{"description":"Maximum number of loadbalancer health monitors","type":"integer"}
		VirtualNetwork: data["virtual_network"].(int),

		//{"description":"Maximum number of virtual networks","type":"integer"}
		Subnet: data["subnet"].(int),

		//{"description":"Maximum number of subnets","type":"integer"}
		VirtualDNSRecord: data["virtual_DNS_record"].(int),

		//{"description":"Maximum number of virtual DNS records","type":"integer"}
		FloatingIP: data["floating_ip"].(int),

		//{"description":"Maximum number of floating ips","type":"integer"}
		VirtualMachineInterface: data["virtual_machine_interface"].(int),

		//{"description":"Maximum number of virtual machine interfaces","type":"integer"}
		GlobalVrouterConfig: data["global_vrouter_config"].(int),

		//{"description":"Maximum number of global vrouter configs","type":"integer"}
		SecurityGroup: data["security_group"].(int),

		//{"description":"Maximum number of security groups","type":"integer"}
		LoadbalancerPool: data["loadbalancer_pool"].(int),

		//{"description":"Maximum number of loadbalancer pools","type":"integer"}
		NetworkPolicy: data["network_policy"].(int),

		//{"description":"Maximum number of network policies","type":"integer"}
		RouteTable: data["route_table"].(int),

		//{"description":"Maximum number of route tables","type":"integer"}
		VirtualIP: data["virtual_ip"].(int),

		//{"description":"Maximum number of virtual ips","type":"integer"}
		Defaults: data["defaults"].(int),

		//{"description":"Need to clarify","type":"integer"}
		VirtualRouter: data["virtual_router"].(int),

		//{"description":"Maximum number of logical routers","type":"integer"}
		SecurityGroupRule: data["security_group_rule"].(int),

		//{"description":"Maximum number of security group rules","type":"integer"}
		ServiceTemplate: data["service_template"].(int),

		//{"description":"Maximum number of service templates","type":"integer"}
		BGPRouter: data["bgp_router"].(int),

		//{"description":"Maximum number of bgp routers","type":"integer"}
		FloatingIPPool: data["floating_ip_pool"].(int),

		//{"description":"Maximum number of floating ip pools","type":"integer"}
		AccessControlList: data["access_control_list"].(int),

		//{"description":"Maximum number of access control lists","type":"integer"}
		SecurityLoggingObject: data["security_logging_object"].(int),

		//{"description":"Maximum number of security logging objects","type":"integer"}
		LogicalRouter: data["logical_router"].(int),

		//{"description":"Maximum number of logical routers","type":"integer"}

	}
}

// InterfaceToQuotaTypeSlice makes a slice of QuotaType from interface
func InterfaceToQuotaTypeSlice(data interface{}) []*QuotaType {
	list := data.([]interface{})
	result := MakeQuotaTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToQuotaType(item))
	}
	return result
}

// MakeQuotaTypeSlice() makes a slice of QuotaType
func MakeQuotaTypeSlice() []*QuotaType {
	return []*QuotaType{}
}
