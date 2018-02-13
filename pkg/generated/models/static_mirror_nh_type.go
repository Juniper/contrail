package models

// StaticMirrorNhType

import "encoding/json"

// StaticMirrorNhType
//proteus:generate
type StaticMirrorNhType struct {
	VtepDSTIPAddress  string                     `json:"vtep_dst_ip_address,omitempty"`
	VtepDSTMacAddress string                     `json:"vtep_dst_mac_address,omitempty"`
	Vni               VxlanNetworkIdentifierType `json:"vni,omitempty"`
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
