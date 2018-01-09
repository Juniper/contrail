package models

// InterfaceMirrorType

import "encoding/json"

// InterfaceMirrorType
type InterfaceMirrorType struct {
	MirrorTo         *MirrorActionType    `json:"mirror_to"`
	TrafficDirection TrafficDirectionType `json:"traffic_direction"`
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

// InterfaceToInterfaceMirrorType makes InterfaceMirrorType from interface
func InterfaceToInterfaceMirrorType(iData interface{}) *InterfaceMirrorType {
	data := iData.(map[string]interface{})
	return &InterfaceMirrorType{
		TrafficDirection: InterfaceToTrafficDirectionType(data["traffic_direction"]),

		//{"description":"Specifies direction of traffic to mirror, Ingress, Egress or both","default":"both","type":"string","enum":["ingress","egress","both"]}
		MirrorTo: InterfaceToMirrorActionType(data["mirror_to"]),

		//{"description":"Mirror destination configuration","type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}}

	}
}

// InterfaceToInterfaceMirrorTypeSlice makes a slice of InterfaceMirrorType from interface
func InterfaceToInterfaceMirrorTypeSlice(data interface{}) []*InterfaceMirrorType {
	list := data.([]interface{})
	result := MakeInterfaceMirrorTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToInterfaceMirrorType(item))
	}
	return result
}

// MakeInterfaceMirrorTypeSlice() makes a slice of InterfaceMirrorType
func MakeInterfaceMirrorTypeSlice() []*InterfaceMirrorType {
	return []*InterfaceMirrorType{}
}
