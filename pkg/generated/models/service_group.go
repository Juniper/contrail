package models

// ServiceGroup

import "encoding/json"

// ServiceGroup
type ServiceGroup struct {
	ParentUUID                      string                    `json:"parent_uuid"`
	ParentType                      string                    `json:"parent_type"`
	IDPerms                         *IdPermsType              `json:"id_perms"`
	Annotations                     *KeyValuePairs            `json:"annotations"`
	Perms2                          *PermType2                `json:"perms2"`
	ServiceGroupFirewallServiceList *FirewallServiceGroupType `json:"service_group_firewall_service_list"`
	UUID                            string                    `json:"uuid"`
	FQName                          []string                  `json:"fq_name"`
	DisplayName                     string                    `json:"display_name"`
}

// String returns json representation of the object
func (model *ServiceGroup) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceGroup makes ServiceGroup
func MakeServiceGroup() *ServiceGroup {
	return &ServiceGroup{
		//TODO(nati): Apply default
		UUID:                            "",
		FQName:                          []string{},
		DisplayName:                     "",
		ServiceGroupFirewallServiceList: MakeFirewallServiceGroupType(),
		ParentType:                      "",
		IDPerms:                         MakeIdPermsType(),
		Annotations:                     MakeKeyValuePairs(),
		Perms2:                          MakePermType2(),
		ParentUUID:                      "",
	}
}

// InterfaceToServiceGroup makes ServiceGroup from interface
func InterfaceToServiceGroup(iData interface{}) *ServiceGroup {
	data := iData.(map[string]interface{})
	return &ServiceGroup{
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ServiceGroupFirewallServiceList: InterfaceToFirewallServiceGroupType(data["service_group_firewall_service_list"]),

		//{"description":"list of service objects (protocol, source port and destination port","type":"object","properties":{"firewall_service":{"type":"array","item":{"type":"object","properties":{"dst_ports":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}},"protocol":{"type":"string"},"protocol_id":{"type":"integer"},"src_ports":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToServiceGroupSlice makes a slice of ServiceGroup from interface
func InterfaceToServiceGroupSlice(data interface{}) []*ServiceGroup {
	list := data.([]interface{})
	result := MakeServiceGroupSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceGroup(item))
	}
	return result
}

// MakeServiceGroupSlice() makes a slice of ServiceGroup
func MakeServiceGroupSlice() []*ServiceGroup {
	return []*ServiceGroup{}
}
