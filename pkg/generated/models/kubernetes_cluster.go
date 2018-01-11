package models

// KubernetesCluster

import "encoding/json"

// KubernetesCluster
type KubernetesCluster struct {
	Annotations          *KeyValuePairs `json:"annotations"`
	KuberunetesDashboard string         `json:"kuberunetes_dashboard"`
	DisplayName          string         `json:"display_name"`
	UUID                 string         `json:"uuid"`
	ParentUUID           string         `json:"parent_uuid"`
	ParentType           string         `json:"parent_type"`
	FQName               []string       `json:"fq_name"`
	IDPerms              *IdPermsType   `json:"id_perms"`
	ContrailClusterID    string         `json:"contrail_cluster_id"`
	Perms2               *PermType2     `json:"perms2"`
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
		Perms2:               MakePermType2(),
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		KuberunetesDashboard: "",
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
