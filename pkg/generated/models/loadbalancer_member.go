package models

// LoadbalancerMember

import "encoding/json"

// LoadbalancerMember
type LoadbalancerMember struct {
	LoadbalancerMemberProperties *LoadbalancerMemberType `json:"loadbalancer_member_properties"`
	Annotations                  *KeyValuePairs          `json:"annotations"`
	Perms2                       *PermType2              `json:"perms2"`
	ParentUUID                   string                  `json:"parent_uuid"`
	ParentType                   string                  `json:"parent_type"`
	IDPerms                      *IdPermsType            `json:"id_perms"`
	DisplayName                  string                  `json:"display_name"`
	UUID                         string                  `json:"uuid"`
	FQName                       []string                `json:"fq_name"`
}

// String returns json representation of the object
func (model *LoadbalancerMember) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerMember makes LoadbalancerMember
func MakeLoadbalancerMember() *LoadbalancerMember {
	return &LoadbalancerMember{
		//TODO(nati): Apply default
		LoadbalancerMemberProperties: MakeLoadbalancerMemberType(),
		Annotations:                  MakeKeyValuePairs(),
		Perms2:                       MakePermType2(),
		ParentUUID:                   "",
		ParentType:                   "",
		IDPerms:                      MakeIdPermsType(),
		DisplayName:                  "",
		UUID:                         "",
		FQName:                       []string{},
	}
}

// InterfaceToLoadbalancerMember makes LoadbalancerMember from interface
func InterfaceToLoadbalancerMember(iData interface{}) *LoadbalancerMember {
	data := iData.(map[string]interface{})
	return &LoadbalancerMember{
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		LoadbalancerMemberProperties: InterfaceToLoadbalancerMemberType(data["loadbalancer_member_properties"]),

		//{"description":"Member configuration like ip address, destination port, weight etc.","type":"object","properties":{"address":{"type":"string"},"admin_state":{"type":"boolean"},"protocol_port":{"type":"integer"},"status":{"type":"string"},"status_description":{"type":"string"},"weight":{"type":"integer"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToLoadbalancerMemberSlice makes a slice of LoadbalancerMember from interface
func InterfaceToLoadbalancerMemberSlice(data interface{}) []*LoadbalancerMember {
	list := data.([]interface{})
	result := MakeLoadbalancerMemberSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerMember(item))
	}
	return result
}

// MakeLoadbalancerMemberSlice() makes a slice of LoadbalancerMember
func MakeLoadbalancerMemberSlice() []*LoadbalancerMember {
	return []*LoadbalancerMember{}
}
