package models

// VirtualIP

import "encoding/json"

// VirtualIP
type VirtualIP struct {
	ParentUUID          string         `json:"parent_uuid"`
	ParentType          string         `json:"parent_type"`
	DisplayName         string         `json:"display_name"`
	UUID                string         `json:"uuid"`
	FQName              []string       `json:"fq_name"`
	IDPerms             *IdPermsType   `json:"id_perms"`
	Annotations         *KeyValuePairs `json:"annotations"`
	Perms2              *PermType2     `json:"perms2"`
	VirtualIPProperties *VirtualIpType `json:"virtual_ip_properties"`

	LoadbalancerPoolRefs        []*VirtualIPLoadbalancerPoolRef        `json:"loadbalancer_pool_refs"`
	VirtualMachineInterfaceRefs []*VirtualIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
}

// VirtualIPLoadbalancerPoolRef references each other
type VirtualIPLoadbalancerPoolRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualIPVirtualMachineInterfaceRef references each other
type VirtualIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *VirtualIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualIP makes VirtualIP
func MakeVirtualIP() *VirtualIP {
	return &VirtualIP{
		//TODO(nati): Apply default
		Perms2:              MakePermType2(),
		VirtualIPProperties: MakeVirtualIpType(),
		FQName:              []string{},
		IDPerms:             MakeIdPermsType(),
		Annotations:         MakeKeyValuePairs(),
		UUID:                "",
		ParentUUID:          "",
		ParentType:          "",
		DisplayName:         "",
	}
}

// InterfaceToVirtualIP makes VirtualIP from interface
func InterfaceToVirtualIP(iData interface{}) *VirtualIP {
	data := iData.(map[string]interface{})
	return &VirtualIP{
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		VirtualIPProperties: InterfaceToVirtualIpType(data["virtual_ip_properties"]),

		//{"description":"Virtual ip configuration like port, protocol, subnet etc.","type":"object","properties":{"address":{"type":"string"},"admin_state":{"type":"boolean"},"connection_limit":{"type":"integer"},"persistence_cookie_name":{"type":"string"},"persistence_type":{"type":"string","enum":["SOURCE_IP","HTTP_COOKIE","APP_COOKIE"]},"protocol":{"type":"string","enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"]},"protocol_port":{"type":"integer"},"status":{"type":"string"},"status_description":{"type":"string"},"subnet_id":{"type":"string"}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToVirtualIPSlice makes a slice of VirtualIP from interface
func InterfaceToVirtualIPSlice(data interface{}) []*VirtualIP {
	list := data.([]interface{})
	result := MakeVirtualIPSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualIP(item))
	}
	return result
}

// MakeVirtualIPSlice() makes a slice of VirtualIP
func MakeVirtualIPSlice() []*VirtualIP {
	return []*VirtualIP{}
}
