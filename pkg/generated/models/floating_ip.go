package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address,omitempty"`
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	ParentType                   string               `json:"parent_type,omitempty"`
	FQName                       []string             `json:"fq_name,omitempty"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`
	UUID                         string               `json:"uuid,omitempty"`
	ParentUUID                   string               `json:"parent_uuid,omitempty"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family,omitempty"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address,omitempty"`
	IDPerms                      *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName                  string               `json:"display_name,omitempty"`
	Annotations                  *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2                       *PermType2           `json:"perms2,omitempty"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings,omitempty"`

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
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		FloatingIPFixedIPAddress:     MakeIpAddressType(),
		IDPerms:                      MakeIdPermsType(),
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		Perms2:                       MakePermType2(),
		FloatingIPPortMappings:       MakePortMappings(),
		FloatingIPIsVirtualIP:        false,
		FloatingIPAddress:            MakeIpAddressType(),
		FloatingIPPortMappingsEnable: false,
		ParentType:                   "",
		FQName:                       []string{},
		FloatingIPTrafficDirection: MakeTrafficDirectionType(),
		UUID:       "",
		ParentUUID: "",
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
