package models

import (
	"github.com/gogo/protobuf/types"
)

type CreateNetworkDeviceConfigRequest struct {
	NetworkDeviceConfig *NetworkDeviceConfig `json:"network-device-config"`
}

type CreateNetworkDeviceConfigResponse struct {
	NetworkDeviceConfig *NetworkDeviceConfig `json:"network-device-config"`
}

type UpdateNetworkDeviceConfigRequest struct {
	NetworkDeviceConfig *NetworkDeviceConfig `json:"network-device-config"`
	FieldMask           types.FieldMask      `json:"field_mask,omitempty"`
}

type UpdateNetworkDeviceConfigResponse struct {
	NetworkDeviceConfig *NetworkDeviceConfig `json:"network-device-config"`
}

type DeleteNetworkDeviceConfigRequest struct {
	ID string `json:"id"`
}

type DeleteNetworkDeviceConfigResponse struct {
	ID string `json:"id"`
}

type ListNetworkDeviceConfigRequest struct {
	Spec *ListSpec
}

type ListNetworkDeviceConfigResponse struct {
	NetworkDeviceConfigs []*NetworkDeviceConfig `json:"network-device-configs"`
}

type GetNetworkDeviceConfigRequest struct {
	ID string `json:"id"`
}

type GetNetworkDeviceConfigResponse struct {
	NetworkDeviceConfig *NetworkDeviceConfig `json:"network-device-config"`
}

func InterfaceToUpdateNetworkDeviceConfigRequest(i interface{}) *UpdateNetworkDeviceConfigRequest {
	//TODO implement
	return &UpdateNetworkDeviceConfigRequest{}
}
