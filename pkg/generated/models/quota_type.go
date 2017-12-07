package models

// QuotaType

import "encoding/json"

// QuotaType
type QuotaType struct {
	Subnet                    int `json:"subnet"`
	VirtualDNSRecord          int `json:"virtual_DNS_record"`
	SecurityGroupRule         int `json:"security_group_rule"`
	VirtualMachineInterface   int `json:"virtual_machine_interface"`
	VirtualIP                 int `json:"virtual_ip"`
	LoadbalancerPool          int `json:"loadbalancer_pool"`
	ServiceTemplate           int `json:"service_template"`
	GlobalVrouterConfig       int `json:"global_vrouter_config"`
	SecurityGroup             int `json:"security_group"`
	AccessControlList         int `json:"access_control_list"`
	SecurityLoggingObject     int `json:"security_logging_object"`
	Defaults                  int `json:"defaults"`
	NetworkPolicy             int `json:"network_policy"`
	RouteTable                int `json:"route_table"`
	NetworkIpam               int `json:"network_ipam"`
	FloatingIPPool            int `json:"floating_ip_pool"`
	LoadbalancerMember        int `json:"loadbalancer_member"`
	FloatingIP                int `json:"floating_ip"`
	InstanceIP                int `json:"instance_ip"`
	LoadbalancerHealthmonitor int `json:"loadbalancer_healthmonitor"`
	VirtualRouter             int `json:"virtual_router"`
	LogicalRouter             int `json:"logical_router"`
	VirtualDNS                int `json:"virtual_DNS"`
	ServiceInstance           int `json:"service_instance"`
	BGPRouter                 int `json:"bgp_router"`
	VirtualNetwork            int `json:"virtual_network"`
}

//  parents relation object

