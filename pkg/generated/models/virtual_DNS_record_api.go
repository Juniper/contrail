package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateVirtualDNSRecordRequest struct {
	VirtualDNSRecord *VirtualDNSRecord `json:"virtual-DNS-record"`
}

type CreateVirtualDNSRecordResponse struct {
	VirtualDNSRecord *VirtualDNSRecord `json:"virtual-DNS-record"`
}

type UpdateVirtualDNSRecordRequest struct {
	VirtualDNSRecord *VirtualDNSRecord `json:"virtual-DNS-record"`
	FieldMask        types.FieldMask   `json:"field_mask,omitempty"`
}

type UpdateVirtualDNSRecordResponse struct {
	VirtualDNSRecord *VirtualDNSRecord `json:"virtual-DNS-record"`
}

type DeleteVirtualDNSRecordRequest struct {
	ID string `json:"id"`
}

type DeleteVirtualDNSRecordResponse struct {
	ID string `json:"id"`
}

type ListVirtualDNSRecordRequest struct {
	Spec *ListSpec
}

type ListVirtualDNSRecordResponse struct {
	VirtualDNSRecords []*VirtualDNSRecord `json:"virtual-DNS-records"`
}

type GetVirtualDNSRecordRequest struct {
	ID string `json:"id"`
}

type GetVirtualDNSRecordResponse struct {
	VirtualDNSRecord *VirtualDNSRecord `json:"virtual-DNS-record"`
}

func InterfaceToUpdateVirtualDNSRecordRequest(i interface{}) *UpdateVirtualDNSRecordRequest {
	//TODO implement
	return &UpdateVirtualDNSRecordRequest{}
}
