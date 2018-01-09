package models

// RoutingPolicy

import "encoding/json"

// RoutingPolicy
type RoutingPolicy struct {
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`

	ServiceInstanceRefs []*RoutingPolicyServiceInstanceRef `json:"service_instance_refs"`
}

// RoutingPolicyServiceInstanceRef references each other
type RoutingPolicyServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *RoutingPolicyServiceInstanceType
}

// String returns json representation of the object
func (model *RoutingPolicy) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRoutingPolicy makes RoutingPolicy
func MakeRoutingPolicy() *RoutingPolicy {
	return &RoutingPolicy{
		//TODO(nati): Apply default
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
	}
}

// InterfaceToRoutingPolicy makes RoutingPolicy from interface
func InterfaceToRoutingPolicy(iData interface{}) *RoutingPolicy {
	data := iData.(map[string]interface{})
	return &RoutingPolicy{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}

	}
}

// InterfaceToRoutingPolicySlice makes a slice of RoutingPolicy from interface
func InterfaceToRoutingPolicySlice(data interface{}) []*RoutingPolicy {
	list := data.([]interface{})
	result := MakeRoutingPolicySlice()
	for _, item := range list {
		result = append(result, InterfaceToRoutingPolicy(item))
	}
	return result
}

// MakeRoutingPolicySlice() makes a slice of RoutingPolicy
func MakeRoutingPolicySlice() []*RoutingPolicy {
	return []*RoutingPolicy{}
}
