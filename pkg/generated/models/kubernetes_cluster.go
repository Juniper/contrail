package models

// KubernetesCluster

import "encoding/json"

// KubernetesCluster
type KubernetesCluster struct {
	Annotations          *KeyValuePairs `json:"annotations"`
	UUID                 string         `json:"uuid"`
	ParentUUID           string         `json:"parent_uuid"`
	FQName               []string       `json:"fq_name"`
	IDPerms              *IdPermsType   `json:"id_perms"`
	ContrailClusterID    string         `json:"contrail_cluster_id"`
	KuberunetesDashboard string         `json:"kuberunetes_dashboard"`
	Perms2               *PermType2     `json:"perms2"`
	ParentType           string         `json:"parent_type"`
	DisplayName          string         `json:"display_name"`
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
		KuberunetesDashboard: "",
		Perms2:               MakePermType2(),
		ParentType:           "",
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		ParentUUID:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
	}
}

// InterfaceToKubernetesCluster makes KubernetesCluster from interface
func InterfaceToKubernetesCluster(iData interface{}) *KubernetesCluster {
	data := iData.(map[string]interface{})
	return &KubernetesCluster{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ContrailClusterID: data["contrail_cluster_id"].(string),

		//{"title":"Contrail Cluster ID","default":"","type":"string","permission":["create","update"]}
		KuberunetesDashboard: data["kuberunetes_dashboard"].(string),

		//{"title":"kubernetes Dashboard","default":"","type":"string","permission":["create","update"]}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToKubernetesClusterSlice makes a slice of KubernetesCluster from interface
func InterfaceToKubernetesClusterSlice(data interface{}) []*KubernetesCluster {
	list := data.([]interface{})
	result := MakeKubernetesClusterSlice()
	for _, item := range list {
		result = append(result, InterfaceToKubernetesCluster(item))
	}
	return result
}

// MakeKubernetesClusterSlice() makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
	return []*KubernetesCluster{}
}
