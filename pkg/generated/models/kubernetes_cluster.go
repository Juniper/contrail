package models

// KubernetesCluster

// KubernetesCluster
//proteus:generate
type KubernetesCluster struct {
	UUID                 string         `json:"uuid,omitempty"`
	ParentUUID           string         `json:"parent_uuid,omitempty"`
	ParentType           string         `json:"parent_type,omitempty"`
	FQName               []string       `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName          string         `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs `json:"annotations,omitempty"`
	Perms2               *PermType2     `json:"perms2,omitempty"`
	ContrailClusterID    string         `json:"contrail_cluster_id,omitempty"`
	KuberunetesDashboard string         `json:"kuberunetes_dashboard,omitempty"`
}

// MakeKubernetesCluster makes KubernetesCluster
func MakeKubernetesCluster() *KubernetesCluster {
	return &KubernetesCluster{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ContrailClusterID:    "",
		KuberunetesDashboard: "",
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
