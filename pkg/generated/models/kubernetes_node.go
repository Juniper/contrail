package models

// KubernetesNode

import "encoding/json"

// KubernetesNode
type KubernetesNode struct {
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
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
		DisplayName:               "",
		Perms2:                    MakePermType2(),
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		ProvisioningStartTime:     "",
		ProvisioningLog:           "",
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		ProvisioningProgressStage: "",
		ProvisioningState:         "",
		ProvisioningProgress:      0,
	}
}

// MakeKubernetesNodeSlice() makes a slice of KubernetesNode
func MakeKubernetesNodeSlice() []*KubernetesNode {
	return []*KubernetesNode{}
}
