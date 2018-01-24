package models

// VirtualIP

import "encoding/json"

// VirtualIP
type VirtualIP struct {
	ParentType          string         `json:"parent_type,omitempty"`
	IDPerms             *IdPermsType   `json:"id_perms,omitempty"`
	Perms2              *PermType2     `json:"perms2,omitempty"`
	UUID                string         `json:"uuid,omitempty"`
	ParentUUID          string         `json:"parent_uuid,omitempty"`
	VirtualIPProperties *VirtualIpType `json:"virtual_ip_properties,omitempty"`
	FQName              []string       `json:"fq_name,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
	Annotations         *KeyValuePairs `json:"annotations,omitempty"`

	VirtualMachineInterfaceRefs []*VirtualIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	LoadbalancerPoolRefs        []*VirtualIPLoadbalancerPoolRef        `json:"loadbalancer_pool_refs,omitempty"`
}

// VirtualIPLoadbalancerPoolRef references each other
type VirtualIPLoadbalancerPoolRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualIPVirtualMachineInterfaceRef references each other
type VirtualIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *VirtualIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualIP makes VirtualIP
func MakeVirtualIP() *VirtualIP {
	return &VirtualIP{
		//TODO(nati): Apply default
		Perms2:              MakePermType2(),
		UUID:                "",
		ParentUUID:          "",
		VirtualIPProperties: MakeVirtualIpType(),
		FQName:              []string{},
		DisplayName:         "",
		Annotations:         MakeKeyValuePairs(),
		ParentType:          "",
		IDPerms:             MakeIdPermsType(),
	}
}

// MakeVirtualIPSlice() makes a slice of VirtualIP
func MakeVirtualIPSlice() []*VirtualIP {
	return []*VirtualIP{}
}
