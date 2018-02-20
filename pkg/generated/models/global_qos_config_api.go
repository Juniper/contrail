package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateGlobalQosConfigRequest struct {
	GlobalQosConfig *GlobalQosConfig `json:"global-qos-config"`
}

type CreateGlobalQosConfigResponse struct {
	GlobalQosConfig *GlobalQosConfig `json:"global-qos-config"`
}

type UpdateGlobalQosConfigRequest struct {
	GlobalQosConfig *GlobalQosConfig `json:"global-qos-config"`
	FieldMask       types.FieldMask  `json:"field_mask,omitempty"`
}

type UpdateGlobalQosConfigResponse struct {
	GlobalQosConfig *GlobalQosConfig `json:"global-qos-config"`
}

type DeleteGlobalQosConfigRequest struct {
	ID string `json:"id"`
}

type DeleteGlobalQosConfigResponse struct {
	ID string `json:"id"`
}

type ListGlobalQosConfigRequest struct {
	Spec *ListSpec
}

type ListGlobalQosConfigResponse struct {
	GlobalQosConfigs []*GlobalQosConfig `json:"global-qos-configs"`
}

type GetGlobalQosConfigRequest struct {
	ID string `json:"id"`
}

type GetGlobalQosConfigResponse struct {
	GlobalQosConfig *GlobalQosConfig `json:"global-qos-config"`
}

func InterfaceToUpdateGlobalQosConfigRequest(i interface{}) *UpdateGlobalQosConfigRequest {
	//TODO implement
	return &UpdateGlobalQosConfigRequest{}
}
