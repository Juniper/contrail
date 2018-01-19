package models

// KubernetesCluster

import "encoding/json"

// KubernetesCluster
type KubernetesCluster struct {
	ParentType           string         `json:"parent_type,omitempty"`
	FQName               []string       `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType   `json:"id_perms,omitempty"`
	ContrailClusterID    string         `json:"contrail_cluster_id,omitempty"`
	Perms2               *PermType2     `json:"perms2,omitempty"`
	UUID                 string         `json:"uuid,omitempty"`
	ParentUUID           string         `json:"parent_uuid,omitempty"`
	KuberunetesDashboard string         `json:"kuberunetes_dashboard,omitempty"`
	DisplayName          string         `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs `json:"annotations,omitempty"`
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
		IDPerms:              MakeIdPermsType(),
		ContrailClusterID:    "",
		Perms2:               MakePermType2(),
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		KuberunetesDashboard: "",
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
