package models

// VirtualRouter

import "encoding/json"

// VirtualRouter
type VirtualRouter struct {
	VirtualRouterIPAddress   IpAddressType     `json:"virtual_router_ip_address"`
	ParentUUID               string            `json:"parent_uuid"`
	ParentType               string            `json:"parent_type"`
	DisplayName              string            `json:"display_name"`
	Annotations              *KeyValuePairs    `json:"annotations"`
	VirtualRouterDPDKEnabled bool              `json:"virtual_router_dpdk_enabled"`
	FQName                   []string          `json:"fq_name"`
	IDPerms                  *IdPermsType      `json:"id_perms"`
	Perms2                   *PermType2        `json:"perms2"`
	UUID                     string            `json:"uuid"`
	VirtualRouterType        VirtualRouterType `json:"virtual_router_type"`

	NetworkIpamRefs    []*VirtualRouterNetworkIpamRef    `json:"network_ipam_refs"`
	VirtualMachineRefs []*VirtualRouterVirtualMachineRef `json:"virtual_machine_refs"`

	VirtualMachineInterfaces []*VirtualMachineInterface `json:"virtual_machine_interfaces"`
}

// VirtualRouterVirtualMachineRef references each other
type VirtualRouterVirtualMachineRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualRouterNetworkIpamRef references each other
type VirtualRouterNetworkIpamRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *VirtualRouterNetworkIpamType
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
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		VirtualRouterDPDKEnabled: false,
		VirtualRouterIPAddress:   MakeIpAddressType(),
		ParentUUID:               "",
		ParentType:               "",
		UUID:                     "",
		VirtualRouterType:        MakeVirtualRouterType(),
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		Perms2:                   MakePermType2(),
	}
}

// InterfaceToVirtualRouter makes VirtualRouter from interface
func InterfaceToVirtualRouter(iData interface{}) *VirtualRouter {
	data := iData.(map[string]interface{})
	return &VirtualRouter{
		VirtualRouterType: InterfaceToVirtualRouterType(data["virtual_router_type"]),

		//{"description":"Different types of the vrouters in the system.","type":"string","enum":["embedded","tor-agent","tor-service-node"]}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		VirtualRouterDPDKEnabled: data["virtual_router_dpdk_enabled"].(bool),

		//{"description":"This vrouter's data path is using DPDK library, Virtual machines interfaces scheduled on this compute node will be tagged with additional flags so that they are spawned with user space virtio driver. It is only applicable for embedded vrouter.","type":"boolean"}
		VirtualRouterIPAddress: InterfaceToIpAddressType(data["virtual_router_ip_address"]),

		//{"description":"Ip address of the virtual router.","type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}

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
