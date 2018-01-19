package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address,omitempty"`
	FQName                       []string             `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType         `json:"id_perms,omitempty"`
	UUID                         string               `json:"uuid,omitempty"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings,omitempty"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	Annotations                  *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2                       *PermType2           `json:"perms2,omitempty"`
	DisplayName                  string               `json:"display_name,omitempty"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family,omitempty"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address,omitempty"`
	ParentUUID                   string               `json:"parent_uuid,omitempty"`
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
		FloatingIPPortMappings:       MakePortMappings(),
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
		FloatingIPPortMappingsEnable: false,
		Annotations:                  MakeKeyValuePairs(),
		Perms2:                       MakePermType2(),
		DisplayName:                  "",
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		FloatingIPIsVirtualIP:        false,
		FloatingIPAddress:            MakeIpAddressType(),
		ParentUUID:                   "",
		ParentType:                   "",
		FloatingIPFixedIPAddress:     MakeIpAddressType(),
		FQName:  []string{},
		IDPerms: MakeIdPermsType(),
		UUID:    "",
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
