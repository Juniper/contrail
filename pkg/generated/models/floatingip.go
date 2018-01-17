package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address,omitempty"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family,omitempty"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address,omitempty"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable,omitempty"`
	DisplayName                  string               `json:"display_name,omitempty"`
	Annotations                  *KeyValuePairs       `json:"annotations,omitempty"`
	ParentType                   string               `json:"parent_type,omitempty"`
	FQName                       []string             `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType         `json:"id_perms,omitempty"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings,omitempty"`
	UUID                         string               `json:"uuid,omitempty"`
	ParentUUID                   string               `json:"parent_uuid,omitempty"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip,omitempty"`
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
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
		FloatingIPFixedIPAddress:     MakeIpAddressType(),
		FloatingIPAddress:            MakeIpAddressType(),
		FloatingIPPortMappingsEnable: false,
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		ParentType:                   "",
		FQName:                       []string{},
		IDPerms:                      MakeIdPermsType(),
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		UUID:                   "",
		ParentUUID:             "",
		FloatingIPPortMappings: MakePortMappings(),
		Perms2:                 MakePermType2(),
		FloatingIPIsVirtualIP:  false,
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
