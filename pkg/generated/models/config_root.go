package models

// ConfigRoot

import "encoding/json"

// ConfigRoot
type ConfigRoot struct {
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`

	TagRefs []*ConfigRootTagRef `json:"tag_refs"`

	Domains             []*Domain             `json:"domains"`
	GlobalSystemConfigs []*GlobalSystemConfig `json:"global_system_configs"`
	Tags                []*Tag                `json:"tags"`
}

// ConfigRootTagRef references each other
type ConfigRootTagRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *ConfigRoot) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeConfigRoot makes ConfigRoot
func MakeConfigRoot() *ConfigRoot {
	return &ConfigRoot{
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

// InterfaceToConfigRoot makes ConfigRoot from interface
func InterfaceToConfigRoot(iData interface{}) *ConfigRoot {
	data := iData.(map[string]interface{})
	return &ConfigRoot{
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

// InterfaceToConfigRootSlice makes a slice of ConfigRoot from interface
func InterfaceToConfigRootSlice(data interface{}) []*ConfigRoot {
	list := data.([]interface{})
	result := MakeConfigRootSlice()
	for _, item := range list {
		result = append(result, InterfaceToConfigRoot(item))
	}
	return result
}

// MakeConfigRootSlice() makes a slice of ConfigRoot
func MakeConfigRootSlice() []*ConfigRoot {
	return []*ConfigRoot{}
}
