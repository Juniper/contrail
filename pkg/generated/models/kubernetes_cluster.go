package models

// KubernetesCluster

import "encoding/json"

// KubernetesCluster
type KubernetesCluster struct {
	Annotations          *KeyValuePairs `json:"annotations,omitempty"`
	UUID                 string         `json:"uuid,omitempty"`
	FQName               []string       `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType   `json:"id_perms,omitempty"`
	ContrailClusterID    string         `json:"contrail_cluster_id,omitempty"`
	DisplayName          string         `json:"display_name,omitempty"`
	Perms2               *PermType2     `json:"perms2,omitempty"`
	ParentUUID           string         `json:"parent_uuid,omitempty"`
	ParentType           string         `json:"parent_type,omitempty"`
	KuberunetesDashboard string         `json:"kuberunetes_dashboard,omitempty"`
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
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		KuberunetesDashboard: "",
		DisplayName:          "",
		Perms2:               MakePermType2(),
		ParentUUID:           "",
		ParentType:           "",
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
