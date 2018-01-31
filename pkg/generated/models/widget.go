package models

// Widget

import "encoding/json"

// Widget
type Widget struct {
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	ContentConfig   string         `json:"content_config,omitempty"`
	UUID            string         `json:"uuid,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	ContainerConfig string         `json:"container_config,omitempty"`
	LayoutConfig    string         `json:"layout_config,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
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
		LayoutConfig:    "",
		ParentUUID:      "",
		ParentType:      "",
		Annotations:     MakeKeyValuePairs(),
		ContentConfig:   "",
		UUID:            "",
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Perms2:          MakePermType2(),
	}
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}
