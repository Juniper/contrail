package models

// KubernetesCluster

import "encoding/json"

// KubernetesCluster
type KubernetesCluster struct {
	KuberunetesDashboard string         `json:"kuberunetes_dashboard,omitempty"`
	Perms2               *PermType2     `json:"perms2,omitempty"`
	UUID                 string         `json:"uuid,omitempty"`
	ParentType           string         `json:"parent_type,omitempty"`
	FQName               []string       `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType   `json:"id_perms,omitempty"`
	Annotations          *KeyValuePairs `json:"annotations,omitempty"`
	ContrailClusterID    string         `json:"contrail_cluster_id,omitempty"`
	ParentUUID           string         `json:"parent_uuid,omitempty"`
	DisplayName          string         `json:"display_name,omitempty"`
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
		ContrailClusterID:    "",
		ParentUUID:           "",
		DisplayName:          "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		Annotations:          MakeKeyValuePairs(),
		KuberunetesDashboard: "",
		Perms2:               MakePermType2(),
		UUID:                 "",
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
