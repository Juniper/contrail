package models

// Widget

import "encoding/json"

// Widget
type Widget struct {
	ContentConfig   string         `json:"content_config"`
	ParentType      string         `json:"parent_type"`
	FQName          []string       `json:"fq_name"`
	IDPerms         *IdPermsType   `json:"id_perms"`
	ParentUUID      string         `json:"parent_uuid"`
	DisplayName     string         `json:"display_name"`
	ContainerConfig string         `json:"container_config"`
	LayoutConfig    string         `json:"layout_config"`
	Annotations     *KeyValuePairs `json:"annotations"`
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
		ContentConfig:   "",
		LayoutConfig:    "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		UUID:            "",
		ParentUUID:      "",
		DisplayName:     "",
		ContainerConfig: "",
	}
}

// InterfaceToWidget makes Widget from interface
func InterfaceToWidget(iData interface{}) *Widget {
	data := iData.(map[string]interface{})
	return &Widget{
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ContainerConfig: data["container_config"].(string),

		//{"title":"Container Config","type":"string","permission":["create","update"]}
		LayoutConfig: data["layout_config"].(string),

		//{"title":"Layout Config","type":"string","permission":["create","update"]}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ContentConfig: data["content_config"].(string),

		//{"title":"Content Config","type":"string","permission":["create","update"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToWidgetSlice makes a slice of Widget from interface
func InterfaceToWidgetSlice(data interface{}) []*Widget {
	list := data.([]interface{})
	result := MakeWidgetSlice()
	for _, item := range list {
		result = append(result, InterfaceToWidget(item))
	}
	return result
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}
