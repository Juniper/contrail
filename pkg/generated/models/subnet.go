package models

// Subnet

import "encoding/json"

// Subnet
type Subnet struct {
	SubnetIPPrefix *SubnetType    `json:"subnet_ip_prefix"`
	Perms2         *PermType2     `json:"perms2"`
	IDPerms        *IdPermsType   `json:"id_perms"`
	DisplayName    string         `json:"display_name"`
	UUID           string         `json:"uuid"`
	ParentUUID     string         `json:"parent_uuid"`
	ParentType     string         `json:"parent_type"`
	FQName         []string       `json:"fq_name"`
	Annotations    *KeyValuePairs `json:"annotations"`

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
		Perms2:         MakePermType2(),
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		SubnetIPPrefix: MakeSubnetType(),
		ParentUUID:     "",
		ParentType:     "",
		FQName:         []string{},
		Annotations:    MakeKeyValuePairs(),
		UUID:           "",
	}
}

// InterfaceToSubnet makes Subnet from interface
func InterfaceToSubnet(iData interface{}) *Subnet {
	data := iData.(map[string]interface{})
	return &Subnet{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		SubnetIPPrefix: InterfaceToSubnetType(data["subnet_ip_prefix"]),

		//{"description":"Ip prefix/length of the subnet.","type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToSubnetSlice makes a slice of Subnet from interface
func InterfaceToSubnetSlice(data interface{}) []*Subnet {
	list := data.([]interface{})
	result := MakeSubnetSlice()
	for _, item := range list {
		result = append(result, InterfaceToSubnet(item))
	}
	return result
}

// MakeSubnetSlice() makes a slice of Subnet
func MakeSubnetSlice() []*Subnet {
	return []*Subnet{}
}
