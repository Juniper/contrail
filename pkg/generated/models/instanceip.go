package models

// InstanceIP

import "encoding/json"

// InstanceIP
type InstanceIP struct {
	InstanceIPFamily      IpAddressFamilyType `json:"instance_ip_family"`
	DisplayName           string              `json:"display_name"`
	ServiceInstanceIP     bool                `json:"service_instance_ip"`
	InstanceIPSecondary   bool                `json:"instance_ip_secondary"`
	InstanceIPLocalIP     bool                `json:"instance_ip_local_ip"`
	IDPerms               *IdPermsType        `json:"id_perms"`
	UUID                  string              `json:"uuid"`
	ParentType            string              `json:"parent_type"`
	InstanceIPAddress     IpAddressType       `json:"instance_ip_address"`
	InstanceIPMode        AddressMode         `json:"instance_ip_mode"`
	SubnetUUID            string              `json:"subnet_uuid"`
	Annotations           *KeyValuePairs      `json:"annotations"`
	Perms2                *PermType2          `json:"perms2"`
	ParentUUID            string              `json:"parent_uuid"`
	FQName                []string            `json:"fq_name"`
	ServiceHealthCheckIP  bool                `json:"service_health_check_ip"`
	SecondaryIPTrackingIP *SubnetType         `json:"secondary_ip_tracking_ip"`

	NetworkIpamRefs             []*InstanceIPNetworkIpamRef             `json:"network_ipam_refs"`
	VirtualNetworkRefs          []*InstanceIPVirtualNetworkRef          `json:"virtual_network_refs"`
	VirtualMachineInterfaceRefs []*InstanceIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
	PhysicalRouterRefs          []*InstanceIPPhysicalRouterRef          `json:"physical_router_refs"`
	VirtualRouterRefs           []*InstanceIPVirtualRouterRef           `json:"virtual_router_refs"`

	FloatingIPs []*FloatingIP `json:"floating_ips"`
}

// InstanceIPVirtualMachineInterfaceRef references each other
type InstanceIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// InstanceIPPhysicalRouterRef references each other
type InstanceIPPhysicalRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// InstanceIPVirtualRouterRef references each other
type InstanceIPVirtualRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// InstanceIPNetworkIpamRef references each other
type InstanceIPNetworkIpamRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// InstanceIPVirtualNetworkRef references each other
type InstanceIPVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *InstanceIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeInstanceIP makes InstanceIP
func MakeInstanceIP() *InstanceIP {
	return &InstanceIP{
		//TODO(nati): Apply default
		ServiceInstanceIP:     false,
		InstanceIPSecondary:   false,
		InstanceIPAddress:     MakeIpAddressType(),
		InstanceIPMode:        MakeAddressMode(),
		InstanceIPLocalIP:     false,
		IDPerms:               MakeIdPermsType(),
		UUID:                  "",
		ParentType:            "",
		ServiceHealthCheckIP:  false,
		SecondaryIPTrackingIP: MakeSubnetType(),
		SubnetUUID:            "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		ParentUUID:            "",
		FQName:                []string{},
		InstanceIPFamily:      MakeIpAddressFamilyType(),
		DisplayName:           "",
	}
}

// InterfaceToInstanceIP makes InstanceIP from interface
func InterfaceToInstanceIP(iData interface{}) *InstanceIP {
	data := iData.(map[string]interface{})
	return &InstanceIP{
		ServiceInstanceIP: data["service_instance_ip"].(bool),

		//{"description":"This instance ip is used as service chain next hop","default":false,"type":"boolean"}
		InstanceIPSecondary: data["instance_ip_secondary"].(bool),

		//{"description":"This instance ip is secondary ip of the interface.","default":false,"type":"boolean"}
		InstanceIPAddress: InterfaceToIpAddressType(data["instance_ip_address"]),

		//{"description":"Ip address value for instance ip.","type":"string"}
		InstanceIPMode: InterfaceToAddressMode(data["instance_ip_mode"]),

		//{"description":"Ip address HA mode in case this instance ip is used in more than one interface, active-Active or active-Standby.","type":"string","enum":["active-active","active-standby"]}
		InstanceIPLocalIP: data["instance_ip_local_ip"].(bool),

		//{"description":"This instance ip is local to compute and will not be exported to other nodes.","default":false,"type":"boolean"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		ServiceHealthCheckIP: data["service_health_check_ip"].(bool),

		//{"description":"This instance ip is used as service health check source ip","default":false,"type":"boolean"}
		SecondaryIPTrackingIP: InterfaceToSubnetType(data["secondary_ip_tracking_ip"]),

		//{"description":"When this instance ip is secondary ip, it can track activeness of another ip.","type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}
		SubnetUUID: data["subnet_uuid"].(string),

		//{"description":"This instance ip was allocated from this Subnet(UUID).","type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		InstanceIPFamily: InterfaceToIpAddressFamilyType(data["instance_ip_family"]),

		//{"description":"Ip address family for instance ip, IPv4(v4) or IPv6(v6).","type":"string","enum":["v4","v6"]}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToInstanceIPSlice makes a slice of InstanceIP from interface
func InterfaceToInstanceIPSlice(data interface{}) []*InstanceIP {
	list := data.([]interface{})
	result := MakeInstanceIPSlice()
	for _, item := range list {
		result = append(result, InterfaceToInstanceIP(item))
	}
	return result
}

// MakeInstanceIPSlice() makes a slice of InstanceIP
func MakeInstanceIPSlice() []*InstanceIP {
	return []*InstanceIP{}
}
