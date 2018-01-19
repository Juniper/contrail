package models

// VirtualIP

import "encoding/json"

// VirtualIP
type VirtualIP struct {
	VirtualIPProperties *VirtualIpType `json:"virtual_ip_properties,omitempty"`
	ParentType          string         `json:"parent_type,omitempty"`
	IDPerms             *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
	UUID                string         `json:"uuid,omitempty"`
	FQName              []string       `json:"fq_name,omitempty"`
	Annotations         *KeyValuePairs `json:"annotations,omitempty"`
	Perms2              *PermType2     `json:"perms2,omitempty"`
	ParentUUID          string         `json:"parent_uuid,omitempty"`

	LoadbalancerPoolRefs        []*VirtualIPLoadbalancerPoolRef        `json:"loadbalancer_pool_refs,omitempty"`
	VirtualMachineInterfaceRefs []*VirtualIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
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
		FQName:              []string{},
		Annotations:         MakeKeyValuePairs(),
		Perms2:              MakePermType2(),
		ParentUUID:          "",
		VirtualIPProperties: MakeVirtualIpType(),
		ParentType:          "",
		IDPerms:             MakeIdPermsType(),
		DisplayName:         "",
		UUID:                "",
	}
}

// MakeVirtualIPSlice() makes a slice of VirtualIP
func MakeVirtualIPSlice() []*VirtualIP {
	return []*VirtualIP{}
}
