package models

// LogicalInterface

import "encoding/json"

// LogicalInterface
type LogicalInterface struct {
	LogicalInterfaceVlanTag int                  `json:"logical_interface_vlan_tag,omitempty"`
	LogicalInterfaceType    LogicalInterfaceType `json:"logical_interface_type,omitempty"`
	Annotations             *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2                  *PermType2           `json:"perms2,omitempty"`
	UUID                    string               `json:"uuid,omitempty"`
	ParentUUID              string               `json:"parent_uuid,omitempty"`
	FQName                  []string             `json:"fq_name,omitempty"`
	IDPerms                 *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName             string               `json:"display_name,omitempty"`
	ParentType              string               `json:"parent_type,omitempty"`

	VirtualMachineInterfaceRefs []*LogicalInterfaceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
}

// LogicalInterfaceVirtualMachineInterfaceRef references each other
type LogicalInterfaceVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *LogicalInterface) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLogicalInterface makes LogicalInterface
func MakeLogicalInterface() *LogicalInterface {
	return &LogicalInterface{
		//TODO(nati): Apply default
		FQName:                  []string{},
		IDPerms:                 MakeIdPermsType(),
		DisplayName:             "",
		ParentType:              "",
		LogicalInterfaceVlanTag: 0,
		LogicalInterfaceType:    MakeLogicalInterfaceType(),
		Annotations:             MakeKeyValuePairs(),
		Perms2:                  MakePermType2(),
		UUID:                    "",
		ParentUUID:              "",
	}
}

// MakeLogicalInterfaceSlice() makes a slice of LogicalInterface
func MakeLogicalInterfaceSlice() []*LogicalInterface {
	return []*LogicalInterface{}
}
