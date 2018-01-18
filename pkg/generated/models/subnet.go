package models

// Subnet

import "encoding/json"

// Subnet
type Subnet struct {
	SubnetIPPrefix *SubnetType    `json:"subnet_ip_prefix,omitempty"`
	FQName         []string       `json:"fq_name,omitempty"`
	IDPerms        *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName    string         `json:"display_name,omitempty"`
	Perms2         *PermType2     `json:"perms2,omitempty"`
	ParentType     string         `json:"parent_type,omitempty"`
	Annotations    *KeyValuePairs `json:"annotations,omitempty"`
	UUID           string         `json:"uuid,omitempty"`
	ParentUUID     string         `json:"parent_uuid,omitempty"`

	VirtualMachineInterfaceRefs []*SubnetVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
}

// SubnetVirtualMachineInterfaceRef references each other
type SubnetVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *Subnet) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSubnet makes Subnet
func MakeSubnet() *Subnet {
	return &Subnet{
		//TODO(nati): Apply default
		Annotations:    MakeKeyValuePairs(),
		UUID:           "",
		ParentUUID:     "",
		Perms2:         MakePermType2(),
		ParentType:     "",
		SubnetIPPrefix: MakeSubnetType(),
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
	}
}

// MakeSubnetSlice() makes a slice of Subnet
func MakeSubnetSlice() []*Subnet {
	return []*Subnet{}
}
