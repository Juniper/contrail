package models

// ContrailAnalyticsDatabaseNodeRole

import "encoding/json"

// ContrailAnalyticsDatabaseNodeRole
type ContrailAnalyticsDatabaseNodeRole struct {
	ProvisioningState         string         `json:"provisioning_state"`
	ParentType                string         `json:"parent_type"`
	IDPerms                   *IdPermsType   `json:"id_perms"`
	Annotations               *KeyValuePairs `json:"annotations"`
	Perms2                    *PermType2     `json:"perms2"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage"`
	ProvisioningStartTime     string         `json:"provisioning_start_time"`
	UUID                      string         `json:"uuid"`
	ParentUUID                string         `json:"parent_uuid"`
	FQName                    []string       `json:"fq_name"`
	DisplayName               string         `json:"display_name"`
	ProvisioningProgress      int            `json:"provisioning_progress"`
	ProvisioningLog           string         `json:"provisioning_log"`
}

// String returns json representation of the object
func (model *ContrailAnalyticsDatabaseNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeContrailAnalyticsDatabaseNodeRole makes ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRole() *ContrailAnalyticsDatabaseNodeRole {
	return &ContrailAnalyticsDatabaseNodeRole{
		//TODO(nati): Apply default
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
		UUID:                      "",
		ParentUUID:                "",
		FQName:                    []string{},
		DisplayName:               "",
		ProvisioningProgress:      0,
		ProvisioningLog:           "",
	}
}

// InterfaceToContrailAnalyticsDatabaseNodeRole makes ContrailAnalyticsDatabaseNodeRole from interface
func InterfaceToContrailAnalyticsDatabaseNodeRole(iData interface{}) *ContrailAnalyticsDatabaseNodeRole {
	data := iData.(map[string]interface{})
	return &ContrailAnalyticsDatabaseNodeRole{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}

	}
}

// InterfaceToContrailAnalyticsDatabaseNodeRoleSlice makes a slice of ContrailAnalyticsDatabaseNodeRole from interface
func InterfaceToContrailAnalyticsDatabaseNodeRoleSlice(data interface{}) []*ContrailAnalyticsDatabaseNodeRole {
	list := data.([]interface{})
	result := MakeContrailAnalyticsDatabaseNodeRoleSlice()
	for _, item := range list {
		result = append(result, InterfaceToContrailAnalyticsDatabaseNodeRole(item))
	}
	return result
}

// MakeContrailAnalyticsDatabaseNodeRoleSlice() makes a slice of ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRoleSlice() []*ContrailAnalyticsDatabaseNodeRole {
	return []*ContrailAnalyticsDatabaseNodeRole{}
}
