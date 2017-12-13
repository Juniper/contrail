package models

// QuotaType

import "encoding/json"

// QuotaType
type QuotaType struct {
	VirtualRouter             int `json:"virtual_router"`
	NetworkPolicy             int `json:"network_policy"`
	RouteTable                int `json:"route_table"`
	SecurityGroupRule         int `json:"security_group_rule"`
	ServiceInstance           int `json:"service_instance"`
	LoadbalancerMember        int `json:"loadbalancer_member"`
	AccessControlList         int `json:"access_control_list"`
	VirtualMachineInterface   int `json:"virtual_machine_interface"`
	SecurityLoggingObject     int `json:"security_logging_object"`
	SecurityGroup             int `json:"security_group"`
	VirtualNetwork            int `json:"virtual_network"`
	VirtualDNSRecord          int `json:"virtual_DNS_record"`
	Defaults                  int `json:"defaults"`
	Subnet                    int `json:"subnet"`
	NetworkIpam               int `json:"network_ipam"`
	VirtualDNS                int `json:"virtual_DNS"`
	ServiceTemplate           int `json:"service_template"`
	FloatingIPPool            int `json:"floating_ip_pool"`
	InstanceIP                int `json:"instance_ip"`
	GlobalVrouterConfig       int `json:"global_vrouter_config"`
	VirtualIP                 int `json:"virtual_ip"`
	LoadbalancerPool          int `json:"loadbalancer_pool"`
	LogicalRouter             int `json:"logical_router"`
	BGPRouter                 int `json:"bgp_router"`
	FloatingIP                int `json:"floating_ip"`
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
		AccessControlList:         0,
		VirtualMachineInterface:   0,
		VirtualRouter:             0,
		NetworkPolicy:             0,
		RouteTable:                0,
		SecurityGroupRule:         0,
		ServiceInstance:           0,
		LoadbalancerMember:        0,
		SecurityLoggingObject:     0,
		SecurityGroup:             0,
		VirtualNetwork:            0,
		VirtualDNSRecord:          0,
		Defaults:                  0,
		GlobalVrouterConfig:       0,
		VirtualIP:                 0,
		Subnet:                    0,
		NetworkIpam:               0,
		VirtualDNS:                0,
		ServiceTemplate:           0,
		FloatingIPPool:            0,
		InstanceIP:                0,
		LoadbalancerPool:          0,
		LogicalRouter:             0,
		BGPRouter:                 0,
		FloatingIP:                0,
		LoadbalancerHealthmonitor: 0,
	}
}

// InterfaceToQuotaType makes QuotaType from interface
func InterfaceToQuotaType(iData interface{}) *QuotaType {
	data := iData.(map[string]interface{})
	return &QuotaType{
		NetworkPolicy: data["network_policy"].(int),

		//{"description":"Maximum number of network policies","type":"integer"}
		RouteTable: data["route_table"].(int),

		//{"description":"Maximum number of route tables","type":"integer"}
		SecurityGroupRule: data["security_group_rule"].(int),

		//{"description":"Maximum number of security group rules","type":"integer"}
		ServiceInstance: data["service_instance"].(int),

		//{"description":"Maximum number of service instances","type":"integer"}
		LoadbalancerMember: data["loadbalancer_member"].(int),

		//{"description":"Maximum number of loadbalancer member","type":"integer"}
		AccessControlList: data["access_control_list"].(int),

		//{"description":"Maximum number of access control lists","type":"integer"}
		VirtualMachineInterface: data["virtual_machine_interface"].(int),

		//{"description":"Maximum number of virtual machine interfaces","type":"integer"}
		VirtualRouter: data["virtual_router"].(int),

		//{"description":"Maximum number of logical routers","type":"integer"}
		SecurityGroup: data["security_group"].(int),

		//{"description":"Maximum number of security groups","type":"integer"}
		VirtualNetwork: data["virtual_network"].(int),

		//{"description":"Maximum number of virtual networks","type":"integer"}
		SecurityLoggingObject: data["security_logging_object"].(int),

		//{"description":"Maximum number of security logging objects","type":"integer"}
		Defaults: data["defaults"].(int),

		//{"description":"Need to clarify","type":"integer"}
		VirtualDNSRecord: data["virtual_DNS_record"].(int),

		//{"description":"Maximum number of virtual DNS records","type":"integer"}
		NetworkIpam: data["network_ipam"].(int),

		//{"description":"Maximum number of network IPAMs","type":"integer"}
		VirtualDNS: data["virtual_DNS"].(int),

		//{"description":"Maximum number of virtual DNS servers","type":"integer"}
		ServiceTemplate: data["service_template"].(int),

		//{"description":"Maximum number of service templates","type":"integer"}
		FloatingIPPool: data["floating_ip_pool"].(int),

		//{"description":"Maximum number of floating ip pools","type":"integer"}
		InstanceIP: data["instance_ip"].(int),

		//{"description":"Maximum number of instance ips","type":"integer"}
		GlobalVrouterConfig: data["global_vrouter_config"].(int),

		//{"description":"Maximum number of global vrouter configs","type":"integer"}
		VirtualIP: data["virtual_ip"].(int),

		//{"description":"Maximum number of virtual ips","type":"integer"}
		Subnet: data["subnet"].(int),

		//{"description":"Maximum number of subnets","type":"integer"}
		LogicalRouter: data["logical_router"].(int),

		//{"description":"Maximum number of logical routers","type":"integer"}
		BGPRouter: data["bgp_router"].(int),

		//{"description":"Maximum number of bgp routers","type":"integer"}
		FloatingIP: data["floating_ip"].(int),

		//{"description":"Maximum number of floating ips","type":"integer"}
		LoadbalancerHealthmonitor: data["loadbalancer_healthmonitor"].(int),

		//{"description":"Maximum number of loadbalancer health monitors","type":"integer"}
		LoadbalancerPool: data["loadbalancer_pool"].(int),

		//{"description":"Maximum number of loadbalancer pools","type":"integer"}

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
