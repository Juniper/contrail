package models

// Dashboard

import "encoding/json"

// Dashboard
type Dashboard struct {
	ParentType      string         `json:"parent_type"`
	FQName          []string       `json:"fq_name"`
	IDPerms         *IdPermsType   `json:"id_perms"`
	Annotations     *KeyValuePairs `json:"annotations"`
	Perms2          *PermType2     `json:"perms2"`
	ContainerConfig string         `json:"container_config"`
	UUID            string         `json:"uuid"`
	ParentUUID      string         `json:"parent_uuid"`
	DisplayName     string         `json:"display_name"`
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
		Perms2:          MakePermType2(),
		ContainerConfig: "",
		UUID:            "",
		ParentUUID:      "",
		ParentType:      "",
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		Annotations:     MakeKeyValuePairs(),
		DisplayName:     "",
	}
}

// InterfaceToDashboard makes Dashboard from interface
func InterfaceToDashboard(iData interface{}) *Dashboard {
	data := iData.(map[string]interface{})
	return &Dashboard{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ContainerConfig: data["container_config"].(string),

		//{"title":"Container Config","type":"string","permission":["create","update"]}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

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
