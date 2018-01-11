package models

// VirtualMachine

import "encoding/json"

// VirtualMachine
type VirtualMachine struct {
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`

	ServiceInstanceRefs []*VirtualMachineServiceInstanceRef `json:"service_instance_refs"`

	VirtualMachineInterfaces []*VirtualMachineInterface `json:"virtual_machine_interfaces"`
}

// VirtualMachineServiceInstanceRef references each other
type VirtualMachineServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *VirtualMachine) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualMachine makes VirtualMachine
func MakeVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		//TODO(nati): Apply default
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
	}
}

// MakeVirtualMachineSlice() makes a slice of VirtualMachine
func MakeVirtualMachineSlice() []*VirtualMachine {
	return []*VirtualMachine{}
}
