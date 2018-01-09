package models

// BGPRouter

import "encoding/json"

// BGPRouter
type BGPRouter struct {
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
}

// String returns json representation of the object
func (model *BGPRouter) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBGPRouter makes BGPRouter
func MakeBGPRouter() *BGPRouter {
	return &BGPRouter{
		//TODO(nati): Apply default
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
	}
}

// InterfaceToBGPRouter makes BGPRouter from interface
func InterfaceToBGPRouter(iData interface{}) *BGPRouter {
	data := iData.(map[string]interface{})
	return &BGPRouter{
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
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToBGPRouterSlice makes a slice of BGPRouter from interface
func InterfaceToBGPRouterSlice(data interface{}) []*BGPRouter {
	list := data.([]interface{})
	result := MakeBGPRouterSlice()
	for _, item := range list {
		result = append(result, InterfaceToBGPRouter(item))
	}
	return result
}

// MakeBGPRouterSlice() makes a slice of BGPRouter
func MakeBGPRouterSlice() []*BGPRouter {
	return []*BGPRouter{}
}
