package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateContrailClusterRequest struct {
	ContrailCluster *ContrailCluster `json:"contrail-cluster"`
}

type CreateContrailClusterResponse struct {
	ContrailCluster *ContrailCluster `json:"contrail-cluster"`
}

type UpdateContrailClusterRequest struct {
	ContrailCluster *ContrailCluster `json:"contrail-cluster"`
	FieldMask       types.FieldMask  `json:"field_mask,omitempty"`
}

type UpdateContrailClusterResponse struct {
	ContrailCluster *ContrailCluster `json:"contrail-cluster"`
}

type DeleteContrailClusterRequest struct {
	ID string `json:"id"`
}

type DeleteContrailClusterResponse struct {
	ID string `json:"id"`
}

type ListContrailClusterRequest struct {
	Spec *ListSpec
}

type ListContrailClusterResponse struct {
	ContrailClusters []*ContrailCluster `json:"contrail-clusters"`
}

type GetContrailClusterRequest struct {
	ID string `json:"id"`
}

type GetContrailClusterResponse struct {
	ContrailCluster *ContrailCluster `json:"contrail-cluster"`
}

func InterfaceToUpdateContrailClusterRequest(i interface{}) *UpdateContrailClusterRequest {
	//TODO implement
	return &UpdateContrailClusterRequest{}
}
