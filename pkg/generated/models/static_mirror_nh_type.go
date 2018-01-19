package models

// StaticMirrorNhType

import "encoding/json"

// StaticMirrorNhType
type StaticMirrorNhType struct {
	Vni               VxlanNetworkIdentifierType `json:"vni,omitempty"`
	VtepDSTIPAddress  string                     `json:"vtep_dst_ip_address,omitempty"`
	VtepDSTMacAddress string                     `json:"vtep_dst_mac_address,omitempty"`
}

// String returns json representation of the object
func (model *StaticMirrorNhType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeStaticMirrorNhType makes StaticMirrorNhType
func MakeStaticMirrorNhType() *StaticMirrorNhType {
	return &StaticMirrorNhType{
		//TODO(nati): Apply default
		VtepDSTIPAddress:  "",
		VtepDSTMacAddress: "",
		Vni:               MakeVxlanNetworkIdentifierType(),
	}
}

// MakeStaticMirrorNhTypeSlice() makes a slice of StaticMirrorNhType
func MakeStaticMirrorNhTypeSlice() []*StaticMirrorNhType {
	return []*StaticMirrorNhType{}
}
