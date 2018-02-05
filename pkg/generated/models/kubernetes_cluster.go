package models

// KubernetesCluster

import "encoding/json"

// KubernetesCluster
type KubernetesCluster struct {
	FQName               []string       `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType   `json:"id_perms,omitempty"`
	Annotations          *KeyValuePairs `json:"annotations,omitempty"`
	Perms2               *PermType2     `json:"perms2,omitempty"`
	UUID                 string         `json:"uuid,omitempty"`
	KuberunetesDashboard string         `json:"kuberunetes_dashboard,omitempty"`
	ParentType           string         `json:"parent_type,omitempty"`
	ParentUUID           string         `json:"parent_uuid,omitempty"`
	ContrailClusterID    string         `json:"contrail_cluster_id,omitempty"`
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
		KuberunetesDashboard: "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		UUID:                 "",
		ContrailClusterID:    "",
		DisplayName:          "",
		ParentUUID:           "",
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
