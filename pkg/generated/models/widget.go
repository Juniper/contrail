package models

// Widget

import "encoding/json"

// Widget
type Widget struct {
	UUID            string         `json:"uuid,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	ContainerConfig string         `json:"container_config,omitempty"`
	ContentConfig   string         `json:"content_config,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	LayoutConfig    string         `json:"layout_config,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
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
		ParentType:      "",
		FQName:          []string{},
		LayoutConfig:    "",
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		ParentUUID:      "",
		IDPerms:         MakeIdPermsType(),
		ContainerConfig: "",
		ContentConfig:   "",
		Perms2:          MakePermType2(),
		UUID:            "",
	}
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}
