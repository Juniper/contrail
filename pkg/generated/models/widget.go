package models

// Widget

import "encoding/json"

// Widget
type Widget struct {
	FQName          []string       `json:"fq_name,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	LayoutConfig    string         `json:"layout_config,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	UUID            string         `json:"uuid,omitempty"`
	ContainerConfig string         `json:"container_config,omitempty"`
	ContentConfig   string         `json:"content_config,omitempty"`
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
		LayoutConfig:    "",
		ParentType:      "",
		FQName:          []string{},
		Annotations:     MakeKeyValuePairs(),
		ParentUUID:      "",
		ContainerConfig: "",
		ContentConfig:   "",
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Perms2:          MakePermType2(),
		UUID:            "",
	}
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}
