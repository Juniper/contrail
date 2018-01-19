package models

// Widget

import "encoding/json"

// Widget
type Widget struct {
	UUID            string         `json:"uuid,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	ContainerConfig string         `json:"container_config,omitempty"`
	ContentConfig   string         `json:"content_config,omitempty"`
	LayoutConfig    string         `json:"layout_config,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
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
		ContentConfig:   "",
		LayoutConfig:    "",
		UUID:            "",
		ParentUUID:      "",
		IDPerms:         MakeIdPermsType(),
		ContainerConfig: "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		ParentType:      "",
		FQName:          []string{},
		DisplayName:     "",
	}
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}
