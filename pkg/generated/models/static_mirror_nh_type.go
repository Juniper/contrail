package models

// StaticMirrorNhType

import "encoding/json"

// StaticMirrorNhType
type StaticMirrorNhType struct {
	VtepDSTIPAddress  string                     `json:"vtep_dst_ip_address"`
	VtepDSTMacAddress string                     `json:"vtep_dst_mac_address"`
	Vni               VxlanNetworkIdentifierType `json:"vni"`
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
		VtepDSTMacAddress: "",
		Vni:               MakeVxlanNetworkIdentifierType(),
		VtepDSTIPAddress:  "",
	}
}

// MakeStaticMirrorNhTypeSlice() makes a slice of StaticMirrorNhType
func MakeStaticMirrorNhTypeSlice() []*StaticMirrorNhType {
	return []*StaticMirrorNhType{}
}
