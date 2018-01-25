package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateKubernetesClusterRequest struct {
	KubernetesCluster *KubernetesCluster `json:"kubernetes-cluster"`
}

type CreateKubernetesClusterResponse struct {
	KubernetesCluster *KubernetesCluster `json:"kubernetes-cluster"`
}

type UpdateKubernetesClusterRequest struct {
	KubernetesCluster *KubernetesCluster `json:"kubernetes-cluster"`
	FieldMask         types.FieldMask    `json:"field_mask,omitempty"`
}

type UpdateKubernetesClusterResponse struct {
	KubernetesCluster *KubernetesCluster `json:"kubernetes-cluster"`
}

type DeleteKubernetesClusterRequest struct {
	ID string `json:"id"`
}

type DeleteKubernetesClusterResponse struct {
	ID string `json:"id"`
}

type ListKubernetesClusterRequest struct {
	Spec *ListSpec
}

type ListKubernetesClusterResponse struct {
	KubernetesClusters []*KubernetesCluster `json:"kubernetes-clusters"`
}

type GetKubernetesClusterRequest struct {
	ID string `json:"id"`
}

type GetKubernetesClusterResponse struct {
	KubernetesCluster *KubernetesCluster `json:"kubernetes-cluster"`
}

func InterfaceToUpdateKubernetesClusterRequest(i interface{}) *UpdateKubernetesClusterRequest {
	//TODO implement
	return &UpdateKubernetesClusterRequest{}
}
