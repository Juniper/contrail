package models

// KubernetesNode

import "encoding/json"

// KubernetesNode
type KubernetesNode struct {
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
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
		ProvisioningProgressStage: "",
		IDPerms:                   MakeIdPermsType(),
		ProvisioningState:         "",
		ProvisioningProgress:      0,
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ParentUUID:                "",
		ProvisioningLog:           "",
		ParentType:                "",
		FQName:                    []string{},
		DisplayName:               "",
		ProvisioningStartTime:     "",
	}
}

// MakeKubernetesNodeSlice() makes a slice of KubernetesNode
func MakeKubernetesNodeSlice() []*KubernetesNode {
	return []*KubernetesNode{}
}
