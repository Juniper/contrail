package models

// Subnet

import "encoding/json"

// Subnet
type Subnet struct {
	ParentType     string         `json:"parent_type,omitempty"`
	IDPerms        *IdPermsType   `json:"id_perms,omitempty"`
	UUID           string         `json:"uuid,omitempty"`
	SubnetIPPrefix *SubnetType    `json:"subnet_ip_prefix,omitempty"`
	FQName         []string       `json:"fq_name,omitempty"`
	DisplayName    string         `json:"display_name,omitempty"`
	Annotations    *KeyValuePairs `json:"annotations,omitempty"`
	Perms2         *PermType2     `json:"perms2,omitempty"`
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
		ParentType:     "",
		IDPerms:        MakeIdPermsType(),
		UUID:           "",
		ParentUUID:     "",
		SubnetIPPrefix: MakeSubnetType(),
		FQName:         []string{},
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
	}
}

// MakeSubnetSlice() makes a slice of Subnet
func MakeSubnetSlice() []*Subnet {
	return []*Subnet{}
}
