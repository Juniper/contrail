package models

// Widget

import "encoding/json"

// Widget
type Widget struct {
	ParentType      string         `json:"parent_type,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	LayoutConfig    string         `json:"layout_config,omitempty"`
	UUID            string         `json:"uuid,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	ContainerConfig string         `json:"container_config,omitempty"`
	ContentConfig   string         `json:"content_config,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
}

// String returns json representation of the object
func (model *Widget) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeWidget makes Widget
func MakeWidget() *Widget {
	return &Widget{
		//TODO(nati): Apply default
		DisplayName:     "",
		ContainerConfig: "",
		ContentConfig:   "",
		IDPerms:         MakeIdPermsType(),
		ParentType:      "",
		FQName:          []string{},
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		LayoutConfig:    "",
		UUID:            "",
		ParentUUID:      "",
	}
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}
