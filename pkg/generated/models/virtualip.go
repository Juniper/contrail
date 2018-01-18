package models

// VirtualIP

import "encoding/json"

// VirtualIP
type VirtualIP struct {
	Annotations         *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID          string         `json:"parent_uuid,omitempty"`
	ParentType          string         `json:"parent_type,omitempty"`
	VirtualIPProperties *VirtualIpType `json:"virtual_ip_properties,omitempty"`
	FQName              []string       `json:"fq_name,omitempty"`
	IDPerms             *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
	Perms2              *PermType2     `json:"perms2,omitempty"`
	UUID                string         `json:"uuid,omitempty"`

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
		Perms2:              MakePermType2(),
		UUID:                "",
		ParentUUID:          "",
		ParentType:          "",
		VirtualIPProperties: MakeVirtualIpType(),
		FQName:              []string{},
		IDPerms:             MakeIdPermsType(),
		DisplayName:         "",
		Annotations:         MakeKeyValuePairs(),
	}
}

// MakeVirtualIPSlice() makes a slice of VirtualIP
func MakeVirtualIPSlice() []*VirtualIP {
	return []*VirtualIP{}
}
