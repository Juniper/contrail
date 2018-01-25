package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateOpenstackClusterRequest struct {
	OpenstackCluster *OpenstackCluster `json:"openstack-cluster"`
}

type CreateOpenstackClusterResponse struct {
	OpenstackCluster *OpenstackCluster `json:"openstack-cluster"`
}

type UpdateOpenstackClusterRequest struct {
	OpenstackCluster *OpenstackCluster `json:"openstack-cluster"`
	FieldMask        types.FieldMask   `json:"field_mask,omitempty"`
}

type UpdateOpenstackClusterResponse struct {
	OpenstackCluster *OpenstackCluster `json:"openstack-cluster"`
}

type DeleteOpenstackClusterRequest struct {
	ID string `json:"id"`
}

type DeleteOpenstackClusterResponse struct {
	ID string `json:"id"`
}

type ListOpenstackClusterRequest struct {
	Spec *ListSpec
}

type ListOpenstackClusterResponse struct {
	OpenstackClusters []*OpenstackCluster `json:"openstack-clusters"`
}

type GetOpenstackClusterRequest struct {
	ID string `json:"id"`
}

type GetOpenstackClusterResponse struct {
	OpenstackCluster *OpenstackCluster `json:"openstack-cluster"`
}

func InterfaceToUpdateOpenstackClusterRequest(i interface{}) *UpdateOpenstackClusterRequest {
	//TODO implement
	return &UpdateOpenstackClusterRequest{}
}
