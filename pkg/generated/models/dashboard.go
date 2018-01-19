package models

// Dashboard

import "encoding/json"

// Dashboard
type Dashboard struct {
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	UUID            string         `json:"uuid,omitempty"`
	ContainerConfig string         `json:"container_config,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
}

// String returns json representation of the object
func (model *Dashboard) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDashboard makes Dashboard
func MakeDashboard() *Dashboard {
	return &Dashboard{
		//TODO(nati): Apply default
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		UUID:            "",
		ParentUUID:      "",
		ParentType:      "",
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		ContainerConfig: "",
		FQName:          []string{},
	}
}

// MakeDashboardSlice() makes a slice of Dashboard
func MakeDashboardSlice() []*Dashboard {
	return []*Dashboard{}
}
