package models

// Subnet

import "encoding/json"

// Subnet
type Subnet struct {
	ParentUUID     string         `json:"parent_uuid,omitempty"`
	IDPerms        *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName    string         `json:"display_name,omitempty"`
	ParentType     string         `json:"parent_type,omitempty"`
	FQName         []string       `json:"fq_name,omitempty"`
	SubnetIPPrefix *SubnetType    `json:"subnet_ip_prefix,omitempty"`
	Annotations    *KeyValuePairs `json:"annotations,omitempty"`
	Perms2         *PermType2     `json:"perms2,omitempty"`
	UUID           string         `json:"uuid,omitempty"`

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
		ParentUUID:     "",
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		SubnetIPPrefix: MakeSubnetType(),
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		UUID:           "",
		ParentType:     "",
		FQName:         []string{},
	}
}

// MakeSubnetSlice() makes a slice of Subnet
func MakeSubnetSlice() []*Subnet {
	return []*Subnet{}
}
