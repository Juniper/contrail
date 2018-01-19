package models

// LogicalInterface

import "encoding/json"

// LogicalInterface
type LogicalInterface struct {
	IDPerms                 *IdPermsType         `json:"id_perms,omitempty"`
	ParentType              string               `json:"parent_type,omitempty"`
	FQName                  []string             `json:"fq_name,omitempty"`
	Perms2                  *PermType2           `json:"perms2,omitempty"`
	UUID                    string               `json:"uuid,omitempty"`
	ParentUUID              string               `json:"parent_uuid,omitempty"`
	LogicalInterfaceVlanTag int                  `json:"logical_interface_vlan_tag,omitempty"`
	LogicalInterfaceType    LogicalInterfaceType `json:"logical_interface_type,omitempty"`
	DisplayName             string               `json:"display_name,omitempty"`
	Annotations             *KeyValuePairs       `json:"annotations,omitempty"`

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
		ParentType:              "",
		FQName:                  []string{},
		IDPerms:                 MakeIdPermsType(),
		LogicalInterfaceType:    MakeLogicalInterfaceType(),
		DisplayName:             "",
		Annotations:             MakeKeyValuePairs(),
		Perms2:                  MakePermType2(),
		UUID:                    "",
		ParentUUID:              "",
		LogicalInterfaceVlanTag: 0,
	}
}

// MakeLogicalInterfaceSlice() makes a slice of LogicalInterface
func MakeLogicalInterfaceSlice() []*LogicalInterface {
	return []*LogicalInterface{}
}