// String returns json representation of the object
func (model *QuotaType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeQuotaType makes QuotaType
func MakeQuotaType() *QuotaType {
	return &QuotaType{
		//TODO(nati): Apply default
		FloatingIPPool:            0,
		LoadbalancerMember:        0,
		AccessControlList:         0,
		SecurityLoggingObject:     0,
		Defaults:                  0,
		NetworkPolicy:             0,
		RouteTable:                0,
		NetworkIpam:               0,
		ServiceInstance:           0,
		BGPRouter:                 0,
		FloatingIP:                0,
		InstanceIP:                0,
		LoadbalancerHealthmonitor: 0,
		VirtualRouter:             0,
		LogicalRouter:             0,
		VirtualDNS:                0,
		VirtualNetwork:            0,
		VirtualMachineInterface:   0,
		VirtualIP:                 0,
		Subnet:                    0,
		VirtualDNSRecord:          0,
		SecurityGroupRule:         0,
		SecurityGroup:             0,
		LoadbalancerPool:          0,
		ServiceTemplate:           0,
		GlobalVrouterConfig:       0,
	}
}

// InterfaceToQuotaType makes QuotaType from interface
func InterfaceToQuotaType(iData interface{}) *QuotaType {
	data := iData.(map[string]interface{})
	return &QuotaType{
		LoadbalancerPool: data["loadbalancer_pool"].(int),

		//{"Title":"","Description":"Maximum number of loadbalancer pools","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LoadbalancerPool","GoType":"int","GoPremitive":true}
		ServiceTemplate: data["service_template"].(int),

		//{"Title":"","Description":"Maximum number of service templates","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServiceTemplate","GoType":"int","GoPremitive":true}
		GlobalVrouterConfig: data["global_vrouter_config"].(int),

		//{"Title":"","Description":"Maximum number of global vrouter configs","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"GlobalVrouterConfig","GoType":"int","GoPremitive":true}
		SecurityGroup: data["security_group"].(int),

		//{"Title":"","Description":"Maximum number of security groups","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"int","GoPremitive":true}
		AccessControlList: data["access_control_list"].(int),

		//{"Title":"","Description":"Maximum number of access control lists","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AccessControlList","GoType":"int","GoPremitive":true}
		SecurityLoggingObject: data["security_logging_object"].(int),

		//{"Title":"","Description":"Maximum number of security logging objects","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityLoggingObject","GoType":"int","GoPremitive":true}
		Defaults: data["defaults"].(int),

		//{"Title":"","Description":"Need to clarify","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Defaults","GoType":"int","GoPremitive":true}
		NetworkPolicy: data["network_policy"].(int),

		//{"Title":"","Description":"Maximum number of network policies","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"int","GoPremitive":true}
		RouteTable: data["route_table"].(int),

		//{"Title":"","Description":"Maximum number of route tables","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RouteTable","GoType":"int","GoPremitive":true}
		NetworkIpam: data["network_ipam"].(int),

		//{"Title":"","Description":"Maximum number of network IPAMs","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkIpam","GoType":"int","GoPremitive":true}
		FloatingIPPool: data["floating_ip_pool"].(int),

		//{"Title":"","Description":"Maximum number of floating ip pools","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FloatingIPPool","GoType":"int","GoPremitive":true}
		LoadbalancerMember: data["loadbalancer_member"].(int),

		//{"Title":"","Description":"Maximum number of loadbalancer member","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LoadbalancerMember","GoType":"int","GoPremitive":true}
		FloatingIP: data["floating_ip"].(int),

		//{"Title":"","Description":"Maximum number of floating ips","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FloatingIP","GoType":"int","GoPremitive":true}
		InstanceIP: data["instance_ip"].(int),

		//{"Title":"","Description":"Maximum number of instance ips","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"InstanceIP","GoType":"int","GoPremitive":true}
		LoadbalancerHealthmonitor: data["loadbalancer_healthmonitor"].(int),

		//{"Title":"","Description":"Maximum number of loadbalancer health monitors","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LoadbalancerHealthmonitor","GoType":"int","GoPremitive":true}
		VirtualRouter: data["virtual_router"].(int),

		//{"Title":"","Description":"Maximum number of logical routers","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualRouter","GoType":"int","GoPremitive":true}
		LogicalRouter: data["logical_router"].(int),

		//{"Title":"","Description":"Maximum number of logical routers","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LogicalRouter","GoType":"int","GoPremitive":true}
		VirtualDNS: data["virtual_DNS"].(int),

		//{"Title":"","Description":"Maximum number of virtual DNS servers","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualDNS","GoType":"int","GoPremitive":true}
		ServiceInstance: data["service_instance"].(int),

		//{"Title":"","Description":"Maximum number of service instances","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServiceInstance","GoType":"int","GoPremitive":true}
		BGPRouter: data["bgp_router"].(int),

		//{"Title":"","Description":"Maximum number of bgp routers","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"BGPRouter","GoType":"int","GoPremitive":true}
		VirtualNetwork: data["virtual_network"].(int),

		//{"Title":"","Description":"Maximum number of virtual networks","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"int","GoPremitive":true}
		Subnet: data["subnet"].(int),

		//{"Title":"","Description":"Maximum number of subnets","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"int","GoPremitive":true}
		VirtualDNSRecord: data["virtual_DNS_record"].(int),

		//{"Title":"","Description":"Maximum number of virtual DNS records","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualDNSRecord","GoType":"int","GoPremitive":true}
		SecurityGroupRule: data["security_group_rule"].(int),

		//{"Title":"","Description":"Maximum number of security group rules","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroupRule","GoType":"int","GoPremitive":true}
		VirtualMachineInterface: data["virtual_machine_interface"].(int),

		//{"Title":"","Description":"Maximum number of virtual machine interfaces","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualMachineInterface","GoType":"int","GoPremitive":true}
		VirtualIP: data["virtual_ip"].(int),

		//{"Title":"","Description":"Maximum number of virtual ips","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualIP","GoType":"int","GoPremitive":true}

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
