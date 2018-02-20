package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateVirtualMachineRequest struct {
	VirtualMachine *VirtualMachine `json:"virtual-machine"`
}

type CreateVirtualMachineResponse struct {
	VirtualMachine *VirtualMachine `json:"virtual-machine"`
}

type UpdateVirtualMachineRequest struct {
	VirtualMachine *VirtualMachine `json:"virtual-machine"`
	FieldMask      types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateVirtualMachineResponse struct {
	VirtualMachine *VirtualMachine `json:"virtual-machine"`
}

type DeleteVirtualMachineRequest struct {
	ID string `json:"id"`
}

type DeleteVirtualMachineResponse struct {
	ID string `json:"id"`
}

type ListVirtualMachineRequest struct {
	Spec *ListSpec
}

type ListVirtualMachineResponse struct {
	VirtualMachines []*VirtualMachine `json:"virtual-machines"`
}

type GetVirtualMachineRequest struct {
	ID string `json:"id"`
}

type GetVirtualMachineResponse struct {
	VirtualMachine *VirtualMachine `json:"virtual-machine"`
}

func InterfaceToUpdateVirtualMachineRequest(i interface{}) *UpdateVirtualMachineRequest {
	//TODO implement
	return &UpdateVirtualMachineRequest{}
}
