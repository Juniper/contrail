package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateLoadbalancerMemberRequest struct {
	LoadbalancerMember *LoadbalancerMember `json:"loadbalancer-member"`
}

type CreateLoadbalancerMemberResponse struct {
	LoadbalancerMember *LoadbalancerMember `json:"loadbalancer-member"`
}

type UpdateLoadbalancerMemberRequest struct {
	LoadbalancerMember *LoadbalancerMember `json:"loadbalancer-member"`
	FieldMask          types.FieldMask     `json:"field_mask,omitempty"`
}

type UpdateLoadbalancerMemberResponse struct {
	LoadbalancerMember *LoadbalancerMember `json:"loadbalancer-member"`
}

type DeleteLoadbalancerMemberRequest struct {
	ID string `json:"id"`
}

type DeleteLoadbalancerMemberResponse struct {
	ID string `json:"id"`
}

type ListLoadbalancerMemberRequest struct {
	Spec *ListSpec
}

type ListLoadbalancerMemberResponse struct {
	LoadbalancerMembers []*LoadbalancerMember `json:"loadbalancer-members"`
}

type GetLoadbalancerMemberRequest struct {
	ID string `json:"id"`
}

type GetLoadbalancerMemberResponse struct {
	LoadbalancerMember *LoadbalancerMember `json:"loadbalancer-member"`
}

func InterfaceToUpdateLoadbalancerMemberRequest(i interface{}) *UpdateLoadbalancerMemberRequest {
	//TODO implement
	return &UpdateLoadbalancerMemberRequest{}
}
