package models

// LogicalInterface

import "encoding/json"

// LogicalInterface
type LogicalInterface struct {
	FQName                  []string             `json:"fq_name"`
	IDPerms                 *IdPermsType         `json:"id_perms"`
	DisplayName             string               `json:"display_name"`
	Annotations             *KeyValuePairs       `json:"annotations"`
	Perms2                  *PermType2           `json:"perms2"`
	UUID                    string               `json:"uuid"`
	ParentUUID              string               `json:"parent_uuid"`
	LogicalInterfaceVlanTag int                  `json:"logical_interface_vlan_tag"`
	LogicalInterfaceType    LogicalInterfaceType `json:"logical_interface_type"`
	ParentType              string               `json:"parent_type"`

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
		DisplayName:             "",
		FQName:                  []string{},
		IDPerms:                 MakeIdPermsType(),
		ParentType:              "",
		Annotations:             MakeKeyValuePairs(),
		Perms2:                  MakePermType2(),
		UUID:                    "",
		ParentUUID:              "",
		LogicalInterfaceVlanTag: 0,
		LogicalInterfaceType:    MakeLogicalInterfaceType(),
	}
}

// InterfaceToLogicalInterface makes LogicalInterface from interface
func InterfaceToLogicalInterface(iData interface{}) *LogicalInterface {
	data := iData.(map[string]interface{})
	return &LogicalInterface{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		LogicalInterfaceVlanTag: data["logical_interface_vlan_tag"].(int),

		//{"description":"VLAN tag (.1Q) classifier for this logical interface.","type":"integer"}
		LogicalInterfaceType: InterfaceToLogicalInterfaceType(data["logical_interface_type"]),

		//{"type":"string","enum":["l2","l3"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}

	}
}

// InterfaceToLogicalInterfaceSlice makes a slice of LogicalInterface from interface
func InterfaceToLogicalInterfaceSlice(data interface{}) []*LogicalInterface {
	list := data.([]interface{})
	result := MakeLogicalInterfaceSlice()
	for _, item := range list {
		result = append(result, InterfaceToLogicalInterface(item))
	}
	return result
}

// MakeLogicalInterfaceSlice() makes a slice of LogicalInterface
func MakeLogicalInterfaceSlice() []*LogicalInterface {
	return []*LogicalInterface{}
}
