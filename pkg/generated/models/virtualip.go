package models

// VirtualIP

import "encoding/json"

// VirtualIP
type VirtualIP struct {
	Perms2              *PermType2     `json:"perms2"`
	ParentType          string         `json:"parent_type"`
	DisplayName         string         `json:"display_name"`
	IDPerms             *IdPermsType   `json:"id_perms"`
	Annotations         *KeyValuePairs `json:"annotations"`
	VirtualIPProperties *VirtualIpType `json:"virtual_ip_properties"`
	UUID                string         `json:"uuid"`
	ParentUUID          string         `json:"parent_uuid"`
	FQName              []string       `json:"fq_name"`

	LoadbalancerPoolRefs        []*VirtualIPLoadbalancerPoolRef        `json:"loadbalancer_pool_refs"`
	VirtualMachineInterfaceRefs []*VirtualIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
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
		ParentType:          "",
		DisplayName:         "",
		FQName:              []string{},
		IDPerms:             MakeIdPermsType(),
		Annotations:         MakeKeyValuePairs(),
		VirtualIPProperties: MakeVirtualIpType(),
		UUID:                "",
		ParentUUID:          "",
	}
}

// MakeVirtualIPSlice() makes a slice of VirtualIP
func MakeVirtualIPSlice() []*VirtualIP {
	return []*VirtualIP{}
}
