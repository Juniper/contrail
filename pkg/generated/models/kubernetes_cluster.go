package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

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

// MakeKubernetesCluster makes KubernetesCluster
func InterfaceToKubernetesCluster(i interface{}) *KubernetesCluster {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &KubernetesCluster{
		//TODO(nati): Apply default
		UUID:                 schema.InterfaceToString(m["uuid"]),
		ParentUUID:           schema.InterfaceToString(m["parent_uuid"]),
		ParentType:           schema.InterfaceToString(m["parent_type"]),
		FQName:               schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          schema.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		ContrailClusterID:    schema.InterfaceToString(m["contrail_cluster_id"]),
		KuberunetesDashboard: schema.InterfaceToString(m["kuberunetes_dashboard"]),
	}
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}

// InterfaceToKubernetesClusterSlice() makes a slice of KubernetesCluster
func InterfaceToKubernetesClusterSlice(i interface{}) []*KubernetesCluster {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*KubernetesCluster{}
	for _, item := range list {
		result = append(result, InterfaceToKubernetesCluster(item))
	}
	return result
}
