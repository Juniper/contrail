package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateVirtualMachineInterfaceRequest struct {
	VirtualMachineInterface *VirtualMachineInterface `json:"virtual-machine-interface"`
}

type CreateVirtualMachineInterfaceResponse struct {
	VirtualMachineInterface *VirtualMachineInterface `json:"virtual-machine-interface"`
}

type UpdateVirtualMachineInterfaceRequest struct {
	VirtualMachineInterface *VirtualMachineInterface `json:"virtual-machine-interface"`
	FieldMask               types.FieldMask          `json:"field_mask,omitempty"`
}

type UpdateVirtualMachineInterfaceResponse struct {
	VirtualMachineInterface *VirtualMachineInterface `json:"virtual-machine-interface"`
}

type DeleteVirtualMachineInterfaceRequest struct {
	ID string `json:"id"`
}

type DeleteVirtualMachineInterfaceResponse struct {
	ID string `json:"id"`
}

type ListVirtualMachineInterfaceRequest struct {
	Spec *ListSpec
}

type ListVirtualMachineInterfaceResponse struct {
	VirtualMachineInterfaces []*VirtualMachineInterface `json:"virtual-machine-interfaces"`
}

type GetVirtualMachineInterfaceRequest struct {
	ID string `json:"id"`
}

type GetVirtualMachineInterfaceResponse struct {
	VirtualMachineInterface *VirtualMachineInterface `json:"virtual-machine-interface"`
}

func InterfaceToUpdateVirtualMachineInterfaceRequest(i interface{}) *UpdateVirtualMachineInterfaceRequest {
	//TODO implement
	return &UpdateVirtualMachineInterfaceRequest{}
}
