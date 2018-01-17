package models

// KubernetesCluster

import "encoding/json"

// KubernetesCluster
type KubernetesCluster struct {
	KuberunetesDashboard string         `json:"kuberunetes_dashboard,omitempty"`
	ParentType           string         `json:"parent_type,omitempty"`
	FQName               []string       `json:"fq_name,omitempty"`
	Perms2               *PermType2     `json:"perms2,omitempty"`
	ContrailClusterID    string         `json:"contrail_cluster_id,omitempty"`
	IDPerms              *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName          string         `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs `json:"annotations,omitempty"`
	UUID                 string         `json:"uuid,omitempty"`
	ParentUUID           string         `json:"parent_uuid,omitempty"`
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
		ParentType:           "",
		FQName:               []string{},
		Perms2:               MakePermType2(),
		ContrailClusterID:    "",
		KuberunetesDashboard: "",
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		ParentUUID:           "",
		IDPerms:              MakeIdPermsType(),
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
