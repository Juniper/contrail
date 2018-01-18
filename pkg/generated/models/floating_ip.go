package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	ParentUUID                   string               `json:"parent_uuid,omitempty"`
	DisplayName                  string               `json:"display_name,omitempty"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings,omitempty"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	UUID                         string               `json:"uuid,omitempty"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address,omitempty"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`
	ParentType                   string               `json:"parent_type,omitempty"`
	FQName                       []string             `json:"fq_name,omitempty"`
	Annotations                  *KeyValuePairs       `json:"annotations,omitempty"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family,omitempty"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address,omitempty"`
	IDPerms                      *IdPermsType         `json:"id_perms,omitempty"`
	Perms2                       *PermType2           `json:"perms2,omitempty"`

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
		FloatingIPPortMappings: MakePortMappings(),
		FloatingIPIsVirtualIP:  false,
		UUID:                         "",
		FloatingIPAddress:            MakeIpAddressType(),
		FloatingIPPortMappingsEnable: false,
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
		ParentType:                   "",
		FQName:                       []string{},
		Annotations:                  MakeKeyValuePairs(),
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		FloatingIPFixedIPAddress:     MakeIpAddressType(),
		IDPerms:                      MakeIdPermsType(),
		Perms2:                       MakePermType2(),
		ParentUUID:                   "",
		DisplayName:                  "",
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
