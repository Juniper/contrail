package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family"`
	UUID                         string               `json:"uuid"`
	FQName                       []string             `json:"fq_name"`
	IDPerms                      *IdPermsType         `json:"id_perms"`
	DisplayName                  string               `json:"display_name"`
	ParentUUID                   string               `json:"parent_uuid"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address"`
	Annotations                  *KeyValuePairs       `json:"annotations"`
	Perms2                       *PermType2           `json:"perms2"`
	ParentType                   string               `json:"parent_type"`

	ProjectRefs                 []*FloatingIPProjectRef                 `json:"project_refs"`
	VirtualMachineInterfaceRefs []*FloatingIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
}

// FloatingIPProjectRef references each other
type FloatingIPProjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// FloatingIPVirtualMachineInterfaceRef references each other
type FloatingIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *FloatingIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFloatingIP makes FloatingIP
func MakeFloatingIP() *FloatingIP {
	return &FloatingIP{
		//TODO(nati): Apply default
		FloatingIPPortMappings:       MakePortMappings(),
		FloatingIPPortMappingsEnable: false,
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
		IDPerms:                      MakeIdPermsType(),
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		UUID:                     "",
		FQName:                   []string{},
		DisplayName:              "",
		ParentUUID:               "",
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		ParentType:               "",
		FloatingIPIsVirtualIP:    false,
		FloatingIPAddress:        MakeIpAddressType(),
		FloatingIPFixedIPAddress: MakeIpAddressType(),
	}
}

// InterfaceToFloatingIP makes FloatingIP from interface
func InterfaceToFloatingIP(iData interface{}) *FloatingIP {
	data := iData.(map[string]interface{})
	return &FloatingIP{
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FloatingIPIsVirtualIP: data["floating_ip_is_virtual_ip"].(bool),

		//{"description":"This floating ip is used as virtual ip (VIP) in case of LBaaS.","type":"boolean"}
		FloatingIPAddress: InterfaceToIpAddressType(data["floating_ip_address"]),

		//{"description":"Floating ip address.","type":"string"}
		FloatingIPFixedIPAddress: InterfaceToIpAddressType(data["floating_ip_fixed_ip_address"]),

		//{"description":"This floating is tracking given fixed ip of the interface. The given fixed ip is used in 1:1 NAT.","type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FloatingIPPortMappings: InterfaceToPortMappings(data["floating_ip_port_mappings"]),

		//{"description":"List of PortMaps for this floating-ip.","type":"object","properties":{"port_mappings":{"type":"array","item":{"type":"object","properties":{"dst_port":{"type":"integer"},"protocol":{"type":"string"},"src_port":{"type":"integer"}}}}}}
		FloatingIPPortMappingsEnable: data["floating_ip_port_mappings_enable"].(bool),

		//{"description":"If it is false, floating-ip Nat is done for all Ports. If it is true, floating-ip Nat is done to the list of PortMaps.","default":false,"type":"boolean"}
		FloatingIPTrafficDirection: InterfaceToTrafficDirectionType(data["floating_ip_traffic_direction"]),

		//{"description":"Specifies direction of traffic for the floating-ip","default":"both","type":"string","enum":["ingress","egress","both"]}
		FloatingIPAddressFamily: InterfaceToIpAddressFamilyType(data["floating_ip_address_family"]),

		//{"description":"Ip address family of the floating ip, IpV4 or IpV6","type":"string","enum":["v4","v6"]}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToFloatingIPSlice makes a slice of FloatingIP from interface
func InterfaceToFloatingIPSlice(data interface{}) []*FloatingIP {
	list := data.([]interface{})
	result := MakeFloatingIPSlice()
	for _, item := range list {
		result = append(result, InterfaceToFloatingIP(item))
	}
	return result
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
