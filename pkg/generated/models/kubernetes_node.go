package models

// KubernetesNode

import "encoding/json"

// KubernetesNode
type KubernetesNode struct {
	IDPerms                   *IdPermsType   `json:"id_perms"`
	Annotations               *KeyValuePairs `json:"annotations"`
	Perms2                    *PermType2     `json:"perms2"`
	ParentType                string         `json:"parent_type"`
	ProvisioningLog           string         `json:"provisioning_log"`
	ProvisioningState         string         `json:"provisioning_state"`
	FQName                    []string       `json:"fq_name"`
	DisplayName               string         `json:"display_name"`
	UUID                      string         `json:"uuid"`
	ParentUUID                string         `json:"parent_uuid"`
	ProvisioningProgress      int            `json:"provisioning_progress"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage"`
	ProvisioningStartTime     string         `json:"provisioning_start_time"`
}

// String returns json representation of the object
func (model *KubernetesNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeKubernetesNode makes KubernetesNode
func MakeKubernetesNode() *KubernetesNode {
	return &KubernetesNode{
		//TODO(nati): Apply default
		IDPerms:                   MakeIdPermsType(),
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		ParentType:                "",
		ProvisioningLog:           "",
		ProvisioningState:         "",
		FQName:                    []string{},
		DisplayName:               "",
		UUID:                      "",
		ParentUUID:                "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
	}
}

// InterfaceToKubernetesNode makes KubernetesNode from interface
func InterfaceToKubernetesNode(iData interface{}) *KubernetesNode {
	data := iData.(map[string]interface{})
	return &KubernetesNode{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}

	}
}

// InterfaceToKubernetesNodeSlice makes a slice of KubernetesNode from interface
func InterfaceToKubernetesNodeSlice(data interface{}) []*KubernetesNode {
	list := data.([]interface{})
	result := MakeKubernetesNodeSlice()
	for _, item := range list {
		result = append(result, InterfaceToKubernetesNode(item))
	}
	return result
}

// MakeKubernetesNodeSlice() makes a slice of KubernetesNode
func MakeKubernetesNodeSlice() []*KubernetesNode {
	return []*KubernetesNode{}
}
