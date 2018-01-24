package models

// Widget

import "encoding/json"

// Widget
type Widget struct {
	ContentConfig   string         `json:"content_config,omitempty"`
	LayoutConfig    string         `json:"layout_config,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	ContainerConfig string         `json:"container_config,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	UUID            string         `json:"uuid,omitempty"`
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
		UUID:            "",
		ContainerConfig: "",
		ParentType:      "",
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		ParentUUID:      "",
		ContentConfig:   "",
		LayoutConfig:    "",
	}
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}
