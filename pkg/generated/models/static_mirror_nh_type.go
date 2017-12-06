package models

// StaticMirrorNhType

import "encoding/json"

// StaticMirrorNhType
type StaticMirrorNhType struct {
	Vni               VxlanNetworkIdentifierType `json:"vni"`
	VtepDSTIPAddress  string                     `json:"vtep_dst_ip_address"`
	VtepDSTMacAddress string                     `json:"vtep_dst_mac_address"`
}

//  parents relation object

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
		VtepDSTMacAddress: data["vtep_dst_mac_address"].(string),

		//{"Title":"","Description":"mac address of destination vtep","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string","GoPremitive":true}
		Vni: InterfaceToVxlanNetworkIdentifierType(data["vni"]),

		//{"Title":"","Description":"Vni of vtep","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType","GoPremitive":false}
		VtepDSTIPAddress: data["vtep_dst_ip_address"].(string),

		//{"Title":"","Description":"ip address of destination vtep","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string","GoPremitive":true}

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
