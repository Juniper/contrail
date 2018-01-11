package models

// ConfigRoot

import "encoding/json"

// ConfigRoot
type ConfigRoot struct {
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`

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
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
	}
}

// MakeConfigRootSlice() makes a slice of ConfigRoot
func MakeConfigRootSlice() []*ConfigRoot {
	return []*ConfigRoot{}
}
