package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings,omitempty"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	UUID                         string               `json:"uuid,omitempty"`
	ParentType                   string               `json:"parent_type,omitempty"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address,omitempty"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address,omitempty"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`
	Perms2                       *PermType2           `json:"perms2,omitempty"`
	ParentUUID                   string               `json:"parent_uuid,omitempty"`
	DisplayName                  string               `json:"display_name,omitempty"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family,omitempty"`
	FQName                       []string             `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType         `json:"id_perms,omitempty"`
	Annotations                  *KeyValuePairs       `json:"annotations,omitempty"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`

	ProjectRefs                 []*FloatingIPProjectRef                 `json:"project_refs,omitempty"`
	VirtualMachineInterfaceRefs []*FloatingIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
}

// FloatingIPVirtualMachineInterfaceRef references each other
type FloatingIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// FloatingIPProjectRef references each other
type FloatingIPProjectRef struct {
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
		FloatingIPPortMappingsEnable: false,
		UUID:                       "",
		ParentType:                 "",
		FloatingIPPortMappings:     MakePortMappings(),
		FloatingIPFixedIPAddress:   MakeIpAddressType(),
		FloatingIPTrafficDirection: MakeTrafficDirectionType(),
		Perms2:                  MakePermType2(),
		ParentUUID:              "",
		DisplayName:             "",
		FloatingIPAddress:       MakeIpAddressType(),
		FQName:                  []string{},
		IDPerms:                 MakeIdPermsType(),
		Annotations:             MakeKeyValuePairs(),
		FloatingIPAddressFamily: MakeIpAddressFamilyType(),
		FloatingIPIsVirtualIP:   false,
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
