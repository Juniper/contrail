package models

// LogicalInterface

import "encoding/json"

// LogicalInterface
type LogicalInterface struct {
	LogicalInterfaceVlanTag int                  `json:"logical_interface_vlan_tag"`
	LogicalInterfaceType    LogicalInterfaceType `json:"logical_interface_type"`
	ParentUUID              string               `json:"parent_uuid"`
	ParentType              string               `json:"parent_type"`
	FQName                  []string             `json:"fq_name"`
	DisplayName             string               `json:"display_name"`
	Annotations             *KeyValuePairs       `json:"annotations"`
	IDPerms                 *IdPermsType         `json:"id_perms"`
	Perms2                  *PermType2           `json:"perms2"`
	UUID                    string               `json:"uuid"`

	VirtualMachineInterfaceRefs []*LogicalInterfaceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
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
		IDPerms:                 MakeIdPermsType(),
		Perms2:                  MakePermType2(),
		UUID:                    "",
		FQName:                  []string{},
		DisplayName:             "",
		Annotations:             MakeKeyValuePairs(),
		LogicalInterfaceVlanTag: 0,
		LogicalInterfaceType:    MakeLogicalInterfaceType(),
		ParentUUID:              "",
		ParentType:              "",
	}
}

// MakeLogicalInterfaceSlice() makes a slice of LogicalInterface
func MakeLogicalInterfaceSlice() []*LogicalInterface {
	return []*LogicalInterface{}
}
