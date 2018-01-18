package models

// Widget

import "encoding/json"

// Widget
type Widget struct {
	ContainerConfig string         `json:"container_config,omitempty"`
	ContentConfig   string         `json:"content_config,omitempty"`
	LayoutConfig    string         `json:"layout_config,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
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
		ContainerConfig: "",
		ContentConfig:   "",
		LayoutConfig:    "",
		IDPerms:         MakeIdPermsType(),
		Annotations:     MakeKeyValuePairs(),
		ParentUUID:      "",
		ParentType:      "",
		FQName:          []string{},
		DisplayName:     "",
		Perms2:          MakePermType2(),
		UUID:            "",
	}
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}
