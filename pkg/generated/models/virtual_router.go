package models

// VirtualRouter

import "encoding/json"

// VirtualRouter
type VirtualRouter struct {
	VirtualRouterType        VirtualRouterType `json:"virtual_router_type"`
	VirtualRouterIPAddress   IpAddressType     `json:"virtual_router_ip_address"`
	IDPerms                  *IdPermsType      `json:"id_perms"`
	UUID                     string            `json:"uuid"`
	VirtualRouterDPDKEnabled bool              `json:"virtual_router_dpdk_enabled"`
	FQName                   []string          `json:"fq_name"`
	DisplayName              string            `json:"display_name"`
	Annotations              *KeyValuePairs    `json:"annotations"`
	Perms2                   *PermType2        `json:"perms2"`
	ParentUUID               string            `json:"parent_uuid"`
	ParentType               string            `json:"parent_type"`

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
		VirtualRouterDPDKEnabled: false,
		FQName:                 []string{},
		DisplayName:            "",
		Annotations:            MakeKeyValuePairs(),
		Perms2:                 MakePermType2(),
		ParentUUID:             "",
		ParentType:             "",
		VirtualRouterType:      MakeVirtualRouterType(),
		VirtualRouterIPAddress: MakeIpAddressType(),
		IDPerms:                MakeIdPermsType(),
		UUID:                   "",
	}
}

// MakeVirtualRouterSlice() makes a slice of VirtualRouter
func MakeVirtualRouterSlice() []*VirtualRouter {
	return []*VirtualRouter{}
}
