package models

// KubernetesNode

import "encoding/json"

// KubernetesNode
type KubernetesNode struct {
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
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
		ProvisioningState:         "",
		UUID:                      "",
		ParentUUID:                "",
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		Perms2:                    MakePermType2(),
		ParentType:                "",
		FQName:                    []string{},
		Annotations:               MakeKeyValuePairs(),
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
	}
}

// MakeKubernetesNodeSlice() makes a slice of KubernetesNode
func MakeKubernetesNodeSlice() []*KubernetesNode {
	return []*KubernetesNode{}
}
