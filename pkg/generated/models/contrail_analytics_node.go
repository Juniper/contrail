package models

// ContrailAnalyticsNode

import "encoding/json"

// ContrailAnalyticsNode
type ContrailAnalyticsNode struct {
	IDPerms                   *IdPermsType   `json:"id_perms"`
	Perms2                    *PermType2     `json:"perms2"`
	UUID                      string         `json:"uuid"`
	ProvisioningLog           string         `json:"provisioning_log"`
	ProvisioningProgress      int            `json:"provisioning_progress"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage"`
	ParentType                string         `json:"parent_type"`
	FQName                    []string       `json:"fq_name"`
	ProvisioningStartTime     string         `json:"provisioning_start_time"`
	ProvisioningState         string         `json:"provisioning_state"`
	Annotations               *KeyValuePairs `json:"annotations"`
	ParentUUID                string         `json:"parent_uuid"`
	DisplayName               string         `json:"display_name"`
}

// String returns json representation of the object
func (model *ContrailAnalyticsNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeContrailAnalyticsNode makes ContrailAnalyticsNode
func MakeContrailAnalyticsNode() *ContrailAnalyticsNode {
	return &ContrailAnalyticsNode{
		//TODO(nati): Apply default
		ParentUUID:                "",
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
	}
}

// InterfaceToContrailAnalyticsNode makes ContrailAnalyticsNode from interface
func InterfaceToContrailAnalyticsNode(iData interface{}) *ContrailAnalyticsNode {
	data := iData.(map[string]interface{})
	return &ContrailAnalyticsNode{
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}

	}
}

// InterfaceToContrailAnalyticsNodeSlice makes a slice of ContrailAnalyticsNode from interface
func InterfaceToContrailAnalyticsNodeSlice(data interface{}) []*ContrailAnalyticsNode {
	list := data.([]interface{})
	result := MakeContrailAnalyticsNodeSlice()
	for _, item := range list {
		result = append(result, InterfaceToContrailAnalyticsNode(item))
	}
	return result
}

// MakeContrailAnalyticsNodeSlice() makes a slice of ContrailAnalyticsNode
func MakeContrailAnalyticsNodeSlice() []*ContrailAnalyticsNode {
	return []*ContrailAnalyticsNode{}
}
