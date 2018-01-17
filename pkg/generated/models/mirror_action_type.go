package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	NHMode                   NHModeType          `json:"nh_mode,omitempty"`
	UDPPort                  int                 `json:"udp_port,omitempty"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address,omitempty"`
	Encapsulation            string              `json:"encapsulation,omitempty"`
	AnalyzerName             string              `json:"analyzer_name,omitempty"`
	JuniperHeader            bool                `json:"juniper_header,omitempty"`
	RoutingInstance          string              `json:"routing_instance,omitempty"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header,omitempty"`
	AnalyzerMacAddress       string              `json:"analyzer_mac_address,omitempty"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring,omitempty"`
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan,omitempty"`
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
		NicAssistedMirroringVlan: MakeVlanIdType(),
		JuniperHeader:            false,
		RoutingInstance:          "",
		StaticNHHeader:           MakeStaticMirrorNhType(),
		AnalyzerMacAddress:       "",
		NicAssistedMirroring:     false,
		AnalyzerName:             "",
		NHMode:                   MakeNHModeType(),
		UDPPort:                  0,
		AnalyzerIPAddress:        "",
		Encapsulation:            "",
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
