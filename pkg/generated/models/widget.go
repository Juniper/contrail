package models

// Widget

import "encoding/json"

// Widget
type Widget struct {
	FQName          []string       `json:"fq_name"`
	IDPerms         *IdPermsType   `json:"id_perms"`
	DisplayName     string         `json:"display_name"`
	Annotations     *KeyValuePairs `json:"annotations"`
	ParentType      string         `json:"parent_type"`
	ParentUUID      string         `json:"parent_uuid"`
	ContainerConfig string         `json:"container_config"`
	ContentConfig   string         `json:"content_config"`
	LayoutConfig    string         `json:"layout_config"`
	Perms2          *PermType2     `json:"perms2"`
	UUID            string         `json:"uuid"`
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
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		UUID:            "",
		ParentUUID:      "",
		ContainerConfig: "",
		ContentConfig:   "",
		LayoutConfig:    "",
		Perms2:          MakePermType2(),
	}
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}
