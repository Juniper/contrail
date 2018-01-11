package models

// Subnet

import "encoding/json"

// Subnet
type Subnet struct {
	DisplayName    string         `json:"display_name"`
	Perms2         *PermType2     `json:"perms2"`
	ParentType     string         `json:"parent_type"`
	FQName         []string       `json:"fq_name"`
	IDPerms        *IdPermsType   `json:"id_perms"`
	SubnetIPPrefix *SubnetType    `json:"subnet_ip_prefix"`
	Annotations    *KeyValuePairs `json:"annotations"`
	UUID           string         `json:"uuid"`
	ParentUUID     string         `json:"parent_uuid"`

	VirtualMachineInterfaceRefs []*SubnetVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
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
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Perms2:         MakePermType2(),
		UUID:           "",
		ParentUUID:     "",
		SubnetIPPrefix: MakeSubnetType(),
		Annotations:    MakeKeyValuePairs(),
	}
}

// MakeSubnetSlice() makes a slice of Subnet
func MakeSubnetSlice() []*Subnet {
	return []*Subnet{}
}
