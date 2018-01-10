package models

// VirtualRouter

import "encoding/json"

// VirtualRouter
type VirtualRouter struct {
	VirtualRouterDPDKEnabled bool              `json:"virtual_router_dpdk_enabled"`
	ParentType               string            `json:"parent_type"`
	Annotations              *KeyValuePairs    `json:"annotations"`
	UUID                     string            `json:"uuid"`
	Perms2                   *PermType2        `json:"perms2"`
	VirtualRouterType        VirtualRouterType `json:"virtual_router_type"`
	VirtualRouterIPAddress   IpAddressType     `json:"virtual_router_ip_address"`
	ParentUUID               string            `json:"parent_uuid"`
	FQName                   []string          `json:"fq_name"`
	IDPerms                  *IdPermsType      `json:"id_perms"`
	DisplayName              string            `json:"display_name"`

	NetworkIpamRefs    []*VirtualRouterNetworkIpamRef    `json:"network_ipam_refs"`
	VirtualMachineRefs []*VirtualRouterVirtualMachineRef `json:"virtual_machine_refs"`

	VirtualMachineInterfaces []*VirtualMachineInterface `json:"virtual_machine_interfaces"`
}

// VirtualRouterNetworkIpamRef references each other
type VirtualRouterNetworkIpamRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *VirtualRouterNetworkIpamType
}

// VirtualRouterVirtualMachineRef references each other
type VirtualRouterVirtualMachineRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *VirtualRouter) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualRouter makes VirtualRouter
func MakeVirtualRouter() *VirtualRouter {
	return &VirtualRouter{
		//TODO(nati): Apply default
		VirtualRouterType:      MakeVirtualRouterType(),
		VirtualRouterIPAddress: MakeIpAddressType(),
		ParentUUID:             "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Perms2:                 MakePermType2(),
		VirtualRouterDPDKEnabled: false,
		ParentType:               "",
		Annotations:              MakeKeyValuePairs(),
		UUID:                     "",
	}
}

// InterfaceToVirtualRouter makes VirtualRouter from interface
func InterfaceToVirtualRouter(iData interface{}) *VirtualRouter {
	data := iData.(map[string]interface{})
	return &VirtualRouter{
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		VirtualRouterDPDKEnabled: data["virtual_router_dpdk_enabled"].(bool),

		//{"description":"This vrouter's data path is using DPDK library, Virtual machines interfaces scheduled on this compute node will be tagged with additional flags so that they are spawned with user space virtio driver. It is only applicable for embedded vrouter.","type":"boolean"}
		VirtualRouterIPAddress: InterfaceToIpAddressType(data["virtual_router_ip_address"]),

		//{"description":"Ip address of the virtual router.","type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		VirtualRouterType: InterfaceToVirtualRouterType(data["virtual_router_type"]),

		//{"description":"Different types of the vrouters in the system.","type":"string","enum":["embedded","tor-agent","tor-service-node"]}

	}
}

// InterfaceToVirtualRouterSlice makes a slice of VirtualRouter from interface
func InterfaceToVirtualRouterSlice(data interface{}) []*VirtualRouter {
	list := data.([]interface{})
	result := MakeVirtualRouterSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualRouter(item))
	}
	return result
}

// MakeVirtualRouterSlice() makes a slice of VirtualRouter
func MakeVirtualRouterSlice() []*VirtualRouter {
	return []*VirtualRouter{}
}
