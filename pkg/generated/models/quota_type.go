package models

// QuotaType

import "encoding/json"

type QuotaType struct {
	NetworkIpam               int `json:"network_ipam"`
	VirtualDNSRecord          int `json:"virtual_DNS_record"`
	ServiceInstance           int `json:"service_instance"`
	SecurityLoggingObject     int `json:"security_logging_object"`
	LoadbalancerHealthmonitor int `json:"loadbalancer_healthmonitor"`
	SecurityGroup             int `json:"security_group"`
	RouteTable                int `json:"route_table"`
	Subnet                    int `json:"subnet"`
	VirtualNetwork            int `json:"virtual_network"`
	ServiceTemplate           int `json:"service_template"`
	BGPRouter                 int `json:"bgp_router"`
	FloatingIP                int `json:"floating_ip"`
	Defaults                  int `json:"defaults"`
	LoadbalancerPool          int `json:"loadbalancer_pool"`
	VirtualDNS                int `json:"virtual_DNS"`
	LogicalRouter             int `json:"logical_router"`
	FloatingIPPool            int `json:"floating_ip_pool"`
	VirtualMachineInterface   int `json:"virtual_machine_interface"`
	VirtualIP                 int `json:"virtual_ip"`
	VirtualRouter             int `json:"virtual_router"`
	NetworkPolicy             int `json:"network_policy"`
	AccessControlList         int `json:"access_control_list"`
	InstanceIP                int `json:"instance_ip"`
	GlobalVrouterConfig       int `json:"global_vrouter_config"`
	SecurityGroupRule         int `json:"security_group_rule"`
	LoadbalancerMember        int `json:"loadbalancer_member"`
}

func (model *QuotaType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeQuotaType() *QuotaType {
	return &QuotaType{
		//TODO(nati): Apply default
		Defaults:                  0,
		LoadbalancerPool:          0,
		VirtualDNS:                0,
		ServiceTemplate:           0,
		BGPRouter:                 0,
		FloatingIP:                0,
		VirtualIP:                 0,
		VirtualRouter:             0,
		NetworkPolicy:             0,
		LogicalRouter:             0,
		FloatingIPPool:            0,
		VirtualMachineInterface:   0,
		SecurityGroupRule:         0,
		LoadbalancerMember:        0,
		AccessControlList:         0,
		InstanceIP:                0,
		GlobalVrouterConfig:       0,
		SecurityLoggingObject:     0,
		LoadbalancerHealthmonitor: 0,
		SecurityGroup:             0,
		RouteTable:                0,
		Subnet:                    0,
		NetworkIpam:               0,
		VirtualDNSRecord:          0,
		ServiceInstance:           0,
		VirtualNetwork:            0,
	}
}

func InterfaceToQuotaType(iData interface{}) *QuotaType {
	data := iData.(map[string]interface{})
	return &QuotaType{
		VirtualRouter: data["virtual_router"].(int),

		//{"Title":"","Description":"Maximum number of logical routers","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualRouter","GoType":"int"}
		NetworkPolicy: data["network_policy"].(int),

		//{"Title":"","Description":"Maximum number of network policies","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"int"}
		LogicalRouter: data["logical_router"].(int),

		//{"Title":"","Description":"Maximum number of logical routers","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LogicalRouter","GoType":"int"}
		FloatingIPPool: data["floating_ip_pool"].(int),

		//{"Title":"","Description":"Maximum number of floating ip pools","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FloatingIPPool","GoType":"int"}
		VirtualMachineInterface: data["virtual_machine_interface"].(int),

		//{"Title":"","Description":"Maximum number of virtual machine interfaces","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualMachineInterface","GoType":"int"}
		VirtualIP: data["virtual_ip"].(int),

		//{"Title":"","Description":"Maximum number of virtual ips","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualIP","GoType":"int"}
		SecurityGroupRule: data["security_group_rule"].(int),

		//{"Title":"","Description":"Maximum number of security group rules","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroupRule","GoType":"int"}
		LoadbalancerMember: data["loadbalancer_member"].(int),

		//{"Title":"","Description":"Maximum number of loadbalancer member","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LoadbalancerMember","GoType":"int"}
		AccessControlList: data["access_control_list"].(int),

		//{"Title":"","Description":"Maximum number of access control lists","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AccessControlList","GoType":"int"}
		InstanceIP: data["instance_ip"].(int),

		//{"Title":"","Description":"Maximum number of instance ips","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"InstanceIP","GoType":"int"}
		GlobalVrouterConfig: data["global_vrouter_config"].(int),

		//{"Title":"","Description":"Maximum number of global vrouter configs","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"GlobalVrouterConfig","GoType":"int"}
		LoadbalancerHealthmonitor: data["loadbalancer_healthmonitor"].(int),

		//{"Title":"","Description":"Maximum number of loadbalancer health monitors","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LoadbalancerHealthmonitor","GoType":"int"}
		SecurityGroup: data["security_group"].(int),

		//{"Title":"","Description":"Maximum number of security groups","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"int"}
		RouteTable: data["route_table"].(int),

		//{"Title":"","Description":"Maximum number of route tables","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RouteTable","GoType":"int"}
		Subnet: data["subnet"].(int),

		//{"Title":"","Description":"Maximum number of subnets","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"int"}
		NetworkIpam: data["network_ipam"].(int),

		//{"Title":"","Description":"Maximum number of network IPAMs","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkIpam","GoType":"int"}
		VirtualDNSRecord: data["virtual_DNS_record"].(int),

		//{"Title":"","Description":"Maximum number of virtual DNS records","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualDNSRecord","GoType":"int"}
		ServiceInstance: data["service_instance"].(int),

		//{"Title":"","Description":"Maximum number of service instances","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServiceInstance","GoType":"int"}
		SecurityLoggingObject: data["security_logging_object"].(int),

		//{"Title":"","Description":"Maximum number of security logging objects","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityLoggingObject","GoType":"int"}
		VirtualNetwork: data["virtual_network"].(int),

		//{"Title":"","Description":"Maximum number of virtual networks","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"int"}
		LoadbalancerPool: data["loadbalancer_pool"].(int),

		//{"Title":"","Description":"Maximum number of loadbalancer pools","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LoadbalancerPool","GoType":"int"}
		VirtualDNS: data["virtual_DNS"].(int),

		//{"Title":"","Description":"Maximum number of virtual DNS servers","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualDNS","GoType":"int"}
		ServiceTemplate: data["service_template"].(int),

		//{"Title":"","Description":"Maximum number of service templates","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServiceTemplate","GoType":"int"}
		BGPRouter: data["bgp_router"].(int),

		//{"Title":"","Description":"Maximum number of bgp routers","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"BGPRouter","GoType":"int"}
		FloatingIP: data["floating_ip"].(int),

		//{"Title":"","Description":"Maximum number of floating ips","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FloatingIP","GoType":"int"}
		Defaults: data["defaults"].(int),

		//{"Title":"","Description":"Need to clarify","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Defaults","GoType":"int"}

	}
}

func InterfaceToQuotaTypeSlice(data interface{}) []*QuotaType {
	list := data.([]interface{})
	result := MakeQuotaTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToQuotaType(item))
	}
	return result
}

func MakeQuotaTypeSlice() []*QuotaType {
	return []*QuotaType{}
}
