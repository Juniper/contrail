package models

// Dashboard

import "encoding/json"

// Dashboard
type Dashboard struct {
	ContainerConfig string         `json:"container_config,omitempty"`
	UUID            string         `json:"uuid,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
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
		ContainerConfig: "",
		UUID:            "",
		ParentUUID:      "",
		FQName:          []string{},
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		ParentType:      "",
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
	}
}

// MakeDashboardSlice() makes a slice of Dashboard
func MakeDashboardSlice() []*Dashboard {
	return []*Dashboard{}
}
