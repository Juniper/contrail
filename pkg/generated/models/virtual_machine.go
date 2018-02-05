package models

// VirtualMachine

import "encoding/json"

// VirtualMachine
type VirtualMachine struct {
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`

	ServiceInstanceRefs []*VirtualMachineServiceInstanceRef `json:"service_instance_refs,omitempty"`

	VirtualMachineInterfaces []*VirtualMachineInterface `json:"virtual_machine_interfaces,omitempty"`
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
