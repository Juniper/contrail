package models

// FloatingIP

// FloatingIP
//proteus:generate
type FloatingIP struct {
	UUID                         string               `json:"uuid,omitempty"`
	ParentUUID                   string               `json:"parent_uuid,omitempty"`
	ParentType                   string               `json:"parent_type,omitempty"`
	FQName                       []string             `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName                  string               `json:"display_name,omitempty"`
	Annotations                  *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2                       *PermType2           `json:"perms2,omitempty"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family,omitempty"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings,omitempty"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address,omitempty"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address,omitempty"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`

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

// MakeFloatingIP makes FloatingIP
func MakeFloatingIP() *FloatingIP {
	return &FloatingIP{
		//TODO(nati): Apply default
		UUID:                         "",
		ParentUUID:                   "",
		ParentType:                   "",
		FQName:                       []string{},
		IDPerms:                      MakeIdPermsType(),
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		Perms2:                       MakePermType2(),
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		FloatingIPPortMappings:       MakePortMappings(),
		FloatingIPIsVirtualIP:        false,
		FloatingIPAddress:            MakeIpAddressType(),
		FloatingIPPortMappingsEnable: false,
		FloatingIPFixedIPAddress:     MakeIpAddressType(),
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
