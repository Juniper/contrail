package models

// KubernetesCluster

import "encoding/json"

// KubernetesCluster
type KubernetesCluster struct {
	DisplayName          string         `json:"display_name,omitempty"`
	Perms2               *PermType2     `json:"perms2,omitempty"`
	UUID                 string         `json:"uuid,omitempty"`
	ContrailClusterID    string         `json:"contrail_cluster_id,omitempty"`
	ParentType           string         `json:"parent_type,omitempty"`
	FQName               []string       `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType   `json:"id_perms,omitempty"`
	Annotations          *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID           string         `json:"parent_uuid,omitempty"`
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
		UUID:                 "",
		ContrailClusterID:    "",
		DisplayName:          "",
		Perms2:               MakePermType2(),
		IDPerms:              MakeIdPermsType(),
		Annotations:          MakeKeyValuePairs(),
		ParentUUID:           "",
		KuberunetesDashboard: "",
		ParentType:           "",
		FQName:               []string{},
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
