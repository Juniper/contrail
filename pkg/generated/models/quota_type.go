package models

// QuotaType

import "encoding/json"

// QuotaType
type QuotaType struct {
	Defaults                  int `json:"defaults"`
	LogicalRouter             int `json:"logical_router"`
	BGPRouter                 int `json:"bgp_router"`
	GlobalVrouterConfig       int `json:"global_vrouter_config"`
	SecurityLoggingObject     int `json:"security_logging_object"`
	NetworkIpam               int `json:"network_ipam"`
	ServiceInstance           int `json:"service_instance"`
	InstanceIP                int `json:"instance_ip"`
	VirtualIP                 int `json:"virtual_ip"`
	VirtualNetwork            int `json:"virtual_network"`
	LoadbalancerPool          int `json:"loadbalancer_pool"`
	Subnet                    int `json:"subnet"`
	SecurityGroupRule         int `json:"security_group_rule"`
	ServiceTemplate           int `json:"service_template"`
	VirtualDNS                int `json:"virtual_DNS"`
	FloatingIP                int `json:"floating_ip"`
	FloatingIPPool            int `json:"floating_ip_pool"`
	LoadbalancerMember        int `json:"loadbalancer_member"`
	VirtualRouter             int `json:"virtual_router"`
	NetworkPolicy             int `json:"network_policy"`
	RouteTable                int `json:"route_table"`
	VirtualDNSRecord          int `json:"virtual_DNS_record"`
	AccessControlList         int `json:"access_control_list"`
	VirtualMachineInterface   int `json:"virtual_machine_interface"`
	LoadbalancerHealthmonitor int `json:"loadbalancer_healthmonitor"`
	SecurityGroup             int `json:"security_group"`
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
		NetworkIpam:               0,
		ServiceInstance:           0,
		VirtualIP:                 0,
		VirtualNetwork:            0,
		LoadbalancerPool:          0,
		Subnet:                    0,
		SecurityGroupRule:         0,
		ServiceTemplate:           0,
		InstanceIP:                0,
		FloatingIP:                0,
		FloatingIPPool:            0,
		LoadbalancerMember:        0,
		VirtualRouter:             0,
		NetworkPolicy:             0,
		RouteTable:                0,
		VirtualDNSRecord:          0,
		VirtualDNS:                0,
		AccessControlList:         0,
		VirtualMachineInterface:   0,
		LoadbalancerHealthmonitor: 0,
		SecurityGroup:             0,
		LogicalRouter:             0,
		BGPRouter:                 0,
		GlobalVrouterConfig:       0,
		SecurityLoggingObject:     0,
		Defaults:                  0,
	}
}

// InterfaceToQuotaType makes QuotaType from interface
func InterfaceToQuotaType(iData interface{}) *QuotaType {
	data := iData.(map[string]interface{})
	return &QuotaType{
		Defaults: data["defaults"].(int),

		//{"description":"Need to clarify","type":"integer"}
		LogicalRouter: data["logical_router"].(int),

		//{"description":"Maximum number of logical routers","type":"integer"}
		BGPRouter: data["bgp_router"].(int),

		//{"description":"Maximum number of bgp routers","type":"integer"}
		GlobalVrouterConfig: data["global_vrouter_config"].(int),

		//{"description":"Maximum number of global vrouter configs","type":"integer"}
		SecurityLoggingObject: data["security_logging_object"].(int),

		//{"description":"Maximum number of security logging objects","type":"integer"}
		NetworkIpam: data["network_ipam"].(int),

		//{"description":"Maximum number of network IPAMs","type":"integer"}
		ServiceInstance: data["service_instance"].(int),

		//{"description":"Maximum number of service instances","type":"integer"}
		InstanceIP: data["instance_ip"].(int),

		//{"description":"Maximum number of instance ips","type":"integer"}
		VirtualIP: data["virtual_ip"].(int),

		//{"description":"Maximum number of virtual ips","type":"integer"}
		VirtualNetwork: data["virtual_network"].(int),

		//{"description":"Maximum number of virtual networks","type":"integer"}
		LoadbalancerPool: data["loadbalancer_pool"].(int),

		//{"description":"Maximum number of loadbalancer pools","type":"integer"}
		Subnet: data["subnet"].(int),

		//{"description":"Maximum number of subnets","type":"integer"}
		SecurityGroupRule: data["security_group_rule"].(int),

		//{"description":"Maximum number of security group rules","type":"integer"}
		ServiceTemplate: data["service_template"].(int),

		//{"description":"Maximum number of service templates","type":"integer"}
		VirtualDNS: data["virtual_DNS"].(int),

		//{"description":"Maximum number of virtual DNS servers","type":"integer"}
		FloatingIP: data["floating_ip"].(int),

		//{"description":"Maximum number of floating ips","type":"integer"}
		FloatingIPPool: data["floating_ip_pool"].(int),

		//{"description":"Maximum number of floating ip pools","type":"integer"}
		LoadbalancerMember: data["loadbalancer_member"].(int),

		//{"description":"Maximum number of loadbalancer member","type":"integer"}
		VirtualRouter: data["virtual_router"].(int),

		//{"description":"Maximum number of logical routers","type":"integer"}
		NetworkPolicy: data["network_policy"].(int),

		//{"description":"Maximum number of network policies","type":"integer"}
		RouteTable: data["route_table"].(int),

		//{"description":"Maximum number of route tables","type":"integer"}
		VirtualDNSRecord: data["virtual_DNS_record"].(int),

		//{"description":"Maximum number of virtual DNS records","type":"integer"}
		AccessControlList: data["access_control_list"].(int),

		//{"description":"Maximum number of access control lists","type":"integer"}
		VirtualMachineInterface: data["virtual_machine_interface"].(int),

		//{"description":"Maximum number of virtual machine interfaces","type":"integer"}
		LoadbalancerHealthmonitor: data["loadbalancer_healthmonitor"].(int),

		//{"description":"Maximum number of loadbalancer health monitors","type":"integer"}
		SecurityGroup: data["security_group"].(int),

		//{"description":"Maximum number of security groups","type":"integer"}

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
