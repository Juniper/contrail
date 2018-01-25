package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateQosConfigRequest struct {
	QosConfig *QosConfig `json:"qos-config"`
}

type CreateQosConfigResponse struct {
	QosConfig *QosConfig `json:"qos-config"`
}

type UpdateQosConfigRequest struct {
	QosConfig *QosConfig      `json:"qos-config"`
	FieldMask types.FieldMask `json:"field_mask,omitempty"`
}

type UpdateQosConfigResponse struct {
	QosConfig *QosConfig `json:"qos-config"`
}

type DeleteQosConfigRequest struct {
	ID string `json:"id"`
}

type DeleteQosConfigResponse struct {
	ID string `json:"id"`
}

type ListQosConfigRequest struct {
	Spec *ListSpec
}

type ListQosConfigResponse struct {
	QosConfigs []*QosConfig `json:"qos-configs"`
}

type GetQosConfigRequest struct {
	ID string `json:"id"`
}

type GetQosConfigResponse struct {
	QosConfig *QosConfig `json:"qos-config"`
}

func InterfaceToUpdateQosConfigRequest(i interface{}) *UpdateQosConfigRequest {
	//TODO implement
	return &UpdateQosConfigRequest{}
}
