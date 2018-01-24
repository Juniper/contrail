package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family,omitempty"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings,omitempty"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	Perms2                       *PermType2           `json:"perms2,omitempty"`
	DisplayName                  string               `json:"display_name,omitempty"`
	Annotations                  *KeyValuePairs       `json:"annotations,omitempty"`
	ParentUUID                   string               `json:"parent_uuid,omitempty"`
	FQName                       []string             `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType         `json:"id_perms,omitempty"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address,omitempty"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address,omitempty"`
	UUID                         string               `json:"uuid,omitempty"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`
	ParentType                   string               `json:"parent_type,omitempty"`

	ProjectRefs                 []*FloatingIPProjectRef                 `json:"project_refs,omitempty"`
	VirtualMachineInterfaceRefs []*FloatingIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
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
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		FloatingIPIsVirtualIP:    false,
		FloatingIPAddress:        MakeIpAddressType(),
		FloatingIPFixedIPAddress: MakeIpAddressType(),
		UUID:                         "",
		ParentUUID:                   "",
		FQName:                       []string{},
		IDPerms:                      MakeIdPermsType(),
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
		ParentType:                   "",
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		FloatingIPPortMappings:       MakePortMappings(),
		FloatingIPPortMappingsEnable: false,
		Perms2: MakePermType2(),
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
