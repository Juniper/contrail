package models

// FloatingIP

import "encoding/json"

// FloatingIP
type FloatingIP struct {
	FloatingIPPortMappingsEnable bool                 `json:"floating_ip_port_mappings_enable"`
	FloatingIPFixedIPAddress     IpAddressType        `json:"floating_ip_fixed_ip_address"`
	FloatingIPTrafficDirection   TrafficDirectionType `json:"floating_ip_traffic_direction"`
	FloatingIPAddressFamily      IpAddressFamilyType  `json:"floating_ip_address_family"`
	Annotations                  *KeyValuePairs       `json:"annotations"`
	ParentUUID                   string               `json:"parent_uuid"`
	IDPerms                      *IdPermsType         `json:"id_perms"`
	FloatingIPPortMappings       *PortMappings        `json:"floating_ip_port_mappings"`
	FloatingIPAddress            IpAddressType        `json:"floating_ip_address"`
	Perms2                       *PermType2           `json:"perms2"`
	UUID                         string               `json:"uuid"`
	FQName                       []string             `json:"fq_name"`
	FloatingIPIsVirtualIP        bool                 `json:"floating_ip_is_virtual_ip"`
	ParentType                   string               `json:"parent_type"`
	DisplayName                  string               `json:"display_name"`

	ProjectRefs                 []*FloatingIPProjectRef                 `json:"project_refs"`
	VirtualMachineInterfaceRefs []*FloatingIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
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
		ParentType:                   "",
		DisplayName:                  "",
		FloatingIPPortMappingsEnable: false,
		FloatingIPFixedIPAddress:     MakeIpAddressType(),
		FloatingIPTrafficDirection:   MakeTrafficDirectionType(),
		FloatingIPAddressFamily:      MakeIpAddressFamilyType(),
		Annotations:                  MakeKeyValuePairs(),
		ParentUUID:                   "",
		UUID:                         "",
		FQName:                       []string{},
		IDPerms:                      MakeIdPermsType(),
		FloatingIPPortMappings:       MakePortMappings(),
		FloatingIPAddress:            MakeIpAddressType(),
		Perms2:                       MakePermType2(),
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}
