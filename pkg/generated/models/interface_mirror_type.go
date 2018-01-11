package models

// InterfaceMirrorType

import "encoding/json"

// InterfaceMirrorType
type InterfaceMirrorType struct {
	TrafficDirection TrafficDirectionType `json:"traffic_direction"`
	MirrorTo         *MirrorActionType    `json:"mirror_to"`
}

// String returns json representation of the object
func (model *InterfaceMirrorType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeInterfaceMirrorType makes InterfaceMirrorType
func MakeInterfaceMirrorType() *InterfaceMirrorType {
	return &InterfaceMirrorType{
		//TODO(nati): Apply default
		TrafficDirection: MakeTrafficDirectionType(),
		MirrorTo:         MakeMirrorActionType(),
	}
}

// MakeInterfaceMirrorTypeSlice() makes a slice of InterfaceMirrorType
func MakeInterfaceMirrorTypeSlice() []*InterfaceMirrorType {
	return []*InterfaceMirrorType{}
}
