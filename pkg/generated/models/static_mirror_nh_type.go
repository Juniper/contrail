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
		VtepDSTIPAddress:  "",
		VtepDSTMacAddress: "",
		Vni:               MakeVxlanNetworkIdentifierType(),
	}
}

// InterfaceToStaticMirrorNhType makes StaticMirrorNhType from interface
func InterfaceToStaticMirrorNhType(iData interface{}) *StaticMirrorNhType {
	data := iData.(map[string]interface{})
	return &StaticMirrorNhType{
		VtepDSTIPAddress: data["vtep_dst_ip_address"].(string),

		//{"description":"ip address of destination vtep","type":"string"}
		VtepDSTMacAddress: data["vtep_dst_mac_address"].(string),

		//{"description":"mac address of destination vtep","type":"string"}
		Vni: InterfaceToVxlanNetworkIdentifierType(data["vni"]),

		//{"description":"Vni of vtep","type":"integer","minimum":1,"maximum":16777215}

	}
}

// InterfaceToStaticMirrorNhTypeSlice makes a slice of StaticMirrorNhType from interface
func InterfaceToStaticMirrorNhTypeSlice(data interface{}) []*StaticMirrorNhType {
	list := data.([]interface{})
	result := MakeStaticMirrorNhTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToStaticMirrorNhType(item))
	}
	return result
}

// MakeStaticMirrorNhTypeSlice() makes a slice of StaticMirrorNhType
func MakeStaticMirrorNhTypeSlice() []*StaticMirrorNhType {
	return []*StaticMirrorNhType{}
}
