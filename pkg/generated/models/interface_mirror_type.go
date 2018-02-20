package models

// InterfaceMirrorType

// InterfaceMirrorType
//proteus:generate
type InterfaceMirrorType struct {
	TrafficDirection TrafficDirectionType `json:"traffic_direction,omitempty"`
	MirrorTo         *MirrorActionType    `json:"mirror_to,omitempty"`
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
