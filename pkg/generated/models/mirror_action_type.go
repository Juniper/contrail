package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	JuniperHeader            bool                `json:"juniper_header"`
	UDPPort                  int                 `json:"udp_port"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address"`
	Encapsulation            string              `json:"encapsulation"`
	AnalyzerMacAddress       string              `json:"analyzer_mac_address"`
	NHMode                   NHModeType          `json:"nh_mode"`
	AnalyzerName             string              `json:"analyzer_name"`
	RoutingInstance          string              `json:"routing_instance"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring"`
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan"`
}

// String returns json representation of the object
func (model *MirrorActionType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMirrorActionType makes MirrorActionType
func MakeMirrorActionType() *MirrorActionType {
	return &MirrorActionType{
		//TODO(nati): Apply default
		NHMode:                   MakeNHModeType(),
		JuniperHeader:            false,
		UDPPort:                  0,
		StaticNHHeader:           MakeStaticMirrorNhType(),
		AnalyzerIPAddress:        "",
		Encapsulation:            "",
		AnalyzerMacAddress:       "",
		NicAssistedMirroringVlan: MakeVlanIdType(),
		AnalyzerName:             "",
		RoutingInstance:          "",
		NicAssistedMirroring:     false,
	}
}

// InterfaceToMirrorActionType makes MirrorActionType from interface
func InterfaceToMirrorActionType(iData interface{}) *MirrorActionType {
	data := iData.(map[string]interface{})
	return &MirrorActionType{
		RoutingInstance: data["routing_instance"].(string),

		//{"description":"Internal use only, should be set to -1","type":"string"}
		NicAssistedMirroring: data["nic_assisted_mirroring"].(bool),

		//{"description":"This flag is used to select nic assisted mirroring","type":"boolean"}
		NicAssistedMirroringVlan: InterfaceToVlanIdType(data["nic_assisted_mirroring_vlan"]),

		//{"description":"The VLAN to be tagged on the traffic for NIC to Mirror","type":"integer","minimum":1,"maximum":4094}
		AnalyzerName: data["analyzer_name"].(string),

		//{"description":"Name of service instance used as analyzer","type":"string"}
		UDPPort: data["udp_port"].(int),

		//{"description":"ip udp port used in contrail default encapsulation for mirroring","type":"integer"}
		StaticNHHeader: InterfaceToStaticMirrorNhType(data["static_nh_header"]),

		//{"description":"vtep details required if static nh enabled","type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}}
		AnalyzerIPAddress: data["analyzer_ip_address"].(string),

		//{"description":"ip address of interface to which mirrored packets are sent","type":"string"}
		Encapsulation: data["encapsulation"].(string),

		//{"description":"Encapsulation for Mirrored packet, not used currently","type":"string"}
		AnalyzerMacAddress: data["analyzer_mac_address"].(string),

		//{"description":"mac address of interface to which mirrored packets are sent ","type":"string"}
		NHMode: InterfaceToNHModeType(data["nh_mode"]),

		//{"description":"This mode used to determine static or dynamic nh ","type":"string","enum":["dynamic","static"]}
		JuniperHeader: data["juniper_header"].(bool),

		//{"description":"This flag is used to determine with/without juniper-header","type":"boolean"}

	}
}

// InterfaceToMirrorActionTypeSlice makes a slice of MirrorActionType from interface
func InterfaceToMirrorActionTypeSlice(data interface{}) []*MirrorActionType {
	list := data.([]interface{})
	result := MakeMirrorActionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMirrorActionType(item))
	}
	return result
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
