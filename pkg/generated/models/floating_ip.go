package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	UUID                         string               `json:"uuid,omitempty"`
	ParentType                   string               `json:"parent_type,omitempty"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family,omitempty"`
	Perms2                       *PermType2           `json:"perms2,omitempty"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address,omitempty"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`
	ParentUUID                   string               `json:"parent_uuid,omitempty"`
	FQName                       []string             `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType         `json:"id_perms,omitempty"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings,omitempty"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	Annotations                  *KeyValuePairs       `json:"annotations,omitempty"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address,omitempty"`
	DisplayName                  string               `json:"display_name,omitempty"`

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
		Perms2:                       MakePermType2(),
		UUID:                         "",
		ParentType:                   "",
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		FloatingIPIsVirtualIP:        false,
		FloatingIPAddress:            MakeIpAddressType(),
		FloatingIPPortMappingsEnable: false,
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
		ParentUUID:                   "",
		FQName:                       []string{},
		IDPerms:                      MakeIdPermsType(),
		FloatingIPPortMappings:       MakePortMappings(),
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		FloatingIPFixedIPAddress:     MakeIpAddressType(),
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
