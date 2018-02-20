package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateKubernetesNodeRequest struct {
	KubernetesNode *KubernetesNode `json:"kubernetes-node"`
}

type CreateKubernetesNodeResponse struct {
	KubernetesNode *KubernetesNode `json:"kubernetes-node"`
}

type UpdateKubernetesNodeRequest struct {
	KubernetesNode *KubernetesNode `json:"kubernetes-node"`
	FieldMask      types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateKubernetesNodeResponse struct {
	KubernetesNode *KubernetesNode `json:"kubernetes-node"`
}

type DeleteKubernetesNodeRequest struct {
	ID string `json:"id"`
}

type DeleteKubernetesNodeResponse struct {
	ID string `json:"id"`
}

type ListKubernetesNodeRequest struct {
	Spec *ListSpec
}

type ListKubernetesNodeResponse struct {
	KubernetesNodes []*KubernetesNode `json:"kubernetes-nodes"`
}

type GetKubernetesNodeRequest struct {
	ID string `json:"id"`
}

type GetKubernetesNodeResponse struct {
	KubernetesNode *KubernetesNode `json:"kubernetes-node"`
}

func InterfaceToUpdateKubernetesNodeRequest(i interface{}) *UpdateKubernetesNodeRequest {
	//TODO implement
	return &UpdateKubernetesNodeRequest{}
}
