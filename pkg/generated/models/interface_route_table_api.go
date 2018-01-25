package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateInterfaceRouteTableRequest struct {
	InterfaceRouteTable *InterfaceRouteTable `json:"interface-route-table"`
}

type CreateInterfaceRouteTableResponse struct {
	InterfaceRouteTable *InterfaceRouteTable `json:"interface-route-table"`
}

type UpdateInterfaceRouteTableRequest struct {
	InterfaceRouteTable *InterfaceRouteTable `json:"interface-route-table"`
	FieldMask           types.FieldMask      `json:"field_mask,omitempty"`
}

type UpdateInterfaceRouteTableResponse struct {
	InterfaceRouteTable *InterfaceRouteTable `json:"interface-route-table"`
}

type DeleteInterfaceRouteTableRequest struct {
	ID string `json:"id"`
}

type DeleteInterfaceRouteTableResponse struct {
	ID string `json:"id"`
}

type ListInterfaceRouteTableRequest struct {
	Spec *ListSpec
}

type ListInterfaceRouteTableResponse struct {
	InterfaceRouteTables []*InterfaceRouteTable `json:"interface-route-tables"`
}

type GetInterfaceRouteTableRequest struct {
	ID string `json:"id"`
}

type GetInterfaceRouteTableResponse struct {
	InterfaceRouteTable *InterfaceRouteTable `json:"interface-route-table"`
}

func InterfaceToUpdateInterfaceRouteTableRequest(i interface{}) *UpdateInterfaceRouteTableRequest {
	//TODO implement
	return &UpdateInterfaceRouteTableRequest{}
}
