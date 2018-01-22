package models

// LogicalInterface

import "encoding/json"

// LogicalInterface
type LogicalInterface struct {
	Perms2                  *PermType2           `json:"perms2,omitempty"`
	ParentType              string               `json:"parent_type,omitempty"`
	LogicalInterfaceType    LogicalInterfaceType `json:"logical_interface_type,omitempty"`
	IDPerms                 *IdPermsType         `json:"id_perms,omitempty"`
	Annotations             *KeyValuePairs       `json:"annotations,omitempty"`
	UUID                    string               `json:"uuid,omitempty"`
	ParentUUID              string               `json:"parent_uuid,omitempty"`
	LogicalInterfaceVlanTag int                  `json:"logical_interface_vlan_tag,omitempty"`
	FQName                  []string             `json:"fq_name,omitempty"`
	DisplayName             string               `json:"display_name,omitempty"`

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
		LogicalInterfaceType:    MakeLogicalInterfaceType(),
		IDPerms:                 MakeIdPermsType(),
		Annotations:             MakeKeyValuePairs(),
		Perms2:                  MakePermType2(),
		ParentType:              "",
		LogicalInterfaceVlanTag: 0,
		FQName:                  []string{},
		DisplayName:             "",
		UUID:                    "",
		ParentUUID:              "",
	}
}

// MakeLogicalInterfaceSlice() makes a slice of LogicalInterface
func MakeLogicalInterfaceSlice() []*LogicalInterface {
	return []*LogicalInterface{}
}
