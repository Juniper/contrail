package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring"`
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan,omitempty"`
	NHMode                   NHModeType          `json:"nh_mode,omitempty"`
	JuniperHeader            bool                `json:"juniper_header"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address,omitempty"`
	AnalyzerMacAddress       string              `json:"analyzer_mac_address,omitempty"`
	AnalyzerName             string              `json:"analyzer_name,omitempty"`
	UDPPort                  int                 `json:"udp_port,omitempty"`
	RoutingInstance          string              `json:"routing_instance,omitempty"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header,omitempty"`
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
		NicAssistedMirroringVlan: MakeVlanIdType(),
		NHMode:               MakeNHModeType(),
		JuniperHeader:        false,
		AnalyzerIPAddress:    "",
		AnalyzerMacAddress:   "",
		NicAssistedMirroring: false,
		AnalyzerName:         "",
		UDPPort:              0,
		RoutingInstance:      "",
		StaticNHHeader:       MakeStaticMirrorNhType(),
		Encapsulation:        "",
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
