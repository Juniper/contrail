package models

// Dashboard

import "encoding/json"

// Dashboard
type Dashboard struct {
	ContainerConfig string         `json:"container_config"`
	Perms2          *PermType2     `json:"perms2"`
	UUID            string         `json:"uuid"`
	ParentType      string         `json:"parent_type"`
	FQName          []string       `json:"fq_name"`
	DisplayName     string         `json:"display_name"`
	Annotations     *KeyValuePairs `json:"annotations"`
	ParentUUID      string         `json:"parent_uuid"`
	IDPerms         *IdPermsType   `json:"id_perms"`
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
		ParentUUID:      "",
		IDPerms:         MakeIdPermsType(),
		ContainerConfig: "",
		Perms2:          MakePermType2(),
		UUID:            "",
		ParentType:      "",
		FQName:          []string{},
		DisplayName:     "",
	}
}

// MakeDashboardSlice() makes a slice of Dashboard
func MakeDashboardSlice() []*Dashboard {
	return []*Dashboard{}
}
