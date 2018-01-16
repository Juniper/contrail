package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	AnalyzerMacAddress       string              `json:"analyzer_mac_address,omitempty"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring,omitempty"`
	AnalyzerName             string              `json:"analyzer_name,omitempty"`
	JuniperHeader            bool                `json:"juniper_header,omitempty"`
	UDPPort                  int                 `json:"udp_port,omitempty"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header,omitempty"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address,omitempty"`
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan,omitempty"`
	NHMode                   NHModeType          `json:"nh_mode,omitempty"`
	RoutingInstance          string              `json:"routing_instance,omitempty"`
	Encapsulation            string              `json:"encapsulation,omitempty"`
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
		UDPPort:                  0,
		StaticNHHeader:           MakeStaticMirrorNhType(),
		AnalyzerIPAddress:        "",
		AnalyzerMacAddress:       "",
		NicAssistedMirroring:     false,
		AnalyzerName:             "",
		JuniperHeader:            false,
		RoutingInstance:          "",
		Encapsulation:            "",
		NicAssistedMirroringVlan: MakeVlanIdType(),
		NHMode: MakeNHModeType(),
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
