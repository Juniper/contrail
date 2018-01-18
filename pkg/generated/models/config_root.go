package models

// ConfigRoot

import "encoding/json"

// ConfigRoot
type ConfigRoot struct {
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`

	TagRefs []*ConfigRootTagRef `json:"tag_refs,omitempty"`

	Domains             []*Domain             `json:"domains,omitempty"`
	GlobalSystemConfigs []*GlobalSystemConfig `json:"global_system_configs,omitempty"`
	Tags                []*Tag                `json:"tags,omitempty"`
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
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeConfigRootSlice() makes a slice of ConfigRoot
func MakeConfigRootSlice() []*ConfigRoot {
	return []*ConfigRoot{}
}
