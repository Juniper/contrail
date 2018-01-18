package models

// VirtualRouter

import "encoding/json"

// VirtualRouter
type VirtualRouter struct {
	UUID                     string            `json:"uuid,omitempty"`
	ParentType               string            `json:"parent_type,omitempty"`
	VirtualRouterDPDKEnabled bool              `json:"virtual_router_dpdk_enabled"`
	VirtualRouterIPAddress   IpAddressType     `json:"virtual_router_ip_address,omitempty"`
	DisplayName              string            `json:"display_name,omitempty"`
	Annotations              *KeyValuePairs    `json:"annotations,omitempty"`
	Perms2                   *PermType2        `json:"perms2,omitempty"`
	VirtualRouterType        VirtualRouterType `json:"virtual_router_type,omitempty"`
	IDPerms                  *IdPermsType      `json:"id_perms,omitempty"`
	ParentUUID               string            `json:"parent_uuid,omitempty"`
	FQName                   []string          `json:"fq_name,omitempty"`

	NetworkIpamRefs    []*VirtualRouterNetworkIpamRef    `json:"network_ipam_refs,omitempty"`
	VirtualMachineRefs []*VirtualRouterVirtualMachineRef `json:"virtual_machine_refs,omitempty"`

	VirtualMachineInterfaces []*VirtualMachineInterface `json:"virtual_machine_interfaces,omitempty"`
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
		VirtualRouterIPAddress:   MakeIpAddressType(),
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		UUID:                     "",
		ParentType:               "",
		VirtualRouterDPDKEnabled: false,
		IDPerms:                  MakeIdPermsType(),
		ParentUUID:               "",
		FQName:                   []string{},
		VirtualRouterType:        MakeVirtualRouterType(),
	}
}

// MakeVirtualRouterSlice() makes a slice of VirtualRouter
func MakeVirtualRouterSlice() []*VirtualRouter {
	return []*VirtualRouter{}
}
