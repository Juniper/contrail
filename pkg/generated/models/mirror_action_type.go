package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	AnalyzerMacAddress       string              `json:"analyzer_mac_address,omitempty"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header,omitempty"`
	Encapsulation            string              `json:"encapsulation,omitempty"`
	NHMode                   NHModeType          `json:"nh_mode,omitempty"`
	JuniperHeader            bool                `json:"juniper_header"`
	UDPPort                  int                 `json:"udp_port,omitempty"`
	RoutingInstance          string              `json:"routing_instance,omitempty"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address,omitempty"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring"`
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan,omitempty"`
	AnalyzerName             string              `json:"analyzer_name,omitempty"`
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
		JuniperHeader:            false,
		UDPPort:                  0,
		RoutingInstance:          "",
		AnalyzerIPAddress:        "",
		NicAssistedMirroring:     false,
		NicAssistedMirroringVlan: MakeVlanIdType(),
		AnalyzerName:             "",
		NHMode:                   MakeNHModeType(),
		StaticNHHeader:           MakeStaticMirrorNhType(),
		Encapsulation:            "",
		AnalyzerMacAddress:       "",
	}
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
