package models

// Dashboard

import "encoding/json"

// Dashboard
type Dashboard struct {
	ParentType      string         `json:"parent_type"`
	FQName          []string       `json:"fq_name"`
	DisplayName     string         `json:"display_name"`
	Annotations     *KeyValuePairs `json:"annotations"`
	UUID            string         `json:"uuid"`
	Perms2          *PermType2     `json:"perms2"`
	ParentUUID      string         `json:"parent_uuid"`
	IDPerms         *IdPermsType   `json:"id_perms"`
	ContainerConfig string         `json:"container_config"`
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
		IDPerms:         MakeIdPermsType(),
		ContainerConfig: "",
		Perms2:          MakePermType2(),
		ParentUUID:      "",
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		UUID:            "",
		ParentType:      "",
		FQName:          []string{},
	}
}

// InterfaceToDashboard makes Dashboard from interface
func InterfaceToDashboard(iData interface{}) *Dashboard {
	data := iData.(map[string]interface{})
	return &Dashboard{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ContainerConfig: data["container_config"].(string),

		//{"title":"Container Config","type":"string","permission":["create","update"]}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToDashboardSlice makes a slice of Dashboard from interface
func InterfaceToDashboardSlice(data interface{}) []*Dashboard {
	list := data.([]interface{})
	result := MakeDashboardSlice()
	for _, item := range list {
		result = append(result, InterfaceToDashboard(item))
	}
	return result
}

// MakeDashboardSlice() makes a slice of Dashboard
func MakeDashboardSlice() []*Dashboard {
	return []*Dashboard{}
}
