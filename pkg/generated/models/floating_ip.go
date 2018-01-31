package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address,omitempty"`
	IDPerms                      *IdPermsType         `json:"id_perms,omitempty"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family,omitempty"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address,omitempty"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`
	UUID                         string               `json:"uuid,omitempty"`
	FQName                       []string             `json:"fq_name,omitempty"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings,omitempty"`
	ParentType                   string               `json:"parent_type,omitempty"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	DisplayName                  string               `json:"display_name,omitempty"`
	Annotations                  *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2                       *PermType2           `json:"perms2,omitempty"`
	ParentUUID                   string               `json:"parent_uuid,omitempty"`

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
		FloatingIPIsVirtualIP:        false,
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		Perms2:                       MakePermType2(),
		ParentUUID:                   "",
		FloatingIPAddress:            MakeIpAddressType(),
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		FloatingIPPortMappingsEnable: false,
		FloatingIPFixedIPAddress:     MakeIpAddressType(),
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
		UUID:                   "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		FloatingIPPortMappings: MakePortMappings(),
		ParentType:             "",
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
