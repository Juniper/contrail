package models

// KubernetesCluster

import "encoding/json"

// KubernetesCluster
type KubernetesCluster struct {
	ParentUUID           string         `json:"parent_uuid,omitempty"`
	ParentType           string         `json:"parent_type,omitempty"`
	IDPerms              *IdPermsType   `json:"id_perms,omitempty"`
	KuberunetesDashboard string         `json:"kuberunetes_dashboard,omitempty"`
	FQName               []string       `json:"fq_name,omitempty"`
	DisplayName          string         `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs `json:"annotations,omitempty"`
	Perms2               *PermType2     `json:"perms2,omitempty"`
	UUID                 string         `json:"uuid,omitempty"`
	ContrailClusterID    string         `json:"contrail_cluster_id,omitempty"`
}

// String returns json representation of the object
func (model *KubernetesCluster) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeKubernetesCluster makes KubernetesCluster
func MakeKubernetesCluster() *KubernetesCluster {
	return &KubernetesCluster{
		//TODO(nati): Apply default
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		UUID:                 "",
		ContrailClusterID:    "",
		FQName:               []string{},
		DisplayName:          "",
		IDPerms:              MakeIdPermsType(),
		KuberunetesDashboard: "",
		ParentUUID:           "",
		ParentType:           "",
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
