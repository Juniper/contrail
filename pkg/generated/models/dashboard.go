package models

// Dashboard

import "encoding/json"

// Dashboard
type Dashboard struct {
	FQName          []string       `json:"fq_name,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	UUID            string         `json:"uuid,omitempty"`
	ContainerConfig string         `json:"container_config,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
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
		FQName:          []string{},
		Perms2:          MakePermType2(),
		UUID:            "",
		ContainerConfig: "",
		ParentUUID:      "",
		ParentType:      "",
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
	}
}

// MakeDashboardSlice() makes a slice of Dashboard
func MakeDashboardSlice() []*Dashboard {
	return []*Dashboard{}
}
