package models

// ContrailControllerNodeRole

import "encoding/json"

// ContrailControllerNodeRole
type ContrailControllerNodeRole struct {
	ParentUUID                string         `json:"parent_uuid"`
	FQName                    []string       `json:"fq_name"`
	IDPerms                   *IdPermsType   `json:"id_perms"`
	DisplayName               string         `json:"display_name"`
	Annotations               *KeyValuePairs `json:"annotations"`
	ProvisioningLog           string         `json:"provisioning_log"`
	ProvisioningProgress      int            `json:"provisioning_progress"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage"`
	ProvisioningStartTime     string         `json:"provisioning_start_time"`
	ParentType                string         `json:"parent_type"`
	Perms2                    *PermType2     `json:"perms2"`
	UUID                      string         `json:"uuid"`
	ProvisioningState         string         `json:"provisioning_state"`
}

// String returns json representation of the object
func (model *ContrailControllerNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeContrailControllerNodeRole makes ContrailControllerNodeRole
func MakeContrailControllerNodeRole() *ContrailControllerNodeRole {
	return &ContrailControllerNodeRole{
		//TODO(nati): Apply default
		ParentType:                "",
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ProvisioningState:         "",
		ProvisioningProgressStage: "",
		ParentUUID:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		ProvisioningStartTime:     "",
	}
}

// InterfaceToContrailControllerNodeRole makes ContrailControllerNodeRole from interface
func InterfaceToContrailControllerNodeRole(iData interface{}) *ContrailControllerNodeRole {
	data := iData.(map[string]interface{})
	return &ContrailControllerNodeRole{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}

	}
}

// InterfaceToContrailControllerNodeRoleSlice makes a slice of ContrailControllerNodeRole from interface
func InterfaceToContrailControllerNodeRoleSlice(data interface{}) []*ContrailControllerNodeRole {
	list := data.([]interface{})
	result := MakeContrailControllerNodeRoleSlice()
	for _, item := range list {
		result = append(result, InterfaceToContrailControllerNodeRole(item))
	}
	return result
}

// MakeContrailControllerNodeRoleSlice() makes a slice of ContrailControllerNodeRole
func MakeContrailControllerNodeRoleSlice() []*ContrailControllerNodeRole {
	return []*ContrailControllerNodeRole{}
}
