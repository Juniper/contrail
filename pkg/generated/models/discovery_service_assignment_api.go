package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateDiscoveryServiceAssignmentRequest struct {
	DiscoveryServiceAssignment *DiscoveryServiceAssignment `json:"discovery-service-assignment"`
}

type CreateDiscoveryServiceAssignmentResponse struct {
	DiscoveryServiceAssignment *DiscoveryServiceAssignment `json:"discovery-service-assignment"`
}

type UpdateDiscoveryServiceAssignmentRequest struct {
	DiscoveryServiceAssignment *DiscoveryServiceAssignment `json:"discovery-service-assignment"`
	FieldMask                  types.FieldMask             `json:"field_mask,omitempty"`
}

type UpdateDiscoveryServiceAssignmentResponse struct {
	DiscoveryServiceAssignment *DiscoveryServiceAssignment `json:"discovery-service-assignment"`
}

type DeleteDiscoveryServiceAssignmentRequest struct {
	ID string `json:"id"`
}

type DeleteDiscoveryServiceAssignmentResponse struct {
	ID string `json:"id"`
}

type ListDiscoveryServiceAssignmentRequest struct {
	Spec *ListSpec
}

type ListDiscoveryServiceAssignmentResponse struct {
	DiscoveryServiceAssignments []*DiscoveryServiceAssignment `json:"discovery-service-assignments"`
}

type GetDiscoveryServiceAssignmentRequest struct {
	ID string `json:"id"`
}

type GetDiscoveryServiceAssignmentResponse struct {
	DiscoveryServiceAssignment *DiscoveryServiceAssignment `json:"discovery-service-assignment"`
}

func InterfaceToUpdateDiscoveryServiceAssignmentRequest(i interface{}) *UpdateDiscoveryServiceAssignmentRequest {
	//TODO implement
	return &UpdateDiscoveryServiceAssignmentRequest{}
}
