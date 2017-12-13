package models

// VirtualMachine

import "encoding/json"

// VirtualMachine
type VirtualMachine struct {
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`

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
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
	}
}

// InterfaceToVirtualMachine makes VirtualMachine from interface
func InterfaceToVirtualMachine(iData interface{}) *VirtualMachine {
	data := iData.(map[string]interface{})
	return &VirtualMachine{
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}

	}
}

// InterfaceToVirtualMachineSlice makes a slice of VirtualMachine from interface
func InterfaceToVirtualMachineSlice(data interface{}) []*VirtualMachine {
	list := data.([]interface{})
	result := MakeVirtualMachineSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualMachine(item))
	}
	return result
}

// MakeVirtualMachineSlice() makes a slice of VirtualMachine
func MakeVirtualMachineSlice() []*VirtualMachine {
	return []*VirtualMachine{}
}
